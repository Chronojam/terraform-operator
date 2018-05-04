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

// AwsOpsworksJavaAppLayerLister helps list AwsOpsworksJavaAppLayers.
type AwsOpsworksJavaAppLayerLister interface {
	// List lists all AwsOpsworksJavaAppLayers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksJavaAppLayer, err error)
	// AwsOpsworksJavaAppLayers returns an object that can list and get AwsOpsworksJavaAppLayers.
	AwsOpsworksJavaAppLayers(namespace string) AwsOpsworksJavaAppLayerNamespaceLister
	AwsOpsworksJavaAppLayerListerExpansion
}

// awsOpsworksJavaAppLayerLister implements the AwsOpsworksJavaAppLayerLister interface.
type awsOpsworksJavaAppLayerLister struct {
	indexer cache.Indexer
}

// NewAwsOpsworksJavaAppLayerLister returns a new AwsOpsworksJavaAppLayerLister.
func NewAwsOpsworksJavaAppLayerLister(indexer cache.Indexer) AwsOpsworksJavaAppLayerLister {
	return &awsOpsworksJavaAppLayerLister{indexer: indexer}
}

// List lists all AwsOpsworksJavaAppLayers in the indexer.
func (s *awsOpsworksJavaAppLayerLister) List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksJavaAppLayer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsOpsworksJavaAppLayer))
	})
	return ret, err
}

// AwsOpsworksJavaAppLayers returns an object that can list and get AwsOpsworksJavaAppLayers.
func (s *awsOpsworksJavaAppLayerLister) AwsOpsworksJavaAppLayers(namespace string) AwsOpsworksJavaAppLayerNamespaceLister {
	return awsOpsworksJavaAppLayerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsOpsworksJavaAppLayerNamespaceLister helps list and get AwsOpsworksJavaAppLayers.
type AwsOpsworksJavaAppLayerNamespaceLister interface {
	// List lists all AwsOpsworksJavaAppLayers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksJavaAppLayer, err error)
	// Get retrieves the AwsOpsworksJavaAppLayer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsOpsworksJavaAppLayer, error)
	AwsOpsworksJavaAppLayerNamespaceListerExpansion
}

// awsOpsworksJavaAppLayerNamespaceLister implements the AwsOpsworksJavaAppLayerNamespaceLister
// interface.
type awsOpsworksJavaAppLayerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsOpsworksJavaAppLayers in the indexer for a given namespace.
func (s awsOpsworksJavaAppLayerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksJavaAppLayer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsOpsworksJavaAppLayer))
	})
	return ret, err
}

// Get retrieves the AwsOpsworksJavaAppLayer from the indexer for a given namespace and name.
func (s awsOpsworksJavaAppLayerNamespaceLister) Get(name string) (*v1alpha1.AwsOpsworksJavaAppLayer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsopsworksjavaapplayer"), name)
	}
	return obj.(*v1alpha1.AwsOpsworksJavaAppLayer), nil
}
