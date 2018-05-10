package 

import (
    csinf_v1 "github.com/trussle/terraform-operator/pkg/client/informers/externalversions/aws/v1"
    "k8s.io/client-go/tools/cache"
)

func DefaultInformer() cache.SharedIndexInformer {
    return csinf_v1.NewSTInformer(
			resClient,
			meta_v1.NamespaceAll,
			0,
			cache.Indexers{},
		)
}
