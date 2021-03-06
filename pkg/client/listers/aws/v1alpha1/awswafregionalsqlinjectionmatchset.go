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

// AwsWafregionalSqlInjectionMatchSetLister helps list AwsWafregionalSqlInjectionMatchSets.
type AwsWafregionalSqlInjectionMatchSetLister interface {
	// List lists all AwsWafregionalSqlInjectionMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalSqlInjectionMatchSet, err error)
	// AwsWafregionalSqlInjectionMatchSets returns an object that can list and get AwsWafregionalSqlInjectionMatchSets.
	AwsWafregionalSqlInjectionMatchSets(namespace string) AwsWafregionalSqlInjectionMatchSetNamespaceLister
	AwsWafregionalSqlInjectionMatchSetListerExpansion
}

// awsWafregionalSqlInjectionMatchSetLister implements the AwsWafregionalSqlInjectionMatchSetLister interface.
type awsWafregionalSqlInjectionMatchSetLister struct {
	indexer cache.Indexer
}

// NewAwsWafregionalSqlInjectionMatchSetLister returns a new AwsWafregionalSqlInjectionMatchSetLister.
func NewAwsWafregionalSqlInjectionMatchSetLister(indexer cache.Indexer) AwsWafregionalSqlInjectionMatchSetLister {
	return &awsWafregionalSqlInjectionMatchSetLister{indexer: indexer}
}

// List lists all AwsWafregionalSqlInjectionMatchSets in the indexer.
func (s *awsWafregionalSqlInjectionMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalSqlInjectionMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsWafregionalSqlInjectionMatchSet))
	})
	return ret, err
}

// AwsWafregionalSqlInjectionMatchSets returns an object that can list and get AwsWafregionalSqlInjectionMatchSets.
func (s *awsWafregionalSqlInjectionMatchSetLister) AwsWafregionalSqlInjectionMatchSets(namespace string) AwsWafregionalSqlInjectionMatchSetNamespaceLister {
	return awsWafregionalSqlInjectionMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsWafregionalSqlInjectionMatchSetNamespaceLister helps list and get AwsWafregionalSqlInjectionMatchSets.
type AwsWafregionalSqlInjectionMatchSetNamespaceLister interface {
	// List lists all AwsWafregionalSqlInjectionMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalSqlInjectionMatchSet, err error)
	// Get retrieves the AwsWafregionalSqlInjectionMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsWafregionalSqlInjectionMatchSet, error)
	AwsWafregionalSqlInjectionMatchSetNamespaceListerExpansion
}

// awsWafregionalSqlInjectionMatchSetNamespaceLister implements the AwsWafregionalSqlInjectionMatchSetNamespaceLister
// interface.
type awsWafregionalSqlInjectionMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsWafregionalSqlInjectionMatchSets in the indexer for a given namespace.
func (s awsWafregionalSqlInjectionMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsWafregionalSqlInjectionMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsWafregionalSqlInjectionMatchSet))
	})
	return ret, err
}

// Get retrieves the AwsWafregionalSqlInjectionMatchSet from the indexer for a given namespace and name.
func (s awsWafregionalSqlInjectionMatchSetNamespaceLister) Get(name string) (*v1alpha1.AwsWafregionalSqlInjectionMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awswafregionalsqlinjectionmatchset"), name)
	}
	return obj.(*v1alpha1.AwsWafregionalSqlInjectionMatchSet), nil
}
