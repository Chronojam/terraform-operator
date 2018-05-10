package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"

	"k8s.io/client-go/tools/cache"

	cs "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned"
	operator "github.com/chronojam/terraform-operator/pkg/operator"
	"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/awsiampolicy"
)

func main() {
	client, resClient := operator.GetKubernetesClient()
	stopCh := make(chan struct{})
	defer close(stopCh)
	informers, err := constructInformers(resClient)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for inf, handler := range informers {
		op := operator.New(client, inf, handler)
		go op.Run(stopCh)
	}

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}

func constructInformers(resClient cs.Interface) (map[cache.SharedIndexInformer]operator.Handler, error) {
	return map[cache.SharedIndexInformer]operator.Handler{
		awsiampolicy.DefaultInformer(resClient): &awsiampolicy.Handler{},
	}, nil

}
