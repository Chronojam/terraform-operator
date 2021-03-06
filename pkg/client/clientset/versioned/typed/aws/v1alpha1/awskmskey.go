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

// AwsKmsKeysGetter has a method to return a AwsKmsKeyInterface.
// A group's client should implement this interface.
type AwsKmsKeysGetter interface {
	AwsKmsKeys(namespace string) AwsKmsKeyInterface
}

// AwsKmsKeyInterface has methods to work with AwsKmsKey resources.
type AwsKmsKeyInterface interface {
	Create(*v1alpha1.AwsKmsKey) (*v1alpha1.AwsKmsKey, error)
	Update(*v1alpha1.AwsKmsKey) (*v1alpha1.AwsKmsKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsKmsKey, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsKmsKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKmsKey, err error)
	AwsKmsKeyExpansion
}

// awsKmsKeys implements AwsKmsKeyInterface
type awsKmsKeys struct {
	client rest.Interface
	ns     string
}

// newAwsKmsKeys returns a AwsKmsKeys
func newAwsKmsKeys(c *ChronojamV1alpha1Client, namespace string) *awsKmsKeys {
	return &awsKmsKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsKmsKey, and returns the corresponding awsKmsKey object, and an error if there is any.
func (c *awsKmsKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsKmsKey, err error) {
	result = &v1alpha1.AwsKmsKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskmskeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsKmsKeys that match those selectors.
func (c *awsKmsKeys) List(opts v1.ListOptions) (result *v1alpha1.AwsKmsKeyList, err error) {
	result = &v1alpha1.AwsKmsKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskmskeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsKmsKeys.
func (c *awsKmsKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awskmskeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsKmsKey and creates it.  Returns the server's representation of the awsKmsKey, and an error, if there is any.
func (c *awsKmsKeys) Create(awsKmsKey *v1alpha1.AwsKmsKey) (result *v1alpha1.AwsKmsKey, err error) {
	result = &v1alpha1.AwsKmsKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awskmskeys").
		Body(awsKmsKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsKmsKey and updates it. Returns the server's representation of the awsKmsKey, and an error, if there is any.
func (c *awsKmsKeys) Update(awsKmsKey *v1alpha1.AwsKmsKey) (result *v1alpha1.AwsKmsKey, err error) {
	result = &v1alpha1.AwsKmsKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awskmskeys").
		Name(awsKmsKey.Name).
		Body(awsKmsKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsKmsKey and deletes it. Returns an error if one occurs.
func (c *awsKmsKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskmskeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsKmsKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskmskeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsKmsKey.
func (c *awsKmsKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKmsKey, err error) {
	result = &v1alpha1.AwsKmsKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awskmskeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
