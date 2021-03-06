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

// AwsRoute53ZonesGetter has a method to return a AwsRoute53ZoneInterface.
// A group's client should implement this interface.
type AwsRoute53ZonesGetter interface {
	AwsRoute53Zones(namespace string) AwsRoute53ZoneInterface
}

// AwsRoute53ZoneInterface has methods to work with AwsRoute53Zone resources.
type AwsRoute53ZoneInterface interface {
	Create(*v1alpha1.AwsRoute53Zone) (*v1alpha1.AwsRoute53Zone, error)
	Update(*v1alpha1.AwsRoute53Zone) (*v1alpha1.AwsRoute53Zone, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsRoute53Zone, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsRoute53ZoneList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRoute53Zone, err error)
	AwsRoute53ZoneExpansion
}

// awsRoute53Zones implements AwsRoute53ZoneInterface
type awsRoute53Zones struct {
	client rest.Interface
	ns     string
}

// newAwsRoute53Zones returns a AwsRoute53Zones
func newAwsRoute53Zones(c *ChronojamV1alpha1Client, namespace string) *awsRoute53Zones {
	return &awsRoute53Zones{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsRoute53Zone, and returns the corresponding awsRoute53Zone object, and an error if there is any.
func (c *awsRoute53Zones) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRoute53Zone, err error) {
	result = &v1alpha1.AwsRoute53Zone{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsroute53zones").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsRoute53Zones that match those selectors.
func (c *awsRoute53Zones) List(opts v1.ListOptions) (result *v1alpha1.AwsRoute53ZoneList, err error) {
	result = &v1alpha1.AwsRoute53ZoneList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsroute53zones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsRoute53Zones.
func (c *awsRoute53Zones) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsroute53zones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsRoute53Zone and creates it.  Returns the server's representation of the awsRoute53Zone, and an error, if there is any.
func (c *awsRoute53Zones) Create(awsRoute53Zone *v1alpha1.AwsRoute53Zone) (result *v1alpha1.AwsRoute53Zone, err error) {
	result = &v1alpha1.AwsRoute53Zone{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsroute53zones").
		Body(awsRoute53Zone).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsRoute53Zone and updates it. Returns the server's representation of the awsRoute53Zone, and an error, if there is any.
func (c *awsRoute53Zones) Update(awsRoute53Zone *v1alpha1.AwsRoute53Zone) (result *v1alpha1.AwsRoute53Zone, err error) {
	result = &v1alpha1.AwsRoute53Zone{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsroute53zones").
		Name(awsRoute53Zone.Name).
		Body(awsRoute53Zone).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsRoute53Zone and deletes it. Returns an error if one occurs.
func (c *awsRoute53Zones) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsroute53zones").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsRoute53Zones) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsroute53zones").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsRoute53Zone.
func (c *awsRoute53Zones) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRoute53Zone, err error) {
	result = &v1alpha1.AwsRoute53Zone{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsroute53zones").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
