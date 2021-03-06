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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsServiceDiscoveryPublicDnsNamespaceLister helps list AwsServiceDiscoveryPublicDnsNamespaces.
type AwsServiceDiscoveryPublicDnsNamespaceLister interface {
	// List lists all AwsServiceDiscoveryPublicDnsNamespaces in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error)
	// AwsServiceDiscoveryPublicDnsNamespaces returns an object that can list and get AwsServiceDiscoveryPublicDnsNamespaces.
	AwsServiceDiscoveryPublicDnsNamespaces(namespace string) AwsServiceDiscoveryPublicDnsNamespaceNamespaceLister
	AwsServiceDiscoveryPublicDnsNamespaceListerExpansion
}

// awsServiceDiscoveryPublicDnsNamespaceLister implements the AwsServiceDiscoveryPublicDnsNamespaceLister interface.
type awsServiceDiscoveryPublicDnsNamespaceLister struct {
	indexer cache.Indexer
}

// NewAwsServiceDiscoveryPublicDnsNamespaceLister returns a new AwsServiceDiscoveryPublicDnsNamespaceLister.
func NewAwsServiceDiscoveryPublicDnsNamespaceLister(indexer cache.Indexer) AwsServiceDiscoveryPublicDnsNamespaceLister {
	return &awsServiceDiscoveryPublicDnsNamespaceLister{indexer: indexer}
}

// List lists all AwsServiceDiscoveryPublicDnsNamespaces in the indexer.
func (s *awsServiceDiscoveryPublicDnsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace))
	})
	return ret, err
}

// AwsServiceDiscoveryPublicDnsNamespaces returns an object that can list and get AwsServiceDiscoveryPublicDnsNamespaces.
func (s *awsServiceDiscoveryPublicDnsNamespaceLister) AwsServiceDiscoveryPublicDnsNamespaces(namespace string) AwsServiceDiscoveryPublicDnsNamespaceNamespaceLister {
	return awsServiceDiscoveryPublicDnsNamespaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsServiceDiscoveryPublicDnsNamespaceNamespaceLister helps list and get AwsServiceDiscoveryPublicDnsNamespaces.
type AwsServiceDiscoveryPublicDnsNamespaceNamespaceLister interface {
	// List lists all AwsServiceDiscoveryPublicDnsNamespaces in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error)
	// Get retrieves the AwsServiceDiscoveryPublicDnsNamespace from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, error)
	AwsServiceDiscoveryPublicDnsNamespaceNamespaceListerExpansion
}

// awsServiceDiscoveryPublicDnsNamespaceNamespaceLister implements the AwsServiceDiscoveryPublicDnsNamespaceNamespaceLister
// interface.
type awsServiceDiscoveryPublicDnsNamespaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsServiceDiscoveryPublicDnsNamespaces in the indexer for a given namespace.
func (s awsServiceDiscoveryPublicDnsNamespaceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace))
	})
	return ret, err
}

// Get retrieves the AwsServiceDiscoveryPublicDnsNamespace from the indexer for a given namespace and name.
func (s awsServiceDiscoveryPublicDnsNamespaceNamespaceLister) Get(name string) (*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsservicediscoverypublicdnsnamespace"), name)
	}
	return obj.(*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace), nil
}
