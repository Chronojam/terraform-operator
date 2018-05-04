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

// AwsApiGatewayResourcesGetter has a method to return a AwsApiGatewayResourceInterface.
// A group's client should implement this interface.
type AwsApiGatewayResourcesGetter interface {
	AwsApiGatewayResources(namespace string) AwsApiGatewayResourceInterface
}

// AwsApiGatewayResourceInterface has methods to work with AwsApiGatewayResource resources.
type AwsApiGatewayResourceInterface interface {
	Create(*v1alpha1.AwsApiGatewayResource) (*v1alpha1.AwsApiGatewayResource, error)
	Update(*v1alpha1.AwsApiGatewayResource) (*v1alpha1.AwsApiGatewayResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayResource, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayResource, err error)
	AwsApiGatewayResourceExpansion
}

// awsApiGatewayResources implements AwsApiGatewayResourceInterface
type awsApiGatewayResources struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayResources returns a AwsApiGatewayResources
func newAwsApiGatewayResources(c *AwsV1alpha1Client, namespace string) *awsApiGatewayResources {
	return &awsApiGatewayResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayResource, and returns the corresponding awsApiGatewayResource object, and an error if there is any.
func (c *awsApiGatewayResources) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayResource, err error) {
	result = &v1alpha1.AwsApiGatewayResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayResources that match those selectors.
func (c *awsApiGatewayResources) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayResourceList, err error) {
	result = &v1alpha1.AwsApiGatewayResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayResources.
func (c *awsApiGatewayResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayResource and creates it.  Returns the server's representation of the awsApiGatewayResource, and an error, if there is any.
func (c *awsApiGatewayResources) Create(awsApiGatewayResource *v1alpha1.AwsApiGatewayResource) (result *v1alpha1.AwsApiGatewayResource, err error) {
	result = &v1alpha1.AwsApiGatewayResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		Body(awsApiGatewayResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayResource and updates it. Returns the server's representation of the awsApiGatewayResource, and an error, if there is any.
func (c *awsApiGatewayResources) Update(awsApiGatewayResource *v1alpha1.AwsApiGatewayResource) (result *v1alpha1.AwsApiGatewayResource, err error) {
	result = &v1alpha1.AwsApiGatewayResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		Name(awsApiGatewayResource.Name).
		Body(awsApiGatewayResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayResource and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayResource.
func (c *awsApiGatewayResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayResource, err error) {
	result = &v1alpha1.AwsApiGatewayResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
