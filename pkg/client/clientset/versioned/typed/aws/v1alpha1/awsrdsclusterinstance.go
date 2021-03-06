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

// AwsRdsClusterInstancesGetter has a method to return a AwsRdsClusterInstanceInterface.
// A group's client should implement this interface.
type AwsRdsClusterInstancesGetter interface {
	AwsRdsClusterInstances(namespace string) AwsRdsClusterInstanceInterface
}

// AwsRdsClusterInstanceInterface has methods to work with AwsRdsClusterInstance resources.
type AwsRdsClusterInstanceInterface interface {
	Create(*v1alpha1.AwsRdsClusterInstance) (*v1alpha1.AwsRdsClusterInstance, error)
	Update(*v1alpha1.AwsRdsClusterInstance) (*v1alpha1.AwsRdsClusterInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsRdsClusterInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsRdsClusterInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRdsClusterInstance, err error)
	AwsRdsClusterInstanceExpansion
}

// awsRdsClusterInstances implements AwsRdsClusterInstanceInterface
type awsRdsClusterInstances struct {
	client rest.Interface
	ns     string
}

// newAwsRdsClusterInstances returns a AwsRdsClusterInstances
func newAwsRdsClusterInstances(c *ChronojamV1alpha1Client, namespace string) *awsRdsClusterInstances {
	return &awsRdsClusterInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsRdsClusterInstance, and returns the corresponding awsRdsClusterInstance object, and an error if there is any.
func (c *awsRdsClusterInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRdsClusterInstance, err error) {
	result = &v1alpha1.AwsRdsClusterInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsRdsClusterInstances that match those selectors.
func (c *awsRdsClusterInstances) List(opts v1.ListOptions) (result *v1alpha1.AwsRdsClusterInstanceList, err error) {
	result = &v1alpha1.AwsRdsClusterInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsRdsClusterInstances.
func (c *awsRdsClusterInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsRdsClusterInstance and creates it.  Returns the server's representation of the awsRdsClusterInstance, and an error, if there is any.
func (c *awsRdsClusterInstances) Create(awsRdsClusterInstance *v1alpha1.AwsRdsClusterInstance) (result *v1alpha1.AwsRdsClusterInstance, err error) {
	result = &v1alpha1.AwsRdsClusterInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		Body(awsRdsClusterInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsRdsClusterInstance and updates it. Returns the server's representation of the awsRdsClusterInstance, and an error, if there is any.
func (c *awsRdsClusterInstances) Update(awsRdsClusterInstance *v1alpha1.AwsRdsClusterInstance) (result *v1alpha1.AwsRdsClusterInstance, err error) {
	result = &v1alpha1.AwsRdsClusterInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		Name(awsRdsClusterInstance.Name).
		Body(awsRdsClusterInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsRdsClusterInstance and deletes it. Returns an error if one occurs.
func (c *awsRdsClusterInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsRdsClusterInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsRdsClusterInstance.
func (c *awsRdsClusterInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRdsClusterInstance, err error) {
	result = &v1alpha1.AwsRdsClusterInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsrdsclusterinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
