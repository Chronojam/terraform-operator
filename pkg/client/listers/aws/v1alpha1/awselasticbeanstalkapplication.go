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

// AwsElasticBeanstalkApplicationLister helps list AwsElasticBeanstalkApplications.
type AwsElasticBeanstalkApplicationLister interface {
	// List lists all AwsElasticBeanstalkApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElasticBeanstalkApplication, err error)
	// AwsElasticBeanstalkApplications returns an object that can list and get AwsElasticBeanstalkApplications.
	AwsElasticBeanstalkApplications(namespace string) AwsElasticBeanstalkApplicationNamespaceLister
	AwsElasticBeanstalkApplicationListerExpansion
}

// awsElasticBeanstalkApplicationLister implements the AwsElasticBeanstalkApplicationLister interface.
type awsElasticBeanstalkApplicationLister struct {
	indexer cache.Indexer
}

// NewAwsElasticBeanstalkApplicationLister returns a new AwsElasticBeanstalkApplicationLister.
func NewAwsElasticBeanstalkApplicationLister(indexer cache.Indexer) AwsElasticBeanstalkApplicationLister {
	return &awsElasticBeanstalkApplicationLister{indexer: indexer}
}

// List lists all AwsElasticBeanstalkApplications in the indexer.
func (s *awsElasticBeanstalkApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElasticBeanstalkApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElasticBeanstalkApplication))
	})
	return ret, err
}

// AwsElasticBeanstalkApplications returns an object that can list and get AwsElasticBeanstalkApplications.
func (s *awsElasticBeanstalkApplicationLister) AwsElasticBeanstalkApplications(namespace string) AwsElasticBeanstalkApplicationNamespaceLister {
	return awsElasticBeanstalkApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsElasticBeanstalkApplicationNamespaceLister helps list and get AwsElasticBeanstalkApplications.
type AwsElasticBeanstalkApplicationNamespaceLister interface {
	// List lists all AwsElasticBeanstalkApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElasticBeanstalkApplication, err error)
	// Get retrieves the AwsElasticBeanstalkApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsElasticBeanstalkApplication, error)
	AwsElasticBeanstalkApplicationNamespaceListerExpansion
}

// awsElasticBeanstalkApplicationNamespaceLister implements the AwsElasticBeanstalkApplicationNamespaceLister
// interface.
type awsElasticBeanstalkApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsElasticBeanstalkApplications in the indexer for a given namespace.
func (s awsElasticBeanstalkApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElasticBeanstalkApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElasticBeanstalkApplication))
	})
	return ret, err
}

// Get retrieves the AwsElasticBeanstalkApplication from the indexer for a given namespace and name.
func (s awsElasticBeanstalkApplicationNamespaceLister) Get(name string) (*v1alpha1.AwsElasticBeanstalkApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awselasticbeanstalkapplication"), name)
	}
	return obj.(*v1alpha1.AwsElasticBeanstalkApplication), nil
}
