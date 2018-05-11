package awsiampolicy

import (
	cs "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned"
	csinf_v1 "github.com/chronojam/terraform-operator/pkg/client/informers/externalversions/aws/v1alpha1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func DefaultInformer(resClient cs.Interface) cache.SharedIndexInformer {
	informer := csinf_v1.NewAwsIamPolicyInformer(
		resClient,
		meta_v1.NamespaceAll,
		0,
		cache.Indexers{},
	)

	return informer
}
