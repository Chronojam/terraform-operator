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

// AwsIotPolicyLister helps list AwsIotPolicies.
type AwsIotPolicyLister interface {
	// List lists all AwsIotPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIotPolicy, err error)
	// AwsIotPolicies returns an object that can list and get AwsIotPolicies.
	AwsIotPolicies(namespace string) AwsIotPolicyNamespaceLister
	AwsIotPolicyListerExpansion
}

// awsIotPolicyLister implements the AwsIotPolicyLister interface.
type awsIotPolicyLister struct {
	indexer cache.Indexer
}

// NewAwsIotPolicyLister returns a new AwsIotPolicyLister.
func NewAwsIotPolicyLister(indexer cache.Indexer) AwsIotPolicyLister {
	return &awsIotPolicyLister{indexer: indexer}
}

// List lists all AwsIotPolicies in the indexer.
func (s *awsIotPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIotPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIotPolicy))
	})
	return ret, err
}

// AwsIotPolicies returns an object that can list and get AwsIotPolicies.
func (s *awsIotPolicyLister) AwsIotPolicies(namespace string) AwsIotPolicyNamespaceLister {
	return awsIotPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIotPolicyNamespaceLister helps list and get AwsIotPolicies.
type AwsIotPolicyNamespaceLister interface {
	// List lists all AwsIotPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIotPolicy, err error)
	// Get retrieves the AwsIotPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIotPolicy, error)
	AwsIotPolicyNamespaceListerExpansion
}

// awsIotPolicyNamespaceLister implements the AwsIotPolicyNamespaceLister
// interface.
type awsIotPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIotPolicies in the indexer for a given namespace.
func (s awsIotPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIotPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIotPolicy))
	})
	return ret, err
}

// Get retrieves the AwsIotPolicy from the indexer for a given namespace and name.
func (s awsIotPolicyNamespaceLister) Get(name string) (*v1alpha1.AwsIotPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiotpolicy"), name)
	}
	return obj.(*v1alpha1.AwsIotPolicy), nil
}
