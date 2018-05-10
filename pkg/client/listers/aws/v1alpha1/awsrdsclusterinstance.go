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

// AwsRdsClusterInstanceLister helps list AwsRdsClusterInstances.
type AwsRdsClusterInstanceLister interface {
	// List lists all AwsRdsClusterInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRdsClusterInstance, err error)
	// AwsRdsClusterInstances returns an object that can list and get AwsRdsClusterInstances.
	AwsRdsClusterInstances(namespace string) AwsRdsClusterInstanceNamespaceLister
	AwsRdsClusterInstanceListerExpansion
}

// awsRdsClusterInstanceLister implements the AwsRdsClusterInstanceLister interface.
type awsRdsClusterInstanceLister struct {
	indexer cache.Indexer
}

// NewAwsRdsClusterInstanceLister returns a new AwsRdsClusterInstanceLister.
func NewAwsRdsClusterInstanceLister(indexer cache.Indexer) AwsRdsClusterInstanceLister {
	return &awsRdsClusterInstanceLister{indexer: indexer}
}

// List lists all AwsRdsClusterInstances in the indexer.
func (s *awsRdsClusterInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRdsClusterInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRdsClusterInstance))
	})
	return ret, err
}

// AwsRdsClusterInstances returns an object that can list and get AwsRdsClusterInstances.
func (s *awsRdsClusterInstanceLister) AwsRdsClusterInstances(namespace string) AwsRdsClusterInstanceNamespaceLister {
	return awsRdsClusterInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsRdsClusterInstanceNamespaceLister helps list and get AwsRdsClusterInstances.
type AwsRdsClusterInstanceNamespaceLister interface {
	// List lists all AwsRdsClusterInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRdsClusterInstance, err error)
	// Get retrieves the AwsRdsClusterInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsRdsClusterInstance, error)
	AwsRdsClusterInstanceNamespaceListerExpansion
}

// awsRdsClusterInstanceNamespaceLister implements the AwsRdsClusterInstanceNamespaceLister
// interface.
type awsRdsClusterInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsRdsClusterInstances in the indexer for a given namespace.
func (s awsRdsClusterInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRdsClusterInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRdsClusterInstance))
	})
	return ret, err
}

// Get retrieves the AwsRdsClusterInstance from the indexer for a given namespace and name.
func (s awsRdsClusterInstanceNamespaceLister) Get(name string) (*v1alpha1.AwsRdsClusterInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsrdsclusterinstance"), name)
	}
	return obj.(*v1alpha1.AwsRdsClusterInstance), nil
}
