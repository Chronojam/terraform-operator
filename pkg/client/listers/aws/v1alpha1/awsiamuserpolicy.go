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

// AwsIamUserPolicyLister helps list AwsIamUserPolicies.
type AwsIamUserPolicyLister interface {
	// List lists all AwsIamUserPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamUserPolicy, err error)
	// AwsIamUserPolicies returns an object that can list and get AwsIamUserPolicies.
	AwsIamUserPolicies(namespace string) AwsIamUserPolicyNamespaceLister
	AwsIamUserPolicyListerExpansion
}

// awsIamUserPolicyLister implements the AwsIamUserPolicyLister interface.
type awsIamUserPolicyLister struct {
	indexer cache.Indexer
}

// NewAwsIamUserPolicyLister returns a new AwsIamUserPolicyLister.
func NewAwsIamUserPolicyLister(indexer cache.Indexer) AwsIamUserPolicyLister {
	return &awsIamUserPolicyLister{indexer: indexer}
}

// List lists all AwsIamUserPolicies in the indexer.
func (s *awsIamUserPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamUserPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamUserPolicy))
	})
	return ret, err
}

// AwsIamUserPolicies returns an object that can list and get AwsIamUserPolicies.
func (s *awsIamUserPolicyLister) AwsIamUserPolicies(namespace string) AwsIamUserPolicyNamespaceLister {
	return awsIamUserPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamUserPolicyNamespaceLister helps list and get AwsIamUserPolicies.
type AwsIamUserPolicyNamespaceLister interface {
	// List lists all AwsIamUserPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamUserPolicy, err error)
	// Get retrieves the AwsIamUserPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamUserPolicy, error)
	AwsIamUserPolicyNamespaceListerExpansion
}

// awsIamUserPolicyNamespaceLister implements the AwsIamUserPolicyNamespaceLister
// interface.
type awsIamUserPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamUserPolicies in the indexer for a given namespace.
func (s awsIamUserPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamUserPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamUserPolicy))
	})
	return ret, err
}

// Get retrieves the AwsIamUserPolicy from the indexer for a given namespace and name.
func (s awsIamUserPolicyNamespaceLister) Get(name string) (*v1alpha1.AwsIamUserPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiamuserpolicy"), name)
	}
	return obj.(*v1alpha1.AwsIamUserPolicy), nil
}
