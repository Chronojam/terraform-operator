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

// AwsApiGatewayStageLister helps list AwsApiGatewayStages.
type AwsApiGatewayStageLister interface {
	// List lists all AwsApiGatewayStages in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayStage, err error)
	// AwsApiGatewayStages returns an object that can list and get AwsApiGatewayStages.
	AwsApiGatewayStages(namespace string) AwsApiGatewayStageNamespaceLister
	AwsApiGatewayStageListerExpansion
}

// awsApiGatewayStageLister implements the AwsApiGatewayStageLister interface.
type awsApiGatewayStageLister struct {
	indexer cache.Indexer
}

// NewAwsApiGatewayStageLister returns a new AwsApiGatewayStageLister.
func NewAwsApiGatewayStageLister(indexer cache.Indexer) AwsApiGatewayStageLister {
	return &awsApiGatewayStageLister{indexer: indexer}
}

// List lists all AwsApiGatewayStages in the indexer.
func (s *awsApiGatewayStageLister) List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayStage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsApiGatewayStage))
	})
	return ret, err
}

// AwsApiGatewayStages returns an object that can list and get AwsApiGatewayStages.
func (s *awsApiGatewayStageLister) AwsApiGatewayStages(namespace string) AwsApiGatewayStageNamespaceLister {
	return awsApiGatewayStageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsApiGatewayStageNamespaceLister helps list and get AwsApiGatewayStages.
type AwsApiGatewayStageNamespaceLister interface {
	// List lists all AwsApiGatewayStages in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayStage, err error)
	// Get retrieves the AwsApiGatewayStage from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsApiGatewayStage, error)
	AwsApiGatewayStageNamespaceListerExpansion
}

// awsApiGatewayStageNamespaceLister implements the AwsApiGatewayStageNamespaceLister
// interface.
type awsApiGatewayStageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsApiGatewayStages in the indexer for a given namespace.
func (s awsApiGatewayStageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayStage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsApiGatewayStage))
	})
	return ret, err
}

// Get retrieves the AwsApiGatewayStage from the indexer for a given namespace and name.
func (s awsApiGatewayStageNamespaceLister) Get(name string) (*v1alpha1.AwsApiGatewayStage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsapigatewaystage"), name)
	}
	return obj.(*v1alpha1.AwsApiGatewayStage), nil
}
