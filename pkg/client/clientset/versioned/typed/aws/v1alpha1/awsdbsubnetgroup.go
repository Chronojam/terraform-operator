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

// AwsDbSubnetGroupsGetter has a method to return a AwsDbSubnetGroupInterface.
// A group's client should implement this interface.
type AwsDbSubnetGroupsGetter interface {
	AwsDbSubnetGroups(namespace string) AwsDbSubnetGroupInterface
}

// AwsDbSubnetGroupInterface has methods to work with AwsDbSubnetGroup resources.
type AwsDbSubnetGroupInterface interface {
	Create(*v1alpha1.AwsDbSubnetGroup) (*v1alpha1.AwsDbSubnetGroup, error)
	Update(*v1alpha1.AwsDbSubnetGroup) (*v1alpha1.AwsDbSubnetGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDbSubnetGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDbSubnetGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDbSubnetGroup, err error)
	AwsDbSubnetGroupExpansion
}

// awsDbSubnetGroups implements AwsDbSubnetGroupInterface
type awsDbSubnetGroups struct {
	client rest.Interface
	ns     string
}

// newAwsDbSubnetGroups returns a AwsDbSubnetGroups
func newAwsDbSubnetGroups(c *ChronojamV1alpha1Client, namespace string) *awsDbSubnetGroups {
	return &awsDbSubnetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDbSubnetGroup, and returns the corresponding awsDbSubnetGroup object, and an error if there is any.
func (c *awsDbSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDbSubnetGroup, err error) {
	result = &v1alpha1.AwsDbSubnetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDbSubnetGroups that match those selectors.
func (c *awsDbSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsDbSubnetGroupList, err error) {
	result = &v1alpha1.AwsDbSubnetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDbSubnetGroups.
func (c *awsDbSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDbSubnetGroup and creates it.  Returns the server's representation of the awsDbSubnetGroup, and an error, if there is any.
func (c *awsDbSubnetGroups) Create(awsDbSubnetGroup *v1alpha1.AwsDbSubnetGroup) (result *v1alpha1.AwsDbSubnetGroup, err error) {
	result = &v1alpha1.AwsDbSubnetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		Body(awsDbSubnetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDbSubnetGroup and updates it. Returns the server's representation of the awsDbSubnetGroup, and an error, if there is any.
func (c *awsDbSubnetGroups) Update(awsDbSubnetGroup *v1alpha1.AwsDbSubnetGroup) (result *v1alpha1.AwsDbSubnetGroup, err error) {
	result = &v1alpha1.AwsDbSubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		Name(awsDbSubnetGroup.Name).
		Body(awsDbSubnetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDbSubnetGroup and deletes it. Returns an error if one occurs.
func (c *awsDbSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDbSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDbSubnetGroup.
func (c *awsDbSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDbSubnetGroup, err error) {
	result = &v1alpha1.AwsDbSubnetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdbsubnetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
