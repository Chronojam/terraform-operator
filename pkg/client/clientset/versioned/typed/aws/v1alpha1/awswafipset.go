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

// AwsWafIpsetsGetter has a method to return a AwsWafIpsetInterface.
// A group's client should implement this interface.
type AwsWafIpsetsGetter interface {
	AwsWafIpsets(namespace string) AwsWafIpsetInterface
}

// AwsWafIpsetInterface has methods to work with AwsWafIpset resources.
type AwsWafIpsetInterface interface {
	Create(*v1alpha1.AwsWafIpset) (*v1alpha1.AwsWafIpset, error)
	Update(*v1alpha1.AwsWafIpset) (*v1alpha1.AwsWafIpset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafIpset, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafIpsetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafIpset, err error)
	AwsWafIpsetExpansion
}

// awsWafIpsets implements AwsWafIpsetInterface
type awsWafIpsets struct {
	client rest.Interface
	ns     string
}

// newAwsWafIpsets returns a AwsWafIpsets
func newAwsWafIpsets(c *ChronojamV1alpha1Client, namespace string) *awsWafIpsets {
	return &awsWafIpsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafIpset, and returns the corresponding awsWafIpset object, and an error if there is any.
func (c *awsWafIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafIpset, err error) {
	result = &v1alpha1.AwsWafIpset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafipsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafIpsets that match those selectors.
func (c *awsWafIpsets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafIpsetList, err error) {
	result = &v1alpha1.AwsWafIpsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafIpsets.
func (c *awsWafIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafIpset and creates it.  Returns the server's representation of the awsWafIpset, and an error, if there is any.
func (c *awsWafIpsets) Create(awsWafIpset *v1alpha1.AwsWafIpset) (result *v1alpha1.AwsWafIpset, err error) {
	result = &v1alpha1.AwsWafIpset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafipsets").
		Body(awsWafIpset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafIpset and updates it. Returns the server's representation of the awsWafIpset, and an error, if there is any.
func (c *awsWafIpsets) Update(awsWafIpset *v1alpha1.AwsWafIpset) (result *v1alpha1.AwsWafIpset, err error) {
	result = &v1alpha1.AwsWafIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafipsets").
		Name(awsWafIpset.Name).
		Body(awsWafIpset).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafIpset and deletes it. Returns an error if one occurs.
func (c *awsWafIpsets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafipsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafipsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafIpset.
func (c *awsWafIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafIpset, err error) {
	result = &v1alpha1.AwsWafIpset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafipsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
