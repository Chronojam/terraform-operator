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

// AwsFlowLogLister helps list AwsFlowLogs.
type AwsFlowLogLister interface {
	// List lists all AwsFlowLogs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsFlowLog, err error)
	// AwsFlowLogs returns an object that can list and get AwsFlowLogs.
	AwsFlowLogs(namespace string) AwsFlowLogNamespaceLister
	AwsFlowLogListerExpansion
}

// awsFlowLogLister implements the AwsFlowLogLister interface.
type awsFlowLogLister struct {
	indexer cache.Indexer
}

// NewAwsFlowLogLister returns a new AwsFlowLogLister.
func NewAwsFlowLogLister(indexer cache.Indexer) AwsFlowLogLister {
	return &awsFlowLogLister{indexer: indexer}
}

// List lists all AwsFlowLogs in the indexer.
func (s *awsFlowLogLister) List(selector labels.Selector) (ret []*v1alpha1.AwsFlowLog, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsFlowLog))
	})
	return ret, err
}

// AwsFlowLogs returns an object that can list and get AwsFlowLogs.
func (s *awsFlowLogLister) AwsFlowLogs(namespace string) AwsFlowLogNamespaceLister {
	return awsFlowLogNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsFlowLogNamespaceLister helps list and get AwsFlowLogs.
type AwsFlowLogNamespaceLister interface {
	// List lists all AwsFlowLogs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsFlowLog, err error)
	// Get retrieves the AwsFlowLog from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsFlowLog, error)
	AwsFlowLogNamespaceListerExpansion
}

// awsFlowLogNamespaceLister implements the AwsFlowLogNamespaceLister
// interface.
type awsFlowLogNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsFlowLogs in the indexer for a given namespace.
func (s awsFlowLogNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsFlowLog, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsFlowLog))
	})
	return ret, err
}

// Get retrieves the AwsFlowLog from the indexer for a given namespace and name.
func (s awsFlowLogNamespaceLister) Get(name string) (*v1alpha1.AwsFlowLog, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsflowlog"), name)
	}
	return obj.(*v1alpha1.AwsFlowLog), nil
}
