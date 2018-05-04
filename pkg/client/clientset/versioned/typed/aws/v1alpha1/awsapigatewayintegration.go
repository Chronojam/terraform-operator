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

// AwsApiGatewayIntegrationsGetter has a method to return a AwsApiGatewayIntegrationInterface.
// A group's client should implement this interface.
type AwsApiGatewayIntegrationsGetter interface {
	AwsApiGatewayIntegrations(namespace string) AwsApiGatewayIntegrationInterface
}

// AwsApiGatewayIntegrationInterface has methods to work with AwsApiGatewayIntegration resources.
type AwsApiGatewayIntegrationInterface interface {
	Create(*v1alpha1.AwsApiGatewayIntegration) (*v1alpha1.AwsApiGatewayIntegration, error)
	Update(*v1alpha1.AwsApiGatewayIntegration) (*v1alpha1.AwsApiGatewayIntegration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayIntegration, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayIntegrationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayIntegration, err error)
	AwsApiGatewayIntegrationExpansion
}

// awsApiGatewayIntegrations implements AwsApiGatewayIntegrationInterface
type awsApiGatewayIntegrations struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayIntegrations returns a AwsApiGatewayIntegrations
func newAwsApiGatewayIntegrations(c *AwsV1alpha1Client, namespace string) *awsApiGatewayIntegrations {
	return &awsApiGatewayIntegrations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayIntegration, and returns the corresponding awsApiGatewayIntegration object, and an error if there is any.
func (c *awsApiGatewayIntegrations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayIntegration, err error) {
	result = &v1alpha1.AwsApiGatewayIntegration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayIntegrations that match those selectors.
func (c *awsApiGatewayIntegrations) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayIntegrationList, err error) {
	result = &v1alpha1.AwsApiGatewayIntegrationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayIntegrations.
func (c *awsApiGatewayIntegrations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayIntegration and creates it.  Returns the server's representation of the awsApiGatewayIntegration, and an error, if there is any.
func (c *awsApiGatewayIntegrations) Create(awsApiGatewayIntegration *v1alpha1.AwsApiGatewayIntegration) (result *v1alpha1.AwsApiGatewayIntegration, err error) {
	result = &v1alpha1.AwsApiGatewayIntegration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		Body(awsApiGatewayIntegration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayIntegration and updates it. Returns the server's representation of the awsApiGatewayIntegration, and an error, if there is any.
func (c *awsApiGatewayIntegrations) Update(awsApiGatewayIntegration *v1alpha1.AwsApiGatewayIntegration) (result *v1alpha1.AwsApiGatewayIntegration, err error) {
	result = &v1alpha1.AwsApiGatewayIntegration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		Name(awsApiGatewayIntegration.Name).
		Body(awsApiGatewayIntegration).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayIntegration and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayIntegrations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayIntegrations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayIntegration.
func (c *awsApiGatewayIntegrations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayIntegration, err error) {
	result = &v1alpha1.AwsApiGatewayIntegration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayintegrations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
