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

// AwsDevicefarmProjectLister helps list AwsDevicefarmProjects.
type AwsDevicefarmProjectLister interface {
	// List lists all AwsDevicefarmProjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDevicefarmProject, err error)
	// AwsDevicefarmProjects returns an object that can list and get AwsDevicefarmProjects.
	AwsDevicefarmProjects(namespace string) AwsDevicefarmProjectNamespaceLister
	AwsDevicefarmProjectListerExpansion
}

// awsDevicefarmProjectLister implements the AwsDevicefarmProjectLister interface.
type awsDevicefarmProjectLister struct {
	indexer cache.Indexer
}

// NewAwsDevicefarmProjectLister returns a new AwsDevicefarmProjectLister.
func NewAwsDevicefarmProjectLister(indexer cache.Indexer) AwsDevicefarmProjectLister {
	return &awsDevicefarmProjectLister{indexer: indexer}
}

// List lists all AwsDevicefarmProjects in the indexer.
func (s *awsDevicefarmProjectLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDevicefarmProject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDevicefarmProject))
	})
	return ret, err
}

// AwsDevicefarmProjects returns an object that can list and get AwsDevicefarmProjects.
func (s *awsDevicefarmProjectLister) AwsDevicefarmProjects(namespace string) AwsDevicefarmProjectNamespaceLister {
	return awsDevicefarmProjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDevicefarmProjectNamespaceLister helps list and get AwsDevicefarmProjects.
type AwsDevicefarmProjectNamespaceLister interface {
	// List lists all AwsDevicefarmProjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDevicefarmProject, err error)
	// Get retrieves the AwsDevicefarmProject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsDevicefarmProject, error)
	AwsDevicefarmProjectNamespaceListerExpansion
}

// awsDevicefarmProjectNamespaceLister implements the AwsDevicefarmProjectNamespaceLister
// interface.
type awsDevicefarmProjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDevicefarmProjects in the indexer for a given namespace.
func (s awsDevicefarmProjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDevicefarmProject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDevicefarmProject))
	})
	return ret, err
}

// Get retrieves the AwsDevicefarmProject from the indexer for a given namespace and name.
func (s awsDevicefarmProjectNamespaceLister) Get(name string) (*v1alpha1.AwsDevicefarmProject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsdevicefarmproject"), name)
	}
	return obj.(*v1alpha1.AwsDevicefarmProject), nil
}
