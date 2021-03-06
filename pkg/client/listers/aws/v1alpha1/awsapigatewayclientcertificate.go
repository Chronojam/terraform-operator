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

// AwsApiGatewayClientCertificateLister helps list AwsApiGatewayClientCertificates.
type AwsApiGatewayClientCertificateLister interface {
	// List lists all AwsApiGatewayClientCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayClientCertificate, err error)
	// AwsApiGatewayClientCertificates returns an object that can list and get AwsApiGatewayClientCertificates.
	AwsApiGatewayClientCertificates(namespace string) AwsApiGatewayClientCertificateNamespaceLister
	AwsApiGatewayClientCertificateListerExpansion
}

// awsApiGatewayClientCertificateLister implements the AwsApiGatewayClientCertificateLister interface.
type awsApiGatewayClientCertificateLister struct {
	indexer cache.Indexer
}

// NewAwsApiGatewayClientCertificateLister returns a new AwsApiGatewayClientCertificateLister.
func NewAwsApiGatewayClientCertificateLister(indexer cache.Indexer) AwsApiGatewayClientCertificateLister {
	return &awsApiGatewayClientCertificateLister{indexer: indexer}
}

// List lists all AwsApiGatewayClientCertificates in the indexer.
func (s *awsApiGatewayClientCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayClientCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsApiGatewayClientCertificate))
	})
	return ret, err
}

// AwsApiGatewayClientCertificates returns an object that can list and get AwsApiGatewayClientCertificates.
func (s *awsApiGatewayClientCertificateLister) AwsApiGatewayClientCertificates(namespace string) AwsApiGatewayClientCertificateNamespaceLister {
	return awsApiGatewayClientCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsApiGatewayClientCertificateNamespaceLister helps list and get AwsApiGatewayClientCertificates.
type AwsApiGatewayClientCertificateNamespaceLister interface {
	// List lists all AwsApiGatewayClientCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayClientCertificate, err error)
	// Get retrieves the AwsApiGatewayClientCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsApiGatewayClientCertificate, error)
	AwsApiGatewayClientCertificateNamespaceListerExpansion
}

// awsApiGatewayClientCertificateNamespaceLister implements the AwsApiGatewayClientCertificateNamespaceLister
// interface.
type awsApiGatewayClientCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsApiGatewayClientCertificates in the indexer for a given namespace.
func (s awsApiGatewayClientCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsApiGatewayClientCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsApiGatewayClientCertificate))
	})
	return ret, err
}

// Get retrieves the AwsApiGatewayClientCertificate from the indexer for a given namespace and name.
func (s awsApiGatewayClientCertificateNamespaceLister) Get(name string) (*v1alpha1.AwsApiGatewayClientCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsapigatewayclientcertificate"), name)
	}
	return obj.(*v1alpha1.AwsApiGatewayClientCertificate), nil
}
