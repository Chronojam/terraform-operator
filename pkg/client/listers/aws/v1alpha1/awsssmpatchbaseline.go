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

// AwsSsmPatchBaselineLister helps list AwsSsmPatchBaselines.
type AwsSsmPatchBaselineLister interface {
	// List lists all AwsSsmPatchBaselines in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSsmPatchBaseline, err error)
	// AwsSsmPatchBaselines returns an object that can list and get AwsSsmPatchBaselines.
	AwsSsmPatchBaselines(namespace string) AwsSsmPatchBaselineNamespaceLister
	AwsSsmPatchBaselineListerExpansion
}

// awsSsmPatchBaselineLister implements the AwsSsmPatchBaselineLister interface.
type awsSsmPatchBaselineLister struct {
	indexer cache.Indexer
}

// NewAwsSsmPatchBaselineLister returns a new AwsSsmPatchBaselineLister.
func NewAwsSsmPatchBaselineLister(indexer cache.Indexer) AwsSsmPatchBaselineLister {
	return &awsSsmPatchBaselineLister{indexer: indexer}
}

// List lists all AwsSsmPatchBaselines in the indexer.
func (s *awsSsmPatchBaselineLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSsmPatchBaseline, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSsmPatchBaseline))
	})
	return ret, err
}

// AwsSsmPatchBaselines returns an object that can list and get AwsSsmPatchBaselines.
func (s *awsSsmPatchBaselineLister) AwsSsmPatchBaselines(namespace string) AwsSsmPatchBaselineNamespaceLister {
	return awsSsmPatchBaselineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSsmPatchBaselineNamespaceLister helps list and get AwsSsmPatchBaselines.
type AwsSsmPatchBaselineNamespaceLister interface {
	// List lists all AwsSsmPatchBaselines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSsmPatchBaseline, err error)
	// Get retrieves the AwsSsmPatchBaseline from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsSsmPatchBaseline, error)
	AwsSsmPatchBaselineNamespaceListerExpansion
}

// awsSsmPatchBaselineNamespaceLister implements the AwsSsmPatchBaselineNamespaceLister
// interface.
type awsSsmPatchBaselineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSsmPatchBaselines in the indexer for a given namespace.
func (s awsSsmPatchBaselineNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSsmPatchBaseline, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSsmPatchBaseline))
	})
	return ret, err
}

// Get retrieves the AwsSsmPatchBaseline from the indexer for a given namespace and name.
func (s awsSsmPatchBaselineNamespaceLister) Get(name string) (*v1alpha1.AwsSsmPatchBaseline, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsssmpatchbaseline"), name)
	}
	return obj.(*v1alpha1.AwsSsmPatchBaseline), nil
}
