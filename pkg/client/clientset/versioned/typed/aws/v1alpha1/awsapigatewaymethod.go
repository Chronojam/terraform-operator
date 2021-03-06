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

// AwsApiGatewayMethodsGetter has a method to return a AwsApiGatewayMethodInterface.
// A group's client should implement this interface.
type AwsApiGatewayMethodsGetter interface {
	AwsApiGatewayMethods(namespace string) AwsApiGatewayMethodInterface
}

// AwsApiGatewayMethodInterface has methods to work with AwsApiGatewayMethod resources.
type AwsApiGatewayMethodInterface interface {
	Create(*v1alpha1.AwsApiGatewayMethod) (*v1alpha1.AwsApiGatewayMethod, error)
	Update(*v1alpha1.AwsApiGatewayMethod) (*v1alpha1.AwsApiGatewayMethod, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayMethod, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayMethodList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayMethod, err error)
	AwsApiGatewayMethodExpansion
}

// awsApiGatewayMethods implements AwsApiGatewayMethodInterface
type awsApiGatewayMethods struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayMethods returns a AwsApiGatewayMethods
func newAwsApiGatewayMethods(c *ChronojamV1alpha1Client, namespace string) *awsApiGatewayMethods {
	return &awsApiGatewayMethods{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayMethod, and returns the corresponding awsApiGatewayMethod object, and an error if there is any.
func (c *awsApiGatewayMethods) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	result = &v1alpha1.AwsApiGatewayMethod{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayMethods that match those selectors.
func (c *awsApiGatewayMethods) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayMethodList, err error) {
	result = &v1alpha1.AwsApiGatewayMethodList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayMethods.
func (c *awsApiGatewayMethods) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayMethod and creates it.  Returns the server's representation of the awsApiGatewayMethod, and an error, if there is any.
func (c *awsApiGatewayMethods) Create(awsApiGatewayMethod *v1alpha1.AwsApiGatewayMethod) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	result = &v1alpha1.AwsApiGatewayMethod{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		Body(awsApiGatewayMethod).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayMethod and updates it. Returns the server's representation of the awsApiGatewayMethod, and an error, if there is any.
func (c *awsApiGatewayMethods) Update(awsApiGatewayMethod *v1alpha1.AwsApiGatewayMethod) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	result = &v1alpha1.AwsApiGatewayMethod{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		Name(awsApiGatewayMethod.Name).
		Body(awsApiGatewayMethod).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayMethod and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayMethods) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayMethods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayMethod.
func (c *awsApiGatewayMethods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	result = &v1alpha1.AwsApiGatewayMethod{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewaymethods").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
