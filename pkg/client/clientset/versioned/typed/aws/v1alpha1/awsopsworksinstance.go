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

// AwsOpsworksInstancesGetter has a method to return a AwsOpsworksInstanceInterface.
// A group's client should implement this interface.
type AwsOpsworksInstancesGetter interface {
	AwsOpsworksInstances(namespace string) AwsOpsworksInstanceInterface
}

// AwsOpsworksInstanceInterface has methods to work with AwsOpsworksInstance resources.
type AwsOpsworksInstanceInterface interface {
	Create(*v1alpha1.AwsOpsworksInstance) (*v1alpha1.AwsOpsworksInstance, error)
	Update(*v1alpha1.AwsOpsworksInstance) (*v1alpha1.AwsOpsworksInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOpsworksInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOpsworksInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksInstance, err error)
	AwsOpsworksInstanceExpansion
}

// awsOpsworksInstances implements AwsOpsworksInstanceInterface
type awsOpsworksInstances struct {
	client rest.Interface
	ns     string
}

// newAwsOpsworksInstances returns a AwsOpsworksInstances
func newAwsOpsworksInstances(c *ChronojamV1alpha1Client, namespace string) *awsOpsworksInstances {
	return &awsOpsworksInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOpsworksInstance, and returns the corresponding awsOpsworksInstance object, and an error if there is any.
func (c *awsOpsworksInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksInstance, err error) {
	result = &v1alpha1.AwsOpsworksInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksInstances that match those selectors.
func (c *awsOpsworksInstances) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksInstanceList, err error) {
	result = &v1alpha1.AwsOpsworksInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksInstances.
func (c *awsOpsworksInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOpsworksInstance and creates it.  Returns the server's representation of the awsOpsworksInstance, and an error, if there is any.
func (c *awsOpsworksInstances) Create(awsOpsworksInstance *v1alpha1.AwsOpsworksInstance) (result *v1alpha1.AwsOpsworksInstance, err error) {
	result = &v1alpha1.AwsOpsworksInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		Body(awsOpsworksInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksInstance and updates it. Returns the server's representation of the awsOpsworksInstance, and an error, if there is any.
func (c *awsOpsworksInstances) Update(awsOpsworksInstance *v1alpha1.AwsOpsworksInstance) (result *v1alpha1.AwsOpsworksInstance, err error) {
	result = &v1alpha1.AwsOpsworksInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		Name(awsOpsworksInstance.Name).
		Body(awsOpsworksInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksInstance and deletes it. Returns an error if one occurs.
func (c *awsOpsworksInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksInstance.
func (c *awsOpsworksInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksInstance, err error) {
	result = &v1alpha1.AwsOpsworksInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsopsworksinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
