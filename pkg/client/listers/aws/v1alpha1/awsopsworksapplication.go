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

// AwsOpsworksApplicationLister helps list AwsOpsworksApplications.
type AwsOpsworksApplicationLister interface {
	// List lists all AwsOpsworksApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksApplication, err error)
	// AwsOpsworksApplications returns an object that can list and get AwsOpsworksApplications.
	AwsOpsworksApplications(namespace string) AwsOpsworksApplicationNamespaceLister
	AwsOpsworksApplicationListerExpansion
}

// awsOpsworksApplicationLister implements the AwsOpsworksApplicationLister interface.
type awsOpsworksApplicationLister struct {
	indexer cache.Indexer
}

// NewAwsOpsworksApplicationLister returns a new AwsOpsworksApplicationLister.
func NewAwsOpsworksApplicationLister(indexer cache.Indexer) AwsOpsworksApplicationLister {
	return &awsOpsworksApplicationLister{indexer: indexer}
}

// List lists all AwsOpsworksApplications in the indexer.
func (s *awsOpsworksApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsOpsworksApplication))
	})
	return ret, err
}

// AwsOpsworksApplications returns an object that can list and get AwsOpsworksApplications.
func (s *awsOpsworksApplicationLister) AwsOpsworksApplications(namespace string) AwsOpsworksApplicationNamespaceLister {
	return awsOpsworksApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsOpsworksApplicationNamespaceLister helps list and get AwsOpsworksApplications.
type AwsOpsworksApplicationNamespaceLister interface {
	// List lists all AwsOpsworksApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksApplication, err error)
	// Get retrieves the AwsOpsworksApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsOpsworksApplication, error)
	AwsOpsworksApplicationNamespaceListerExpansion
}

// awsOpsworksApplicationNamespaceLister implements the AwsOpsworksApplicationNamespaceLister
// interface.
type awsOpsworksApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsOpsworksApplications in the indexer for a given namespace.
func (s awsOpsworksApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsOpsworksApplication))
	})
	return ret, err
}

// Get retrieves the AwsOpsworksApplication from the indexer for a given namespace and name.
func (s awsOpsworksApplicationNamespaceLister) Get(name string) (*v1alpha1.AwsOpsworksApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsopsworksapplication"), name)
	}
	return obj.(*v1alpha1.AwsOpsworksApplication), nil
}
