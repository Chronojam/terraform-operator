package main

import (
	"fmt"

	"github.com/chronojam/terraform-operator/pkg/generator"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

func main() {
	prov := aws.Provider().(*schema.Provider)
	f, err := generator.NewFile("v1alpha1")
	if err != nil {
		panic(err)
	}

	err = f.WriteToBuffer("import (\nmeta_v1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"\n)")
	if err != nil {
		panic(err)
	}

	for k, v := range prov.ResourcesMap {
		base := generator.NewStructType(k, v.Schema)
		baseBytes, err := base.GenerateWithFields([]string{
			"meta_v1.TypeMeta `json\",inline\"`",
			"meta_v1.ObjectMeta `json\"metadata,omitempty\"`",
			fmt.Sprintf("Spec %s `json\"spec\"`", generator.CamelAndTitle(k)+"Spec"),
		},
			[]string{
				"+genclient",
				"+genclient:noStatus",
				"+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
			})
		if err != nil {
			panic(err)
		}
		spec := generator.NewStructType(k+"Spec", v.Schema)
		specBytes, err := spec.Generate([]string{})
		if err != nil {
			panic(err)
		}

		list := generator.NewStructType(k+"List", v.Schema)
		listBytes, err := list.GenerateWithFields([]string{
			"meta_v1.TypeMeta `json\",inline\"`",
			"meta_v1.ObjectMeta `json\"metadata,omitempty\"`",
			fmt.Sprintf("Items []%s `json\"items\"`", generator.CamelAndTitle(k)),
		}, []string{
			"+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		})

		err = f.WriteToBuffer(string(baseBytes))
		if err != nil {
			panic(err)
		}
		err = f.WriteToBuffer(string(specBytes))
		if err != nil {
			panic(err)
		}
		err = f.WriteToBuffer(string(listBytes))
		if err != nil {
			panic(err)
		}
	}
	err = f.WriteToFile("pkg/apis/aws/v1alpha1/types.go")
	if err != nil {
		panic(err)
	}
}
