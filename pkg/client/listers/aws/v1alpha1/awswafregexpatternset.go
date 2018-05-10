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

// AwsWafRegexPatternSetLister helps list AwsWafRegexPatternSets.
type AwsWafRegexPatternSetLister interface {
	// List lists all AwsWafRegexPatternSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsWafRegexPatternSet, err error)
	// AwsWafRegexPatternSets returns an object that can list and get AwsWafRegexPatternSets.
	AwsWafRegexPatternSets(namespace string) AwsWafRegexPatternSetNamespaceLister
	AwsWafRegexPatternSetListerExpansion
}

// awsWafRegexPatternSetLister implements the AwsWafRegexPatternSetLister interface.
type awsWafRegexPatternSetLister struct {
	indexer cache.Indexer
}

// NewAwsWafRegexPatternSetLister returns a new AwsWafRegexPatternSetLister.
func NewAwsWafRegexPatternSetLister(indexer cache.Indexer) AwsWafRegexPatternSetLister {
	return &awsWafRegexPatternSetLister{indexer: indexer}
}

// List lists all AwsWafRegexPatternSets in the indexer.
func (s *awsWafRegexPatternSetLister) List(selector labels.Selector) (ret []*v1alpha1.AwsWafRegexPatternSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsWafRegexPatternSet))
	})
	return ret, err
}

// AwsWafRegexPatternSets returns an object that can list and get AwsWafRegexPatternSets.
func (s *awsWafRegexPatternSetLister) AwsWafRegexPatternSets(namespace string) AwsWafRegexPatternSetNamespaceLister {
	return awsWafRegexPatternSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsWafRegexPatternSetNamespaceLister helps list and get AwsWafRegexPatternSets.
type AwsWafRegexPatternSetNamespaceLister interface {
	// List lists all AwsWafRegexPatternSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsWafRegexPatternSet, err error)
	// Get retrieves the AwsWafRegexPatternSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsWafRegexPatternSet, error)
	AwsWafRegexPatternSetNamespaceListerExpansion
}

// awsWafRegexPatternSetNamespaceLister implements the AwsWafRegexPatternSetNamespaceLister
// interface.
type awsWafRegexPatternSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsWafRegexPatternSets in the indexer for a given namespace.
func (s awsWafRegexPatternSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsWafRegexPatternSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsWafRegexPatternSet))
	})
	return ret, err
}

// Get retrieves the AwsWafRegexPatternSet from the indexer for a given namespace and name.
func (s awsWafRegexPatternSetNamespaceLister) Get(name string) (*v1alpha1.AwsWafRegexPatternSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awswafregexpatternset"), name)
	}
	return obj.(*v1alpha1.AwsWafRegexPatternSet), nil
}
