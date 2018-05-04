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

// AwsVpcEndpointServiceLister helps list AwsVpcEndpointServices.
type AwsVpcEndpointServiceLister interface {
	// List lists all AwsVpcEndpointServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointService, err error)
	// AwsVpcEndpointServices returns an object that can list and get AwsVpcEndpointServices.
	AwsVpcEndpointServices(namespace string) AwsVpcEndpointServiceNamespaceLister
	AwsVpcEndpointServiceListerExpansion
}

// awsVpcEndpointServiceLister implements the AwsVpcEndpointServiceLister interface.
type awsVpcEndpointServiceLister struct {
	indexer cache.Indexer
}

// NewAwsVpcEndpointServiceLister returns a new AwsVpcEndpointServiceLister.
func NewAwsVpcEndpointServiceLister(indexer cache.Indexer) AwsVpcEndpointServiceLister {
	return &awsVpcEndpointServiceLister{indexer: indexer}
}

// List lists all AwsVpcEndpointServices in the indexer.
func (s *awsVpcEndpointServiceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsVpcEndpointService))
	})
	return ret, err
}

// AwsVpcEndpointServices returns an object that can list and get AwsVpcEndpointServices.
func (s *awsVpcEndpointServiceLister) AwsVpcEndpointServices(namespace string) AwsVpcEndpointServiceNamespaceLister {
	return awsVpcEndpointServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsVpcEndpointServiceNamespaceLister helps list and get AwsVpcEndpointServices.
type AwsVpcEndpointServiceNamespaceLister interface {
	// List lists all AwsVpcEndpointServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointService, err error)
	// Get retrieves the AwsVpcEndpointService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsVpcEndpointService, error)
	AwsVpcEndpointServiceNamespaceListerExpansion
}

// awsVpcEndpointServiceNamespaceLister implements the AwsVpcEndpointServiceNamespaceLister
// interface.
type awsVpcEndpointServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsVpcEndpointServices in the indexer for a given namespace.
func (s awsVpcEndpointServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsVpcEndpointService))
	})
	return ret, err
}

// Get retrieves the AwsVpcEndpointService from the indexer for a given namespace and name.
func (s awsVpcEndpointServiceNamespaceLister) Get(name string) (*v1alpha1.AwsVpcEndpointService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsvpcendpointservice"), name)
	}
	return obj.(*v1alpha1.AwsVpcEndpointService), nil
}
