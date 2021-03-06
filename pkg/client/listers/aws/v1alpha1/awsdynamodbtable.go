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

// AwsDynamodbTableLister helps list AwsDynamodbTables.
type AwsDynamodbTableLister interface {
	// List lists all AwsDynamodbTables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDynamodbTable, err error)
	// AwsDynamodbTables returns an object that can list and get AwsDynamodbTables.
	AwsDynamodbTables(namespace string) AwsDynamodbTableNamespaceLister
	AwsDynamodbTableListerExpansion
}

// awsDynamodbTableLister implements the AwsDynamodbTableLister interface.
type awsDynamodbTableLister struct {
	indexer cache.Indexer
}

// NewAwsDynamodbTableLister returns a new AwsDynamodbTableLister.
func NewAwsDynamodbTableLister(indexer cache.Indexer) AwsDynamodbTableLister {
	return &awsDynamodbTableLister{indexer: indexer}
}

// List lists all AwsDynamodbTables in the indexer.
func (s *awsDynamodbTableLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDynamodbTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDynamodbTable))
	})
	return ret, err
}

// AwsDynamodbTables returns an object that can list and get AwsDynamodbTables.
func (s *awsDynamodbTableLister) AwsDynamodbTables(namespace string) AwsDynamodbTableNamespaceLister {
	return awsDynamodbTableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDynamodbTableNamespaceLister helps list and get AwsDynamodbTables.
type AwsDynamodbTableNamespaceLister interface {
	// List lists all AwsDynamodbTables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDynamodbTable, err error)
	// Get retrieves the AwsDynamodbTable from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsDynamodbTable, error)
	AwsDynamodbTableNamespaceListerExpansion
}

// awsDynamodbTableNamespaceLister implements the AwsDynamodbTableNamespaceLister
// interface.
type awsDynamodbTableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDynamodbTables in the indexer for a given namespace.
func (s awsDynamodbTableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDynamodbTable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDynamodbTable))
	})
	return ret, err
}

// Get retrieves the AwsDynamodbTable from the indexer for a given namespace and name.
func (s awsDynamodbTableNamespaceLister) Get(name string) (*v1alpha1.AwsDynamodbTable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsdynamodbtable"), name)
	}
	return obj.(*v1alpha1.AwsDynamodbTable), nil
}
