package main

import (
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"sync"
)

type FieldParser struct {
	service   *Service
	action    *Action
	overrides map[string]Field
}

func NewFieldParser(service *Service, action *Action, overrides map[string]Field) *FieldParser {
	parser := &FieldParser{service: service, action: action, overrides: overrides}

	if action.hasPaging() {
		parser.overrides["paging"] = NewStatementField("paging", Qual(qualifier("paging"), "Paging"))
	}

	return parser
}

func (p FieldParser) parse(name string, value interface{}) Field {
	if override, ok := p.overrides[name]; ok {
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
		return p.NewMapField(name, value.(map[string]interface{}))
	case []interface{}:
		return p.NewSliceField(name, value.([]interface{}))
	}
	return &EmptyField{}
}

type Field interface {
	Name() string
	Render(tags bool) *Statement
}

type StringField struct {
	name string
}

func (f *StringField) Name() string {
	return f.name
}

func (f *StringField) Render(tags bool) *Statement {
	output := renderId(f.name).String()

	if tags {
		output.Add(Tag(map[string]string{"json": f.name + ",omitempty"}))
	}

	return output
}

type FloatField struct {
	name string
}

func (f *FloatField) Name() string {
	return f.name
}

func (f *FloatField) Render(tags bool) *Statement {
	output := renderId(f.name).Float64()

	if tags {
		output.Add(Tag(map[string]string{"json": f.name + ",omitempty"}))
	}

	return output
}

type BoolField struct {
	name string
}

func (f *BoolField) Name() string {
	return f.name
}

func (f *BoolField) Render(tags bool) *Statement {
	output := renderId(f.name).Bool()

	if tags {
		output.Add(Tag(map[string]string{"json": f.name + ",omitempty"}))
	}

	return output
}

type MapField struct {
	name      string
	fields    []Field
	overrides map[string]*Statement
	mutex     sync.Mutex
}

func (p *FieldParser) NewMapField(name string, values map[string]interface{}) *MapField {
	fields := make([]Field, len(values))
	keys := sortedKeys(values)

	for i, k := range keys {
		v := values[k]
		field := p.parse(k, v)

		fields[i] = field
	}

	return &MapField{name: name, fields: fields}
}

func (f *MapField) Name() string {
	return f.name
}

func (f *MapField) Render(tags bool) *Statement {
	code := make([]Code, len(f.fields))
	for i, field := range f.fields {
		code[i] = field.Render(true)
	}

	output := renderId(f.name).Struct(code...)
	if tags {
		output.Add(Tag(map[string]string{"json": f.name + ",omitempty"}))
	}

	return output
}

func (f *MapField) Accessors() []string {
	keys := make([]string, len(f.fields))
	for i, field := range f.fields {
		keys[i] = strcase.ToCamel(field.Name())
	}

	return keys
}

// CombineWith adds the fields from other to this field if they do not exist yet
func (f *MapField) CombineWith(other *MapField) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	existing := make(map[string]int, len(f.fields))

	for i, field := range f.fields {
		existing[field.Name()] = i
	}

	for _, field := range other.fields {
		if _, ok := existing[field.Name()]; !ok {
			f.fields = append(f.fields, field)
		}
	}
}

type SliceField struct {
	name string
	elem Field
}

func (p *FieldParser) NewSliceField(name string, values []interface{}) *SliceField {
	var elem Field

	if len(values) > 0 {
		value := values[0]
		elem = p.parse("", value)

		// Some example arrays have multiple entries, with not every entry containing every field.
		// If the first entry is an object (MapField), we collect all possibles object keys.
		if v, ok := elem.(*MapField); ok && len(values) > 1 {
			for _, other := range values[1:] {
				v.CombineWith(p.parse("dummy", other).(*MapField))
			}
		}
	} else {
		// Assume []string
		elem = &StringField{}
	}

	return &SliceField{name: name, elem: elem}
}

func (f *SliceField) Name() string {
	return f.name
}

func (f *SliceField) Render(tags bool) *Statement {
	output := renderId(f.name).Index().Add(f.elem.Render(false))

	if tags {
		output.Add(Tag(map[string]string{"json": f.name + ",omitempty"}))
	}

	return output
}

type EmptyField struct{}

func (f *EmptyField) Name() string {
	return ""
}

func (f *EmptyField) Render(_ bool) *Statement {
	return Empty()
}

type StatementField struct {
	name      string
	statement *Statement
}

func NewStatementField(name string, statement *Statement) *StatementField {
	return &StatementField{name: name, statement: statement}
}

func (f *StatementField) Name() string {
	return f.name
}

func (f *StatementField) Render(tags bool) *Statement {
	output := renderId(f.name).Add(f.statement)

	if tags {
		output.Add(Tag(map[string]string{"json": f.name + ",omitempty"}))
	}

	return output
}
