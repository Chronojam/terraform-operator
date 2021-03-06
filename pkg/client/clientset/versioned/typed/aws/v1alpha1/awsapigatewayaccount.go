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

// AwsApiGatewayAccountsGetter has a method to return a AwsApiGatewayAccountInterface.
// A group's client should implement this interface.
type AwsApiGatewayAccountsGetter interface {
	AwsApiGatewayAccounts(namespace string) AwsApiGatewayAccountInterface
}

// AwsApiGatewayAccountInterface has methods to work with AwsApiGatewayAccount resources.
type AwsApiGatewayAccountInterface interface {
	Create(*v1alpha1.AwsApiGatewayAccount) (*v1alpha1.AwsApiGatewayAccount, error)
	Update(*v1alpha1.AwsApiGatewayAccount) (*v1alpha1.AwsApiGatewayAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayAccount, err error)
	AwsApiGatewayAccountExpansion
}

// awsApiGatewayAccounts implements AwsApiGatewayAccountInterface
type awsApiGatewayAccounts struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayAccounts returns a AwsApiGatewayAccounts
func newAwsApiGatewayAccounts(c *ChronojamV1alpha1Client, namespace string) *awsApiGatewayAccounts {
	return &awsApiGatewayAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayAccount, and returns the corresponding awsApiGatewayAccount object, and an error if there is any.
func (c *awsApiGatewayAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayAccount, err error) {
	result = &v1alpha1.AwsApiGatewayAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayAccounts that match those selectors.
func (c *awsApiGatewayAccounts) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayAccountList, err error) {
	result = &v1alpha1.AwsApiGatewayAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayAccounts.
func (c *awsApiGatewayAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayAccount and creates it.  Returns the server's representation of the awsApiGatewayAccount, and an error, if there is any.
func (c *awsApiGatewayAccounts) Create(awsApiGatewayAccount *v1alpha1.AwsApiGatewayAccount) (result *v1alpha1.AwsApiGatewayAccount, err error) {
	result = &v1alpha1.AwsApiGatewayAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		Body(awsApiGatewayAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayAccount and updates it. Returns the server's representation of the awsApiGatewayAccount, and an error, if there is any.
func (c *awsApiGatewayAccounts) Update(awsApiGatewayAccount *v1alpha1.AwsApiGatewayAccount) (result *v1alpha1.AwsApiGatewayAccount, err error) {
	result = &v1alpha1.AwsApiGatewayAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		Name(awsApiGatewayAccount.Name).
		Body(awsApiGatewayAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayAccount and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayAccount.
func (c *awsApiGatewayAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayAccount, err error) {
	result = &v1alpha1.AwsApiGatewayAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
