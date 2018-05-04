/*
Copyright The Kubernetes Authors.

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

// AwsLightsailKeyPairLister helps list AwsLightsailKeyPairs.
type AwsLightsailKeyPairLister interface {
	// List lists all AwsLightsailKeyPairs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLightsailKeyPair, err error)
	// AwsLightsailKeyPairs returns an object that can list and get AwsLightsailKeyPairs.
	AwsLightsailKeyPairs(namespace string) AwsLightsailKeyPairNamespaceLister
	AwsLightsailKeyPairListerExpansion
}

// awsLightsailKeyPairLister implements the AwsLightsailKeyPairLister interface.
type awsLightsailKeyPairLister struct {
	indexer cache.Indexer
}

// NewAwsLightsailKeyPairLister returns a new AwsLightsailKeyPairLister.
func NewAwsLightsailKeyPairLister(indexer cache.Indexer) AwsLightsailKeyPairLister {
	return &awsLightsailKeyPairLister{indexer: indexer}
}

// List lists all AwsLightsailKeyPairs in the indexer.
func (s *awsLightsailKeyPairLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLightsailKeyPair, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLightsailKeyPair))
	})
	return ret, err
}

// AwsLightsailKeyPairs returns an object that can list and get AwsLightsailKeyPairs.
func (s *awsLightsailKeyPairLister) AwsLightsailKeyPairs(namespace string) AwsLightsailKeyPairNamespaceLister {
	return awsLightsailKeyPairNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLightsailKeyPairNamespaceLister helps list and get AwsLightsailKeyPairs.
type AwsLightsailKeyPairNamespaceLister interface {
	// List lists all AwsLightsailKeyPairs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLightsailKeyPair, err error)
	// Get retrieves the AwsLightsailKeyPair from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsLightsailKeyPair, error)
	AwsLightsailKeyPairNamespaceListerExpansion
}

// awsLightsailKeyPairNamespaceLister implements the AwsLightsailKeyPairNamespaceLister
// interface.
type awsLightsailKeyPairNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLightsailKeyPairs in the indexer for a given namespace.
func (s awsLightsailKeyPairNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLightsailKeyPair, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLightsailKeyPair))
	})
	return ret, err
}

// Get retrieves the AwsLightsailKeyPair from the indexer for a given namespace and name.
func (s awsLightsailKeyPairNamespaceLister) Get(name string) (*v1alpha1.AwsLightsailKeyPair, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awslightsailkeypair"), name)
	}
	return obj.(*v1alpha1.AwsLightsailKeyPair), nil
}
