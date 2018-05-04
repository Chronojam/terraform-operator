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

// AwsSubnetsGetter has a method to return a AwsSubnetInterface.
// A group's client should implement this interface.
type AwsSubnetsGetter interface {
	AwsSubnets(namespace string) AwsSubnetInterface
}

// AwsSubnetInterface has methods to work with AwsSubnet resources.
type AwsSubnetInterface interface {
	Create(*v1alpha1.AwsSubnet) (*v1alpha1.AwsSubnet, error)
	Update(*v1alpha1.AwsSubnet) (*v1alpha1.AwsSubnet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSubnet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSubnetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSubnet, err error)
	AwsSubnetExpansion
}

// awsSubnets implements AwsSubnetInterface
type awsSubnets struct {
	client rest.Interface
	ns     string
}

// newAwsSubnets returns a AwsSubnets
func newAwsSubnets(c *ChronojamV1alpha1Client, namespace string) *awsSubnets {
	return &awsSubnets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSubnet, and returns the corresponding awsSubnet object, and an error if there is any.
func (c *awsSubnets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSubnet, err error) {
	result = &v1alpha1.AwsSubnet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssubnets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSubnets that match those selectors.
func (c *awsSubnets) List(opts v1.ListOptions) (result *v1alpha1.AwsSubnetList, err error) {
	result = &v1alpha1.AwsSubnetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssubnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSubnets.
func (c *awsSubnets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssubnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSubnet and creates it.  Returns the server's representation of the awsSubnet, and an error, if there is any.
func (c *awsSubnets) Create(awsSubnet *v1alpha1.AwsSubnet) (result *v1alpha1.AwsSubnet, err error) {
	result = &v1alpha1.AwsSubnet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssubnets").
		Body(awsSubnet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSubnet and updates it. Returns the server's representation of the awsSubnet, and an error, if there is any.
func (c *awsSubnets) Update(awsSubnet *v1alpha1.AwsSubnet) (result *v1alpha1.AwsSubnet, err error) {
	result = &v1alpha1.AwsSubnet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssubnets").
		Name(awsSubnet.Name).
		Body(awsSubnet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSubnet and deletes it. Returns an error if one occurs.
func (c *awsSubnets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssubnets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSubnets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssubnets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSubnet.
func (c *awsSubnets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSubnet, err error) {
	result = &v1alpha1.AwsSubnet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssubnets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
