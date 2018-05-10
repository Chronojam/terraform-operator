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

// AwsSubnetLister helps list AwsSubnets.
type AwsSubnetLister interface {
	// List lists all AwsSubnets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSubnet, err error)
	// AwsSubnets returns an object that can list and get AwsSubnets.
	AwsSubnets(namespace string) AwsSubnetNamespaceLister
	AwsSubnetListerExpansion
}

// awsSubnetLister implements the AwsSubnetLister interface.
type awsSubnetLister struct {
	indexer cache.Indexer
}

// NewAwsSubnetLister returns a new AwsSubnetLister.
func NewAwsSubnetLister(indexer cache.Indexer) AwsSubnetLister {
	return &awsSubnetLister{indexer: indexer}
}

// List lists all AwsSubnets in the indexer.
func (s *awsSubnetLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSubnet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSubnet))
	})
	return ret, err
}

// AwsSubnets returns an object that can list and get AwsSubnets.
func (s *awsSubnetLister) AwsSubnets(namespace string) AwsSubnetNamespaceLister {
	return awsSubnetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSubnetNamespaceLister helps list and get AwsSubnets.
type AwsSubnetNamespaceLister interface {
	// List lists all AwsSubnets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSubnet, err error)
	// Get retrieves the AwsSubnet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSubnet, error)
	AwsSubnetNamespaceListerExpansion
}

// awsSubnetNamespaceLister implements the AwsSubnetNamespaceLister
// interface.
type awsSubnetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSubnets in the indexer for a given namespace.
func (s awsSubnetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSubnet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSubnet))
	})
	return ret, err
}

// Get retrieves the AwsSubnet from the indexer for a given namespace and name.
func (s awsSubnetNamespaceLister) Get(name string) (*v1alpha1.AwsSubnet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awssubnet"), name)
	}
	return obj.(*v1alpha1.AwsSubnet), nil
}
