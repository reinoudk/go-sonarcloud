package main

import (
	"encoding/json"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"net/http"
	"reflect"
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

func (act *Action) id() string {
	return strcase.ToCamel(act.Key)
}

func (act *Action) requestType() string {
	return fmt.Sprintf("%s%s", act.id(), "Request")
}

func (act *Action) responseType() string {
	return fmt.Sprintf("%s%s", act.id(), "Response")
}

func (act *Action) responseTypeAll() string {
	return fmt.Sprintf("%s%s", act.id(), "ResponseAll")
}

func (act *Action) responseTypePaging() string {
	return "GetPaging"
}

func (act *Action) serviceHandler() string {
	return act.id()
}

func (act *Action) serviceHandlerAll() string {
	return fmt.Sprintf("%s%s", act.id(), "All")
}

func (act *Action) hasPaging() bool {
	hasP := false
	hasPs := false

	for _, param := range act.Params {
		switch param.Key {
		case "p": hasP = true
		case "ps": hasPs = true
		}
	}

	return hasP && hasPs
}

func (act *Action) responseFields(example map[string]interface{}) (Field, error) {
	overrides := make(map[string]Field, 0)
	if act.hasPaging() {
		overrides = map[string]Field{
			"paging": NewStatementField("paging", Qual(qualifier("paging"), "Paging")),
		}
	}

	return NewMapField(act.responseType(), example, overrides), nil
}

func (act *Action) responseFieldsWithoutPaging(example map[string]interface{}) (Field, error) {
	delete(example, "paging")

	// Remove flattened paging as well
	delete(example, "p")
	delete(example, "ps")
	delete(example, "total")

	return NewMapField(act.responseTypeAll(), example, nil), nil
}

func (act *Action) renderRequestStruct() *Statement {
	fields := make([]Code, len(act.Params))
	for i, param := range act.Params {
		// filter out unwanted fields and paging parameters
		if contains(param.Key, append(skippedRequestFields, "p", "ps")) {
			continue
		}

		fields[i] = param.render()
	}

	statement := Commentf("%s %s", act.requestType(), act.Description)
	if act.DeprecatedSince != "" {
		statement.Line()
		statement.Commentf("Deprecated: this action has been deprecated since version %s", act.DeprecatedSince)
	}
	statement.Line()

	statement.Type().Id(act.requestType()).Struct(fields...)

	return statement
}

func (act *Action) renderResponseStruct(response Field) *Statement {
	// EmptyField should not be rendered
	if reflect.TypeOf(response) != reflect.TypeOf(&EmptyField{}) {
		fields := response.Render(false)
		statement := Commentf("%s is the response for %s", act.responseType(), act.requestType())
		statement.Line()
		statement.Type().Add(fields)
		return statement
	}

	return Empty()
}

func (act *Action) renderResponseStructPaging(collection Field) *Statement {
	// EmptyField should not have a Paging
	if reflect.TypeOf(collection) == reflect.TypeOf(&MapField{}) {
		statement := Commentf("%s extracts the paging from %s", act.responseTypePaging(), act.responseType())
		statement.Line()
		statement.Func().Parens(Id("r").Op("*").Id(act.responseType())).Id(act.responseTypePaging()).Call().Op("*").Qual(qualifier("paging"), "Paging")

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

func (act *Action) renderResponseAllStruct(collection Field) *Statement {
	// EmptyField should not be rendered
	if reflect.TypeOf(collection) != reflect.TypeOf(&EmptyField{}) {
		fields := collection.Render(false)
		statement := Commentf("%s is the collection for %s", act.responseTypeAll(), act.requestType())
		statement.Line()
		statement.Type().Add(fields)
		return statement
	}

	return Empty()
}

func (act *Action) fetchExample(endpoint string) (map[string]interface{}, error) {
	controller := fmt.Sprintf("api/%s", endpoint)
	request := ResponseExampleRequest{ID: act.responseType(), RequestID: act.id(), Controller: controller, Action: act.Key}

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
		return &StringField{name: name}
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
