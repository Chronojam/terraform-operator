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

// AwsDefaultVpcDhcpOptionsesGetter has a method to return a AwsDefaultVpcDhcpOptionsInterface.
// A group's client should implement this interface.
type AwsDefaultVpcDhcpOptionsesGetter interface {
	AwsDefaultVpcDhcpOptionses(namespace string) AwsDefaultVpcDhcpOptionsInterface
}

// AwsDefaultVpcDhcpOptionsInterface has methods to work with AwsDefaultVpcDhcpOptions resources.
type AwsDefaultVpcDhcpOptionsInterface interface {
	Create(*v1alpha1.AwsDefaultVpcDhcpOptions) (*v1alpha1.AwsDefaultVpcDhcpOptions, error)
	Update(*v1alpha1.AwsDefaultVpcDhcpOptions) (*v1alpha1.AwsDefaultVpcDhcpOptions, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDefaultVpcDhcpOptions, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDefaultVpcDhcpOptionsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultVpcDhcpOptions, err error)
	AwsDefaultVpcDhcpOptionsExpansion
}

// awsDefaultVpcDhcpOptionses implements AwsDefaultVpcDhcpOptionsInterface
type awsDefaultVpcDhcpOptionses struct {
	client rest.Interface
	ns     string
}

// newAwsDefaultVpcDhcpOptionses returns a AwsDefaultVpcDhcpOptionses
func newAwsDefaultVpcDhcpOptionses(c *ChronojamV1alpha1Client, namespace string) *awsDefaultVpcDhcpOptionses {
	return &awsDefaultVpcDhcpOptionses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDefaultVpcDhcpOptions, and returns the corresponding awsDefaultVpcDhcpOptions object, and an error if there is any.
func (c *awsDefaultVpcDhcpOptionses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDefaultVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsDefaultVpcDhcpOptions{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDefaultVpcDhcpOptionses that match those selectors.
func (c *awsDefaultVpcDhcpOptionses) List(opts v1.ListOptions) (result *v1alpha1.AwsDefaultVpcDhcpOptionsList, err error) {
	result = &v1alpha1.AwsDefaultVpcDhcpOptionsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDefaultVpcDhcpOptionses.
func (c *awsDefaultVpcDhcpOptionses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDefaultVpcDhcpOptions and creates it.  Returns the server's representation of the awsDefaultVpcDhcpOptions, and an error, if there is any.
func (c *awsDefaultVpcDhcpOptionses) Create(awsDefaultVpcDhcpOptions *v1alpha1.AwsDefaultVpcDhcpOptions) (result *v1alpha1.AwsDefaultVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsDefaultVpcDhcpOptions{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		Body(awsDefaultVpcDhcpOptions).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDefaultVpcDhcpOptions and updates it. Returns the server's representation of the awsDefaultVpcDhcpOptions, and an error, if there is any.
func (c *awsDefaultVpcDhcpOptionses) Update(awsDefaultVpcDhcpOptions *v1alpha1.AwsDefaultVpcDhcpOptions) (result *v1alpha1.AwsDefaultVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsDefaultVpcDhcpOptions{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		Name(awsDefaultVpcDhcpOptions.Name).
		Body(awsDefaultVpcDhcpOptions).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDefaultVpcDhcpOptions and deletes it. Returns an error if one occurs.
func (c *awsDefaultVpcDhcpOptionses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDefaultVpcDhcpOptionses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDefaultVpcDhcpOptions.
func (c *awsDefaultVpcDhcpOptionses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultVpcDhcpOptions, err error) {
	result = &v1alpha1.AwsDefaultVpcDhcpOptions{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdefaultvpcdhcpoptionses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
