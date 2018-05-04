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

// AwsIamAccountAliasLister helps list AwsIamAccountAliases.
type AwsIamAccountAliasLister interface {
	// List lists all AwsIamAccountAliases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccountAlias, err error)
	// AwsIamAccountAliases returns an object that can list and get AwsIamAccountAliases.
	AwsIamAccountAliases(namespace string) AwsIamAccountAliasNamespaceLister
	AwsIamAccountAliasListerExpansion
}

// awsIamAccountAliasLister implements the AwsIamAccountAliasLister interface.
type awsIamAccountAliasLister struct {
	indexer cache.Indexer
}

// NewAwsIamAccountAliasLister returns a new AwsIamAccountAliasLister.
func NewAwsIamAccountAliasLister(indexer cache.Indexer) AwsIamAccountAliasLister {
	return &awsIamAccountAliasLister{indexer: indexer}
}

// List lists all AwsIamAccountAliases in the indexer.
func (s *awsIamAccountAliasLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccountAlias, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamAccountAlias))
	})
	return ret, err
}

// AwsIamAccountAliases returns an object that can list and get AwsIamAccountAliases.
func (s *awsIamAccountAliasLister) AwsIamAccountAliases(namespace string) AwsIamAccountAliasNamespaceLister {
	return awsIamAccountAliasNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamAccountAliasNamespaceLister helps list and get AwsIamAccountAliases.
type AwsIamAccountAliasNamespaceLister interface {
	// List lists all AwsIamAccountAliases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccountAlias, err error)
	// Get retrieves the AwsIamAccountAlias from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamAccountAlias, error)
	AwsIamAccountAliasNamespaceListerExpansion
}

// awsIamAccountAliasNamespaceLister implements the AwsIamAccountAliasNamespaceLister
// interface.
type awsIamAccountAliasNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamAccountAliases in the indexer for a given namespace.
func (s awsIamAccountAliasNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccountAlias, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamAccountAlias))
	})
	return ret, err
}

// Get retrieves the AwsIamAccountAlias from the indexer for a given namespace and name.
func (s awsIamAccountAliasNamespaceLister) Get(name string) (*v1alpha1.AwsIamAccountAlias, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiamaccountalias"), name)
	}
	return obj.(*v1alpha1.AwsIamAccountAlias), nil
}
