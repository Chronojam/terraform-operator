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

// AwsOrganizationsPoliciesGetter has a method to return a AwsOrganizationsPolicyInterface.
// A group's client should implement this interface.
type AwsOrganizationsPoliciesGetter interface {
	AwsOrganizationsPolicies(namespace string) AwsOrganizationsPolicyInterface
}

// AwsOrganizationsPolicyInterface has methods to work with AwsOrganizationsPolicy resources.
type AwsOrganizationsPolicyInterface interface {
	Create(*v1alpha1.AwsOrganizationsPolicy) (*v1alpha1.AwsOrganizationsPolicy, error)
	Update(*v1alpha1.AwsOrganizationsPolicy) (*v1alpha1.AwsOrganizationsPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOrganizationsPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOrganizationsPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOrganizationsPolicy, err error)
	AwsOrganizationsPolicyExpansion
}

// awsOrganizationsPolicies implements AwsOrganizationsPolicyInterface
type awsOrganizationsPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsOrganizationsPolicies returns a AwsOrganizationsPolicies
func newAwsOrganizationsPolicies(c *AwsV1alpha1Client, namespace string) *awsOrganizationsPolicies {
	return &awsOrganizationsPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOrganizationsPolicy, and returns the corresponding awsOrganizationsPolicy object, and an error if there is any.
func (c *awsOrganizationsPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOrganizationsPolicy, err error) {
	result = &v1alpha1.AwsOrganizationsPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOrganizationsPolicies that match those selectors.
func (c *awsOrganizationsPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsOrganizationsPolicyList, err error) {
	result = &v1alpha1.AwsOrganizationsPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOrganizationsPolicies.
func (c *awsOrganizationsPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOrganizationsPolicy and creates it.  Returns the server's representation of the awsOrganizationsPolicy, and an error, if there is any.
func (c *awsOrganizationsPolicies) Create(awsOrganizationsPolicy *v1alpha1.AwsOrganizationsPolicy) (result *v1alpha1.AwsOrganizationsPolicy, err error) {
	result = &v1alpha1.AwsOrganizationsPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		Body(awsOrganizationsPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOrganizationsPolicy and updates it. Returns the server's representation of the awsOrganizationsPolicy, and an error, if there is any.
func (c *awsOrganizationsPolicies) Update(awsOrganizationsPolicy *v1alpha1.AwsOrganizationsPolicy) (result *v1alpha1.AwsOrganizationsPolicy, err error) {
	result = &v1alpha1.AwsOrganizationsPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		Name(awsOrganizationsPolicy.Name).
		Body(awsOrganizationsPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOrganizationsPolicy and deletes it. Returns an error if one occurs.
func (c *awsOrganizationsPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOrganizationsPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOrganizationsPolicy.
func (c *awsOrganizationsPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOrganizationsPolicy, err error) {
	result = &v1alpha1.AwsOrganizationsPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsorganizationspolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
