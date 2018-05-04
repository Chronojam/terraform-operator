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

// AwsSecurityGroupRuleLister helps list AwsSecurityGroupRules.
type AwsSecurityGroupRuleLister interface {
	// List lists all AwsSecurityGroupRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSecurityGroupRule, err error)
	// AwsSecurityGroupRules returns an object that can list and get AwsSecurityGroupRules.
	AwsSecurityGroupRules(namespace string) AwsSecurityGroupRuleNamespaceLister
	AwsSecurityGroupRuleListerExpansion
}

// awsSecurityGroupRuleLister implements the AwsSecurityGroupRuleLister interface.
type awsSecurityGroupRuleLister struct {
	indexer cache.Indexer
}

// NewAwsSecurityGroupRuleLister returns a new AwsSecurityGroupRuleLister.
func NewAwsSecurityGroupRuleLister(indexer cache.Indexer) AwsSecurityGroupRuleLister {
	return &awsSecurityGroupRuleLister{indexer: indexer}
}

// List lists all AwsSecurityGroupRules in the indexer.
func (s *awsSecurityGroupRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSecurityGroupRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSecurityGroupRule))
	})
	return ret, err
}

// AwsSecurityGroupRules returns an object that can list and get AwsSecurityGroupRules.
func (s *awsSecurityGroupRuleLister) AwsSecurityGroupRules(namespace string) AwsSecurityGroupRuleNamespaceLister {
	return awsSecurityGroupRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSecurityGroupRuleNamespaceLister helps list and get AwsSecurityGroupRules.
type AwsSecurityGroupRuleNamespaceLister interface {
	// List lists all AwsSecurityGroupRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSecurityGroupRule, err error)
	// Get retrieves the AwsSecurityGroupRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSecurityGroupRule, error)
	AwsSecurityGroupRuleNamespaceListerExpansion
}

// awsSecurityGroupRuleNamespaceLister implements the AwsSecurityGroupRuleNamespaceLister
// interface.
type awsSecurityGroupRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSecurityGroupRules in the indexer for a given namespace.
func (s awsSecurityGroupRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSecurityGroupRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSecurityGroupRule))
	})
	return ret, err
}

// Get retrieves the AwsSecurityGroupRule from the indexer for a given namespace and name.
func (s awsSecurityGroupRuleNamespaceLister) Get(name string) (*v1alpha1.AwsSecurityGroupRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awssecuritygrouprule"), name)
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), nil
}
