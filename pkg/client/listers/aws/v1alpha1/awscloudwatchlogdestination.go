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

// AwsCloudwatchLogDestinationLister helps list AwsCloudwatchLogDestinations.
type AwsCloudwatchLogDestinationLister interface {
	// List lists all AwsCloudwatchLogDestinations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCloudwatchLogDestination, err error)
	// AwsCloudwatchLogDestinations returns an object that can list and get AwsCloudwatchLogDestinations.
	AwsCloudwatchLogDestinations(namespace string) AwsCloudwatchLogDestinationNamespaceLister
	AwsCloudwatchLogDestinationListerExpansion
}

// awsCloudwatchLogDestinationLister implements the AwsCloudwatchLogDestinationLister interface.
type awsCloudwatchLogDestinationLister struct {
	indexer cache.Indexer
}

// NewAwsCloudwatchLogDestinationLister returns a new AwsCloudwatchLogDestinationLister.
func NewAwsCloudwatchLogDestinationLister(indexer cache.Indexer) AwsCloudwatchLogDestinationLister {
	return &awsCloudwatchLogDestinationLister{indexer: indexer}
}

// List lists all AwsCloudwatchLogDestinations in the indexer.
func (s *awsCloudwatchLogDestinationLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCloudwatchLogDestination, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCloudwatchLogDestination))
	})
	return ret, err
}

// AwsCloudwatchLogDestinations returns an object that can list and get AwsCloudwatchLogDestinations.
func (s *awsCloudwatchLogDestinationLister) AwsCloudwatchLogDestinations(namespace string) AwsCloudwatchLogDestinationNamespaceLister {
	return awsCloudwatchLogDestinationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCloudwatchLogDestinationNamespaceLister helps list and get AwsCloudwatchLogDestinations.
type AwsCloudwatchLogDestinationNamespaceLister interface {
	// List lists all AwsCloudwatchLogDestinations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCloudwatchLogDestination, err error)
	// Get retrieves the AwsCloudwatchLogDestination from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsCloudwatchLogDestination, error)
	AwsCloudwatchLogDestinationNamespaceListerExpansion
}

// awsCloudwatchLogDestinationNamespaceLister implements the AwsCloudwatchLogDestinationNamespaceLister
// interface.
type awsCloudwatchLogDestinationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCloudwatchLogDestinations in the indexer for a given namespace.
func (s awsCloudwatchLogDestinationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCloudwatchLogDestination, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCloudwatchLogDestination))
	})
	return ret, err
}

// Get retrieves the AwsCloudwatchLogDestination from the indexer for a given namespace and name.
func (s awsCloudwatchLogDestinationNamespaceLister) Get(name string) (*v1alpha1.AwsCloudwatchLogDestination, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awscloudwatchlogdestination"), name)
	}
	return obj.(*v1alpha1.AwsCloudwatchLogDestination), nil
}
