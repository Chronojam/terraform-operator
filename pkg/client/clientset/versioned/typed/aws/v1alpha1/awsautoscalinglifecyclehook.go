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

// AwsAutoscalingLifecycleHooksGetter has a method to return a AwsAutoscalingLifecycleHookInterface.
// A group's client should implement this interface.
type AwsAutoscalingLifecycleHooksGetter interface {
	AwsAutoscalingLifecycleHooks(namespace string) AwsAutoscalingLifecycleHookInterface
}

// AwsAutoscalingLifecycleHookInterface has methods to work with AwsAutoscalingLifecycleHook resources.
type AwsAutoscalingLifecycleHookInterface interface {
	Create(*v1alpha1.AwsAutoscalingLifecycleHook) (*v1alpha1.AwsAutoscalingLifecycleHook, error)
	Update(*v1alpha1.AwsAutoscalingLifecycleHook) (*v1alpha1.AwsAutoscalingLifecycleHook, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAutoscalingLifecycleHook, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAutoscalingLifecycleHookList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error)
	AwsAutoscalingLifecycleHookExpansion
}

// awsAutoscalingLifecycleHooks implements AwsAutoscalingLifecycleHookInterface
type awsAutoscalingLifecycleHooks struct {
	client rest.Interface
	ns     string
}

// newAwsAutoscalingLifecycleHooks returns a AwsAutoscalingLifecycleHooks
func newAwsAutoscalingLifecycleHooks(c *ChronojamV1alpha1Client, namespace string) *awsAutoscalingLifecycleHooks {
	return &awsAutoscalingLifecycleHooks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAutoscalingLifecycleHook, and returns the corresponding awsAutoscalingLifecycleHook object, and an error if there is any.
func (c *awsAutoscalingLifecycleHooks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	result = &v1alpha1.AwsAutoscalingLifecycleHook{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAutoscalingLifecycleHooks that match those selectors.
func (c *awsAutoscalingLifecycleHooks) List(opts v1.ListOptions) (result *v1alpha1.AwsAutoscalingLifecycleHookList, err error) {
	result = &v1alpha1.AwsAutoscalingLifecycleHookList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingLifecycleHooks.
func (c *awsAutoscalingLifecycleHooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAutoscalingLifecycleHook and creates it.  Returns the server's representation of the awsAutoscalingLifecycleHook, and an error, if there is any.
func (c *awsAutoscalingLifecycleHooks) Create(awsAutoscalingLifecycleHook *v1alpha1.AwsAutoscalingLifecycleHook) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	result = &v1alpha1.AwsAutoscalingLifecycleHook{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		Body(awsAutoscalingLifecycleHook).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAutoscalingLifecycleHook and updates it. Returns the server's representation of the awsAutoscalingLifecycleHook, and an error, if there is any.
func (c *awsAutoscalingLifecycleHooks) Update(awsAutoscalingLifecycleHook *v1alpha1.AwsAutoscalingLifecycleHook) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	result = &v1alpha1.AwsAutoscalingLifecycleHook{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		Name(awsAutoscalingLifecycleHook.Name).
		Body(awsAutoscalingLifecycleHook).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAutoscalingLifecycleHook and deletes it. Returns an error if one occurs.
func (c *awsAutoscalingLifecycleHooks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAutoscalingLifecycleHooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAutoscalingLifecycleHook.
func (c *awsAutoscalingLifecycleHooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	result = &v1alpha1.AwsAutoscalingLifecycleHook{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsautoscalinglifecyclehooks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
