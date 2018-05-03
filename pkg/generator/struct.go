package generator

import (
	"strings"
)

type goStruct interface {
	Generate() ([]byte, error)
}

func camelAndTitle(s string) string {
	p := strings.Replace(s, "_", " ", -1)
	t := strings.Title(p)
	return strings.Replace(t, " ", "", -1)
}
