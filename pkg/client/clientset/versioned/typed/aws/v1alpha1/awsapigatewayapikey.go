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

// AwsApiGatewayApiKeysGetter has a method to return a AwsApiGatewayApiKeyInterface.
// A group's client should implement this interface.
type AwsApiGatewayApiKeysGetter interface {
	AwsApiGatewayApiKeys(namespace string) AwsApiGatewayApiKeyInterface
}

// AwsApiGatewayApiKeyInterface has methods to work with AwsApiGatewayApiKey resources.
type AwsApiGatewayApiKeyInterface interface {
	Create(*v1alpha1.AwsApiGatewayApiKey) (*v1alpha1.AwsApiGatewayApiKey, error)
	Update(*v1alpha1.AwsApiGatewayApiKey) (*v1alpha1.AwsApiGatewayApiKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayApiKey, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayApiKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayApiKey, err error)
	AwsApiGatewayApiKeyExpansion
}

// awsApiGatewayApiKeys implements AwsApiGatewayApiKeyInterface
type awsApiGatewayApiKeys struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayApiKeys returns a AwsApiGatewayApiKeys
func newAwsApiGatewayApiKeys(c *AwsV1alpha1Client, namespace string) *awsApiGatewayApiKeys {
	return &awsApiGatewayApiKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayApiKey, and returns the corresponding awsApiGatewayApiKey object, and an error if there is any.
func (c *awsApiGatewayApiKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayApiKey, err error) {
	result = &v1alpha1.AwsApiGatewayApiKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayApiKeys that match those selectors.
func (c *awsApiGatewayApiKeys) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayApiKeyList, err error) {
	result = &v1alpha1.AwsApiGatewayApiKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayApiKeys.
func (c *awsApiGatewayApiKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayApiKey and creates it.  Returns the server's representation of the awsApiGatewayApiKey, and an error, if there is any.
func (c *awsApiGatewayApiKeys) Create(awsApiGatewayApiKey *v1alpha1.AwsApiGatewayApiKey) (result *v1alpha1.AwsApiGatewayApiKey, err error) {
	result = &v1alpha1.AwsApiGatewayApiKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		Body(awsApiGatewayApiKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayApiKey and updates it. Returns the server's representation of the awsApiGatewayApiKey, and an error, if there is any.
func (c *awsApiGatewayApiKeys) Update(awsApiGatewayApiKey *v1alpha1.AwsApiGatewayApiKey) (result *v1alpha1.AwsApiGatewayApiKey, err error) {
	result = &v1alpha1.AwsApiGatewayApiKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		Name(awsApiGatewayApiKey.Name).
		Body(awsApiGatewayApiKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayApiKey and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayApiKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayApiKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayApiKey.
func (c *awsApiGatewayApiKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayApiKey, err error) {
	result = &v1alpha1.AwsApiGatewayApiKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayapikeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
