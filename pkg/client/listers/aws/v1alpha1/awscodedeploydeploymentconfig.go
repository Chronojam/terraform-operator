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

// AwsCodedeployDeploymentConfigLister helps list AwsCodedeployDeploymentConfigs.
type AwsCodedeployDeploymentConfigLister interface {
	// List lists all AwsCodedeployDeploymentConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCodedeployDeploymentConfig, err error)
	// AwsCodedeployDeploymentConfigs returns an object that can list and get AwsCodedeployDeploymentConfigs.
	AwsCodedeployDeploymentConfigs(namespace string) AwsCodedeployDeploymentConfigNamespaceLister
	AwsCodedeployDeploymentConfigListerExpansion
}

// awsCodedeployDeploymentConfigLister implements the AwsCodedeployDeploymentConfigLister interface.
type awsCodedeployDeploymentConfigLister struct {
	indexer cache.Indexer
}

// NewAwsCodedeployDeploymentConfigLister returns a new AwsCodedeployDeploymentConfigLister.
func NewAwsCodedeployDeploymentConfigLister(indexer cache.Indexer) AwsCodedeployDeploymentConfigLister {
	return &awsCodedeployDeploymentConfigLister{indexer: indexer}
}

// List lists all AwsCodedeployDeploymentConfigs in the indexer.
func (s *awsCodedeployDeploymentConfigLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCodedeployDeploymentConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCodedeployDeploymentConfig))
	})
	return ret, err
}

// AwsCodedeployDeploymentConfigs returns an object that can list and get AwsCodedeployDeploymentConfigs.
func (s *awsCodedeployDeploymentConfigLister) AwsCodedeployDeploymentConfigs(namespace string) AwsCodedeployDeploymentConfigNamespaceLister {
	return awsCodedeployDeploymentConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCodedeployDeploymentConfigNamespaceLister helps list and get AwsCodedeployDeploymentConfigs.
type AwsCodedeployDeploymentConfigNamespaceLister interface {
	// List lists all AwsCodedeployDeploymentConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCodedeployDeploymentConfig, err error)
	// Get retrieves the AwsCodedeployDeploymentConfig from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsCodedeployDeploymentConfig, error)
	AwsCodedeployDeploymentConfigNamespaceListerExpansion
}

// awsCodedeployDeploymentConfigNamespaceLister implements the AwsCodedeployDeploymentConfigNamespaceLister
// interface.
type awsCodedeployDeploymentConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCodedeployDeploymentConfigs in the indexer for a given namespace.
func (s awsCodedeployDeploymentConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCodedeployDeploymentConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCodedeployDeploymentConfig))
	})
	return ret, err
}

// Get retrieves the AwsCodedeployDeploymentConfig from the indexer for a given namespace and name.
func (s awsCodedeployDeploymentConfigNamespaceLister) Get(name string) (*v1alpha1.AwsCodedeployDeploymentConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awscodedeploydeploymentconfig"), name)
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentConfig), nil
}
