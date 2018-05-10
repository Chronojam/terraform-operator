package awsroute53delegationset

import (
    meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    csinf_v1 "github.com/chronojam/terraform-operator/pkg/client/informers/externalversions/aws/v1alpha1"
    cs "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned"
    "k8s.io/client-go/tools/cache"
)

func DefaultInformer(resClient cs.Interface) cache.SharedIndexInformer {
    return csinf_v1.NewAwsRoute53DelegationSetInformer(
			resClient,
			meta_v1.NamespaceAll,
			0,
			cache.Indexers{},
		)
}
