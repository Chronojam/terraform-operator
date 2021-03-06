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

// AwsAmiCopyLister helps list AwsAmiCopies.
type AwsAmiCopyLister interface {
	// List lists all AwsAmiCopies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAmiCopy, err error)
	// AwsAmiCopies returns an object that can list and get AwsAmiCopies.
	AwsAmiCopies(namespace string) AwsAmiCopyNamespaceLister
	AwsAmiCopyListerExpansion
}

// awsAmiCopyLister implements the AwsAmiCopyLister interface.
type awsAmiCopyLister struct {
	indexer cache.Indexer
}

// NewAwsAmiCopyLister returns a new AwsAmiCopyLister.
func NewAwsAmiCopyLister(indexer cache.Indexer) AwsAmiCopyLister {
	return &awsAmiCopyLister{indexer: indexer}
}

// List lists all AwsAmiCopies in the indexer.
func (s *awsAmiCopyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAmiCopy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAmiCopy))
	})
	return ret, err
}

// AwsAmiCopies returns an object that can list and get AwsAmiCopies.
func (s *awsAmiCopyLister) AwsAmiCopies(namespace string) AwsAmiCopyNamespaceLister {
	return awsAmiCopyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsAmiCopyNamespaceLister helps list and get AwsAmiCopies.
type AwsAmiCopyNamespaceLister interface {
	// List lists all AwsAmiCopies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAmiCopy, err error)
	// Get retrieves the AwsAmiCopy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsAmiCopy, error)
	AwsAmiCopyNamespaceListerExpansion
}

// awsAmiCopyNamespaceLister implements the AwsAmiCopyNamespaceLister
// interface.
type awsAmiCopyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsAmiCopies in the indexer for a given namespace.
func (s awsAmiCopyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAmiCopy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAmiCopy))
	})
	return ret, err
}

// Get retrieves the AwsAmiCopy from the indexer for a given namespace and name.
func (s awsAmiCopyNamespaceLister) Get(name string) (*v1alpha1.AwsAmiCopy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsamicopy"), name)
	}
	return obj.(*v1alpha1.AwsAmiCopy), nil
}
