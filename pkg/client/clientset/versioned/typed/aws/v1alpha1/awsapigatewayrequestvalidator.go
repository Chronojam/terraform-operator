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

// AwsApiGatewayRequestValidatorsGetter has a method to return a AwsApiGatewayRequestValidatorInterface.
// A group's client should implement this interface.
type AwsApiGatewayRequestValidatorsGetter interface {
	AwsApiGatewayRequestValidators(namespace string) AwsApiGatewayRequestValidatorInterface
}

// AwsApiGatewayRequestValidatorInterface has methods to work with AwsApiGatewayRequestValidator resources.
type AwsApiGatewayRequestValidatorInterface interface {
	Create(*v1alpha1.AwsApiGatewayRequestValidator) (*v1alpha1.AwsApiGatewayRequestValidator, error)
	Update(*v1alpha1.AwsApiGatewayRequestValidator) (*v1alpha1.AwsApiGatewayRequestValidator, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayRequestValidator, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayRequestValidatorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayRequestValidator, err error)
	AwsApiGatewayRequestValidatorExpansion
}

// awsApiGatewayRequestValidators implements AwsApiGatewayRequestValidatorInterface
type awsApiGatewayRequestValidators struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayRequestValidators returns a AwsApiGatewayRequestValidators
func newAwsApiGatewayRequestValidators(c *ChronojamV1alpha1Client, namespace string) *awsApiGatewayRequestValidators {
	return &awsApiGatewayRequestValidators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayRequestValidator, and returns the corresponding awsApiGatewayRequestValidator object, and an error if there is any.
func (c *awsApiGatewayRequestValidators) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayRequestValidator, err error) {
	result = &v1alpha1.AwsApiGatewayRequestValidator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayRequestValidators that match those selectors.
func (c *awsApiGatewayRequestValidators) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayRequestValidatorList, err error) {
	result = &v1alpha1.AwsApiGatewayRequestValidatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayRequestValidators.
func (c *awsApiGatewayRequestValidators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayRequestValidator and creates it.  Returns the server's representation of the awsApiGatewayRequestValidator, and an error, if there is any.
func (c *awsApiGatewayRequestValidators) Create(awsApiGatewayRequestValidator *v1alpha1.AwsApiGatewayRequestValidator) (result *v1alpha1.AwsApiGatewayRequestValidator, err error) {
	result = &v1alpha1.AwsApiGatewayRequestValidator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		Body(awsApiGatewayRequestValidator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayRequestValidator and updates it. Returns the server's representation of the awsApiGatewayRequestValidator, and an error, if there is any.
func (c *awsApiGatewayRequestValidators) Update(awsApiGatewayRequestValidator *v1alpha1.AwsApiGatewayRequestValidator) (result *v1alpha1.AwsApiGatewayRequestValidator, err error) {
	result = &v1alpha1.AwsApiGatewayRequestValidator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		Name(awsApiGatewayRequestValidator.Name).
		Body(awsApiGatewayRequestValidator).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayRequestValidator and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayRequestValidators) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayRequestValidators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayRequestValidator.
func (c *awsApiGatewayRequestValidators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayRequestValidator, err error) {
	result = &v1alpha1.AwsApiGatewayRequestValidator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayrequestvalidators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
