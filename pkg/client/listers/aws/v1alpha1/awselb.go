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

// AwsElbLister helps list AwsElbs.
type AwsElbLister interface {
	// List lists all AwsElbs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElb, err error)
	// AwsElbs returns an object that can list and get AwsElbs.
	AwsElbs(namespace string) AwsElbNamespaceLister
	AwsElbListerExpansion
}

// awsElbLister implements the AwsElbLister interface.
type awsElbLister struct {
	indexer cache.Indexer
}

// NewAwsElbLister returns a new AwsElbLister.
func NewAwsElbLister(indexer cache.Indexer) AwsElbLister {
	return &awsElbLister{indexer: indexer}
}

// List lists all AwsElbs in the indexer.
func (s *awsElbLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElb, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElb))
	})
	return ret, err
}

// AwsElbs returns an object that can list and get AwsElbs.
func (s *awsElbLister) AwsElbs(namespace string) AwsElbNamespaceLister {
	return awsElbNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsElbNamespaceLister helps list and get AwsElbs.
type AwsElbNamespaceLister interface {
	// List lists all AwsElbs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElb, err error)
	// Get retrieves the AwsElb from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsElb, error)
	AwsElbNamespaceListerExpansion
}

// awsElbNamespaceLister implements the AwsElbNamespaceLister
// interface.
type awsElbNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsElbs in the indexer for a given namespace.
func (s awsElbNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElb, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElb))
	})
	return ret, err
}

// Get retrieves the AwsElb from the indexer for a given namespace and name.
func (s awsElbNamespaceLister) Get(name string) (*v1alpha1.AwsElb, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awselb"), name)
	}
	return obj.(*v1alpha1.AwsElb), nil
}
