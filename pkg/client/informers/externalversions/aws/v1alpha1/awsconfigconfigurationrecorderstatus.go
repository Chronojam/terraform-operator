/*
Copyright The Kubernetes Authors.

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

// AwsConfigConfigurationRecorderStatusInformer provides access to a shared informer and lister for
// AwsConfigConfigurationRecorderStatuses.
type AwsConfigConfigurationRecorderStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AwsConfigConfigurationRecorderStatusLister
}

type awsConfigConfigurationRecorderStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAwsConfigConfigurationRecorderStatusInformer constructs a new informer for AwsConfigConfigurationRecorderStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsConfigConfigurationRecorderStatusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsConfigConfigurationRecorderStatusInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAwsConfigConfigurationRecorderStatusInformer constructs a new informer for AwsConfigConfigurationRecorderStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsConfigConfigurationRecorderStatusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsConfigConfigurationRecorderStatuses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsConfigConfigurationRecorderStatuses(namespace).Watch(options)
			},
		},
		&aws_v1alpha1.AwsConfigConfigurationRecorderStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsConfigConfigurationRecorderStatusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsConfigConfigurationRecorderStatusInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsConfigConfigurationRecorderStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aws_v1alpha1.AwsConfigConfigurationRecorderStatus{}, f.defaultInformer)
}

func (f *awsConfigConfigurationRecorderStatusInformer) Lister() v1alpha1.AwsConfigConfigurationRecorderStatusLister {
	return v1alpha1.NewAwsConfigConfigurationRecorderStatusLister(f.Informer().GetIndexer())
}
