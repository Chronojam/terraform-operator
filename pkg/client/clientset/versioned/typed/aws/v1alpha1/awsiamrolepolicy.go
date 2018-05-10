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

// AwsIamRolePoliciesGetter has a method to return a AwsIamRolePolicyInterface.
// A group's client should implement this interface.
type AwsIamRolePoliciesGetter interface {
	AwsIamRolePolicies(namespace string) AwsIamRolePolicyInterface
}

// AwsIamRolePolicyInterface has methods to work with AwsIamRolePolicy resources.
type AwsIamRolePolicyInterface interface {
	Create(*v1alpha1.AwsIamRolePolicy) (*v1alpha1.AwsIamRolePolicy, error)
	Update(*v1alpha1.AwsIamRolePolicy) (*v1alpha1.AwsIamRolePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIamRolePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIamRolePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamRolePolicy, err error)
	AwsIamRolePolicyExpansion
}

// awsIamRolePolicies implements AwsIamRolePolicyInterface
type awsIamRolePolicies struct {
	client rest.Interface
	ns     string
}

// newAwsIamRolePolicies returns a AwsIamRolePolicies
func newAwsIamRolePolicies(c *ChronojamV1alpha1Client, namespace string) *awsIamRolePolicies {
	return &awsIamRolePolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamRolePolicy, and returns the corresponding awsIamRolePolicy object, and an error if there is any.
func (c *awsIamRolePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamRolePolicy, err error) {
	result = &v1alpha1.AwsIamRolePolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamRolePolicies that match those selectors.
func (c *awsIamRolePolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsIamRolePolicyList, err error) {
	result = &v1alpha1.AwsIamRolePolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamRolePolicies.
func (c *awsIamRolePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamRolePolicy and creates it.  Returns the server's representation of the awsIamRolePolicy, and an error, if there is any.
func (c *awsIamRolePolicies) Create(awsIamRolePolicy *v1alpha1.AwsIamRolePolicy) (result *v1alpha1.AwsIamRolePolicy, err error) {
	result = &v1alpha1.AwsIamRolePolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		Body(awsIamRolePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamRolePolicy and updates it. Returns the server's representation of the awsIamRolePolicy, and an error, if there is any.
func (c *awsIamRolePolicies) Update(awsIamRolePolicy *v1alpha1.AwsIamRolePolicy) (result *v1alpha1.AwsIamRolePolicy, err error) {
	result = &v1alpha1.AwsIamRolePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		Name(awsIamRolePolicy.Name).
		Body(awsIamRolePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamRolePolicy and deletes it. Returns an error if one occurs.
func (c *awsIamRolePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamRolePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamRolePolicy.
func (c *awsIamRolePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamRolePolicy, err error) {
	result = &v1alpha1.AwsIamRolePolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamrolepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
