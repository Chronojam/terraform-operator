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

// AwsInspectorResourceGroupLister helps list AwsInspectorResourceGroups.
type AwsInspectorResourceGroupLister interface {
	// List lists all AwsInspectorResourceGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorResourceGroup, err error)
	// AwsInspectorResourceGroups returns an object that can list and get AwsInspectorResourceGroups.
	AwsInspectorResourceGroups(namespace string) AwsInspectorResourceGroupNamespaceLister
	AwsInspectorResourceGroupListerExpansion
}

// awsInspectorResourceGroupLister implements the AwsInspectorResourceGroupLister interface.
type awsInspectorResourceGroupLister struct {
	indexer cache.Indexer
}

// NewAwsInspectorResourceGroupLister returns a new AwsInspectorResourceGroupLister.
func NewAwsInspectorResourceGroupLister(indexer cache.Indexer) AwsInspectorResourceGroupLister {
	return &awsInspectorResourceGroupLister{indexer: indexer}
}

// List lists all AwsInspectorResourceGroups in the indexer.
func (s *awsInspectorResourceGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorResourceGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsInspectorResourceGroup))
	})
	return ret, err
}

// AwsInspectorResourceGroups returns an object that can list and get AwsInspectorResourceGroups.
func (s *awsInspectorResourceGroupLister) AwsInspectorResourceGroups(namespace string) AwsInspectorResourceGroupNamespaceLister {
	return awsInspectorResourceGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsInspectorResourceGroupNamespaceLister helps list and get AwsInspectorResourceGroups.
type AwsInspectorResourceGroupNamespaceLister interface {
	// List lists all AwsInspectorResourceGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorResourceGroup, err error)
	// Get retrieves the AwsInspectorResourceGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsInspectorResourceGroup, error)
	AwsInspectorResourceGroupNamespaceListerExpansion
}

// awsInspectorResourceGroupNamespaceLister implements the AwsInspectorResourceGroupNamespaceLister
// interface.
type awsInspectorResourceGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsInspectorResourceGroups in the indexer for a given namespace.
func (s awsInspectorResourceGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorResourceGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsInspectorResourceGroup))
	})
	return ret, err
}

// Get retrieves the AwsInspectorResourceGroup from the indexer for a given namespace and name.
func (s awsInspectorResourceGroupNamespaceLister) Get(name string) (*v1alpha1.AwsInspectorResourceGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsinspectorresourcegroup"), name)
	}
	return obj.(*v1alpha1.AwsInspectorResourceGroup), nil
}
