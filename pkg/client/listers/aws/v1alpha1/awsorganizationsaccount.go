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

// AwsOrganizationsAccountLister helps list AwsOrganizationsAccounts.
type AwsOrganizationsAccountLister interface {
	// List lists all AwsOrganizationsAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsOrganizationsAccount, err error)
	// AwsOrganizationsAccounts returns an object that can list and get AwsOrganizationsAccounts.
	AwsOrganizationsAccounts(namespace string) AwsOrganizationsAccountNamespaceLister
	AwsOrganizationsAccountListerExpansion
}

// awsOrganizationsAccountLister implements the AwsOrganizationsAccountLister interface.
type awsOrganizationsAccountLister struct {
	indexer cache.Indexer
}

// NewAwsOrganizationsAccountLister returns a new AwsOrganizationsAccountLister.
func NewAwsOrganizationsAccountLister(indexer cache.Indexer) AwsOrganizationsAccountLister {
	return &awsOrganizationsAccountLister{indexer: indexer}
}

// List lists all AwsOrganizationsAccounts in the indexer.
func (s *awsOrganizationsAccountLister) List(selector labels.Selector) (ret []*v1alpha1.AwsOrganizationsAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsOrganizationsAccount))
	})
	return ret, err
}

// AwsOrganizationsAccounts returns an object that can list and get AwsOrganizationsAccounts.
func (s *awsOrganizationsAccountLister) AwsOrganizationsAccounts(namespace string) AwsOrganizationsAccountNamespaceLister {
	return awsOrganizationsAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsOrganizationsAccountNamespaceLister helps list and get AwsOrganizationsAccounts.
type AwsOrganizationsAccountNamespaceLister interface {
	// List lists all AwsOrganizationsAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsOrganizationsAccount, err error)
	// Get retrieves the AwsOrganizationsAccount from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsOrganizationsAccount, error)
	AwsOrganizationsAccountNamespaceListerExpansion
}

// awsOrganizationsAccountNamespaceLister implements the AwsOrganizationsAccountNamespaceLister
// interface.
type awsOrganizationsAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsOrganizationsAccounts in the indexer for a given namespace.
func (s awsOrganizationsAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsOrganizationsAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsOrganizationsAccount))
	})
	return ret, err
}

// Get retrieves the AwsOrganizationsAccount from the indexer for a given namespace and name.
func (s awsOrganizationsAccountNamespaceLister) Get(name string) (*v1alpha1.AwsOrganizationsAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsorganizationsaccount"), name)
	}
	return obj.(*v1alpha1.AwsOrganizationsAccount), nil
}
