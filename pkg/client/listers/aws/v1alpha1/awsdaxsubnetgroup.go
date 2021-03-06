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

// AwsDaxSubnetGroupLister helps list AwsDaxSubnetGroups.
type AwsDaxSubnetGroupLister interface {
	// List lists all AwsDaxSubnetGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDaxSubnetGroup, err error)
	// AwsDaxSubnetGroups returns an object that can list and get AwsDaxSubnetGroups.
	AwsDaxSubnetGroups(namespace string) AwsDaxSubnetGroupNamespaceLister
	AwsDaxSubnetGroupListerExpansion
}

// awsDaxSubnetGroupLister implements the AwsDaxSubnetGroupLister interface.
type awsDaxSubnetGroupLister struct {
	indexer cache.Indexer
}

// NewAwsDaxSubnetGroupLister returns a new AwsDaxSubnetGroupLister.
func NewAwsDaxSubnetGroupLister(indexer cache.Indexer) AwsDaxSubnetGroupLister {
	return &awsDaxSubnetGroupLister{indexer: indexer}
}

// List lists all AwsDaxSubnetGroups in the indexer.
func (s *awsDaxSubnetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDaxSubnetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDaxSubnetGroup))
	})
	return ret, err
}

// AwsDaxSubnetGroups returns an object that can list and get AwsDaxSubnetGroups.
func (s *awsDaxSubnetGroupLister) AwsDaxSubnetGroups(namespace string) AwsDaxSubnetGroupNamespaceLister {
	return awsDaxSubnetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDaxSubnetGroupNamespaceLister helps list and get AwsDaxSubnetGroups.
type AwsDaxSubnetGroupNamespaceLister interface {
	// List lists all AwsDaxSubnetGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDaxSubnetGroup, err error)
	// Get retrieves the AwsDaxSubnetGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsDaxSubnetGroup, error)
	AwsDaxSubnetGroupNamespaceListerExpansion
}

// awsDaxSubnetGroupNamespaceLister implements the AwsDaxSubnetGroupNamespaceLister
// interface.
type awsDaxSubnetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDaxSubnetGroups in the indexer for a given namespace.
func (s awsDaxSubnetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDaxSubnetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDaxSubnetGroup))
	})
	return ret, err
}

// Get retrieves the AwsDaxSubnetGroup from the indexer for a given namespace and name.
func (s awsDaxSubnetGroupNamespaceLister) Get(name string) (*v1alpha1.AwsDaxSubnetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsdaxsubnetgroup"), name)
	}
	return obj.(*v1alpha1.AwsDaxSubnetGroup), nil
}
