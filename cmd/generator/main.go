package main

import (
	"fmt"

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

	register, err := generator.NewFile("v1alpha1")
	if err != nil {
		panic(err)
	}

	register.WriteToBuffer(`
	import (
		meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		"k8s.io/apimachinery/pkg/runtime"
		"k8s.io/apimachinery/pkg/runtime/schema"

		"github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
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

	for k, v := range prov.ResourcesMap {
		register.WriteToBuffer(generator.CamelAndTitle(k) + "{},")
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
		err = f.WriteToFile(fmt.Sprintf("pkg/apis/aws/v1alpha1/%s.go", k))
		if err != nil {
			panic(err)
		}
	}

	register.WriteToBuffer(`
	)
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}`)
	register.WriteToFile("pkg/apis/aws/v1alpha1/register.go")
}
