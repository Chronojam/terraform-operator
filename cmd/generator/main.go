package main

import (
	"fmt"
	"strings"

	"github.com/chronojam/terraform-operator/pkg/generator"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

func main() {
	prov := aws.Provider().(*schema.Provider)
	doc, err := generator.NewFile("")
	if err != nil {
		panic(err)
	}

	doc.WriteToBuffer("// +k8s:deepcopy-gen=package\n")
	doc.WriteToBuffer("// +groupName=chronojam.co.uk\n")
	doc.WriteToBuffer("package v1alpha1")
	doc.WriteToFile("pkg/apis/aws/v1alpha1/doc.go")

	crGenerator, err := generator.NewFile("main")
	if err != nil {
		panic(err)
	}

	groupName, err := generator.NewFile("aws")
	if err != nil {
		panic(err)
	}

	groupName.WriteToBuffer(`
	const GroupName = "chronojam.co.uk"	
	`)

	groupName.WriteToFile("pkg/apis/aws/group.go")

	register, err := generator.NewFile("v1alpha1")
	if err != nil {
		panic(err)
	}

	register.WriteToBuffer(`
	import (
		meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		"k8s.io/apimachinery/pkg/runtime"
		"k8s.io/apimachinery/pkg/runtime/schema"

		"github.com/chronojam/terraform-operator/pkg/apis/aws"
	)

	var SchemeGroupVersion = schema.GroupVersion{
		Group:   aws.GroupName,
		Version: "v1alpha1",
	}
	func AddToScheme(s *runtime.Scheme) error {
		sb := runtime.NewSchemeBuilder(addKnownTypes)
		return sb.AddToScheme(s)
	}
	func Resource(resource string) schema.GroupResource {
		return SchemeGroupVersion.WithResource(resource).GroupResource()
	}
	func addKnownTypes(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(
			SchemeGroupVersion,
	`)

	crGenerator.WriteToBuffer(fmt.Sprintf(`
	import (
		"github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
		"github.com/chronojam/terraform-operator/pkg/generator"
	)

	type CustomResource struct {
		%s
		%s
		%s

		%s
	}

	func main() {
	`, "ApiVersion	string	`json:\"apiVersion\"`",
		"Kind     string `json:\"kind\"`",
		"Metadata	map[string]interface{} `json:\"metadata\"`",
		"Spec interface{} `json:\"spec\"`"))

	for k, v := range prov.ResourcesMap {
		register.WriteToBuffer("&" + generator.CamelAndTitle(k) + "{},")
		register.WriteToBuffer("&" + generator.CamelAndTitle(k) + "List{},")
		crGenerator.WriteToBuffer(fmt.Sprintf(`
			cr%s := CustomResource{
				Kind: "%s",
				ApiVersion: "chronojam.co.uk/v1alpha1",
				Metadata: map[string]interface{}{
					"name": "example-%s",
				},
				Spec: v1alpha1.%sSpec{}, 
			}
		`, generator.CamelAndTitle(k), generator.CamelAndTitle(k), strings.Replace(k, "_", "-", -1), generator.CamelAndTitle(k)))
		crGenerator.WriteToBuffer(fmt.Sprintf(`
		generator.ToFile(cr%s, "examples/crs/%s.yaml")`, generator.CamelAndTitle(k), strings.Replace(k, "_", "-", -1)))
		f, err := generator.NewFile("v1alpha1")
		if err != nil {
			panic(err)
		}
		err = f.WriteToBuffer("import (\nmeta_v1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"\n)")
		if err != nil {
			panic(err)
		}
		base := generator.NewStructType(k, v.Schema)
		baseBytes, err := base.GenerateWithFields([]string{
			"meta_v1.TypeMeta `json:\",inline\"`",
			"meta_v1.ObjectMeta `json:\"metadata,omitempty\"`",
			fmt.Sprintf("Spec %s `json:\"spec\"`", generator.CamelAndTitle(k)+"Spec"),
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
			"meta_v1.TypeMeta `json:\",inline\"`",
			"meta_v1.ListMeta `json:\"metadata,omitempty\"`",
			fmt.Sprintf("Items []%s `json:\"items\"`", generator.CamelAndTitle(k)),
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
		err = f.WriteToFile(fmt.Sprintf("pkg/apis/aws/v1alpha1/%s.go", k))
		if err != nil {
			panic(err)
		}
	}

	crGenerator.WriteToBuffer("}")
	register.WriteToBuffer(`
	)
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}`)
	register.WriteToFile("pkg/apis/aws/v1alpha1/register.go")
	crGenerator.WriteToFile("hack/crGenerator.go")
}
