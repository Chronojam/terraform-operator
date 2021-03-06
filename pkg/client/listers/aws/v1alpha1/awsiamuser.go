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

// AwsIamUserLister helps list AwsIamUsers.
type AwsIamUserLister interface {
	// List lists all AwsIamUsers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamUser, err error)
	// AwsIamUsers returns an object that can list and get AwsIamUsers.
	AwsIamUsers(namespace string) AwsIamUserNamespaceLister
	AwsIamUserListerExpansion
}

// awsIamUserLister implements the AwsIamUserLister interface.
type awsIamUserLister struct {
	indexer cache.Indexer
}

// NewAwsIamUserLister returns a new AwsIamUserLister.
func NewAwsIamUserLister(indexer cache.Indexer) AwsIamUserLister {
	return &awsIamUserLister{indexer: indexer}
}

// List lists all AwsIamUsers in the indexer.
func (s *awsIamUserLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamUser))
	})
	return ret, err
}

// AwsIamUsers returns an object that can list and get AwsIamUsers.
func (s *awsIamUserLister) AwsIamUsers(namespace string) AwsIamUserNamespaceLister {
	return awsIamUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamUserNamespaceLister helps list and get AwsIamUsers.
type AwsIamUserNamespaceLister interface {
	// List lists all AwsIamUsers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamUser, err error)
	// Get retrieves the AwsIamUser from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamUser, error)
	AwsIamUserNamespaceListerExpansion
}

// awsIamUserNamespaceLister implements the AwsIamUserNamespaceLister
// interface.
type awsIamUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamUsers in the indexer for a given namespace.
func (s awsIamUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamUser))
	})
	return ret, err
}

// Get retrieves the AwsIamUser from the indexer for a given namespace and name.
func (s awsIamUserNamespaceLister) Get(name string) (*v1alpha1.AwsIamUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiamuser"), name)
	}
	return obj.(*v1alpha1.AwsIamUser), nil
}
