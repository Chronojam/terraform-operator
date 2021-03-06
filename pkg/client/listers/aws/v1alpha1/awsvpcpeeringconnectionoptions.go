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

// AwsVpcPeeringConnectionOptionsLister helps list AwsVpcPeeringConnectionOptionses.
type AwsVpcPeeringConnectionOptionsLister interface {
	// List lists all AwsVpcPeeringConnectionOptionses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsVpcPeeringConnectionOptions, err error)
	// AwsVpcPeeringConnectionOptionses returns an object that can list and get AwsVpcPeeringConnectionOptionses.
	AwsVpcPeeringConnectionOptionses(namespace string) AwsVpcPeeringConnectionOptionsNamespaceLister
	AwsVpcPeeringConnectionOptionsListerExpansion
}

// awsVpcPeeringConnectionOptionsLister implements the AwsVpcPeeringConnectionOptionsLister interface.
type awsVpcPeeringConnectionOptionsLister struct {
	indexer cache.Indexer
}

// NewAwsVpcPeeringConnectionOptionsLister returns a new AwsVpcPeeringConnectionOptionsLister.
func NewAwsVpcPeeringConnectionOptionsLister(indexer cache.Indexer) AwsVpcPeeringConnectionOptionsLister {
	return &awsVpcPeeringConnectionOptionsLister{indexer: indexer}
}

// List lists all AwsVpcPeeringConnectionOptionses in the indexer.
func (s *awsVpcPeeringConnectionOptionsLister) List(selector labels.Selector) (ret []*v1alpha1.AwsVpcPeeringConnectionOptions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsVpcPeeringConnectionOptions))
	})
	return ret, err
}

// AwsVpcPeeringConnectionOptionses returns an object that can list and get AwsVpcPeeringConnectionOptionses.
func (s *awsVpcPeeringConnectionOptionsLister) AwsVpcPeeringConnectionOptionses(namespace string) AwsVpcPeeringConnectionOptionsNamespaceLister {
	return awsVpcPeeringConnectionOptionsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsVpcPeeringConnectionOptionsNamespaceLister helps list and get AwsVpcPeeringConnectionOptionses.
type AwsVpcPeeringConnectionOptionsNamespaceLister interface {
	// List lists all AwsVpcPeeringConnectionOptionses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsVpcPeeringConnectionOptions, err error)
	// Get retrieves the AwsVpcPeeringConnectionOptions from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsVpcPeeringConnectionOptions, error)
	AwsVpcPeeringConnectionOptionsNamespaceListerExpansion
}

// awsVpcPeeringConnectionOptionsNamespaceLister implements the AwsVpcPeeringConnectionOptionsNamespaceLister
// interface.
type awsVpcPeeringConnectionOptionsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsVpcPeeringConnectionOptionses in the indexer for a given namespace.
func (s awsVpcPeeringConnectionOptionsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsVpcPeeringConnectionOptions, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsVpcPeeringConnectionOptions))
	})
	return ret, err
}

// Get retrieves the AwsVpcPeeringConnectionOptions from the indexer for a given namespace and name.
func (s awsVpcPeeringConnectionOptionsNamespaceLister) Get(name string) (*v1alpha1.AwsVpcPeeringConnectionOptions, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsvpcpeeringconnectionoptions"), name)
	}
	return obj.(*v1alpha1.AwsVpcPeeringConnectionOptions), nil
}
