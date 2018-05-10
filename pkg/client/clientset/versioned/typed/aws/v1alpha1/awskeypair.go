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

// AwsKeyPairsGetter has a method to return a AwsKeyPairInterface.
// A group's client should implement this interface.
type AwsKeyPairsGetter interface {
	AwsKeyPairs(namespace string) AwsKeyPairInterface
}

// AwsKeyPairInterface has methods to work with AwsKeyPair resources.
type AwsKeyPairInterface interface {
	Create(*v1alpha1.AwsKeyPair) (*v1alpha1.AwsKeyPair, error)
	Update(*v1alpha1.AwsKeyPair) (*v1alpha1.AwsKeyPair, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsKeyPair, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsKeyPairList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKeyPair, err error)
	AwsKeyPairExpansion
}

// awsKeyPairs implements AwsKeyPairInterface
type awsKeyPairs struct {
	client rest.Interface
	ns     string
}

// newAwsKeyPairs returns a AwsKeyPairs
func newAwsKeyPairs(c *ChronojamV1alpha1Client, namespace string) *awsKeyPairs {
	return &awsKeyPairs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsKeyPair, and returns the corresponding awsKeyPair object, and an error if there is any.
func (c *awsKeyPairs) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsKeyPair, err error) {
	result = &v1alpha1.AwsKeyPair{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskeypairs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsKeyPairs that match those selectors.
func (c *awsKeyPairs) List(opts v1.ListOptions) (result *v1alpha1.AwsKeyPairList, err error) {
	result = &v1alpha1.AwsKeyPairList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskeypairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsKeyPairs.
func (c *awsKeyPairs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awskeypairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsKeyPair and creates it.  Returns the server's representation of the awsKeyPair, and an error, if there is any.
func (c *awsKeyPairs) Create(awsKeyPair *v1alpha1.AwsKeyPair) (result *v1alpha1.AwsKeyPair, err error) {
	result = &v1alpha1.AwsKeyPair{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awskeypairs").
		Body(awsKeyPair).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsKeyPair and updates it. Returns the server's representation of the awsKeyPair, and an error, if there is any.
func (c *awsKeyPairs) Update(awsKeyPair *v1alpha1.AwsKeyPair) (result *v1alpha1.AwsKeyPair, err error) {
	result = &v1alpha1.AwsKeyPair{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awskeypairs").
		Name(awsKeyPair.Name).
		Body(awsKeyPair).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsKeyPair and deletes it. Returns an error if one occurs.
func (c *awsKeyPairs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskeypairs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsKeyPairs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskeypairs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsKeyPair.
func (c *awsKeyPairs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKeyPair, err error) {
	result = &v1alpha1.AwsKeyPair{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awskeypairs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
