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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	scheme "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsAcmCertificateValidationsGetter has a method to return a AwsAcmCertificateValidationInterface.
// A group's client should implement this interface.
type AwsAcmCertificateValidationsGetter interface {
	AwsAcmCertificateValidations(namespace string) AwsAcmCertificateValidationInterface
}

// AwsAcmCertificateValidationInterface has methods to work with AwsAcmCertificateValidation resources.
type AwsAcmCertificateValidationInterface interface {
	Create(*v1alpha1.AwsAcmCertificateValidation) (*v1alpha1.AwsAcmCertificateValidation, error)
	Update(*v1alpha1.AwsAcmCertificateValidation) (*v1alpha1.AwsAcmCertificateValidation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAcmCertificateValidation, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAcmCertificateValidationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAcmCertificateValidation, err error)
	AwsAcmCertificateValidationExpansion
}

// awsAcmCertificateValidations implements AwsAcmCertificateValidationInterface
type awsAcmCertificateValidations struct {
	client rest.Interface
	ns     string
}

// newAwsAcmCertificateValidations returns a AwsAcmCertificateValidations
func newAwsAcmCertificateValidations(c *ChronojamV1alpha1Client, namespace string) *awsAcmCertificateValidations {
	return &awsAcmCertificateValidations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAcmCertificateValidation, and returns the corresponding awsAcmCertificateValidation object, and an error if there is any.
func (c *awsAcmCertificateValidations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAcmCertificateValidation, err error) {
	result = &v1alpha1.AwsAcmCertificateValidation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAcmCertificateValidations that match those selectors.
func (c *awsAcmCertificateValidations) List(opts v1.ListOptions) (result *v1alpha1.AwsAcmCertificateValidationList, err error) {
	result = &v1alpha1.AwsAcmCertificateValidationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAcmCertificateValidations.
func (c *awsAcmCertificateValidations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAcmCertificateValidation and creates it.  Returns the server's representation of the awsAcmCertificateValidation, and an error, if there is any.
func (c *awsAcmCertificateValidations) Create(awsAcmCertificateValidation *v1alpha1.AwsAcmCertificateValidation) (result *v1alpha1.AwsAcmCertificateValidation, err error) {
	result = &v1alpha1.AwsAcmCertificateValidation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		Body(awsAcmCertificateValidation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAcmCertificateValidation and updates it. Returns the server's representation of the awsAcmCertificateValidation, and an error, if there is any.
func (c *awsAcmCertificateValidations) Update(awsAcmCertificateValidation *v1alpha1.AwsAcmCertificateValidation) (result *v1alpha1.AwsAcmCertificateValidation, err error) {
	result = &v1alpha1.AwsAcmCertificateValidation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		Name(awsAcmCertificateValidation.Name).
		Body(awsAcmCertificateValidation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAcmCertificateValidation and deletes it. Returns an error if one occurs.
func (c *awsAcmCertificateValidations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAcmCertificateValidations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAcmCertificateValidation.
func (c *awsAcmCertificateValidations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAcmCertificateValidation, err error) {
	result = &v1alpha1.AwsAcmCertificateValidation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsacmcertificatevalidations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
