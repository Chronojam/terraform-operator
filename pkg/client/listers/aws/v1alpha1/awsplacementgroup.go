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

// AwsPlacementGroupLister helps list AwsPlacementGroups.
type AwsPlacementGroupLister interface {
	// List lists all AwsPlacementGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsPlacementGroup, err error)
	// AwsPlacementGroups returns an object that can list and get AwsPlacementGroups.
	AwsPlacementGroups(namespace string) AwsPlacementGroupNamespaceLister
	AwsPlacementGroupListerExpansion
}

// awsPlacementGroupLister implements the AwsPlacementGroupLister interface.
type awsPlacementGroupLister struct {
	indexer cache.Indexer
}

// NewAwsPlacementGroupLister returns a new AwsPlacementGroupLister.
func NewAwsPlacementGroupLister(indexer cache.Indexer) AwsPlacementGroupLister {
	return &awsPlacementGroupLister{indexer: indexer}
}

// List lists all AwsPlacementGroups in the indexer.
func (s *awsPlacementGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AwsPlacementGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsPlacementGroup))
	})
	return ret, err
}

// AwsPlacementGroups returns an object that can list and get AwsPlacementGroups.
func (s *awsPlacementGroupLister) AwsPlacementGroups(namespace string) AwsPlacementGroupNamespaceLister {
	return awsPlacementGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsPlacementGroupNamespaceLister helps list and get AwsPlacementGroups.
type AwsPlacementGroupNamespaceLister interface {
	// List lists all AwsPlacementGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsPlacementGroup, err error)
	// Get retrieves the AwsPlacementGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsPlacementGroup, error)
	AwsPlacementGroupNamespaceListerExpansion
}

// awsPlacementGroupNamespaceLister implements the AwsPlacementGroupNamespaceLister
// interface.
type awsPlacementGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsPlacementGroups in the indexer for a given namespace.
func (s awsPlacementGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsPlacementGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsPlacementGroup))
	})
	return ret, err
}

// Get retrieves the AwsPlacementGroup from the indexer for a given namespace and name.
func (s awsPlacementGroupNamespaceLister) Get(name string) (*v1alpha1.AwsPlacementGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsplacementgroup"), name)
	}
	return obj.(*v1alpha1.AwsPlacementGroup), nil
}
