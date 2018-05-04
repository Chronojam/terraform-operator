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

// AwsElastictranscoderPresetLister helps list AwsElastictranscoderPresets.
type AwsElastictranscoderPresetLister interface {
	// List lists all AwsElastictranscoderPresets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElastictranscoderPreset, err error)
	// AwsElastictranscoderPresets returns an object that can list and get AwsElastictranscoderPresets.
	AwsElastictranscoderPresets(namespace string) AwsElastictranscoderPresetNamespaceLister
	AwsElastictranscoderPresetListerExpansion
}

// awsElastictranscoderPresetLister implements the AwsElastictranscoderPresetLister interface.
type awsElastictranscoderPresetLister struct {
	indexer cache.Indexer
}

// NewAwsElastictranscoderPresetLister returns a new AwsElastictranscoderPresetLister.
func NewAwsElastictranscoderPresetLister(indexer cache.Indexer) AwsElastictranscoderPresetLister {
	return &awsElastictranscoderPresetLister{indexer: indexer}
}

// List lists all AwsElastictranscoderPresets in the indexer.
func (s *awsElastictranscoderPresetLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElastictranscoderPreset, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElastictranscoderPreset))
	})
	return ret, err
}

// AwsElastictranscoderPresets returns an object that can list and get AwsElastictranscoderPresets.
func (s *awsElastictranscoderPresetLister) AwsElastictranscoderPresets(namespace string) AwsElastictranscoderPresetNamespaceLister {
	return awsElastictranscoderPresetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsElastictranscoderPresetNamespaceLister helps list and get AwsElastictranscoderPresets.
type AwsElastictranscoderPresetNamespaceLister interface {
	// List lists all AwsElastictranscoderPresets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElastictranscoderPreset, err error)
	// Get retrieves the AwsElastictranscoderPreset from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsElastictranscoderPreset, error)
	AwsElastictranscoderPresetNamespaceListerExpansion
}

// awsElastictranscoderPresetNamespaceLister implements the AwsElastictranscoderPresetNamespaceLister
// interface.
type awsElastictranscoderPresetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsElastictranscoderPresets in the indexer for a given namespace.
func (s awsElastictranscoderPresetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElastictranscoderPreset, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElastictranscoderPreset))
	})
	return ret, err
}

// Get retrieves the AwsElastictranscoderPreset from the indexer for a given namespace and name.
func (s awsElastictranscoderPresetNamespaceLister) Get(name string) (*v1alpha1.AwsElastictranscoderPreset, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awselastictranscoderpreset"), name)
	}
	return obj.(*v1alpha1.AwsElastictranscoderPreset), nil
}
