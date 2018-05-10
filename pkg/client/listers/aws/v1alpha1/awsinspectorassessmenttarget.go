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

// AwsInspectorAssessmentTargetLister helps list AwsInspectorAssessmentTargets.
type AwsInspectorAssessmentTargetLister interface {
	// List lists all AwsInspectorAssessmentTargets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorAssessmentTarget, err error)
	// AwsInspectorAssessmentTargets returns an object that can list and get AwsInspectorAssessmentTargets.
	AwsInspectorAssessmentTargets(namespace string) AwsInspectorAssessmentTargetNamespaceLister
	AwsInspectorAssessmentTargetListerExpansion
}

// awsInspectorAssessmentTargetLister implements the AwsInspectorAssessmentTargetLister interface.
type awsInspectorAssessmentTargetLister struct {
	indexer cache.Indexer
}

// NewAwsInspectorAssessmentTargetLister returns a new AwsInspectorAssessmentTargetLister.
func NewAwsInspectorAssessmentTargetLister(indexer cache.Indexer) AwsInspectorAssessmentTargetLister {
	return &awsInspectorAssessmentTargetLister{indexer: indexer}
}

// List lists all AwsInspectorAssessmentTargets in the indexer.
func (s *awsInspectorAssessmentTargetLister) List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorAssessmentTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsInspectorAssessmentTarget))
	})
	return ret, err
}

// AwsInspectorAssessmentTargets returns an object that can list and get AwsInspectorAssessmentTargets.
func (s *awsInspectorAssessmentTargetLister) AwsInspectorAssessmentTargets(namespace string) AwsInspectorAssessmentTargetNamespaceLister {
	return awsInspectorAssessmentTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsInspectorAssessmentTargetNamespaceLister helps list and get AwsInspectorAssessmentTargets.
type AwsInspectorAssessmentTargetNamespaceLister interface {
	// List lists all AwsInspectorAssessmentTargets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorAssessmentTarget, err error)
	// Get retrieves the AwsInspectorAssessmentTarget from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsInspectorAssessmentTarget, error)
	AwsInspectorAssessmentTargetNamespaceListerExpansion
}

// awsInspectorAssessmentTargetNamespaceLister implements the AwsInspectorAssessmentTargetNamespaceLister
// interface.
type awsInspectorAssessmentTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsInspectorAssessmentTargets in the indexer for a given namespace.
func (s awsInspectorAssessmentTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsInspectorAssessmentTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsInspectorAssessmentTarget))
	})
	return ret, err
}

// Get retrieves the AwsInspectorAssessmentTarget from the indexer for a given namespace and name.
func (s awsInspectorAssessmentTargetNamespaceLister) Get(name string) (*v1alpha1.AwsInspectorAssessmentTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsinspectorassessmenttarget"), name)
	}
	return obj.(*v1alpha1.AwsInspectorAssessmentTarget), nil
}
