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

// AwsApiGatewayAuthorizersGetter has a method to return a AwsApiGatewayAuthorizerInterface.
// A group's client should implement this interface.
type AwsApiGatewayAuthorizersGetter interface {
	AwsApiGatewayAuthorizers(namespace string) AwsApiGatewayAuthorizerInterface
}

// AwsApiGatewayAuthorizerInterface has methods to work with AwsApiGatewayAuthorizer resources.
type AwsApiGatewayAuthorizerInterface interface {
	Create(*v1alpha1.AwsApiGatewayAuthorizer) (*v1alpha1.AwsApiGatewayAuthorizer, error)
	Update(*v1alpha1.AwsApiGatewayAuthorizer) (*v1alpha1.AwsApiGatewayAuthorizer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayAuthorizer, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayAuthorizerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayAuthorizer, err error)
	AwsApiGatewayAuthorizerExpansion
}

// awsApiGatewayAuthorizers implements AwsApiGatewayAuthorizerInterface
type awsApiGatewayAuthorizers struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayAuthorizers returns a AwsApiGatewayAuthorizers
func newAwsApiGatewayAuthorizers(c *AwsV1alpha1Client, namespace string) *awsApiGatewayAuthorizers {
	return &awsApiGatewayAuthorizers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayAuthorizer, and returns the corresponding awsApiGatewayAuthorizer object, and an error if there is any.
func (c *awsApiGatewayAuthorizers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayAuthorizer, err error) {
	result = &v1alpha1.AwsApiGatewayAuthorizer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayAuthorizers that match those selectors.
func (c *awsApiGatewayAuthorizers) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayAuthorizerList, err error) {
	result = &v1alpha1.AwsApiGatewayAuthorizerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayAuthorizers.
func (c *awsApiGatewayAuthorizers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayAuthorizer and creates it.  Returns the server's representation of the awsApiGatewayAuthorizer, and an error, if there is any.
func (c *awsApiGatewayAuthorizers) Create(awsApiGatewayAuthorizer *v1alpha1.AwsApiGatewayAuthorizer) (result *v1alpha1.AwsApiGatewayAuthorizer, err error) {
	result = &v1alpha1.AwsApiGatewayAuthorizer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		Body(awsApiGatewayAuthorizer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayAuthorizer and updates it. Returns the server's representation of the awsApiGatewayAuthorizer, and an error, if there is any.
func (c *awsApiGatewayAuthorizers) Update(awsApiGatewayAuthorizer *v1alpha1.AwsApiGatewayAuthorizer) (result *v1alpha1.AwsApiGatewayAuthorizer, err error) {
	result = &v1alpha1.AwsApiGatewayAuthorizer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		Name(awsApiGatewayAuthorizer.Name).
		Body(awsApiGatewayAuthorizer).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayAuthorizer and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayAuthorizers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayAuthorizers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayAuthorizer.
func (c *awsApiGatewayAuthorizers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayAuthorizer, err error) {
	result = &v1alpha1.AwsApiGatewayAuthorizer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayauthorizers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
