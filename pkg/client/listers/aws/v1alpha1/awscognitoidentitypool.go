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

// AwsCognitoIdentityPoolLister helps list AwsCognitoIdentityPools.
type AwsCognitoIdentityPoolLister interface {
	// List lists all AwsCognitoIdentityPools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCognitoIdentityPool, err error)
	// AwsCognitoIdentityPools returns an object that can list and get AwsCognitoIdentityPools.
	AwsCognitoIdentityPools(namespace string) AwsCognitoIdentityPoolNamespaceLister
	AwsCognitoIdentityPoolListerExpansion
}

// awsCognitoIdentityPoolLister implements the AwsCognitoIdentityPoolLister interface.
type awsCognitoIdentityPoolLister struct {
	indexer cache.Indexer
}

// NewAwsCognitoIdentityPoolLister returns a new AwsCognitoIdentityPoolLister.
func NewAwsCognitoIdentityPoolLister(indexer cache.Indexer) AwsCognitoIdentityPoolLister {
	return &awsCognitoIdentityPoolLister{indexer: indexer}
}

// List lists all AwsCognitoIdentityPools in the indexer.
func (s *awsCognitoIdentityPoolLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCognitoIdentityPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCognitoIdentityPool))
	})
	return ret, err
}

// AwsCognitoIdentityPools returns an object that can list and get AwsCognitoIdentityPools.
func (s *awsCognitoIdentityPoolLister) AwsCognitoIdentityPools(namespace string) AwsCognitoIdentityPoolNamespaceLister {
	return awsCognitoIdentityPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCognitoIdentityPoolNamespaceLister helps list and get AwsCognitoIdentityPools.
type AwsCognitoIdentityPoolNamespaceLister interface {
	// List lists all AwsCognitoIdentityPools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCognitoIdentityPool, err error)
	// Get retrieves the AwsCognitoIdentityPool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsCognitoIdentityPool, error)
	AwsCognitoIdentityPoolNamespaceListerExpansion
}

// awsCognitoIdentityPoolNamespaceLister implements the AwsCognitoIdentityPoolNamespaceLister
// interface.
type awsCognitoIdentityPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCognitoIdentityPools in the indexer for a given namespace.
func (s awsCognitoIdentityPoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCognitoIdentityPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCognitoIdentityPool))
	})
	return ret, err
}

// Get retrieves the AwsCognitoIdentityPool from the indexer for a given namespace and name.
func (s awsCognitoIdentityPoolNamespaceLister) Get(name string) (*v1alpha1.AwsCognitoIdentityPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awscognitoidentitypool"), name)
	}
	return obj.(*v1alpha1.AwsCognitoIdentityPool), nil
}
