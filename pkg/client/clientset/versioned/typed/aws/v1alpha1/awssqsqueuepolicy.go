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

// AwsSqsQueuePoliciesGetter has a method to return a AwsSqsQueuePolicyInterface.
// A group's client should implement this interface.
type AwsSqsQueuePoliciesGetter interface {
	AwsSqsQueuePolicies(namespace string) AwsSqsQueuePolicyInterface
}

// AwsSqsQueuePolicyInterface has methods to work with AwsSqsQueuePolicy resources.
type AwsSqsQueuePolicyInterface interface {
	Create(*v1alpha1.AwsSqsQueuePolicy) (*v1alpha1.AwsSqsQueuePolicy, error)
	Update(*v1alpha1.AwsSqsQueuePolicy) (*v1alpha1.AwsSqsQueuePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSqsQueuePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSqsQueuePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSqsQueuePolicy, err error)
	AwsSqsQueuePolicyExpansion
}

// awsSqsQueuePolicies implements AwsSqsQueuePolicyInterface
type awsSqsQueuePolicies struct {
	client rest.Interface
	ns     string
}

// newAwsSqsQueuePolicies returns a AwsSqsQueuePolicies
func newAwsSqsQueuePolicies(c *ChronojamV1alpha1Client, namespace string) *awsSqsQueuePolicies {
	return &awsSqsQueuePolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSqsQueuePolicy, and returns the corresponding awsSqsQueuePolicy object, and an error if there is any.
func (c *awsSqsQueuePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSqsQueuePolicy, err error) {
	result = &v1alpha1.AwsSqsQueuePolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSqsQueuePolicies that match those selectors.
func (c *awsSqsQueuePolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsSqsQueuePolicyList, err error) {
	result = &v1alpha1.AwsSqsQueuePolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSqsQueuePolicies.
func (c *awsSqsQueuePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSqsQueuePolicy and creates it.  Returns the server's representation of the awsSqsQueuePolicy, and an error, if there is any.
func (c *awsSqsQueuePolicies) Create(awsSqsQueuePolicy *v1alpha1.AwsSqsQueuePolicy) (result *v1alpha1.AwsSqsQueuePolicy, err error) {
	result = &v1alpha1.AwsSqsQueuePolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		Body(awsSqsQueuePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSqsQueuePolicy and updates it. Returns the server's representation of the awsSqsQueuePolicy, and an error, if there is any.
func (c *awsSqsQueuePolicies) Update(awsSqsQueuePolicy *v1alpha1.AwsSqsQueuePolicy) (result *v1alpha1.AwsSqsQueuePolicy, err error) {
	result = &v1alpha1.AwsSqsQueuePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		Name(awsSqsQueuePolicy.Name).
		Body(awsSqsQueuePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSqsQueuePolicy and deletes it. Returns an error if one occurs.
func (c *awsSqsQueuePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSqsQueuePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSqsQueuePolicy.
func (c *awsSqsQueuePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSqsQueuePolicy, err error) {
	result = &v1alpha1.AwsSqsQueuePolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssqsqueuepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
