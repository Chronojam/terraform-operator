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

// AwsLaunchConfigurationLister helps list AwsLaunchConfigurations.
type AwsLaunchConfigurationLister interface {
	// List lists all AwsLaunchConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLaunchConfiguration, err error)
	// AwsLaunchConfigurations returns an object that can list and get AwsLaunchConfigurations.
	AwsLaunchConfigurations(namespace string) AwsLaunchConfigurationNamespaceLister
	AwsLaunchConfigurationListerExpansion
}

// awsLaunchConfigurationLister implements the AwsLaunchConfigurationLister interface.
type awsLaunchConfigurationLister struct {
	indexer cache.Indexer
}

// NewAwsLaunchConfigurationLister returns a new AwsLaunchConfigurationLister.
func NewAwsLaunchConfigurationLister(indexer cache.Indexer) AwsLaunchConfigurationLister {
	return &awsLaunchConfigurationLister{indexer: indexer}
}

// List lists all AwsLaunchConfigurations in the indexer.
func (s *awsLaunchConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLaunchConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLaunchConfiguration))
	})
	return ret, err
}

// AwsLaunchConfigurations returns an object that can list and get AwsLaunchConfigurations.
func (s *awsLaunchConfigurationLister) AwsLaunchConfigurations(namespace string) AwsLaunchConfigurationNamespaceLister {
	return awsLaunchConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLaunchConfigurationNamespaceLister helps list and get AwsLaunchConfigurations.
type AwsLaunchConfigurationNamespaceLister interface {
	// List lists all AwsLaunchConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLaunchConfiguration, err error)
	// Get retrieves the AwsLaunchConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsLaunchConfiguration, error)
	AwsLaunchConfigurationNamespaceListerExpansion
}

// awsLaunchConfigurationNamespaceLister implements the AwsLaunchConfigurationNamespaceLister
// interface.
type awsLaunchConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLaunchConfigurations in the indexer for a given namespace.
func (s awsLaunchConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLaunchConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLaunchConfiguration))
	})
	return ret, err
}

// Get retrieves the AwsLaunchConfiguration from the indexer for a given namespace and name.
func (s awsLaunchConfigurationNamespaceLister) Get(name string) (*v1alpha1.AwsLaunchConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awslaunchconfiguration"), name)
	}
	return obj.(*v1alpha1.AwsLaunchConfiguration), nil
}
