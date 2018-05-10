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

// AwsInstanceLister helps list AwsInstances.
type AwsInstanceLister interface {
	// List lists all AwsInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsInstance, err error)
	// AwsInstances returns an object that can list and get AwsInstances.
	AwsInstances(namespace string) AwsInstanceNamespaceLister
	AwsInstanceListerExpansion
}

// awsInstanceLister implements the AwsInstanceLister interface.
type awsInstanceLister struct {
	indexer cache.Indexer
}

// NewAwsInstanceLister returns a new AwsInstanceLister.
func NewAwsInstanceLister(indexer cache.Indexer) AwsInstanceLister {
	return &awsInstanceLister{indexer: indexer}
}

// List lists all AwsInstances in the indexer.
func (s *awsInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsInstance))
	})
	return ret, err
}

// AwsInstances returns an object that can list and get AwsInstances.
func (s *awsInstanceLister) AwsInstances(namespace string) AwsInstanceNamespaceLister {
	return awsInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsInstanceNamespaceLister helps list and get AwsInstances.
type AwsInstanceNamespaceLister interface {
	// List lists all AwsInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsInstance, err error)
	// Get retrieves the AwsInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsInstance, error)
	AwsInstanceNamespaceListerExpansion
}

// awsInstanceNamespaceLister implements the AwsInstanceNamespaceLister
// interface.
type awsInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsInstances in the indexer for a given namespace.
func (s awsInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsInstance))
	})
	return ret, err
}

// Get retrieves the AwsInstance from the indexer for a given namespace and name.
func (s awsInstanceNamespaceLister) Get(name string) (*v1alpha1.AwsInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsinstance"), name)
	}
	return obj.(*v1alpha1.AwsInstance), nil
}
