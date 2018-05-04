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
	// Lets check if we've seen this before
	for _, x := range s.dependentStructs {
		if st.BaseName == x.BaseName {
			fmt.Printf("Seen %s\n", x.BaseName)
			// Yeah we've seen it before, just drop rather than overriding for now.
			return nil
		}
	}
	s.dependentStructs = append(s.dependentStructs, st)
	return nil
}
func (s *structType) WriteDependentStructs() {
	for _, v := range s.dependentStructs {
		_, err := fmt.Fprintf(s.buf, string(v.buf.Bytes())+"\n")
		if err != nil {
			panic(err)
		}
	}
}

func (s *structType) Begin(name string, generationComments []string) error {
	for _, l := range generationComments {
		_, err := fmt.Fprintf(s.buf, "// %s\n", l)
		if err != nil {
			return err
		}
	}
	_, err := fmt.Fprintln(s.buf, "")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(s.buf, "type %s struct\n", name)
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

func (s *structType) GenerateWithFields(fields, generationTags []string) ([]byte, error) {
	// name = FooBar
	name := CamelAndTitle(s.BaseName)
	err := s.Begin(name, generationTags)
	if err != nil {
		return s.b, err
	}

	for _, v := range fields {
		fmt.Fprintf(s.buf, v+"\n")
	}

	err = s.End()
	if err != nil {
		return s.b, err
	}

	return s.buf.Bytes(), nil
}

// Generate ...
func (s *structType) Generate(generationTags []string) ([]byte, error) {
	// name = FooBar
	name := CamelAndTitle(s.BaseName)
	err := s.Begin(name, generationTags)
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
	return s.buf.Bytes(), nil
}

func (s *structType) GenerateField(name string, sch *schema.Schema) error {
	t := s.schemaTypeAsString(CamelAndTitle(name), sch)
	fmt.Fprintf(s.buf, "%s %s `json:\"%s\"`\n", CamelAndTitle(name), t, name)
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
		_, err := p.Generate([]string{})
		if err != nil {
			panic(err)
		}
		return CamelAndTitle(s.BaseName + name)
	}

	return "string"
}
