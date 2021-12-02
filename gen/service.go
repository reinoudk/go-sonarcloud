package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"os"
	"reflect"
	"strings"
)

type Service struct {
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Actions     []Action `json:"actions"`
}

func (svc *Service) id() string {
	return strcase.ToCamel(svc.endpoint())
}

func (svc *Service) endpoint() string {
	path := strings.Split(svc.Path, "/")
	return path[len(path)-1]
}

func (svc *Service) process(output string) error {
	endpoint := svc.endpoint()
	if contains(endpoint, skippedEndpoints) {
		fmt.Printf("Skipping endpoint '%s'\n", endpoint)
		return nil
	}

	typesFile := NewFile(endpoint)
	typesFile.Commentf("// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!\n")

	serviceFile := NewFile(packageName)
	serviceFile.Commentf("// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!\n")

	serviceType := Type().Id(svc.id()).Id("service")
	serviceFile.Add(serviceType)

	for _, action := range svc.Actions {
		requestOutput := action.renderRequestStruct()
		typesFile.Add(requestOutput)

		var responseField Field = &EmptyField{}
		var collectionField Field = &EmptyField{}
		if action.HasResponseExample {
			example, err := action.fetchExample(endpoint)
			if err != nil {
				return fmt.Errorf("could not fetch example: %+v", err)
			}

			responseField, err = action.responseFields(example)
			if err != nil {
				return fmt.Errorf("could not collect response fields: %+v", err)
			}

			if action.hasPaging() {
				collectionField, err = action.responseFieldsWithoutPaging(example)
				if err != nil {
					return fmt.Errorf("could not extract collection field: %+v", err)
				}
			}
		}

		responseOutput := action.renderResponseStruct(responseField)
		typesFile.Add(responseOutput)

		if action.hasPaging() {
			pagingOutput := action.renderResponseStructPaging(responseField)
			typesFile.Add(pagingOutput)
		}

		collectionOutput := action.renderResponseAllStruct(collectionField)
		typesFile.Add(collectionOutput)

		if action.Post {
			svc.postActionHandler(action, serviceFile, endpoint)
		} else {
			svc.getActionHandler(action, serviceFile, endpoint)
		}

		if action.hasPaging() {
			svc.getAllActionHandler(action, serviceFile, endpoint, collectionField)
		}
	}

	dir := fmt.Sprintf("%s/%s", output, endpoint)
	_ = os.Mkdir(dir, 0755)

	typesFileName := fmt.Sprintf("%s/%s/%s.go", output, endpoint, endpoint)
	err := typesFile.Save(typesFileName)
	if err != nil {
		return fmt.Errorf("could not save generated source file for types: %+v\n", err)
	}

	serviceFileName := fmt.Sprintf("%s/%s.go", output, endpoint)
	err = serviceFile.Save(serviceFileName)
	if err != nil {
		return fmt.Errorf("could not save generated source file for service: %+v\n", err)
	}

	return nil
}

