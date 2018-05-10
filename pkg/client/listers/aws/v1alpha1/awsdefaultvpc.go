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

// AwsDefaultVpcLister helps list AwsDefaultVpcs.
type AwsDefaultVpcLister interface {
	// List lists all AwsDefaultVpcs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultVpc, err error)
	// AwsDefaultVpcs returns an object that can list and get AwsDefaultVpcs.
	AwsDefaultVpcs(namespace string) AwsDefaultVpcNamespaceLister
	AwsDefaultVpcListerExpansion
}

// awsDefaultVpcLister implements the AwsDefaultVpcLister interface.
type awsDefaultVpcLister struct {
	indexer cache.Indexer
}

// NewAwsDefaultVpcLister returns a new AwsDefaultVpcLister.
func NewAwsDefaultVpcLister(indexer cache.Indexer) AwsDefaultVpcLister {
	return &awsDefaultVpcLister{indexer: indexer}
}

// List lists all AwsDefaultVpcs in the indexer.
func (s *awsDefaultVpcLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultVpc, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDefaultVpc))
	})
	return ret, err
}

// AwsDefaultVpcs returns an object that can list and get AwsDefaultVpcs.
func (s *awsDefaultVpcLister) AwsDefaultVpcs(namespace string) AwsDefaultVpcNamespaceLister {
	return awsDefaultVpcNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDefaultVpcNamespaceLister helps list and get AwsDefaultVpcs.
type AwsDefaultVpcNamespaceLister interface {
	// List lists all AwsDefaultVpcs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultVpc, err error)
	// Get retrieves the AwsDefaultVpc from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsDefaultVpc, error)
	AwsDefaultVpcNamespaceListerExpansion
}

// awsDefaultVpcNamespaceLister implements the AwsDefaultVpcNamespaceLister
// interface.
type awsDefaultVpcNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDefaultVpcs in the indexer for a given namespace.
func (s awsDefaultVpcNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultVpc, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDefaultVpc))
	})
	return ret, err
}

// Get retrieves the AwsDefaultVpc from the indexer for a given namespace and name.
func (s awsDefaultVpcNamespaceLister) Get(name string) (*v1alpha1.AwsDefaultVpc, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsdefaultvpc"), name)
	}
	return obj.(*v1alpha1.AwsDefaultVpc), nil
}
