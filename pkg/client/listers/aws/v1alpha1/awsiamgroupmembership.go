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

// AwsIamGroupMembershipLister helps list AwsIamGroupMemberships.
type AwsIamGroupMembershipLister interface {
	// List lists all AwsIamGroupMemberships in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamGroupMembership, err error)
	// AwsIamGroupMemberships returns an object that can list and get AwsIamGroupMemberships.
	AwsIamGroupMemberships(namespace string) AwsIamGroupMembershipNamespaceLister
	AwsIamGroupMembershipListerExpansion
}

// awsIamGroupMembershipLister implements the AwsIamGroupMembershipLister interface.
type awsIamGroupMembershipLister struct {
	indexer cache.Indexer
}

// NewAwsIamGroupMembershipLister returns a new AwsIamGroupMembershipLister.
func NewAwsIamGroupMembershipLister(indexer cache.Indexer) AwsIamGroupMembershipLister {
	return &awsIamGroupMembershipLister{indexer: indexer}
}

// List lists all AwsIamGroupMemberships in the indexer.
func (s *awsIamGroupMembershipLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamGroupMembership, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamGroupMembership))
	})
	return ret, err
}

// AwsIamGroupMemberships returns an object that can list and get AwsIamGroupMemberships.
func (s *awsIamGroupMembershipLister) AwsIamGroupMemberships(namespace string) AwsIamGroupMembershipNamespaceLister {
	return awsIamGroupMembershipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamGroupMembershipNamespaceLister helps list and get AwsIamGroupMemberships.
type AwsIamGroupMembershipNamespaceLister interface {
	// List lists all AwsIamGroupMemberships in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamGroupMembership, err error)
	// Get retrieves the AwsIamGroupMembership from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamGroupMembership, error)
	AwsIamGroupMembershipNamespaceListerExpansion
}

// awsIamGroupMembershipNamespaceLister implements the AwsIamGroupMembershipNamespaceLister
// interface.
type awsIamGroupMembershipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamGroupMemberships in the indexer for a given namespace.
func (s awsIamGroupMembershipNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamGroupMembership, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamGroupMembership))
	})
	return ret, err
}

// Get retrieves the AwsIamGroupMembership from the indexer for a given namespace and name.
func (s awsIamGroupMembershipNamespaceLister) Get(name string) (*v1alpha1.AwsIamGroupMembership, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiamgroupmembership"), name)
	}
	return obj.(*v1alpha1.AwsIamGroupMembership), nil
}
