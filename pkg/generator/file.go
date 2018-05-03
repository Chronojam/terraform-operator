package generator

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

type file struct {
	buffer *bytes.Buffer
	raw    []byte
}

func NewFile(packageName string) (*file, error) {
	ret := []byte{}
	buf := bytes.NewBuffer(ret)

	resource := &file{
		buffer: buf,
		raw:    ret,
	}

	// Write out package name to the top of our file.
	err := resource.writeToBuffer(fmt.Sprintf("package %s", packageName))
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (res *file) generateField(name string, s *schema.Schema) (string, error) {

	return "", nil
}

func (res *file) writeToBuffer(s string) error {
	_, err := fmt.Fprintln(res.buffer, s)
	if err != nil {
		return err
	}

	return nil
}

// GenerateFromSchema ...
func (res *file) AppendFromTerraformSchema(name string, s *schema.Schema) error {

	return nil
}
