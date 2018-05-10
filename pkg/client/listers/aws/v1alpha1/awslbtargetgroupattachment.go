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

// AwsLbTargetGroupAttachmentLister helps list AwsLbTargetGroupAttachments.
type AwsLbTargetGroupAttachmentLister interface {
	// List lists all AwsLbTargetGroupAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLbTargetGroupAttachment, err error)
	// AwsLbTargetGroupAttachments returns an object that can list and get AwsLbTargetGroupAttachments.
	AwsLbTargetGroupAttachments(namespace string) AwsLbTargetGroupAttachmentNamespaceLister
	AwsLbTargetGroupAttachmentListerExpansion
}

// awsLbTargetGroupAttachmentLister implements the AwsLbTargetGroupAttachmentLister interface.
type awsLbTargetGroupAttachmentLister struct {
	indexer cache.Indexer
}

// NewAwsLbTargetGroupAttachmentLister returns a new AwsLbTargetGroupAttachmentLister.
func NewAwsLbTargetGroupAttachmentLister(indexer cache.Indexer) AwsLbTargetGroupAttachmentLister {
	return &awsLbTargetGroupAttachmentLister{indexer: indexer}
}

// List lists all AwsLbTargetGroupAttachments in the indexer.
func (s *awsLbTargetGroupAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLbTargetGroupAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLbTargetGroupAttachment))
	})
	return ret, err
}

// AwsLbTargetGroupAttachments returns an object that can list and get AwsLbTargetGroupAttachments.
func (s *awsLbTargetGroupAttachmentLister) AwsLbTargetGroupAttachments(namespace string) AwsLbTargetGroupAttachmentNamespaceLister {
	return awsLbTargetGroupAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLbTargetGroupAttachmentNamespaceLister helps list and get AwsLbTargetGroupAttachments.
type AwsLbTargetGroupAttachmentNamespaceLister interface {
	// List lists all AwsLbTargetGroupAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLbTargetGroupAttachment, err error)
	// Get retrieves the AwsLbTargetGroupAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsLbTargetGroupAttachment, error)
	AwsLbTargetGroupAttachmentNamespaceListerExpansion
}

// awsLbTargetGroupAttachmentNamespaceLister implements the AwsLbTargetGroupAttachmentNamespaceLister
// interface.
type awsLbTargetGroupAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLbTargetGroupAttachments in the indexer for a given namespace.
func (s awsLbTargetGroupAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLbTargetGroupAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLbTargetGroupAttachment))
	})
	return ret, err
}

// Get retrieves the AwsLbTargetGroupAttachment from the indexer for a given namespace and name.
func (s awsLbTargetGroupAttachmentNamespaceLister) Get(name string) (*v1alpha1.AwsLbTargetGroupAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awslbtargetgroupattachment"), name)
	}
	return obj.(*v1alpha1.AwsLbTargetGroupAttachment), nil
}
