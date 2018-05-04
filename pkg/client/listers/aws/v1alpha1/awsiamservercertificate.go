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

// AwsIamServerCertificateLister helps list AwsIamServerCertificates.
type AwsIamServerCertificateLister interface {
	// List lists all AwsIamServerCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamServerCertificate, err error)
	// AwsIamServerCertificates returns an object that can list and get AwsIamServerCertificates.
	AwsIamServerCertificates(namespace string) AwsIamServerCertificateNamespaceLister
	AwsIamServerCertificateListerExpansion
}

// awsIamServerCertificateLister implements the AwsIamServerCertificateLister interface.
type awsIamServerCertificateLister struct {
	indexer cache.Indexer
}

// NewAwsIamServerCertificateLister returns a new AwsIamServerCertificateLister.
func NewAwsIamServerCertificateLister(indexer cache.Indexer) AwsIamServerCertificateLister {
	return &awsIamServerCertificateLister{indexer: indexer}
}

// List lists all AwsIamServerCertificates in the indexer.
func (s *awsIamServerCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamServerCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamServerCertificate))
	})
	return ret, err
}

// AwsIamServerCertificates returns an object that can list and get AwsIamServerCertificates.
func (s *awsIamServerCertificateLister) AwsIamServerCertificates(namespace string) AwsIamServerCertificateNamespaceLister {
	return awsIamServerCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamServerCertificateNamespaceLister helps list and get AwsIamServerCertificates.
type AwsIamServerCertificateNamespaceLister interface {
	// List lists all AwsIamServerCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AwsIamServerCertificate, err error)
	// Get retrieves the AwsIamServerCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AwsIamServerCertificate, error)
	AwsIamServerCertificateNamespaceListerExpansion
}

// awsIamServerCertificateNamespaceLister implements the AwsIamServerCertificateNamespaceLister
// interface.
type awsIamServerCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamServerCertificates in the indexer for a given namespace.
func (s awsIamServerCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsIamServerCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsIamServerCertificate))
	})
	return ret, err
}

// Get retrieves the AwsIamServerCertificate from the indexer for a given namespace and name.
func (s awsIamServerCertificateNamespaceLister) Get(name string) (*v1alpha1.AwsIamServerCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsiamservercertificate"), name)
	}
	return obj.(*v1alpha1.AwsIamServerCertificate), nil
}
