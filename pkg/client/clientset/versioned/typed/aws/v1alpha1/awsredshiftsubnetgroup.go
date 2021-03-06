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

// AwsRedshiftSubnetGroupsGetter has a method to return a AwsRedshiftSubnetGroupInterface.
// A group's client should implement this interface.
type AwsRedshiftSubnetGroupsGetter interface {
	AwsRedshiftSubnetGroups(namespace string) AwsRedshiftSubnetGroupInterface
}

// AwsRedshiftSubnetGroupInterface has methods to work with AwsRedshiftSubnetGroup resources.
type AwsRedshiftSubnetGroupInterface interface {
	Create(*v1alpha1.AwsRedshiftSubnetGroup) (*v1alpha1.AwsRedshiftSubnetGroup, error)
	Update(*v1alpha1.AwsRedshiftSubnetGroup) (*v1alpha1.AwsRedshiftSubnetGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsRedshiftSubnetGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsRedshiftSubnetGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRedshiftSubnetGroup, err error)
	AwsRedshiftSubnetGroupExpansion
}

// awsRedshiftSubnetGroups implements AwsRedshiftSubnetGroupInterface
type awsRedshiftSubnetGroups struct {
	client rest.Interface
	ns     string
}

// newAwsRedshiftSubnetGroups returns a AwsRedshiftSubnetGroups
func newAwsRedshiftSubnetGroups(c *ChronojamV1alpha1Client, namespace string) *awsRedshiftSubnetGroups {
	return &awsRedshiftSubnetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsRedshiftSubnetGroup, and returns the corresponding awsRedshiftSubnetGroup object, and an error if there is any.
func (c *awsRedshiftSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	result = &v1alpha1.AwsRedshiftSubnetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsRedshiftSubnetGroups that match those selectors.
func (c *awsRedshiftSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsRedshiftSubnetGroupList, err error) {
	result = &v1alpha1.AwsRedshiftSubnetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsRedshiftSubnetGroups.
func (c *awsRedshiftSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsRedshiftSubnetGroup and creates it.  Returns the server's representation of the awsRedshiftSubnetGroup, and an error, if there is any.
func (c *awsRedshiftSubnetGroups) Create(awsRedshiftSubnetGroup *v1alpha1.AwsRedshiftSubnetGroup) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	result = &v1alpha1.AwsRedshiftSubnetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		Body(awsRedshiftSubnetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsRedshiftSubnetGroup and updates it. Returns the server's representation of the awsRedshiftSubnetGroup, and an error, if there is any.
func (c *awsRedshiftSubnetGroups) Update(awsRedshiftSubnetGroup *v1alpha1.AwsRedshiftSubnetGroup) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	result = &v1alpha1.AwsRedshiftSubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		Name(awsRedshiftSubnetGroup.Name).
		Body(awsRedshiftSubnetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsRedshiftSubnetGroup and deletes it. Returns an error if one occurs.
func (c *awsRedshiftSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsRedshiftSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsRedshiftSubnetGroup.
func (c *awsRedshiftSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	result = &v1alpha1.AwsRedshiftSubnetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsredshiftsubnetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
