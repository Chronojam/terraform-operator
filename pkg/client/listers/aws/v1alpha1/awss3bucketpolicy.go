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

// AwsS3BucketPolicyLister helps list AwsS3BucketPolicies.
type AwsS3BucketPolicyLister interface {
	// List lists all AwsS3BucketPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsS3BucketPolicy, err error)
	// AwsS3BucketPolicies returns an object that can list and get AwsS3BucketPolicies.
	AwsS3BucketPolicies(namespace string) AwsS3BucketPolicyNamespaceLister
	AwsS3BucketPolicyListerExpansion
}

// awsS3BucketPolicyLister implements the AwsS3BucketPolicyLister interface.
type awsS3BucketPolicyLister struct {
	indexer cache.Indexer
}

// NewAwsS3BucketPolicyLister returns a new AwsS3BucketPolicyLister.
func NewAwsS3BucketPolicyLister(indexer cache.Indexer) AwsS3BucketPolicyLister {
	return &awsS3BucketPolicyLister{indexer: indexer}
}

// List lists all AwsS3BucketPolicies in the indexer.
func (s *awsS3BucketPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsS3BucketPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsS3BucketPolicy))
	})
	return ret, err
}

// AwsS3BucketPolicies returns an object that can list and get AwsS3BucketPolicies.
func (s *awsS3BucketPolicyLister) AwsS3BucketPolicies(namespace string) AwsS3BucketPolicyNamespaceLister {
	return awsS3BucketPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsS3BucketPolicyNamespaceLister helps list and get AwsS3BucketPolicies.
type AwsS3BucketPolicyNamespaceLister interface {
	// List lists all AwsS3BucketPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsS3BucketPolicy, err error)
	// Get retrieves the AwsS3BucketPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsS3BucketPolicy, error)
	AwsS3BucketPolicyNamespaceListerExpansion
}

// awsS3BucketPolicyNamespaceLister implements the AwsS3BucketPolicyNamespaceLister
// interface.
type awsS3BucketPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsS3BucketPolicies in the indexer for a given namespace.
func (s awsS3BucketPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsS3BucketPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsS3BucketPolicy))
	})
	return ret, err
}

// Get retrieves the AwsS3BucketPolicy from the indexer for a given namespace and name.
func (s awsS3BucketPolicyNamespaceLister) Get(name string) (*v1alpha1.AwsS3BucketPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awss3bucketpolicy"), name)
	}
	return obj.(*v1alpha1.AwsS3BucketPolicy), nil
}
