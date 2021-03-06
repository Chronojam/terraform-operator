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

// AwsApiGatewayMethodResponsesGetter has a method to return a AwsApiGatewayMethodResponseInterface.
// A group's client should implement this interface.
type AwsApiGatewayMethodResponsesGetter interface {
	AwsApiGatewayMethodResponses(namespace string) AwsApiGatewayMethodResponseInterface
}

// AwsApiGatewayMethodResponseInterface has methods to work with AwsApiGatewayMethodResponse resources.
type AwsApiGatewayMethodResponseInterface interface {
	Create(*v1alpha1.AwsApiGatewayMethodResponse) (*v1alpha1.AwsApiGatewayMethodResponse, error)
	Update(*v1alpha1.AwsApiGatewayMethodResponse) (*v1alpha1.AwsApiGatewayMethodResponse, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayMethodResponse, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayMethodResponseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayMethodResponse, err error)
	AwsApiGatewayMethodResponseExpansion
}

// awsApiGatewayMethodResponses implements AwsApiGatewayMethodResponseInterface
type awsApiGatewayMethodResponses struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayMethodResponses returns a AwsApiGatewayMethodResponses
func newAwsApiGatewayMethodResponses(c *ChronojamV1alpha1Client, namespace string) *awsApiGatewayMethodResponses {
	return &awsApiGatewayMethodResponses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayMethodResponse, and returns the corresponding awsApiGatewayMethodResponse object, and an error if there is any.
func (c *awsApiGatewayMethodResponses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayMethodResponse, err error) {
	result = &v1alpha1.AwsApiGatewayMethodResponse{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayMethodResponses that match those selectors.
func (c *awsApiGatewayMethodResponses) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayMethodResponseList, err error) {
	result = &v1alpha1.AwsApiGatewayMethodResponseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayMethodResponses.
func (c *awsApiGatewayMethodResponses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayMethodResponse and creates it.  Returns the server's representation of the awsApiGatewayMethodResponse, and an error, if there is any.
func (c *awsApiGatewayMethodResponses) Create(awsApiGatewayMethodResponse *v1alpha1.AwsApiGatewayMethodResponse) (result *v1alpha1.AwsApiGatewayMethodResponse, err error) {
	result = &v1alpha1.AwsApiGatewayMethodResponse{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		Body(awsApiGatewayMethodResponse).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayMethodResponse and updates it. Returns the server's representation of the awsApiGatewayMethodResponse, and an error, if there is any.
func (c *awsApiGatewayMethodResponses) Update(awsApiGatewayMethodResponse *v1alpha1.AwsApiGatewayMethodResponse) (result *v1alpha1.AwsApiGatewayMethodResponse, err error) {
	result = &v1alpha1.AwsApiGatewayMethodResponse{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		Name(awsApiGatewayMethodResponse.Name).
		Body(awsApiGatewayMethodResponse).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayMethodResponse and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayMethodResponses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayMethodResponses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayMethodResponse.
func (c *awsApiGatewayMethodResponses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayMethodResponse, err error) {
	result = &v1alpha1.AwsApiGatewayMethodResponse{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewaymethodresponses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
