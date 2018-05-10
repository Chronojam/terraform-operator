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

// AwsWafregionalXssMatchSetLister helps list AwsWafregionalXssMatchSets.
type AwsWafregionalXssMatchSetLister interface {
	// List lists all AwsWafregionalXssMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalXssMatchSet, err error)
	// AwsWafregionalXssMatchSets returns an object that can list and get AwsWafregionalXssMatchSets.
	AwsWafregionalXssMatchSets(namespace string) AwsWafregionalXssMatchSetNamespaceLister
	AwsWafregionalXssMatchSetListerExpansion
}

// awsWafregionalXssMatchSetLister implements the AwsWafregionalXssMatchSetLister interface.
type awsWafregionalXssMatchSetLister struct {
	indexer cache.Indexer
}

// NewAwsWafregionalXssMatchSetLister returns a new AwsWafregionalXssMatchSetLister.
func NewAwsWafregionalXssMatchSetLister(indexer cache.Indexer) AwsWafregionalXssMatchSetLister {
	return &awsWafregionalXssMatchSetLister{indexer: indexer}
}

// List lists all AwsWafregionalXssMatchSets in the indexer.
func (s *awsWafregionalXssMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalXssMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsWafregionalXssMatchSet))
	})
	return ret, err
}

// AwsWafregionalXssMatchSets returns an object that can list and get AwsWafregionalXssMatchSets.
func (s *awsWafregionalXssMatchSetLister) AwsWafregionalXssMatchSets(namespace string) AwsWafregionalXssMatchSetNamespaceLister {
	return awsWafregionalXssMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsWafregionalXssMatchSetNamespaceLister helps list and get AwsWafregionalXssMatchSets.
type AwsWafregionalXssMatchSetNamespaceLister interface {
	// List lists all AwsWafregionalXssMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalXssMatchSet, err error)
	// Get retrieves the AwsWafregionalXssMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsWafregionalXssMatchSet, error)
	AwsWafregionalXssMatchSetNamespaceListerExpansion
}

// awsWafregionalXssMatchSetNamespaceLister implements the AwsWafregionalXssMatchSetNamespaceLister
// interface.
type awsWafregionalXssMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsWafregionalXssMatchSets in the indexer for a given namespace.
func (s awsWafregionalXssMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalXssMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsWafregionalXssMatchSet))
	})
	return ret, err
}

// Get retrieves the AwsWafregionalXssMatchSet from the indexer for a given namespace and name.
func (s awsWafregionalXssMatchSetNamespaceLister) Get(name string) (*v1alpha1.AwsWafregionalXssMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awswafregionalxssmatchset"), name)
	}
	return obj.(*v1alpha1.AwsWafregionalXssMatchSet), nil
}
