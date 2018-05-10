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

// AwsKinesisStreamLister helps list AwsKinesisStreams.
type AwsKinesisStreamLister interface {
	// List lists all AwsKinesisStreams in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsKinesisStream, err error)
	// AwsKinesisStreams returns an object that can list and get AwsKinesisStreams.
	AwsKinesisStreams(namespace string) AwsKinesisStreamNamespaceLister
	AwsKinesisStreamListerExpansion
}

// awsKinesisStreamLister implements the AwsKinesisStreamLister interface.
type awsKinesisStreamLister struct {
	indexer cache.Indexer
}

// NewAwsKinesisStreamLister returns a new AwsKinesisStreamLister.
func NewAwsKinesisStreamLister(indexer cache.Indexer) AwsKinesisStreamLister {
	return &awsKinesisStreamLister{indexer: indexer}
}

// List lists all AwsKinesisStreams in the indexer.
func (s *awsKinesisStreamLister) List(selector labels.Selector) (ret []*v1alpha1.AwsKinesisStream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsKinesisStream))
	})
	return ret, err
}

// AwsKinesisStreams returns an object that can list and get AwsKinesisStreams.
func (s *awsKinesisStreamLister) AwsKinesisStreams(namespace string) AwsKinesisStreamNamespaceLister {
	return awsKinesisStreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsKinesisStreamNamespaceLister helps list and get AwsKinesisStreams.
type AwsKinesisStreamNamespaceLister interface {
	// List lists all AwsKinesisStreams in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsKinesisStream, err error)
	// Get retrieves the AwsKinesisStream from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsKinesisStream, error)
	AwsKinesisStreamNamespaceListerExpansion
}

// awsKinesisStreamNamespaceLister implements the AwsKinesisStreamNamespaceLister
// interface.
type awsKinesisStreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsKinesisStreams in the indexer for a given namespace.
func (s awsKinesisStreamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsKinesisStream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsKinesisStream))
	})
	return ret, err
}

// Get retrieves the AwsKinesisStream from the indexer for a given namespace and name.
func (s awsKinesisStreamNamespaceLister) Get(name string) (*v1alpha1.AwsKinesisStream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awskinesisstream"), name)
	}
	return obj.(*v1alpha1.AwsKinesisStream), nil
}
