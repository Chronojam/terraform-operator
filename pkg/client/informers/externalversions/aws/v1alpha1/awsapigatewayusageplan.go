/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	aws_v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	versioned "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/chronojam/terraform-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/client/listers/aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AwsApiGatewayUsagePlanInformer provides access to a shared informer and lister for
// AwsApiGatewayUsagePlans.
type AwsApiGatewayUsagePlanInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AwsApiGatewayUsagePlanLister
}

type awsApiGatewayUsagePlanInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAwsApiGatewayUsagePlanInformer constructs a new informer for AwsApiGatewayUsagePlan type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsApiGatewayUsagePlanInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsApiGatewayUsagePlanInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAwsApiGatewayUsagePlanInformer constructs a new informer for AwsApiGatewayUsagePlan type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsApiGatewayUsagePlanInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChronojamV1alpha1().AwsApiGatewayUsagePlans(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChronojamV1alpha1().AwsApiGatewayUsagePlans(namespace).Watch(options)
			},
		},
		&aws_v1alpha1.AwsApiGatewayUsagePlan{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsApiGatewayUsagePlanInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsApiGatewayUsagePlanInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsApiGatewayUsagePlanInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aws_v1alpha1.AwsApiGatewayUsagePlan{}, f.defaultInformer)
}

func (f *awsApiGatewayUsagePlanInformer) Lister() v1alpha1.AwsApiGatewayUsagePlanLister {
	return v1alpha1.NewAwsApiGatewayUsagePlanLister(f.Informer().GetIndexer())
}
