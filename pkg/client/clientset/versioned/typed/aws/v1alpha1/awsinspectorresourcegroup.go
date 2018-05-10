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

// AwsInspectorResourceGroupsGetter has a method to return a AwsInspectorResourceGroupInterface.
// A group's client should implement this interface.
type AwsInspectorResourceGroupsGetter interface {
	AwsInspectorResourceGroups(namespace string) AwsInspectorResourceGroupInterface
}

// AwsInspectorResourceGroupInterface has methods to work with AwsInspectorResourceGroup resources.
type AwsInspectorResourceGroupInterface interface {
	Create(*v1alpha1.AwsInspectorResourceGroup) (*v1alpha1.AwsInspectorResourceGroup, error)
	Update(*v1alpha1.AwsInspectorResourceGroup) (*v1alpha1.AwsInspectorResourceGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsInspectorResourceGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsInspectorResourceGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsInspectorResourceGroup, err error)
	AwsInspectorResourceGroupExpansion
}

// awsInspectorResourceGroups implements AwsInspectorResourceGroupInterface
type awsInspectorResourceGroups struct {
	client rest.Interface
	ns     string
}

// newAwsInspectorResourceGroups returns a AwsInspectorResourceGroups
func newAwsInspectorResourceGroups(c *ChronojamV1alpha1Client, namespace string) *awsInspectorResourceGroups {
	return &awsInspectorResourceGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsInspectorResourceGroup, and returns the corresponding awsInspectorResourceGroup object, and an error if there is any.
func (c *awsInspectorResourceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	result = &v1alpha1.AwsInspectorResourceGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsInspectorResourceGroups that match those selectors.
func (c *awsInspectorResourceGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsInspectorResourceGroupList, err error) {
	result = &v1alpha1.AwsInspectorResourceGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsInspectorResourceGroups.
func (c *awsInspectorResourceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsInspectorResourceGroup and creates it.  Returns the server's representation of the awsInspectorResourceGroup, and an error, if there is any.
func (c *awsInspectorResourceGroups) Create(awsInspectorResourceGroup *v1alpha1.AwsInspectorResourceGroup) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	result = &v1alpha1.AwsInspectorResourceGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		Body(awsInspectorResourceGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsInspectorResourceGroup and updates it. Returns the server's representation of the awsInspectorResourceGroup, and an error, if there is any.
func (c *awsInspectorResourceGroups) Update(awsInspectorResourceGroup *v1alpha1.AwsInspectorResourceGroup) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	result = &v1alpha1.AwsInspectorResourceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		Name(awsInspectorResourceGroup.Name).
		Body(awsInspectorResourceGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsInspectorResourceGroup and deletes it. Returns an error if one occurs.
func (c *awsInspectorResourceGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsInspectorResourceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsInspectorResourceGroup.
func (c *awsInspectorResourceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	result = &v1alpha1.AwsInspectorResourceGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsinspectorresourcegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
