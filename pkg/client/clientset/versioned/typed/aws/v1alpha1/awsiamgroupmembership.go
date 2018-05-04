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

// AwsIamGroupMembershipsGetter has a method to return a AwsIamGroupMembershipInterface.
// A group's client should implement this interface.
type AwsIamGroupMembershipsGetter interface {
	AwsIamGroupMemberships(namespace string) AwsIamGroupMembershipInterface
}

// AwsIamGroupMembershipInterface has methods to work with AwsIamGroupMembership resources.
type AwsIamGroupMembershipInterface interface {
	Create(*v1alpha1.AwsIamGroupMembership) (*v1alpha1.AwsIamGroupMembership, error)
	Update(*v1alpha1.AwsIamGroupMembership) (*v1alpha1.AwsIamGroupMembership, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIamGroupMembership, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIamGroupMembershipList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroupMembership, err error)
	AwsIamGroupMembershipExpansion
}

// awsIamGroupMemberships implements AwsIamGroupMembershipInterface
type awsIamGroupMemberships struct {
	client rest.Interface
	ns     string
}

// newAwsIamGroupMemberships returns a AwsIamGroupMemberships
func newAwsIamGroupMemberships(c *ChronojamV1alpha1Client, namespace string) *awsIamGroupMemberships {
	return &awsIamGroupMemberships{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamGroupMembership, and returns the corresponding awsIamGroupMembership object, and an error if there is any.
func (c *awsIamGroupMemberships) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamGroupMembership, err error) {
	result = &v1alpha1.AwsIamGroupMembership{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamGroupMemberships that match those selectors.
func (c *awsIamGroupMemberships) List(opts v1.ListOptions) (result *v1alpha1.AwsIamGroupMembershipList, err error) {
	result = &v1alpha1.AwsIamGroupMembershipList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamGroupMemberships.
func (c *awsIamGroupMemberships) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamGroupMembership and creates it.  Returns the server's representation of the awsIamGroupMembership, and an error, if there is any.
func (c *awsIamGroupMemberships) Create(awsIamGroupMembership *v1alpha1.AwsIamGroupMembership) (result *v1alpha1.AwsIamGroupMembership, err error) {
	result = &v1alpha1.AwsIamGroupMembership{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		Body(awsIamGroupMembership).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamGroupMembership and updates it. Returns the server's representation of the awsIamGroupMembership, and an error, if there is any.
func (c *awsIamGroupMemberships) Update(awsIamGroupMembership *v1alpha1.AwsIamGroupMembership) (result *v1alpha1.AwsIamGroupMembership, err error) {
	result = &v1alpha1.AwsIamGroupMembership{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		Name(awsIamGroupMembership.Name).
		Body(awsIamGroupMembership).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamGroupMembership and deletes it. Returns an error if one occurs.
func (c *awsIamGroupMemberships) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamGroupMemberships) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamGroupMembership.
func (c *awsIamGroupMemberships) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroupMembership, err error) {
	result = &v1alpha1.AwsIamGroupMembership{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamgroupmemberships").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
