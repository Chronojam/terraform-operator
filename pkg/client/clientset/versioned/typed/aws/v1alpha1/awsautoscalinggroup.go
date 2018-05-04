/*
Copyright The Kubernetes Authors.

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

// AwsAutoscalingGroupsGetter has a method to return a AwsAutoscalingGroupInterface.
// A group's client should implement this interface.
type AwsAutoscalingGroupsGetter interface {
	AwsAutoscalingGroups(namespace string) AwsAutoscalingGroupInterface
}

// AwsAutoscalingGroupInterface has methods to work with AwsAutoscalingGroup resources.
type AwsAutoscalingGroupInterface interface {
	Create(*v1alpha1.AwsAutoscalingGroup) (*v1alpha1.AwsAutoscalingGroup, error)
	Update(*v1alpha1.AwsAutoscalingGroup) (*v1alpha1.AwsAutoscalingGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAutoscalingGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAutoscalingGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingGroup, err error)
	AwsAutoscalingGroupExpansion
}

// awsAutoscalingGroups implements AwsAutoscalingGroupInterface
type awsAutoscalingGroups struct {
	client rest.Interface
	ns     string
}

// newAwsAutoscalingGroups returns a AwsAutoscalingGroups
func newAwsAutoscalingGroups(c *AwsV1alpha1Client, namespace string) *awsAutoscalingGroups {
	return &awsAutoscalingGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAutoscalingGroup, and returns the corresponding awsAutoscalingGroup object, and an error if there is any.
func (c *awsAutoscalingGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	result = &v1alpha1.AwsAutoscalingGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAutoscalingGroups that match those selectors.
func (c *awsAutoscalingGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsAutoscalingGroupList, err error) {
	result = &v1alpha1.AwsAutoscalingGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingGroups.
func (c *awsAutoscalingGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAutoscalingGroup and creates it.  Returns the server's representation of the awsAutoscalingGroup, and an error, if there is any.
func (c *awsAutoscalingGroups) Create(awsAutoscalingGroup *v1alpha1.AwsAutoscalingGroup) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	result = &v1alpha1.AwsAutoscalingGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		Body(awsAutoscalingGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAutoscalingGroup and updates it. Returns the server's representation of the awsAutoscalingGroup, and an error, if there is any.
func (c *awsAutoscalingGroups) Update(awsAutoscalingGroup *v1alpha1.AwsAutoscalingGroup) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	result = &v1alpha1.AwsAutoscalingGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		Name(awsAutoscalingGroup.Name).
		Body(awsAutoscalingGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAutoscalingGroup and deletes it. Returns an error if one occurs.
func (c *awsAutoscalingGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAutoscalingGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAutoscalingGroup.
func (c *awsAutoscalingGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	result = &v1alpha1.AwsAutoscalingGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsautoscalinggroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
