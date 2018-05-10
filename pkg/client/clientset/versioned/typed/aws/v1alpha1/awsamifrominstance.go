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

// AwsAmiFromInstancesGetter has a method to return a AwsAmiFromInstanceInterface.
// A group's client should implement this interface.
type AwsAmiFromInstancesGetter interface {
	AwsAmiFromInstances(namespace string) AwsAmiFromInstanceInterface
}

// AwsAmiFromInstanceInterface has methods to work with AwsAmiFromInstance resources.
type AwsAmiFromInstanceInterface interface {
	Create(*v1alpha1.AwsAmiFromInstance) (*v1alpha1.AwsAmiFromInstance, error)
	Update(*v1alpha1.AwsAmiFromInstance) (*v1alpha1.AwsAmiFromInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAmiFromInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAmiFromInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmiFromInstance, err error)
	AwsAmiFromInstanceExpansion
}

// awsAmiFromInstances implements AwsAmiFromInstanceInterface
type awsAmiFromInstances struct {
	client rest.Interface
	ns     string
}

// newAwsAmiFromInstances returns a AwsAmiFromInstances
func newAwsAmiFromInstances(c *ChronojamV1alpha1Client, namespace string) *awsAmiFromInstances {
	return &awsAmiFromInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAmiFromInstance, and returns the corresponding awsAmiFromInstance object, and an error if there is any.
func (c *awsAmiFromInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAmiFromInstance, err error) {
	result = &v1alpha1.AwsAmiFromInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsamifrominstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAmiFromInstances that match those selectors.
func (c *awsAmiFromInstances) List(opts v1.ListOptions) (result *v1alpha1.AwsAmiFromInstanceList, err error) {
	result = &v1alpha1.AwsAmiFromInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsamifrominstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAmiFromInstances.
func (c *awsAmiFromInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsamifrominstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAmiFromInstance and creates it.  Returns the server's representation of the awsAmiFromInstance, and an error, if there is any.
func (c *awsAmiFromInstances) Create(awsAmiFromInstance *v1alpha1.AwsAmiFromInstance) (result *v1alpha1.AwsAmiFromInstance, err error) {
	result = &v1alpha1.AwsAmiFromInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsamifrominstances").
		Body(awsAmiFromInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAmiFromInstance and updates it. Returns the server's representation of the awsAmiFromInstance, and an error, if there is any.
func (c *awsAmiFromInstances) Update(awsAmiFromInstance *v1alpha1.AwsAmiFromInstance) (result *v1alpha1.AwsAmiFromInstance, err error) {
	result = &v1alpha1.AwsAmiFromInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsamifrominstances").
		Name(awsAmiFromInstance.Name).
		Body(awsAmiFromInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAmiFromInstance and deletes it. Returns an error if one occurs.
func (c *awsAmiFromInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsamifrominstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAmiFromInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsamifrominstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAmiFromInstance.
func (c *awsAmiFromInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmiFromInstance, err error) {
	result = &v1alpha1.AwsAmiFromInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsamifrominstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
