package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	if packageName != "" {
		// Write out package name to the top of our file.
		err := resource.WriteToBuffer(fmt.Sprintf("package %s", packageName))
		if err != nil {
			return nil, err
		}
	}

	return resource, nil
}

func (res *file) WriteToBuffer(s string) error {
	_, err := fmt.Fprintln(res.buffer, s)
	if err != nil {
		return err
	}

	return nil
}

func (res *file) ToString() string {
	return string(res.buffer.Bytes())
}

func (res *file) WriteToFile(path string) error {
	return ioutil.WriteFile(path, res.buffer.Bytes(), 0755)
}
