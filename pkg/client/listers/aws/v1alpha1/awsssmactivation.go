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

// AwsSsmActivationLister helps list AwsSsmActivations.
type AwsSsmActivationLister interface {
	// List lists all AwsSsmActivations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSsmActivation, err error)
	// AwsSsmActivations returns an object that can list and get AwsSsmActivations.
	AwsSsmActivations(namespace string) AwsSsmActivationNamespaceLister
	AwsSsmActivationListerExpansion
}

// awsSsmActivationLister implements the AwsSsmActivationLister interface.
type awsSsmActivationLister struct {
	indexer cache.Indexer
}

// NewAwsSsmActivationLister returns a new AwsSsmActivationLister.
func NewAwsSsmActivationLister(indexer cache.Indexer) AwsSsmActivationLister {
	return &awsSsmActivationLister{indexer: indexer}
}

// List lists all AwsSsmActivations in the indexer.
func (s *awsSsmActivationLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSsmActivation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSsmActivation))
	})
	return ret, err
}

// AwsSsmActivations returns an object that can list and get AwsSsmActivations.
func (s *awsSsmActivationLister) AwsSsmActivations(namespace string) AwsSsmActivationNamespaceLister {
	return awsSsmActivationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSsmActivationNamespaceLister helps list and get AwsSsmActivations.
type AwsSsmActivationNamespaceLister interface {
	// List lists all AwsSsmActivations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSsmActivation, err error)
	// Get retrieves the AwsSsmActivation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSsmActivation, error)
	AwsSsmActivationNamespaceListerExpansion
}

// awsSsmActivationNamespaceLister implements the AwsSsmActivationNamespaceLister
// interface.
type awsSsmActivationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSsmActivations in the indexer for a given namespace.
func (s awsSsmActivationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSsmActivation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSsmActivation))
	})
	return ret, err
}

// Get retrieves the AwsSsmActivation from the indexer for a given namespace and name.
func (s awsSsmActivationNamespaceLister) Get(name string) (*v1alpha1.AwsSsmActivation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsssmactivation"), name)
	}
	return obj.(*v1alpha1.AwsSsmActivation), nil
}
