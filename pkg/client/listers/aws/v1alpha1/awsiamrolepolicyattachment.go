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

// AwsIamRolePolicyAttachmentLister helps list AwsIamRolePolicyAttachments.
type AwsIamRolePolicyAttachmentLister interface {
	// List lists all AwsIamRolePolicyAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamRolePolicyAttachment, err error)
	// AwsIamRolePolicyAttachments returns an object that can list and get AwsIamRolePolicyAttachments.
	AwsIamRolePolicyAttachments(namespace string) AwsIamRolePolicyAttachmentNamespaceLister
	AwsIamRolePolicyAttachmentListerExpansion
}

// awsIamRolePolicyAttachmentLister implements the AwsIamRolePolicyAttachmentLister interface.
type awsIamRolePolicyAttachmentLister struct {
	indexer cache.Indexer
}

// NewAwsIamRolePolicyAttachmentLister returns a new AwsIamRolePolicyAttachmentLister.
func NewAwsIamRolePolicyAttachmentLister(indexer cache.Indexer) AwsIamRolePolicyAttachmentLister {
	return &awsIamRolePolicyAttachmentLister{indexer: indexer}
}

// List lists all AwsIamRolePolicyAttachments in the indexer.
func (s *awsIamRolePolicyAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamRolePolicyAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamRolePolicyAttachment))
	})
	return ret, err
}

// AwsIamRolePolicyAttachments returns an object that can list and get AwsIamRolePolicyAttachments.
func (s *awsIamRolePolicyAttachmentLister) AwsIamRolePolicyAttachments(namespace string) AwsIamRolePolicyAttachmentNamespaceLister {
	return awsIamRolePolicyAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamRolePolicyAttachmentNamespaceLister helps list and get AwsIamRolePolicyAttachments.
type AwsIamRolePolicyAttachmentNamespaceLister interface {
	// List lists all AwsIamRolePolicyAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamRolePolicyAttachment, err error)
	// Get retrieves the AwsIamRolePolicyAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamRolePolicyAttachment, error)
	AwsIamRolePolicyAttachmentNamespaceListerExpansion
}

// awsIamRolePolicyAttachmentNamespaceLister implements the AwsIamRolePolicyAttachmentNamespaceLister
// interface.
type awsIamRolePolicyAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamRolePolicyAttachments in the indexer for a given namespace.
func (s awsIamRolePolicyAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamRolePolicyAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamRolePolicyAttachment))
	})
	return ret, err
}

// Get retrieves the AwsIamRolePolicyAttachment from the indexer for a given namespace and name.
func (s awsIamRolePolicyAttachmentNamespaceLister) Get(name string) (*v1alpha1.AwsIamRolePolicyAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiamrolepolicyattachment"), name)
	}
	return obj.(*v1alpha1.AwsIamRolePolicyAttachment), nil
}
