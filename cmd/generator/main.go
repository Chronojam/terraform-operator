package main

import (
	"fmt"

	"github.com/chronojam/terraform-operator/pkg/generator"
	"github.com/hashicorp/terraform-provider-aws/aws"
	"github.com/hashicorp/terraform/helper/schema"
)

func main() {
	prov := aws.Provider().(*schema.Provider)
	for k, v := range prov.ResourcesMap {
		s := generator.NewStructType(k, v.Schema)
		b, err := s.Generate()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))
	}
}
