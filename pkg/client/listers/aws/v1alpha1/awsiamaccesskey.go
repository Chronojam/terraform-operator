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

// AwsIamAccessKeyLister helps list AwsIamAccessKeys.
type AwsIamAccessKeyLister interface {
	// List lists all AwsIamAccessKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccessKey, err error)
	// AwsIamAccessKeys returns an object that can list and get AwsIamAccessKeys.
	AwsIamAccessKeys(namespace string) AwsIamAccessKeyNamespaceLister
	AwsIamAccessKeyListerExpansion
}

// awsIamAccessKeyLister implements the AwsIamAccessKeyLister interface.
type awsIamAccessKeyLister struct {
	indexer cache.Indexer
}

// NewAwsIamAccessKeyLister returns a new AwsIamAccessKeyLister.
func NewAwsIamAccessKeyLister(indexer cache.Indexer) AwsIamAccessKeyLister {
	return &awsIamAccessKeyLister{indexer: indexer}
}

// List lists all AwsIamAccessKeys in the indexer.
func (s *awsIamAccessKeyLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccessKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamAccessKey))
	})
	return ret, err
}

// AwsIamAccessKeys returns an object that can list and get AwsIamAccessKeys.
func (s *awsIamAccessKeyLister) AwsIamAccessKeys(namespace string) AwsIamAccessKeyNamespaceLister {
	return awsIamAccessKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamAccessKeyNamespaceLister helps list and get AwsIamAccessKeys.
type AwsIamAccessKeyNamespaceLister interface {
	// List lists all AwsIamAccessKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccessKey, err error)
	// Get retrieves the AwsIamAccessKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamAccessKey, error)
	AwsIamAccessKeyNamespaceListerExpansion
}

// awsIamAccessKeyNamespaceLister implements the AwsIamAccessKeyNamespaceLister
// interface.
type awsIamAccessKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamAccessKeys in the indexer for a given namespace.
func (s awsIamAccessKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamAccessKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamAccessKey))
	})
	return ret, err
}

// Get retrieves the AwsIamAccessKey from the indexer for a given namespace and name.
func (s awsIamAccessKeyNamespaceLister) Get(name string) (*v1alpha1.AwsIamAccessKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiamaccesskey"), name)
	}
	return obj.(*v1alpha1.AwsIamAccessKey), nil
}
