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

// AwsIamInstanceProfileLister helps list AwsIamInstanceProfiles.
type AwsIamInstanceProfileLister interface {
	// List lists all AwsIamInstanceProfiles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamInstanceProfile, err error)
	// AwsIamInstanceProfiles returns an object that can list and get AwsIamInstanceProfiles.
	AwsIamInstanceProfiles(namespace string) AwsIamInstanceProfileNamespaceLister
	AwsIamInstanceProfileListerExpansion
}

// awsIamInstanceProfileLister implements the AwsIamInstanceProfileLister interface.
type awsIamInstanceProfileLister struct {
	indexer cache.Indexer
}

// NewAwsIamInstanceProfileLister returns a new AwsIamInstanceProfileLister.
func NewAwsIamInstanceProfileLister(indexer cache.Indexer) AwsIamInstanceProfileLister {
	return &awsIamInstanceProfileLister{indexer: indexer}
}

// List lists all AwsIamInstanceProfiles in the indexer.
func (s *awsIamInstanceProfileLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamInstanceProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamInstanceProfile))
	})
	return ret, err
}

// AwsIamInstanceProfiles returns an object that can list and get AwsIamInstanceProfiles.
func (s *awsIamInstanceProfileLister) AwsIamInstanceProfiles(namespace string) AwsIamInstanceProfileNamespaceLister {
	return awsIamInstanceProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamInstanceProfileNamespaceLister helps list and get AwsIamInstanceProfiles.
type AwsIamInstanceProfileNamespaceLister interface {
	// List lists all AwsIamInstanceProfiles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamInstanceProfile, err error)
	// Get retrieves the AwsIamInstanceProfile from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamInstanceProfile, error)
	AwsIamInstanceProfileNamespaceListerExpansion
}

// awsIamInstanceProfileNamespaceLister implements the AwsIamInstanceProfileNamespaceLister
// interface.
type awsIamInstanceProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamInstanceProfiles in the indexer for a given namespace.
func (s awsIamInstanceProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamInstanceProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamInstanceProfile))
	})
	return ret, err
}

// Get retrieves the AwsIamInstanceProfile from the indexer for a given namespace and name.
func (s awsIamInstanceProfileNamespaceLister) Get(name string) (*v1alpha1.AwsIamInstanceProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiaminstanceprofile"), name)
	}
	return obj.(*v1alpha1.AwsIamInstanceProfile), nil
}
