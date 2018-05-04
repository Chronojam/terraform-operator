package main

import (
	"os"
	"os/signal"
	"syscall"

	"k8s.io/client-go/tools/cache"

	"github.com/chronojam/terraform-operator/pkg/controller"
	csinf_v1 "github.com/trussle/terraform-operator/pkg/client/informers/externalversions/aws/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	client, resClient := controller.GetKubernetesClient()

	informer := csinf_v1.NewAwsIamPolicyInformer(
		resClient,
		meta_v1.NamespaceAll,
		0,
		cache.Indexers{},
	)

	stopCh := make(chan struct{})
	defer close(stopCh)

	go controller.Run(stopCh)

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}
