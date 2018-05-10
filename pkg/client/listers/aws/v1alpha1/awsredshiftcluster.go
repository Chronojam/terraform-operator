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

// AwsRedshiftClusterLister helps list AwsRedshiftClusters.
type AwsRedshiftClusterLister interface {
	// List lists all AwsRedshiftClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftCluster, err error)
	// AwsRedshiftClusters returns an object that can list and get AwsRedshiftClusters.
	AwsRedshiftClusters(namespace string) AwsRedshiftClusterNamespaceLister
	AwsRedshiftClusterListerExpansion
}

// awsRedshiftClusterLister implements the AwsRedshiftClusterLister interface.
type awsRedshiftClusterLister struct {
	indexer cache.Indexer
}

// NewAwsRedshiftClusterLister returns a new AwsRedshiftClusterLister.
func NewAwsRedshiftClusterLister(indexer cache.Indexer) AwsRedshiftClusterLister {
	return &awsRedshiftClusterLister{indexer: indexer}
}

// List lists all AwsRedshiftClusters in the indexer.
func (s *awsRedshiftClusterLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRedshiftCluster))
	})
	return ret, err
}

// AwsRedshiftClusters returns an object that can list and get AwsRedshiftClusters.
func (s *awsRedshiftClusterLister) AwsRedshiftClusters(namespace string) AwsRedshiftClusterNamespaceLister {
	return awsRedshiftClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsRedshiftClusterNamespaceLister helps list and get AwsRedshiftClusters.
type AwsRedshiftClusterNamespaceLister interface {
	// List lists all AwsRedshiftClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftCluster, err error)
	// Get retrieves the AwsRedshiftCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsRedshiftCluster, error)
	AwsRedshiftClusterNamespaceListerExpansion
}

// awsRedshiftClusterNamespaceLister implements the AwsRedshiftClusterNamespaceLister
// interface.
type awsRedshiftClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsRedshiftClusters in the indexer for a given namespace.
func (s awsRedshiftClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRedshiftCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRedshiftCluster))
	})
	return ret, err
}

// Get retrieves the AwsRedshiftCluster from the indexer for a given namespace and name.
func (s awsRedshiftClusterNamespaceLister) Get(name string) (*v1alpha1.AwsRedshiftCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsredshiftcluster"), name)
	}
	return obj.(*v1alpha1.AwsRedshiftCluster), nil
}
