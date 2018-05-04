package generator

import (
	"strings"
)

type goStruct interface {
	Generate() ([]byte, error)
}

func CamelAndTitle(s string) string {
	p := strings.Replace(s, "_", " ", -1)
	t := strings.Title(p)
	return strings.Replace(t, " ", "", -1)
}
