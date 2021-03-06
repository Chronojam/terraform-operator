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

// AwsDefaultNetworkAclLister helps list AwsDefaultNetworkAcls.
type AwsDefaultNetworkAclLister interface {
	// List lists all AwsDefaultNetworkAcls in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultNetworkAcl, err error)
	// AwsDefaultNetworkAcls returns an object that can list and get AwsDefaultNetworkAcls.
	AwsDefaultNetworkAcls(namespace string) AwsDefaultNetworkAclNamespaceLister
	AwsDefaultNetworkAclListerExpansion
}

// awsDefaultNetworkAclLister implements the AwsDefaultNetworkAclLister interface.
type awsDefaultNetworkAclLister struct {
	indexer cache.Indexer
}

// NewAwsDefaultNetworkAclLister returns a new AwsDefaultNetworkAclLister.
func NewAwsDefaultNetworkAclLister(indexer cache.Indexer) AwsDefaultNetworkAclLister {
	return &awsDefaultNetworkAclLister{indexer: indexer}
}

// List lists all AwsDefaultNetworkAcls in the indexer.
func (s *awsDefaultNetworkAclLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultNetworkAcl, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDefaultNetworkAcl))
	})
	return ret, err
}

// AwsDefaultNetworkAcls returns an object that can list and get AwsDefaultNetworkAcls.
func (s *awsDefaultNetworkAclLister) AwsDefaultNetworkAcls(namespace string) AwsDefaultNetworkAclNamespaceLister {
	return awsDefaultNetworkAclNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDefaultNetworkAclNamespaceLister helps list and get AwsDefaultNetworkAcls.
type AwsDefaultNetworkAclNamespaceLister interface {
	// List lists all AwsDefaultNetworkAcls in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultNetworkAcl, err error)
	// Get retrieves the AwsDefaultNetworkAcl from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsDefaultNetworkAcl, error)
	AwsDefaultNetworkAclNamespaceListerExpansion
}

// awsDefaultNetworkAclNamespaceLister implements the AwsDefaultNetworkAclNamespaceLister
// interface.
type awsDefaultNetworkAclNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDefaultNetworkAcls in the indexer for a given namespace.
func (s awsDefaultNetworkAclNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultNetworkAcl, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDefaultNetworkAcl))
	})
	return ret, err
}

// Get retrieves the AwsDefaultNetworkAcl from the indexer for a given namespace and name.
func (s awsDefaultNetworkAclNamespaceLister) Get(name string) (*v1alpha1.AwsDefaultNetworkAcl, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsdefaultnetworkacl"), name)
	}
	return obj.(*v1alpha1.AwsDefaultNetworkAcl), nil
}
