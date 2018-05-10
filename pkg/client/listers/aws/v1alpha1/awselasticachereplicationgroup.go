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

// AwsElasticacheReplicationGroupLister helps list AwsElasticacheReplicationGroups.
type AwsElasticacheReplicationGroupLister interface {
	// List lists all AwsElasticacheReplicationGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElasticacheReplicationGroup, err error)
	// AwsElasticacheReplicationGroups returns an object that can list and get AwsElasticacheReplicationGroups.
	AwsElasticacheReplicationGroups(namespace string) AwsElasticacheReplicationGroupNamespaceLister
	AwsElasticacheReplicationGroupListerExpansion
}

// awsElasticacheReplicationGroupLister implements the AwsElasticacheReplicationGroupLister interface.
type awsElasticacheReplicationGroupLister struct {
	indexer cache.Indexer
}

// NewAwsElasticacheReplicationGroupLister returns a new AwsElasticacheReplicationGroupLister.
func NewAwsElasticacheReplicationGroupLister(indexer cache.Indexer) AwsElasticacheReplicationGroupLister {
	return &awsElasticacheReplicationGroupLister{indexer: indexer}
}

// List lists all AwsElasticacheReplicationGroups in the indexer.
func (s *awsElasticacheReplicationGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElasticacheReplicationGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElasticacheReplicationGroup))
	})
	return ret, err
}

// AwsElasticacheReplicationGroups returns an object that can list and get AwsElasticacheReplicationGroups.
func (s *awsElasticacheReplicationGroupLister) AwsElasticacheReplicationGroups(namespace string) AwsElasticacheReplicationGroupNamespaceLister {
	return awsElasticacheReplicationGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsElasticacheReplicationGroupNamespaceLister helps list and get AwsElasticacheReplicationGroups.
type AwsElasticacheReplicationGroupNamespaceLister interface {
	// List lists all AwsElasticacheReplicationGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElasticacheReplicationGroup, err error)
	// Get retrieves the AwsElasticacheReplicationGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsElasticacheReplicationGroup, error)
	AwsElasticacheReplicationGroupNamespaceListerExpansion
}

// awsElasticacheReplicationGroupNamespaceLister implements the AwsElasticacheReplicationGroupNamespaceLister
// interface.
type awsElasticacheReplicationGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsElasticacheReplicationGroups in the indexer for a given namespace.
func (s awsElasticacheReplicationGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElasticacheReplicationGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElasticacheReplicationGroup))
	})
	return ret, err
}

// Get retrieves the AwsElasticacheReplicationGroup from the indexer for a given namespace and name.
func (s awsElasticacheReplicationGroupNamespaceLister) Get(name string) (*v1alpha1.AwsElasticacheReplicationGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awselasticachereplicationgroup"), name)
	}
	return obj.(*v1alpha1.AwsElasticacheReplicationGroup), nil
}
