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

// AwsCodecommitTriggerLister helps list AwsCodecommitTriggers.
type AwsCodecommitTriggerLister interface {
	// List lists all AwsCodecommitTriggers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCodecommitTrigger, err error)
	// AwsCodecommitTriggers returns an object that can list and get AwsCodecommitTriggers.
	AwsCodecommitTriggers(namespace string) AwsCodecommitTriggerNamespaceLister
	AwsCodecommitTriggerListerExpansion
}

// awsCodecommitTriggerLister implements the AwsCodecommitTriggerLister interface.
type awsCodecommitTriggerLister struct {
	indexer cache.Indexer
}

// NewAwsCodecommitTriggerLister returns a new AwsCodecommitTriggerLister.
func NewAwsCodecommitTriggerLister(indexer cache.Indexer) AwsCodecommitTriggerLister {
	return &awsCodecommitTriggerLister{indexer: indexer}
}

// List lists all AwsCodecommitTriggers in the indexer.
func (s *awsCodecommitTriggerLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCodecommitTrigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCodecommitTrigger))
	})
	return ret, err
}

// AwsCodecommitTriggers returns an object that can list and get AwsCodecommitTriggers.
func (s *awsCodecommitTriggerLister) AwsCodecommitTriggers(namespace string) AwsCodecommitTriggerNamespaceLister {
	return awsCodecommitTriggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCodecommitTriggerNamespaceLister helps list and get AwsCodecommitTriggers.
type AwsCodecommitTriggerNamespaceLister interface {
	// List lists all AwsCodecommitTriggers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCodecommitTrigger, err error)
	// Get retrieves the AwsCodecommitTrigger from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsCodecommitTrigger, error)
	AwsCodecommitTriggerNamespaceListerExpansion
}

// awsCodecommitTriggerNamespaceLister implements the AwsCodecommitTriggerNamespaceLister
// interface.
type awsCodecommitTriggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCodecommitTriggers in the indexer for a given namespace.
func (s awsCodecommitTriggerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCodecommitTrigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCodecommitTrigger))
	})
	return ret, err
}

// Get retrieves the AwsCodecommitTrigger from the indexer for a given namespace and name.
func (s awsCodecommitTriggerNamespaceLister) Get(name string) (*v1alpha1.AwsCodecommitTrigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awscodecommittrigger"), name)
	}
	return obj.(*v1alpha1.AwsCodecommitTrigger), nil
}
