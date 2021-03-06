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

// AwsAppautoscalingPolicyLister helps list AwsAppautoscalingPolicies.
type AwsAppautoscalingPolicyLister interface {
	// List lists all AwsAppautoscalingPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAppautoscalingPolicy, err error)
	// AwsAppautoscalingPolicies returns an object that can list and get AwsAppautoscalingPolicies.
	AwsAppautoscalingPolicies(namespace string) AwsAppautoscalingPolicyNamespaceLister
	AwsAppautoscalingPolicyListerExpansion
}

// awsAppautoscalingPolicyLister implements the AwsAppautoscalingPolicyLister interface.
type awsAppautoscalingPolicyLister struct {
	indexer cache.Indexer
}

// NewAwsAppautoscalingPolicyLister returns a new AwsAppautoscalingPolicyLister.
func NewAwsAppautoscalingPolicyLister(indexer cache.Indexer) AwsAppautoscalingPolicyLister {
	return &awsAppautoscalingPolicyLister{indexer: indexer}
}

// List lists all AwsAppautoscalingPolicies in the indexer.
func (s *awsAppautoscalingPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAppautoscalingPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAppautoscalingPolicy))
	})
	return ret, err
}

// AwsAppautoscalingPolicies returns an object that can list and get AwsAppautoscalingPolicies.
func (s *awsAppautoscalingPolicyLister) AwsAppautoscalingPolicies(namespace string) AwsAppautoscalingPolicyNamespaceLister {
	return awsAppautoscalingPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsAppautoscalingPolicyNamespaceLister helps list and get AwsAppautoscalingPolicies.
type AwsAppautoscalingPolicyNamespaceLister interface {
	// List lists all AwsAppautoscalingPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAppautoscalingPolicy, err error)
	// Get retrieves the AwsAppautoscalingPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsAppautoscalingPolicy, error)
	AwsAppautoscalingPolicyNamespaceListerExpansion
}

// awsAppautoscalingPolicyNamespaceLister implements the AwsAppautoscalingPolicyNamespaceLister
// interface.
type awsAppautoscalingPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsAppautoscalingPolicies in the indexer for a given namespace.
func (s awsAppautoscalingPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAppautoscalingPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAppautoscalingPolicy))
	})
	return ret, err
}

// Get retrieves the AwsAppautoscalingPolicy from the indexer for a given namespace and name.
func (s awsAppautoscalingPolicyNamespaceLister) Get(name string) (*v1alpha1.AwsAppautoscalingPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsappautoscalingpolicy"), name)
	}
	return obj.(*v1alpha1.AwsAppautoscalingPolicy), nil
}
