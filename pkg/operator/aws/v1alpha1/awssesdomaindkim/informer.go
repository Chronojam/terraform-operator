package awssesdomaindkim

import (
    meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    csinf_v1 "github.com/trussle/terraform-operator/pkg/client/informers/externalversions/aws/v1"
    cs "github.com/trussle/terraform-operator/pkg/client/clientset/versioned"
    "k8s.io/client-go/tools/cache"
)

func DefaultInformer(resClient cs.Interface) cache.SharedIndexInformer {
    return csinf_v1.NewAwsSesDomainDkimInformer(
			resClient,
			meta_v1.NamespaceAll,
			0,
			cache.Indexers{},
		)
}
