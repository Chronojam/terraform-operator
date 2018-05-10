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

// AwsRedshiftSecurityGroupLister helps list AwsRedshiftSecurityGroups.
type AwsRedshiftSecurityGroupLister interface {
	// List lists all AwsRedshiftSecurityGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftSecurityGroup, err error)
	// AwsRedshiftSecurityGroups returns an object that can list and get AwsRedshiftSecurityGroups.
	AwsRedshiftSecurityGroups(namespace string) AwsRedshiftSecurityGroupNamespaceLister
	AwsRedshiftSecurityGroupListerExpansion
}

// awsRedshiftSecurityGroupLister implements the AwsRedshiftSecurityGroupLister interface.
type awsRedshiftSecurityGroupLister struct {
	indexer cache.Indexer
}

// NewAwsRedshiftSecurityGroupLister returns a new AwsRedshiftSecurityGroupLister.
func NewAwsRedshiftSecurityGroupLister(indexer cache.Indexer) AwsRedshiftSecurityGroupLister {
	return &awsRedshiftSecurityGroupLister{indexer: indexer}
}

// List lists all AwsRedshiftSecurityGroups in the indexer.
func (s *awsRedshiftSecurityGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftSecurityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRedshiftSecurityGroup))
	})
	return ret, err
}

// AwsRedshiftSecurityGroups returns an object that can list and get AwsRedshiftSecurityGroups.
func (s *awsRedshiftSecurityGroupLister) AwsRedshiftSecurityGroups(namespace string) AwsRedshiftSecurityGroupNamespaceLister {
	return awsRedshiftSecurityGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsRedshiftSecurityGroupNamespaceLister helps list and get AwsRedshiftSecurityGroups.
type AwsRedshiftSecurityGroupNamespaceLister interface {
	// List lists all AwsRedshiftSecurityGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftSecurityGroup, err error)
	// Get retrieves the AwsRedshiftSecurityGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsRedshiftSecurityGroup, error)
	AwsRedshiftSecurityGroupNamespaceListerExpansion
}

// awsRedshiftSecurityGroupNamespaceLister implements the AwsRedshiftSecurityGroupNamespaceLister
// interface.
type awsRedshiftSecurityGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsRedshiftSecurityGroups in the indexer for a given namespace.
func (s awsRedshiftSecurityGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftSecurityGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRedshiftSecurityGroup))
	})
	return ret, err
}

// Get retrieves the AwsRedshiftSecurityGroup from the indexer for a given namespace and name.
func (s awsRedshiftSecurityGroupNamespaceLister) Get(name string) (*v1alpha1.AwsRedshiftSecurityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsredshiftsecuritygroup"), name)
	}
	return obj.(*v1alpha1.AwsRedshiftSecurityGroup), nil
}