func (svc *Service) postActionHandler(action Action, f *File, endpoint string) {
	// import "github.com/go-playground/form/v4"
	f.ImportName("github.com/go-playground/form/v4", "form")
	// import "github.com/reinoudk/go-sonarcloud/sonarcloud/<endpoint>"
	f.ImportName(qualifier(endpoint), endpoint)
	// import "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
	f.ImportName(qualifier("paging"), "paging")

	// start function signature without return type
	// func(s *<service id>) <action id>(r <request type>)
	statement := Func().Parens(Id("s").Op("*").Id(svc.id())).Id(action.serviceHandler())
	statement.Params(Id("r").Qual(qualifier(endpoint), action.requestType()))

	// add return type based on whether we expect a response
	if action.HasResponseExample {
		// (*<response type>, error)
		statement.Parens(
			Op("*").Qual(qualifier(endpoint), action.responseType()).Op(",").Error(),
		)
	} else {
		// error
		statement.Error()
	}

	// function body
	statement.Block(
		// encoder := form.NewEncoder()
		Id("encoder").Op(":=").Qual("github.com/go-playground/form/v4", "NewEncoder").Call(),
		// values, err := encoder.Encode(r)
		Id("values").Op(",").Id("err").Op(":=").Id("encoder").Dot("Encode").Call(Id("r")),
		// if err != nil...
		ifError(action, "could not encode form values: %+v"),
		Line(),

		// req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/<endpoint>", API), strings.NewReader(values.Encode()))
		Id("req").Op(",").Err().Op(":=").Id("s").Dot("client").Dot("NewRequest").Call(
			Lit("POST"),
			Qual("fmt", "Sprintf").Call(
				Lit(fmt.Sprintf("%%s/%s/%s", svc.endpoint(), action.Key)),
				Id("API"),
			),
			Qual("strings", "NewReader").Call(
				Id("values").Dot("Encode").Call(),
			),
		),
		// if err != nil...
		ifError(action, "could not create request: %+v"),
		Line(),

		// resp, err := s.client.Do(req)
		Id("resp").Op(",").Err().Op(":=").Id("s").Dot("client").Dot("Do").Call(Id("req")),
		// if err != nil...
		ifError(action, "error trying to execute request: %+v"),
		// defer resp.Body.Close()
		Defer().Id("resp").Dot("Body").Dot("Close").Call(),
		Line(),

		// if resp.StatusCode >= 300 {
		//		// TODO: parse error message
		//		return fmt.Errorf("received non 2xx status code: %d", resp.StatusCode)
		//	}
		If().Id("resp").Dot("StatusCode").Op(">=").Lit(300).Block(
			Comment("TODO: parse error message"),
			errResult(
				action,
				Qual("fmt", "Errorf").Call(
					Lit("received non 2xx status code: %d"),
					Id("resp").Dot("StatusCode"),
				),
			),
		),
		Line(),

		ifTrueGen(
			action.HasResponseExample,
			// 	response := &projects.BulkUpdateKeyResponse{}
			Id("response").Op(":=").Op("&").Qual(qualifier(endpoint), action.responseType()).Block(),
		),
		ifTrueGen(
			action.HasResponseExample,
			// err = json.NewDecoder(resp.Body).Decode(&response)
			Err().Op("=").Qual("encoding/json", "NewDecoder").Call(
				Id("resp").Dot("Body"),
			).Dot("Decode").Call(
				Op("&").Id("response"),
			),
		),
		ifTrueGen(
			action.HasResponseExample,
			// if err != nil...
			ifError(action, "could not decode response: %+v"),
		),

		// return response, nil
		// OR
		// return nil
		genReturnWithError(action.HasResponseExample, "response"),
	)

	// Spacing
	statement.Line()

	f.Add(statement)
}

func (svc *Service) getActionHandler(action Action, f *File, endpoint string) {
	// import "github.com/reinoudk/go-sonarcloud/sonarcloud/<endpoint>"
	f.ImportName(qualifier(endpoint), endpoint)
	// import "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
	f.ImportName(qualifier("paging"), "paging")

	// start function signature without return type
	// func(s *<service id>) <action id>(r <request type>)
	statement := Func().Parens(Id("s").Op("*").Id(svc.id())).Id(action.serviceHandler())
	statement.Params(
		Id("r").Qual(qualifier(endpoint), action.requestType()),
		ifTrueGen(action.hasPaging(), Id("p").Qual(qualifier("paging"), "PagingParams")),
	)

	// add return type based on whether we expect a response
	if action.HasResponseExample {
		// (*<response type>, error)
		statement.Parens(
			Op("*").Qual(qualifier(endpoint), action.responseType()).Op(",").Error(),
		)
	} else {
		// error
		statement.Error()
	}

	// params := paramsFrom(r [, p])
	params := Id("params").Op(":=").Id("paramsFrom").Call(
		Id("r"),
		ifTrueGen(action.hasPaging(), Id("p")),
	)

	// function body
	statement.Block(
		params,
		Line(),

		// req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/<endpoint>", API), params...)
		Id("req").Op(",").Err().Op(":=").Id("s").Dot("client").Dot("NewRequestWithParameters").Call(
			Lit("GET"),
			Qual("fmt", "Sprintf").Call(
				Lit(fmt.Sprintf("%%s/%s/%s", svc.endpoint(), action.Key)),
				Id("API"),
			),
			Id("params").Op("..."),
		),
		// if err != nil...
		ifError(action, "could not create request: %+v"),
		Line(),

		// resp, err := s.client.Do(req)
		Id("resp").Op(",").Err().Op(":=").Id("s").Dot("client").Dot("Do").Call(Id("req")),
		// if err != nil...
		ifError(action, "error trying to execute request: %+v"),
		// defer resp.Body.Close()
		Defer().Id("resp").Dot("Body").Dot("Close").Call(),
		Line(),

		// if resp.StatusCode >= 300 {
		//		// TODO: parse error message
		//		return fmt.Errorf("received non 2xx status code: %d", resp.StatusCode)
		//	}
		If().Id("resp").Dot("StatusCode").Op(">=").Lit(300).Block(
			Comment("TODO: parse error message"),
			errResult(
				action,
				Qual("fmt", "Errorf").Call(
					Lit("received non 2xx status code: %d"),
					Id("resp").Dot("StatusCode"),
				),
			),
		),
		Line(),

		ifTrueGen(
			action.HasResponseExample,
			// 	response := &projects.BulkUpdateKeyResponse{}
			Id("response").Op(":=").Op("&").Qual(qualifier(endpoint), action.responseType()).Block(),
		),
		ifTrueGen(
			action.HasResponseExample,
			// err = json.NewDecoder(resp.Body).Decode(&response)
			Err().Op("=").Qual("encoding/json", "NewDecoder").Call(
				Id("resp").Dot("Body"),
			).Dot("Decode").Call(
				Op("&").Id("response"),
			),
		),
		ifTrueGen(
			action.HasResponseExample,
			// if err != nil...
			ifError(action, "could not decode response: %+v"),
		),

		// return response, nil
		// OR
		// return nil
		genReturnWithError(action.HasResponseExample, "response"),
	)

	// Spacing
	statement.Line()

	f.Add(statement)
}

