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

// AwsAmiFromInstanceLister helps list AwsAmiFromInstances.
type AwsAmiFromInstanceLister interface {
	// List lists all AwsAmiFromInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAmiFromInstance, err error)
	// AwsAmiFromInstances returns an object that can list and get AwsAmiFromInstances.
	AwsAmiFromInstances(namespace string) AwsAmiFromInstanceNamespaceLister
	AwsAmiFromInstanceListerExpansion
}

// awsAmiFromInstanceLister implements the AwsAmiFromInstanceLister interface.
type awsAmiFromInstanceLister struct {
	indexer cache.Indexer
}

// NewAwsAmiFromInstanceLister returns a new AwsAmiFromInstanceLister.
func NewAwsAmiFromInstanceLister(indexer cache.Indexer) AwsAmiFromInstanceLister {
	return &awsAmiFromInstanceLister{indexer: indexer}
}

// List lists all AwsAmiFromInstances in the indexer.
func (s *awsAmiFromInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAmiFromInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAmiFromInstance))
	})
	return ret, err
}

// AwsAmiFromInstances returns an object that can list and get AwsAmiFromInstances.
func (s *awsAmiFromInstanceLister) AwsAmiFromInstances(namespace string) AwsAmiFromInstanceNamespaceLister {
	return awsAmiFromInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsAmiFromInstanceNamespaceLister helps list and get AwsAmiFromInstances.
type AwsAmiFromInstanceNamespaceLister interface {
	// List lists all AwsAmiFromInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAmiFromInstance, err error)
	// Get retrieves the AwsAmiFromInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsAmiFromInstance, error)
	AwsAmiFromInstanceNamespaceListerExpansion
}

// awsAmiFromInstanceNamespaceLister implements the AwsAmiFromInstanceNamespaceLister
// interface.
type awsAmiFromInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsAmiFromInstances in the indexer for a given namespace.
func (s awsAmiFromInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAmiFromInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAmiFromInstance))
	})
	return ret, err
}

// Get retrieves the AwsAmiFromInstance from the indexer for a given namespace and name.
func (s awsAmiFromInstanceNamespaceLister) Get(name string) (*v1alpha1.AwsAmiFromInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsamifrominstance"), name)
	}
	return obj.(*v1alpha1.AwsAmiFromInstance), nil
}
