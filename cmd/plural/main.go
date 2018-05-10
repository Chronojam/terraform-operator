package main

import (
	"fmt"
	"os"

	"github.com/chronojam/terraform-operator/pkg/plural"
)

func main() {
	if len(os.Args) != 2 {
		panic("Expected 1 arguement")
	}

	fmt.Println(plural.Plural(os.Args[1]))
}
