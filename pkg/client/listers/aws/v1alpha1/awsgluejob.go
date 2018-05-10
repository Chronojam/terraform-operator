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

// AwsGlueJobLister helps list AwsGlueJobs.
type AwsGlueJobLister interface {
	// List lists all AwsGlueJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsGlueJob, err error)
	// AwsGlueJobs returns an object that can list and get AwsGlueJobs.
	AwsGlueJobs(namespace string) AwsGlueJobNamespaceLister
	AwsGlueJobListerExpansion
}

// awsGlueJobLister implements the AwsGlueJobLister interface.
type awsGlueJobLister struct {
	indexer cache.Indexer
}

// NewAwsGlueJobLister returns a new AwsGlueJobLister.
func NewAwsGlueJobLister(indexer cache.Indexer) AwsGlueJobLister {
	return &awsGlueJobLister{indexer: indexer}
}

// List lists all AwsGlueJobs in the indexer.
func (s *awsGlueJobLister) List(selector labels.Selector) (ret []*v1alpha1.AwsGlueJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsGlueJob))
	})
	return ret, err
}

// AwsGlueJobs returns an object that can list and get AwsGlueJobs.
func (s *awsGlueJobLister) AwsGlueJobs(namespace string) AwsGlueJobNamespaceLister {
	return awsGlueJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsGlueJobNamespaceLister helps list and get AwsGlueJobs.
type AwsGlueJobNamespaceLister interface {
	// List lists all AwsGlueJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsGlueJob, err error)
	// Get retrieves the AwsGlueJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsGlueJob, error)
	AwsGlueJobNamespaceListerExpansion
}

// awsGlueJobNamespaceLister implements the AwsGlueJobNamespaceLister
// interface.
type awsGlueJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsGlueJobs in the indexer for a given namespace.
func (s awsGlueJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsGlueJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsGlueJob))
	})
	return ret, err
}

// Get retrieves the AwsGlueJob from the indexer for a given namespace and name.
func (s awsGlueJobNamespaceLister) Get(name string) (*v1alpha1.AwsGlueJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsgluejob"), name)
	}
	return obj.(*v1alpha1.AwsGlueJob), nil
}
