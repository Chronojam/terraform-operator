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

// AwsDmsEndpointsGetter has a method to return a AwsDmsEndpointInterface.
// A group's client should implement this interface.
type AwsDmsEndpointsGetter interface {
	AwsDmsEndpoints(namespace string) AwsDmsEndpointInterface
}

// AwsDmsEndpointInterface has methods to work with AwsDmsEndpoint resources.
type AwsDmsEndpointInterface interface {
	Create(*v1alpha1.AwsDmsEndpoint) (*v1alpha1.AwsDmsEndpoint, error)
	Update(*v1alpha1.AwsDmsEndpoint) (*v1alpha1.AwsDmsEndpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDmsEndpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDmsEndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDmsEndpoint, err error)
	AwsDmsEndpointExpansion
}

// awsDmsEndpoints implements AwsDmsEndpointInterface
type awsDmsEndpoints struct {
	client rest.Interface
	ns     string
}

// newAwsDmsEndpoints returns a AwsDmsEndpoints
func newAwsDmsEndpoints(c *ChronojamV1alpha1Client, namespace string) *awsDmsEndpoints {
	return &awsDmsEndpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDmsEndpoint, and returns the corresponding awsDmsEndpoint object, and an error if there is any.
func (c *awsDmsEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDmsEndpoint, err error) {
	result = &v1alpha1.AwsDmsEndpoint{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDmsEndpoints that match those selectors.
func (c *awsDmsEndpoints) List(opts v1.ListOptions) (result *v1alpha1.AwsDmsEndpointList, err error) {
	result = &v1alpha1.AwsDmsEndpointList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDmsEndpoints.
func (c *awsDmsEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDmsEndpoint and creates it.  Returns the server's representation of the awsDmsEndpoint, and an error, if there is any.
func (c *awsDmsEndpoints) Create(awsDmsEndpoint *v1alpha1.AwsDmsEndpoint) (result *v1alpha1.AwsDmsEndpoint, err error) {
	result = &v1alpha1.AwsDmsEndpoint{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		Body(awsDmsEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDmsEndpoint and updates it. Returns the server's representation of the awsDmsEndpoint, and an error, if there is any.
func (c *awsDmsEndpoints) Update(awsDmsEndpoint *v1alpha1.AwsDmsEndpoint) (result *v1alpha1.AwsDmsEndpoint, err error) {
	result = &v1alpha1.AwsDmsEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		Name(awsDmsEndpoint.Name).
		Body(awsDmsEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDmsEndpoint and deletes it. Returns an error if one occurs.
func (c *awsDmsEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDmsEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDmsEndpoint.
func (c *awsDmsEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDmsEndpoint, err error) {
	result = &v1alpha1.AwsDmsEndpoint{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdmsendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
