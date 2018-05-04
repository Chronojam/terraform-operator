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

// AwsEcsClusterLister helps list AwsEcsClusters.
type AwsEcsClusterLister interface {
	// List lists all AwsEcsClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsEcsCluster, err error)
	// AwsEcsClusters returns an object that can list and get AwsEcsClusters.
	AwsEcsClusters(namespace string) AwsEcsClusterNamespaceLister
	AwsEcsClusterListerExpansion
}

// awsEcsClusterLister implements the AwsEcsClusterLister interface.
type awsEcsClusterLister struct {
	indexer cache.Indexer
}

// NewAwsEcsClusterLister returns a new AwsEcsClusterLister.
func NewAwsEcsClusterLister(indexer cache.Indexer) AwsEcsClusterLister {
	return &awsEcsClusterLister{indexer: indexer}
}

// List lists all AwsEcsClusters in the indexer.
func (s *awsEcsClusterLister) List(selector labels.Selector) (ret []*v1alpha1.AwsEcsCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsEcsCluster))
	})
	return ret, err
}

// AwsEcsClusters returns an object that can list and get AwsEcsClusters.
func (s *awsEcsClusterLister) AwsEcsClusters(namespace string) AwsEcsClusterNamespaceLister {
	return awsEcsClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsEcsClusterNamespaceLister helps list and get AwsEcsClusters.
type AwsEcsClusterNamespaceLister interface {
	// List lists all AwsEcsClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsEcsCluster, err error)
	// Get retrieves the AwsEcsCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsEcsCluster, error)
	AwsEcsClusterNamespaceListerExpansion
}

// awsEcsClusterNamespaceLister implements the AwsEcsClusterNamespaceLister
// interface.
type awsEcsClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsEcsClusters in the indexer for a given namespace.
func (s awsEcsClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsEcsCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsEcsCluster))
	})
	return ret, err
}

// Get retrieves the AwsEcsCluster from the indexer for a given namespace and name.
func (s awsEcsClusterNamespaceLister) Get(name string) (*v1alpha1.AwsEcsCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsecscluster"), name)
	}
	return obj.(*v1alpha1.AwsEcsCluster), nil
}
