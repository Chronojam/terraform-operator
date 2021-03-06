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

// AwsSesReceiptRuleLister helps list AwsSesReceiptRules.
type AwsSesReceiptRuleLister interface {
	// List lists all AwsSesReceiptRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSesReceiptRule, err error)
	// AwsSesReceiptRules returns an object that can list and get AwsSesReceiptRules.
	AwsSesReceiptRules(namespace string) AwsSesReceiptRuleNamespaceLister
	AwsSesReceiptRuleListerExpansion
}

// awsSesReceiptRuleLister implements the AwsSesReceiptRuleLister interface.
type awsSesReceiptRuleLister struct {
	indexer cache.Indexer
}

// NewAwsSesReceiptRuleLister returns a new AwsSesReceiptRuleLister.
func NewAwsSesReceiptRuleLister(indexer cache.Indexer) AwsSesReceiptRuleLister {
	return &awsSesReceiptRuleLister{indexer: indexer}
}

// List lists all AwsSesReceiptRules in the indexer.
func (s *awsSesReceiptRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSesReceiptRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSesReceiptRule))
	})
	return ret, err
}

// AwsSesReceiptRules returns an object that can list and get AwsSesReceiptRules.
func (s *awsSesReceiptRuleLister) AwsSesReceiptRules(namespace string) AwsSesReceiptRuleNamespaceLister {
	return awsSesReceiptRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSesReceiptRuleNamespaceLister helps list and get AwsSesReceiptRules.
type AwsSesReceiptRuleNamespaceLister interface {
	// List lists all AwsSesReceiptRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSesReceiptRule, err error)
	// Get retrieves the AwsSesReceiptRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSesReceiptRule, error)
	AwsSesReceiptRuleNamespaceListerExpansion
}

// awsSesReceiptRuleNamespaceLister implements the AwsSesReceiptRuleNamespaceLister
// interface.
type awsSesReceiptRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSesReceiptRules in the indexer for a given namespace.
func (s awsSesReceiptRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSesReceiptRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSesReceiptRule))
	})
	return ret, err
}

// Get retrieves the AwsSesReceiptRule from the indexer for a given namespace and name.
func (s awsSesReceiptRuleNamespaceLister) Get(name string) (*v1alpha1.AwsSesReceiptRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awssesreceiptrule"), name)
	}
	return obj.(*v1alpha1.AwsSesReceiptRule), nil
}
