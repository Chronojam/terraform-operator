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

// AwsSsmMaintenanceWindowLister helps list AwsSsmMaintenanceWindows.
type AwsSsmMaintenanceWindowLister interface {
	// List lists all AwsSsmMaintenanceWindows in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSsmMaintenanceWindow, err error)
	// AwsSsmMaintenanceWindows returns an object that can list and get AwsSsmMaintenanceWindows.
	AwsSsmMaintenanceWindows(namespace string) AwsSsmMaintenanceWindowNamespaceLister
	AwsSsmMaintenanceWindowListerExpansion
}

// awsSsmMaintenanceWindowLister implements the AwsSsmMaintenanceWindowLister interface.
type awsSsmMaintenanceWindowLister struct {
	indexer cache.Indexer
}

// NewAwsSsmMaintenanceWindowLister returns a new AwsSsmMaintenanceWindowLister.
func NewAwsSsmMaintenanceWindowLister(indexer cache.Indexer) AwsSsmMaintenanceWindowLister {
	return &awsSsmMaintenanceWindowLister{indexer: indexer}
}

// List lists all AwsSsmMaintenanceWindows in the indexer.
func (s *awsSsmMaintenanceWindowLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSsmMaintenanceWindow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSsmMaintenanceWindow))
	})
	return ret, err
}

// AwsSsmMaintenanceWindows returns an object that can list and get AwsSsmMaintenanceWindows.
func (s *awsSsmMaintenanceWindowLister) AwsSsmMaintenanceWindows(namespace string) AwsSsmMaintenanceWindowNamespaceLister {
	return awsSsmMaintenanceWindowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSsmMaintenanceWindowNamespaceLister helps list and get AwsSsmMaintenanceWindows.
type AwsSsmMaintenanceWindowNamespaceLister interface {
	// List lists all AwsSsmMaintenanceWindows in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSsmMaintenanceWindow, err error)
	// Get retrieves the AwsSsmMaintenanceWindow from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSsmMaintenanceWindow, error)
	AwsSsmMaintenanceWindowNamespaceListerExpansion
}

// awsSsmMaintenanceWindowNamespaceLister implements the AwsSsmMaintenanceWindowNamespaceLister
// interface.
type awsSsmMaintenanceWindowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSsmMaintenanceWindows in the indexer for a given namespace.
func (s awsSsmMaintenanceWindowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSsmMaintenanceWindow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSsmMaintenanceWindow))
	})
	return ret, err
}

// Get retrieves the AwsSsmMaintenanceWindow from the indexer for a given namespace and name.
func (s awsSsmMaintenanceWindowNamespaceLister) Get(name string) (*v1alpha1.AwsSsmMaintenanceWindow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsssmmaintenancewindow"), name)
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindow), nil
}
