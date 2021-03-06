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

// AwsApiGatewayDeploymentsGetter has a method to return a AwsApiGatewayDeploymentInterface.
// A group's client should implement this interface.
type AwsApiGatewayDeploymentsGetter interface {
	AwsApiGatewayDeployments(namespace string) AwsApiGatewayDeploymentInterface
}

// AwsApiGatewayDeploymentInterface has methods to work with AwsApiGatewayDeployment resources.
type AwsApiGatewayDeploymentInterface interface {
	Create(*v1alpha1.AwsApiGatewayDeployment) (*v1alpha1.AwsApiGatewayDeployment, error)
	Update(*v1alpha1.AwsApiGatewayDeployment) (*v1alpha1.AwsApiGatewayDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayDeployment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayDeployment, err error)
	AwsApiGatewayDeploymentExpansion
}

// awsApiGatewayDeployments implements AwsApiGatewayDeploymentInterface
type awsApiGatewayDeployments struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayDeployments returns a AwsApiGatewayDeployments
func newAwsApiGatewayDeployments(c *ChronojamV1alpha1Client, namespace string) *awsApiGatewayDeployments {
	return &awsApiGatewayDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayDeployment, and returns the corresponding awsApiGatewayDeployment object, and an error if there is any.
func (c *awsApiGatewayDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayDeployment, err error) {
	result = &v1alpha1.AwsApiGatewayDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayDeployments that match those selectors.
func (c *awsApiGatewayDeployments) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayDeploymentList, err error) {
	result = &v1alpha1.AwsApiGatewayDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayDeployments.
func (c *awsApiGatewayDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayDeployment and creates it.  Returns the server's representation of the awsApiGatewayDeployment, and an error, if there is any.
func (c *awsApiGatewayDeployments) Create(awsApiGatewayDeployment *v1alpha1.AwsApiGatewayDeployment) (result *v1alpha1.AwsApiGatewayDeployment, err error) {
	result = &v1alpha1.AwsApiGatewayDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		Body(awsApiGatewayDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayDeployment and updates it. Returns the server's representation of the awsApiGatewayDeployment, and an error, if there is any.
func (c *awsApiGatewayDeployments) Update(awsApiGatewayDeployment *v1alpha1.AwsApiGatewayDeployment) (result *v1alpha1.AwsApiGatewayDeployment, err error) {
	result = &v1alpha1.AwsApiGatewayDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		Name(awsApiGatewayDeployment.Name).
		Body(awsApiGatewayDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayDeployment and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayDeployment.
func (c *awsApiGatewayDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayDeployment, err error) {
	result = &v1alpha1.AwsApiGatewayDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewaydeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
