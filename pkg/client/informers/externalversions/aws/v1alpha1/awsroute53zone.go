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

// AwsRoute53ZoneInformer provides access to a shared informer and lister for
// AwsRoute53Zones.
type AwsRoute53ZoneInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AwsRoute53ZoneLister
}

type awsRoute53ZoneInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAwsRoute53ZoneInformer constructs a new informer for AwsRoute53Zone type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsRoute53ZoneInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsRoute53ZoneInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAwsRoute53ZoneInformer constructs a new informer for AwsRoute53Zone type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsRoute53ZoneInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChronojamV1alpha1().AwsRoute53Zones(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChronojamV1alpha1().AwsRoute53Zones(namespace).Watch(options)
			},
		},
		&aws_v1alpha1.AwsRoute53Zone{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsRoute53ZoneInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsRoute53ZoneInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsRoute53ZoneInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aws_v1alpha1.AwsRoute53Zone{}, f.defaultInformer)
}

func (f *awsRoute53ZoneInformer) Lister() v1alpha1.AwsRoute53ZoneLister {
	return v1alpha1.NewAwsRoute53ZoneLister(f.Informer().GetIndexer())
}
