package generator

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

type structType struct {
	BaseName string

	b   []byte
	buf *bytes.Buffer

	dependentStructs []*structType
	schema           map[string]*schema.Schema
}

func NewStructType(name string, s map[string]*schema.Schema) *structType {
	b := []byte{}
	buf := bytes.NewBuffer(b)

	return &structType{
		BaseName: name,

		b:      b,
		buf:    buf,
		schema: s,
	}
}

func (s *structType) AddDependentStruct(st *structType) error {
	s.dependentStructs = append(s.dependentStructs, st)
	return nil
}
func (s *structType) WriteDependentStructs() {
	for _, v := range s.dependentStructs {
		b, err := v.Generate()
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprintf(s.buf, string(b)+"\n")
		if err != nil {
			panic(err)
		}
	}
}

func (s *structType) Begin(name string) error {
	_, err := fmt.Fprintf(s.buf, "type %s struct\n", name)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(s.buf, "{\n")
	if err != nil {
		return err
	}

	return nil
}

func (s *structType) End() error {
	_, err := fmt.Fprintf(s.buf, "}\n")
	if err != nil {
		return err
	}

	return nil
}

// Generate ...
func (s *structType) Generate() ([]byte, error) {
	// name = FooBar
	name := camelAndTitle(s.BaseName)
	err := s.Begin(name)
	if err != nil {
		return s.b, err
	}

	for k, v := range s.schema {
		err := s.GenerateField(k, v)
		if err != nil {
			return s.b, err
		}
	}

	err = s.End()
	if err != nil {
		return s.b, err
	}

	s.WriteDependentStructs()
	return s.b, nil
}

func (s *structType) GenerateField(name string, sch *schema.Schema) error {
	t := s.schemaTypeAsString(camelAndTitle(name), sch)
	fmt.Fprintf(s.buf, "%s %s `json:\"%s\"`", camelAndTitle(name), t, name)
	return nil
}

func (s *structType) schemaTypeAsString(fieldName string, t *schema.Schema) string {
	switch t.Type {
	case schema.TypeString:
		return "string"
	case schema.TypeBool:
		return "bool"
	case schema.TypeFloat:
		return "float64"
	case schema.TypeInt:
		return "int"
	case schema.TypeList:
		return "[]" + s.typeElemAsString(fieldName, t.Elem)
	case schema.TypeMap:
		return "map[string]" + s.typeElemAsString(fieldName, t.Elem)
	case schema.TypeSet:
		return s.typeElemAsString(fieldName, t.Elem)
	}
	return "string"
}

func (s *structType) typeElemAsString(name string, i interface{}) string {
	switch i.(type) {
	case *schema.Schema:
		return s.schemaTypeAsString(name, i.(*schema.Schema))
	case *schema.Resource:
		p := NewStructType(s.BaseName+name, i.(*schema.Resource).Schema)
		s.AddDependentStruct(p)
		_, err := s.Generate()
		if err != nil {
			panic(err)
		}
		return s.BaseName + name
	}

	return "string"
}
