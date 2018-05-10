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

// AwsSqsQueuePolicyLister helps list AwsSqsQueuePolicies.
type AwsSqsQueuePolicyLister interface {
	// List lists all AwsSqsQueuePolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSqsQueuePolicy, err error)
	// AwsSqsQueuePolicies returns an object that can list and get AwsSqsQueuePolicies.
	AwsSqsQueuePolicies(namespace string) AwsSqsQueuePolicyNamespaceLister
	AwsSqsQueuePolicyListerExpansion
}

// awsSqsQueuePolicyLister implements the AwsSqsQueuePolicyLister interface.
type awsSqsQueuePolicyLister struct {
	indexer cache.Indexer
}

// NewAwsSqsQueuePolicyLister returns a new AwsSqsQueuePolicyLister.
func NewAwsSqsQueuePolicyLister(indexer cache.Indexer) AwsSqsQueuePolicyLister {
	return &awsSqsQueuePolicyLister{indexer: indexer}
}

// List lists all AwsSqsQueuePolicies in the indexer.
func (s *awsSqsQueuePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSqsQueuePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSqsQueuePolicy))
	})
	return ret, err
}

// AwsSqsQueuePolicies returns an object that can list and get AwsSqsQueuePolicies.
func (s *awsSqsQueuePolicyLister) AwsSqsQueuePolicies(namespace string) AwsSqsQueuePolicyNamespaceLister {
	return awsSqsQueuePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSqsQueuePolicyNamespaceLister helps list and get AwsSqsQueuePolicies.
type AwsSqsQueuePolicyNamespaceLister interface {
	// List lists all AwsSqsQueuePolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSqsQueuePolicy, err error)
	// Get retrieves the AwsSqsQueuePolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSqsQueuePolicy, error)
	AwsSqsQueuePolicyNamespaceListerExpansion
}

// awsSqsQueuePolicyNamespaceLister implements the AwsSqsQueuePolicyNamespaceLister
// interface.
type awsSqsQueuePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSqsQueuePolicies in the indexer for a given namespace.
func (s awsSqsQueuePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSqsQueuePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSqsQueuePolicy))
	})
	return ret, err
}

// Get retrieves the AwsSqsQueuePolicy from the indexer for a given namespace and name.
func (s awsSqsQueuePolicyNamespaceLister) Get(name string) (*v1alpha1.AwsSqsQueuePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awssqsqueuepolicy"), name)
	}
	return obj.(*v1alpha1.AwsSqsQueuePolicy), nil
}
