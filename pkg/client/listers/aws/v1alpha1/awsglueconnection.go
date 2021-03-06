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

// AwsGlueConnectionLister helps list AwsGlueConnections.
type AwsGlueConnectionLister interface {
	// List lists all AwsGlueConnections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsGlueConnection, err error)
	// AwsGlueConnections returns an object that can list and get AwsGlueConnections.
	AwsGlueConnections(namespace string) AwsGlueConnectionNamespaceLister
	AwsGlueConnectionListerExpansion
}

// awsGlueConnectionLister implements the AwsGlueConnectionLister interface.
type awsGlueConnectionLister struct {
	indexer cache.Indexer
}

// NewAwsGlueConnectionLister returns a new AwsGlueConnectionLister.
func NewAwsGlueConnectionLister(indexer cache.Indexer) AwsGlueConnectionLister {
	return &awsGlueConnectionLister{indexer: indexer}
}

// List lists all AwsGlueConnections in the indexer.
func (s *awsGlueConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.AwsGlueConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsGlueConnection))
	})
	return ret, err
}

// AwsGlueConnections returns an object that can list and get AwsGlueConnections.
func (s *awsGlueConnectionLister) AwsGlueConnections(namespace string) AwsGlueConnectionNamespaceLister {
	return awsGlueConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsGlueConnectionNamespaceLister helps list and get AwsGlueConnections.
type AwsGlueConnectionNamespaceLister interface {
	// List lists all AwsGlueConnections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsGlueConnection, err error)
	// Get retrieves the AwsGlueConnection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsGlueConnection, error)
	AwsGlueConnectionNamespaceListerExpansion
}

// awsGlueConnectionNamespaceLister implements the AwsGlueConnectionNamespaceLister
// interface.
type awsGlueConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsGlueConnections in the indexer for a given namespace.
func (s awsGlueConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsGlueConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsGlueConnection))
	})
	return ret, err
}

// Get retrieves the AwsGlueConnection from the indexer for a given namespace and name.
func (s awsGlueConnectionNamespaceLister) Get(name string) (*v1alpha1.AwsGlueConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsglueconnection"), name)
	}
	return obj.(*v1alpha1.AwsGlueConnection), nil
}