func (svc *Service) getAllActionHandler(action Action, f *File, endpoint string, field Field) {
	// import "github.com/reinoudk/go-sonarcloud/sonarcloud/<endpoint>"
	f.ImportName(qualifier(endpoint), endpoint)
	// import "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
	f.ImportName(qualifier("paging"), "paging")

	// start function signature without return type
	// func(s *<service id>) <action id>(r <request type all>)
	statement := Func().Parens(Id("s").Op("*").Id(svc.id())).Id(action.serviceHandlerAll())
	statement.Params(
		Id("r").Qual(qualifier(endpoint), action.requestType()),
	)

	// Just to be safe, check field type
	mapField, ok := field.(*MapField)
	if !ok {
		fmt.Printf("Not generating 'All' handler for %s/%s, only map fields supported, got: %+v\n", svc.endpoint(), action.Key, reflect.TypeOf(field))
		return
	}

	// Create an update statement for all fields of the response structure
	accessors := mapField.Accessors()
	updateStatements := make([]Code, len(accessors))
	for i, accessor := range accessors {
		switch mapField.fields[i].(type) {
		case *SliceField:
			// response.<accessor> = append(response.<accessor>, res.<accessor>...)
			updateStatements[i] =  Id("response").Dot(accessor).Op("=").Id("append").Call(
				Id("response").Dot(accessor),
				Id("res").Dot(accessor).Op("..."),
			)
		default:
			fmt.Printf("Skipping field '%s' for %s.%s, only slices are supported.\n", mapField.fields[i].Name(), action.Key, action.responseTypeAll())
		}
	}

	// Paged requests always have a response
	// (*<response type>, error)
	statement.Parens(
		Op("*").Qual(qualifier(endpoint), action.responseTypeAll()).Op(",").Error(),
	)

	// function body
	funcBody := make([]Code, 0)

	//	p := paging.PagingParams{
	//		P:  1,
	//		Ps: 100,
	//	}
	funcBody = append(funcBody,
		Id("p").Op(":=").Qual(qualifier("paging"), "PagingParams").Values(Dict{
			Id("P"): Lit(1),
			Id("Ps"): Lit(100),
		}),

		Id("response").Op(":=").Op("&").Qual(qualifier(endpoint), action.responseTypeAll()).Block(),
	)

	loopBody := make([]Code, 0)
	loopBody = append(loopBody,
		//	res, err := s.Search(r, p)
		//	if err != nil {
		//		return nil, fmt.Errorf("could not search projects: %+v", err)
		//	}
		Id("res").Op(",").Err().Op(":=").Id("s").Dot(action.serviceHandler()).Call(Id("r"),  Id("p")),
		ifError(action, "could not search all projects: %+v"),
	)

	// Add update statements for each accessor
	loopBody = append(loopBody, updateStatements...)

	loopBody = append(loopBody,
		//	if res.Paging.End() {
		//		break
		//	} else {
		//		p.P++
		//	}
		If(Id("res").Dot(action.responseTypePaging()).Call().Dot("End").Call()).Block(
			Break(),
		).Else().Block(
			Id("p").Dot("P").Op("++"),
		),
	)

	//	for {
	funcBody = append(funcBody,
			For(nil).Block(loopBody...))
	//	}

	funcBody = append(funcBody,
		// return response, nil
		Return(Id("response"), Nil()),
	)

	statement.Block(funcBody...)

	// Spacing
	statement.Line()

	f.Add(statement)
}
