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

// AwsApiGatewayDeploymentLister helps list AwsApiGatewayDeployments.
type AwsApiGatewayDeploymentLister interface {
	// List lists all AwsApiGatewayDeployments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayDeployment, err error)
	// AwsApiGatewayDeployments returns an object that can list and get AwsApiGatewayDeployments.
	AwsApiGatewayDeployments(namespace string) AwsApiGatewayDeploymentNamespaceLister
	AwsApiGatewayDeploymentListerExpansion
}

// awsApiGatewayDeploymentLister implements the AwsApiGatewayDeploymentLister interface.
type awsApiGatewayDeploymentLister struct {
	indexer cache.Indexer
}

// NewAwsApiGatewayDeploymentLister returns a new AwsApiGatewayDeploymentLister.
func NewAwsApiGatewayDeploymentLister(indexer cache.Indexer) AwsApiGatewayDeploymentLister {
	return &awsApiGatewayDeploymentLister{indexer: indexer}
}

// List lists all AwsApiGatewayDeployments in the indexer.
func (s *awsApiGatewayDeploymentLister) List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsApiGatewayDeployment))
	})
	return ret, err
}

// AwsApiGatewayDeployments returns an object that can list and get AwsApiGatewayDeployments.
func (s *awsApiGatewayDeploymentLister) AwsApiGatewayDeployments(namespace string) AwsApiGatewayDeploymentNamespaceLister {
	return awsApiGatewayDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsApiGatewayDeploymentNamespaceLister helps list and get AwsApiGatewayDeployments.
type AwsApiGatewayDeploymentNamespaceLister interface {
	// List lists all AwsApiGatewayDeployments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayDeployment, err error)
	// Get retrieves the AwsApiGatewayDeployment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsApiGatewayDeployment, error)
	AwsApiGatewayDeploymentNamespaceListerExpansion
}

// awsApiGatewayDeploymentNamespaceLister implements the AwsApiGatewayDeploymentNamespaceLister
// interface.
type awsApiGatewayDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsApiGatewayDeployments in the indexer for a given namespace.
func (s awsApiGatewayDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsApiGatewayDeployment))
	})
	return ret, err
}

// Get retrieves the AwsApiGatewayDeployment from the indexer for a given namespace and name.
func (s awsApiGatewayDeploymentNamespaceLister) Get(name string) (*v1alpha1.AwsApiGatewayDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsapigatewaydeployment"), name)
	}
	return obj.(*v1alpha1.AwsApiGatewayDeployment), nil
}
