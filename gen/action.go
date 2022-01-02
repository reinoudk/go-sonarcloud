package main

import (
	"encoding/json"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Action struct {
	Key                string  `json:"key"`
	Description        string  `json:"description"`
	Internal           bool    `json:"internal"`
	Post               bool    `json:"post"`
	HasResponseExample bool    `json:"hasResponseExample"`
	Params             []Param `json:"params"`
	DeprecatedSince    string  `json:"deprecatedSince"`
}

type Param struct {
	Key             string `json:"key"`
	Description     string `json:"description"`
	Internal        bool   `json:"internal"`
	Required        bool   `json:"required"`
	DeprecatedSince string `json:"deprecatedSince"`
}

func (p *Param) render() *Statement {
	return renderId(p.Key).String().Tag(map[string]string{"form": p.Key + ",omitempty"}).Comment(p.Description)
}

type ResponseExampleRequest struct {
	ID         string
	RequestID  string
	Action     string
	Controller string
}

type Example struct {
	Format  string `json:"format"`
	Example string `json:"example"` // yes, it's a string...
}

func (a *Action) id() string {
	return strcase.ToCamel(a.Key)
}

func (a *Action) requestTypeName() string {
	return fmt.Sprintf("%s%s", a.id(), "Request")
}

func (a *Action) responseTypeName() string {
	return fmt.Sprintf("%s%s", a.id(), "Response")
}

func (a *Action) responseAllTypeName() string {
	return fmt.Sprintf("%s%s", a.id(), "ResponseAll")
}

func (a *Action) pagingFuncName() string {
	return "GetPaging"
}

func (a *Action) serviceFuncName() string {
	return a.id()
}

func (a *Action) serviceAllFuncName() string {
	return fmt.Sprintf("%s%s", a.id(), "All")
}

func (a *Action) hasPaging() bool {
	hasP := false
	hasPs := false

	for _, param := range a.Params {
		switch param.Key {
		case "p": hasP = true
		case "ps": hasPs = true
		}
	}

	return hasP && hasPs
}

func (a *Action) responseField(example map[string]interface{}) (Field, error) {
	overrides := make(map[string]Field)
	if a.hasPaging() {
		overrides["paging"] = NewStatementField("paging", Qual(qualifier("paging"), "Paging"))
	}

	// TODO: overrides should be defined globally and not just by action key, but by service path as well.
	if a.Key == "generate" {
		// The token example is "1234567", which we 'falsely' interpret as an integer. Real tokens have letters...
		overrides["token"] = &StringField{name: "Token"}
	}

	return NewMapField(a.responseTypeName(), example, overrides), nil
}

func (a *Action) responseFieldWithoutPaging(example map[string]interface{}) (Field, error) {
	delete(example, "paging")

	// Remove flattened paging as well
	delete(example, "p")
	delete(example, "ps")
	delete(example, "total")

	return NewMapField(a.responseAllTypeName(), example, nil), nil
}

func (a *Action) requestStruct() *Statement {
	fields := make([]Code, len(a.Params))
	for i, param := range a.Params {
		// filter out unwanted fields and paging parameters
		if contains(param.Key, append(skippedRequestFields, "p", "ps")) {
			continue
		}

		fields[i] = param.render()
	}

	statement := Commentf("%s %s", a.requestTypeName(), a.Description)
	if a.DeprecatedSince != "" {
		statement.Line()
		statement.Commentf("Deprecated: this action has been deprecated since version %s", a.DeprecatedSince)
	}
	statement.Line()

	statement.Type().Id(a.requestTypeName()).Struct(fields...)

	return statement
}

func (a *Action) responseStruct(response Field) *Statement {
	// EmptyField should not be rendered
	if reflect.TypeOf(response) != reflect.TypeOf(&EmptyField{}) {
		fields := response.Render(false)
		statement := Commentf("%s is the response for %s", a.responseTypeName(), a.requestTypeName())
		statement.Line()
		statement.Type().Add(fields)
		return statement
	}

	return Empty()
}

func (a *Action) responseStructPagingFunc(collection Field) *Statement {
	// EmptyField should not have a Paging
	if reflect.TypeOf(collection) == reflect.TypeOf(&MapField{}) {
		statement := Commentf("%s extracts the paging from %s", a.pagingFuncName(), a.responseTypeName())
		statement.Line()
		statement.Func().Parens(Id("r").Op("*").Id(a.responseTypeName())).Id(a.pagingFuncName()).Call().Op("*").Qual(qualifier("paging"), "Paging")

		if contains("Paging", collection.(*MapField).Accessors()) {
			statement.Block(Return(Op("&").Id("r").Dot("Paging")))
		} else {
			statement.Block(Return(
				Op("&").Qual(qualifier("paging"), "Paging").Block(Dict{
					Id("PageIndex"): Int().Parens(Id("r").Dot("P")),
					Id("PageSize"): Int().Parens(Id("r").Dot("Ps")),
					Id("Total"): Int().Parens(Id("r").Dot("Total")),
				}),
			))
		}

		return statement
	}

	return Empty()
}

func (a *Action) responseAllStruct(collection Field) *Statement {
	// EmptyField should not be rendered
	if reflect.TypeOf(collection) != reflect.TypeOf(&EmptyField{}) {
		fields := collection.Render(false)
		statement := Commentf("%s is the collection for %s", a.responseAllTypeName(), a.requestTypeName())
		statement.Line()
		statement.Type().Add(fields)
		return statement
	}

	return Empty()
}

func (a *Action) fetchExample(endpoint string) (map[string]interface{}, error) {
	controller := fmt.Sprintf("api/%s", endpoint)
	request := ResponseExampleRequest{ID: a.responseTypeName(), RequestID: a.id(), Controller: controller, Action: a.Key}

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := newRequest(request)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %+v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending reqeust: %+v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %+v", err)
	}

	var responseExample Example
	err = json.Unmarshal(body, &responseExample)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall body: %+v", err)
	}

	if responseExample.Format == "json" {
		// Convert the example JSON string (!!) to a map
		var example map[string]interface{}
		err := json.Unmarshal([]byte(responseExample.Example), &example)
		if err != nil {
			return nil, fmt.Errorf("could not marshall example: %+v", err)
		}

		return example, nil
	} else {
		return nil, fmt.Errorf("unsupported response format %s", responseExample.Format)
	}
}

func parseField(name string, value interface{}, overrides map[string]Field) Field {
	if override, ok := overrides[name]; ok {
		return override
	}
	switch value.(type) {
	case string:
		// Numbers are represented as strings in the examples, while being floats in the real world responses...
		if _, err := strconv.ParseFloat(value.(string), 64); err == nil {
			return &FloatField{name: name}
		} else {
			return &StringField{name: name}
		}
	case float64:
		return &FloatField{name: name}
	case bool:
		return &BoolField{name: name}
	case map[string]interface{}:
		return NewMapField(name, value.(map[string]interface{}), overrides)
	case []interface{}:
		return NewSliceField(name, value.([]interface{}))
	}
	return &EmptyField{}
}
