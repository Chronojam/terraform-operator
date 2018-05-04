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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsServiceDiscoveryServiceLister helps list AwsServiceDiscoveryServices.
type AwsServiceDiscoveryServiceLister interface {
	// List lists all AwsServiceDiscoveryServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryService, err error)
	// AwsServiceDiscoveryServices returns an object that can list and get AwsServiceDiscoveryServices.
	AwsServiceDiscoveryServices(namespace string) AwsServiceDiscoveryServiceNamespaceLister
	AwsServiceDiscoveryServiceListerExpansion
}

// awsServiceDiscoveryServiceLister implements the AwsServiceDiscoveryServiceLister interface.
type awsServiceDiscoveryServiceLister struct {
	indexer cache.Indexer
}

// NewAwsServiceDiscoveryServiceLister returns a new AwsServiceDiscoveryServiceLister.
func NewAwsServiceDiscoveryServiceLister(indexer cache.Indexer) AwsServiceDiscoveryServiceLister {
	return &awsServiceDiscoveryServiceLister{indexer: indexer}
}

// List lists all AwsServiceDiscoveryServices in the indexer.
func (s *awsServiceDiscoveryServiceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsServiceDiscoveryService))
	})
	return ret, err
}

// AwsServiceDiscoveryServices returns an object that can list and get AwsServiceDiscoveryServices.
func (s *awsServiceDiscoveryServiceLister) AwsServiceDiscoveryServices(namespace string) AwsServiceDiscoveryServiceNamespaceLister {
	return awsServiceDiscoveryServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsServiceDiscoveryServiceNamespaceLister helps list and get AwsServiceDiscoveryServices.
type AwsServiceDiscoveryServiceNamespaceLister interface {
	// List lists all AwsServiceDiscoveryServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryService, err error)
	// Get retrieves the AwsServiceDiscoveryService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsServiceDiscoveryService, error)
	AwsServiceDiscoveryServiceNamespaceListerExpansion
}

// awsServiceDiscoveryServiceNamespaceLister implements the AwsServiceDiscoveryServiceNamespaceLister
// interface.
type awsServiceDiscoveryServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsServiceDiscoveryServices in the indexer for a given namespace.
func (s awsServiceDiscoveryServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsServiceDiscoveryService))
	})
	return ret, err
}

// Get retrieves the AwsServiceDiscoveryService from the indexer for a given namespace and name.
func (s awsServiceDiscoveryServiceNamespaceLister) Get(name string) (*v1alpha1.AwsServiceDiscoveryService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsservicediscoveryservice"), name)
	}
	return obj.(*v1alpha1.AwsServiceDiscoveryService), nil
}
