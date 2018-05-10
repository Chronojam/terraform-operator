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

// AwsSsmParametersGetter has a method to return a AwsSsmParameterInterface.
// A group's client should implement this interface.
type AwsSsmParametersGetter interface {
	AwsSsmParameters(namespace string) AwsSsmParameterInterface
}

// AwsSsmParameterInterface has methods to work with AwsSsmParameter resources.
type AwsSsmParameterInterface interface {
	Create(*v1alpha1.AwsSsmParameter) (*v1alpha1.AwsSsmParameter, error)
	Update(*v1alpha1.AwsSsmParameter) (*v1alpha1.AwsSsmParameter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSsmParameter, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSsmParameterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSsmParameter, err error)
	AwsSsmParameterExpansion
}

// awsSsmParameters implements AwsSsmParameterInterface
type awsSsmParameters struct {
	client rest.Interface
	ns     string
}

// newAwsSsmParameters returns a AwsSsmParameters
func newAwsSsmParameters(c *ChronojamV1alpha1Client, namespace string) *awsSsmParameters {
	return &awsSsmParameters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSsmParameter, and returns the corresponding awsSsmParameter object, and an error if there is any.
func (c *awsSsmParameters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSsmParameter, err error) {
	result = &v1alpha1.AwsSsmParameter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsssmparameters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSsmParameters that match those selectors.
func (c *awsSsmParameters) List(opts v1.ListOptions) (result *v1alpha1.AwsSsmParameterList, err error) {
	result = &v1alpha1.AwsSsmParameterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsssmparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSsmParameters.
func (c *awsSsmParameters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsssmparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSsmParameter and creates it.  Returns the server's representation of the awsSsmParameter, and an error, if there is any.
func (c *awsSsmParameters) Create(awsSsmParameter *v1alpha1.AwsSsmParameter) (result *v1alpha1.AwsSsmParameter, err error) {
	result = &v1alpha1.AwsSsmParameter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsssmparameters").
		Body(awsSsmParameter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSsmParameter and updates it. Returns the server's representation of the awsSsmParameter, and an error, if there is any.
func (c *awsSsmParameters) Update(awsSsmParameter *v1alpha1.AwsSsmParameter) (result *v1alpha1.AwsSsmParameter, err error) {
	result = &v1alpha1.AwsSsmParameter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsssmparameters").
		Name(awsSsmParameter.Name).
		Body(awsSsmParameter).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSsmParameter and deletes it. Returns an error if one occurs.
func (c *awsSsmParameters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsssmparameters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSsmParameters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsssmparameters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSsmParameter.
func (c *awsSsmParameters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSsmParameter, err error) {
	result = &v1alpha1.AwsSsmParameter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsssmparameters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
