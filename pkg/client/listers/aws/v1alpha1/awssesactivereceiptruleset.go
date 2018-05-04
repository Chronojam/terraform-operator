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

// AwsSesActiveReceiptRuleSetLister helps list AwsSesActiveReceiptRuleSets.
type AwsSesActiveReceiptRuleSetLister interface {
	// List lists all AwsSesActiveReceiptRuleSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSesActiveReceiptRuleSet, err error)
	// AwsSesActiveReceiptRuleSets returns an object that can list and get AwsSesActiveReceiptRuleSets.
	AwsSesActiveReceiptRuleSets(namespace string) AwsSesActiveReceiptRuleSetNamespaceLister
	AwsSesActiveReceiptRuleSetListerExpansion
}

// awsSesActiveReceiptRuleSetLister implements the AwsSesActiveReceiptRuleSetLister interface.
type awsSesActiveReceiptRuleSetLister struct {
	indexer cache.Indexer
}

// NewAwsSesActiveReceiptRuleSetLister returns a new AwsSesActiveReceiptRuleSetLister.
func NewAwsSesActiveReceiptRuleSetLister(indexer cache.Indexer) AwsSesActiveReceiptRuleSetLister {
	return &awsSesActiveReceiptRuleSetLister{indexer: indexer}
}

// List lists all AwsSesActiveReceiptRuleSets in the indexer.
func (s *awsSesActiveReceiptRuleSetLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSesActiveReceiptRuleSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSesActiveReceiptRuleSet))
	})
	return ret, err
}

// AwsSesActiveReceiptRuleSets returns an object that can list and get AwsSesActiveReceiptRuleSets.
func (s *awsSesActiveReceiptRuleSetLister) AwsSesActiveReceiptRuleSets(namespace string) AwsSesActiveReceiptRuleSetNamespaceLister {
	return awsSesActiveReceiptRuleSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSesActiveReceiptRuleSetNamespaceLister helps list and get AwsSesActiveReceiptRuleSets.
type AwsSesActiveReceiptRuleSetNamespaceLister interface {
	// List lists all AwsSesActiveReceiptRuleSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSesActiveReceiptRuleSet, err error)
	// Get retrieves the AwsSesActiveReceiptRuleSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSesActiveReceiptRuleSet, error)
	AwsSesActiveReceiptRuleSetNamespaceListerExpansion
}

// awsSesActiveReceiptRuleSetNamespaceLister implements the AwsSesActiveReceiptRuleSetNamespaceLister
// interface.
type awsSesActiveReceiptRuleSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSesActiveReceiptRuleSets in the indexer for a given namespace.
func (s awsSesActiveReceiptRuleSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSesActiveReceiptRuleSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSesActiveReceiptRuleSet))
	})
	return ret, err
}

// Get retrieves the AwsSesActiveReceiptRuleSet from the indexer for a given namespace and name.
func (s awsSesActiveReceiptRuleSetNamespaceLister) Get(name string) (*v1alpha1.AwsSesActiveReceiptRuleSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awssesactivereceiptruleset"), name)
	}
	return obj.(*v1alpha1.AwsSesActiveReceiptRuleSet), nil
}
