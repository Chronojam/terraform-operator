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

// AwsElasticacheReplicationGroupsGetter has a method to return a AwsElasticacheReplicationGroupInterface.
// A group's client should implement this interface.
type AwsElasticacheReplicationGroupsGetter interface {
	AwsElasticacheReplicationGroups(namespace string) AwsElasticacheReplicationGroupInterface
}

// AwsElasticacheReplicationGroupInterface has methods to work with AwsElasticacheReplicationGroup resources.
type AwsElasticacheReplicationGroupInterface interface {
	Create(*v1alpha1.AwsElasticacheReplicationGroup) (*v1alpha1.AwsElasticacheReplicationGroup, error)
	Update(*v1alpha1.AwsElasticacheReplicationGroup) (*v1alpha1.AwsElasticacheReplicationGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsElasticacheReplicationGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsElasticacheReplicationGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticacheReplicationGroup, err error)
	AwsElasticacheReplicationGroupExpansion
}

// awsElasticacheReplicationGroups implements AwsElasticacheReplicationGroupInterface
type awsElasticacheReplicationGroups struct {
	client rest.Interface
	ns     string
}

// newAwsElasticacheReplicationGroups returns a AwsElasticacheReplicationGroups
func newAwsElasticacheReplicationGroups(c *ChronojamV1alpha1Client, namespace string) *awsElasticacheReplicationGroups {
	return &awsElasticacheReplicationGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsElasticacheReplicationGroup, and returns the corresponding awsElasticacheReplicationGroup object, and an error if there is any.
func (c *awsElasticacheReplicationGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticacheReplicationGroup, err error) {
	result = &v1alpha1.AwsElasticacheReplicationGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsElasticacheReplicationGroups that match those selectors.
func (c *awsElasticacheReplicationGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticacheReplicationGroupList, err error) {
	result = &v1alpha1.AwsElasticacheReplicationGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsElasticacheReplicationGroups.
func (c *awsElasticacheReplicationGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsElasticacheReplicationGroup and creates it.  Returns the server's representation of the awsElasticacheReplicationGroup, and an error, if there is any.
func (c *awsElasticacheReplicationGroups) Create(awsElasticacheReplicationGroup *v1alpha1.AwsElasticacheReplicationGroup) (result *v1alpha1.AwsElasticacheReplicationGroup, err error) {
	result = &v1alpha1.AwsElasticacheReplicationGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		Body(awsElasticacheReplicationGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsElasticacheReplicationGroup and updates it. Returns the server's representation of the awsElasticacheReplicationGroup, and an error, if there is any.
func (c *awsElasticacheReplicationGroups) Update(awsElasticacheReplicationGroup *v1alpha1.AwsElasticacheReplicationGroup) (result *v1alpha1.AwsElasticacheReplicationGroup, err error) {
	result = &v1alpha1.AwsElasticacheReplicationGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		Name(awsElasticacheReplicationGroup.Name).
		Body(awsElasticacheReplicationGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsElasticacheReplicationGroup and deletes it. Returns an error if one occurs.
func (c *awsElasticacheReplicationGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsElasticacheReplicationGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsElasticacheReplicationGroup.
func (c *awsElasticacheReplicationGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticacheReplicationGroup, err error) {
	result = &v1alpha1.AwsElasticacheReplicationGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awselasticachereplicationgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
