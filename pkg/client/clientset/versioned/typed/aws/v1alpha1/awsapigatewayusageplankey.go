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

// AwsApiGatewayUsagePlanKeysGetter has a method to return a AwsApiGatewayUsagePlanKeyInterface.
// A group's client should implement this interface.
type AwsApiGatewayUsagePlanKeysGetter interface {
	AwsApiGatewayUsagePlanKeys(namespace string) AwsApiGatewayUsagePlanKeyInterface
}

// AwsApiGatewayUsagePlanKeyInterface has methods to work with AwsApiGatewayUsagePlanKey resources.
type AwsApiGatewayUsagePlanKeyInterface interface {
	Create(*v1alpha1.AwsApiGatewayUsagePlanKey) (*v1alpha1.AwsApiGatewayUsagePlanKey, error)
	Update(*v1alpha1.AwsApiGatewayUsagePlanKey) (*v1alpha1.AwsApiGatewayUsagePlanKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayUsagePlanKey, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayUsagePlanKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayUsagePlanKey, err error)
	AwsApiGatewayUsagePlanKeyExpansion
}

// awsApiGatewayUsagePlanKeys implements AwsApiGatewayUsagePlanKeyInterface
type awsApiGatewayUsagePlanKeys struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayUsagePlanKeys returns a AwsApiGatewayUsagePlanKeys
func newAwsApiGatewayUsagePlanKeys(c *ChronojamV1alpha1Client, namespace string) *awsApiGatewayUsagePlanKeys {
	return &awsApiGatewayUsagePlanKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayUsagePlanKey, and returns the corresponding awsApiGatewayUsagePlanKey object, and an error if there is any.
func (c *awsApiGatewayUsagePlanKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlanKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayUsagePlanKeys that match those selectors.
func (c *awsApiGatewayUsagePlanKeys) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayUsagePlanKeyList, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlanKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayUsagePlanKeys.
func (c *awsApiGatewayUsagePlanKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayUsagePlanKey and creates it.  Returns the server's representation of the awsApiGatewayUsagePlanKey, and an error, if there is any.
func (c *awsApiGatewayUsagePlanKeys) Create(awsApiGatewayUsagePlanKey *v1alpha1.AwsApiGatewayUsagePlanKey) (result *v1alpha1.AwsApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlanKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		Body(awsApiGatewayUsagePlanKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayUsagePlanKey and updates it. Returns the server's representation of the awsApiGatewayUsagePlanKey, and an error, if there is any.
func (c *awsApiGatewayUsagePlanKeys) Update(awsApiGatewayUsagePlanKey *v1alpha1.AwsApiGatewayUsagePlanKey) (result *v1alpha1.AwsApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlanKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		Name(awsApiGatewayUsagePlanKey.Name).
		Body(awsApiGatewayUsagePlanKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayUsagePlanKey and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayUsagePlanKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayUsagePlanKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayUsagePlanKey.
func (c *awsApiGatewayUsagePlanKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.AwsApiGatewayUsagePlanKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewayusageplankeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
