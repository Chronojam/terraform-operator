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

// AwsVpcEndpointSubnetAssociationLister helps list AwsVpcEndpointSubnetAssociations.
type AwsVpcEndpointSubnetAssociationLister interface {
	// List lists all AwsVpcEndpointSubnetAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointSubnetAssociation, err error)
	// AwsVpcEndpointSubnetAssociations returns an object that can list and get AwsVpcEndpointSubnetAssociations.
	AwsVpcEndpointSubnetAssociations(namespace string) AwsVpcEndpointSubnetAssociationNamespaceLister
	AwsVpcEndpointSubnetAssociationListerExpansion
}

// awsVpcEndpointSubnetAssociationLister implements the AwsVpcEndpointSubnetAssociationLister interface.
type awsVpcEndpointSubnetAssociationLister struct {
	indexer cache.Indexer
}

// NewAwsVpcEndpointSubnetAssociationLister returns a new AwsVpcEndpointSubnetAssociationLister.
func NewAwsVpcEndpointSubnetAssociationLister(indexer cache.Indexer) AwsVpcEndpointSubnetAssociationLister {
	return &awsVpcEndpointSubnetAssociationLister{indexer: indexer}
}

// List lists all AwsVpcEndpointSubnetAssociations in the indexer.
func (s *awsVpcEndpointSubnetAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointSubnetAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsVpcEndpointSubnetAssociation))
	})
	return ret, err
}

// AwsVpcEndpointSubnetAssociations returns an object that can list and get AwsVpcEndpointSubnetAssociations.
func (s *awsVpcEndpointSubnetAssociationLister) AwsVpcEndpointSubnetAssociations(namespace string) AwsVpcEndpointSubnetAssociationNamespaceLister {
	return awsVpcEndpointSubnetAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsVpcEndpointSubnetAssociationNamespaceLister helps list and get AwsVpcEndpointSubnetAssociations.
type AwsVpcEndpointSubnetAssociationNamespaceLister interface {
	// List lists all AwsVpcEndpointSubnetAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointSubnetAssociation, err error)
	// Get retrieves the AwsVpcEndpointSubnetAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsVpcEndpointSubnetAssociation, error)
	AwsVpcEndpointSubnetAssociationNamespaceListerExpansion
}

// awsVpcEndpointSubnetAssociationNamespaceLister implements the AwsVpcEndpointSubnetAssociationNamespaceLister
// interface.
type awsVpcEndpointSubnetAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsVpcEndpointSubnetAssociations in the indexer for a given namespace.
func (s awsVpcEndpointSubnetAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsVpcEndpointSubnetAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsVpcEndpointSubnetAssociation))
	})
	return ret, err
}

// Get retrieves the AwsVpcEndpointSubnetAssociation from the indexer for a given namespace and name.
func (s awsVpcEndpointSubnetAssociationNamespaceLister) Get(name string) (*v1alpha1.AwsVpcEndpointSubnetAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsvpcendpointsubnetassociation"), name)
	}
	return obj.(*v1alpha1.AwsVpcEndpointSubnetAssociation), nil
}
