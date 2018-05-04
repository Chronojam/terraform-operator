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

// AwsAutoscalingPolicyLister helps list AwsAutoscalingPolicies.
type AwsAutoscalingPolicyLister interface {
	// List lists all AwsAutoscalingPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingPolicy, err error)
	// AwsAutoscalingPolicies returns an object that can list and get AwsAutoscalingPolicies.
	AwsAutoscalingPolicies(namespace string) AwsAutoscalingPolicyNamespaceLister
	AwsAutoscalingPolicyListerExpansion
}

// awsAutoscalingPolicyLister implements the AwsAutoscalingPolicyLister interface.
type awsAutoscalingPolicyLister struct {
	indexer cache.Indexer
}

// NewAwsAutoscalingPolicyLister returns a new AwsAutoscalingPolicyLister.
func NewAwsAutoscalingPolicyLister(indexer cache.Indexer) AwsAutoscalingPolicyLister {
	return &awsAutoscalingPolicyLister{indexer: indexer}
}

// List lists all AwsAutoscalingPolicies in the indexer.
func (s *awsAutoscalingPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAutoscalingPolicy))
	})
	return ret, err
}

// AwsAutoscalingPolicies returns an object that can list and get AwsAutoscalingPolicies.
func (s *awsAutoscalingPolicyLister) AwsAutoscalingPolicies(namespace string) AwsAutoscalingPolicyNamespaceLister {
	return awsAutoscalingPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsAutoscalingPolicyNamespaceLister helps list and get AwsAutoscalingPolicies.
type AwsAutoscalingPolicyNamespaceLister interface {
	// List lists all AwsAutoscalingPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingPolicy, err error)
	// Get retrieves the AwsAutoscalingPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsAutoscalingPolicy, error)
	AwsAutoscalingPolicyNamespaceListerExpansion
}

// awsAutoscalingPolicyNamespaceLister implements the AwsAutoscalingPolicyNamespaceLister
// interface.
type awsAutoscalingPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsAutoscalingPolicies in the indexer for a given namespace.
func (s awsAutoscalingPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAutoscalingPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAutoscalingPolicy))
	})
	return ret, err
}

// Get retrieves the AwsAutoscalingPolicy from the indexer for a given namespace and name.
func (s awsAutoscalingPolicyNamespaceLister) Get(name string) (*v1alpha1.AwsAutoscalingPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsautoscalingpolicy"), name)
	}
	return obj.(*v1alpha1.AwsAutoscalingPolicy), nil
}
