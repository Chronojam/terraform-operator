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

// AwsGuarddutyIpsetsGetter has a method to return a AwsGuarddutyIpsetInterface.
// A group's client should implement this interface.
type AwsGuarddutyIpsetsGetter interface {
	AwsGuarddutyIpsets(namespace string) AwsGuarddutyIpsetInterface
}

// AwsGuarddutyIpsetInterface has methods to work with AwsGuarddutyIpset resources.
type AwsGuarddutyIpsetInterface interface {
	Create(*v1alpha1.AwsGuarddutyIpset) (*v1alpha1.AwsGuarddutyIpset, error)
	Update(*v1alpha1.AwsGuarddutyIpset) (*v1alpha1.AwsGuarddutyIpset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsGuarddutyIpset, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsGuarddutyIpsetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyIpset, err error)
	AwsGuarddutyIpsetExpansion
}

// awsGuarddutyIpsets implements AwsGuarddutyIpsetInterface
type awsGuarddutyIpsets struct {
	client rest.Interface
	ns     string
}

// newAwsGuarddutyIpsets returns a AwsGuarddutyIpsets
func newAwsGuarddutyIpsets(c *ChronojamV1alpha1Client, namespace string) *awsGuarddutyIpsets {
	return &awsGuarddutyIpsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsGuarddutyIpset, and returns the corresponding awsGuarddutyIpset object, and an error if there is any.
func (c *awsGuarddutyIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	result = &v1alpha1.AwsGuarddutyIpset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGuarddutyIpsets that match those selectors.
func (c *awsGuarddutyIpsets) List(opts v1.ListOptions) (result *v1alpha1.AwsGuarddutyIpsetList, err error) {
	result = &v1alpha1.AwsGuarddutyIpsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGuarddutyIpsets.
func (c *awsGuarddutyIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsGuarddutyIpset and creates it.  Returns the server's representation of the awsGuarddutyIpset, and an error, if there is any.
func (c *awsGuarddutyIpsets) Create(awsGuarddutyIpset *v1alpha1.AwsGuarddutyIpset) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	result = &v1alpha1.AwsGuarddutyIpset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		Body(awsGuarddutyIpset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGuarddutyIpset and updates it. Returns the server's representation of the awsGuarddutyIpset, and an error, if there is any.
func (c *awsGuarddutyIpsets) Update(awsGuarddutyIpset *v1alpha1.AwsGuarddutyIpset) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	result = &v1alpha1.AwsGuarddutyIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		Name(awsGuarddutyIpset.Name).
		Body(awsGuarddutyIpset).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGuarddutyIpset and deletes it. Returns an error if one occurs.
func (c *awsGuarddutyIpsets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGuarddutyIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGuarddutyIpset.
func (c *awsGuarddutyIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	result = &v1alpha1.AwsGuarddutyIpset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsguarddutyipsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
