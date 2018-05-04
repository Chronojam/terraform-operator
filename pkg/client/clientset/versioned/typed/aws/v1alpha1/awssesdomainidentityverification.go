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

// AwsSesDomainIdentityVerificationsGetter has a method to return a AwsSesDomainIdentityVerificationInterface.
// A group's client should implement this interface.
type AwsSesDomainIdentityVerificationsGetter interface {
	AwsSesDomainIdentityVerifications(namespace string) AwsSesDomainIdentityVerificationInterface
}

// AwsSesDomainIdentityVerificationInterface has methods to work with AwsSesDomainIdentityVerification resources.
type AwsSesDomainIdentityVerificationInterface interface {
	Create(*v1alpha1.AwsSesDomainIdentityVerification) (*v1alpha1.AwsSesDomainIdentityVerification, error)
	Update(*v1alpha1.AwsSesDomainIdentityVerification) (*v1alpha1.AwsSesDomainIdentityVerification, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSesDomainIdentityVerification, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSesDomainIdentityVerificationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesDomainIdentityVerification, err error)
	AwsSesDomainIdentityVerificationExpansion
}

// awsSesDomainIdentityVerifications implements AwsSesDomainIdentityVerificationInterface
type awsSesDomainIdentityVerifications struct {
	client rest.Interface
	ns     string
}

// newAwsSesDomainIdentityVerifications returns a AwsSesDomainIdentityVerifications
func newAwsSesDomainIdentityVerifications(c *ChronojamV1alpha1Client, namespace string) *awsSesDomainIdentityVerifications {
	return &awsSesDomainIdentityVerifications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSesDomainIdentityVerification, and returns the corresponding awsSesDomainIdentityVerification object, and an error if there is any.
func (c *awsSesDomainIdentityVerifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	result = &v1alpha1.AwsSesDomainIdentityVerification{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSesDomainIdentityVerifications that match those selectors.
func (c *awsSesDomainIdentityVerifications) List(opts v1.ListOptions) (result *v1alpha1.AwsSesDomainIdentityVerificationList, err error) {
	result = &v1alpha1.AwsSesDomainIdentityVerificationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSesDomainIdentityVerifications.
func (c *awsSesDomainIdentityVerifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSesDomainIdentityVerification and creates it.  Returns the server's representation of the awsSesDomainIdentityVerification, and an error, if there is any.
func (c *awsSesDomainIdentityVerifications) Create(awsSesDomainIdentityVerification *v1alpha1.AwsSesDomainIdentityVerification) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	result = &v1alpha1.AwsSesDomainIdentityVerification{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		Body(awsSesDomainIdentityVerification).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSesDomainIdentityVerification and updates it. Returns the server's representation of the awsSesDomainIdentityVerification, and an error, if there is any.
func (c *awsSesDomainIdentityVerifications) Update(awsSesDomainIdentityVerification *v1alpha1.AwsSesDomainIdentityVerification) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	result = &v1alpha1.AwsSesDomainIdentityVerification{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		Name(awsSesDomainIdentityVerification.Name).
		Body(awsSesDomainIdentityVerification).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSesDomainIdentityVerification and deletes it. Returns an error if one occurs.
func (c *awsSesDomainIdentityVerifications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSesDomainIdentityVerifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSesDomainIdentityVerification.
func (c *awsSesDomainIdentityVerifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	result = &v1alpha1.AwsSesDomainIdentityVerification{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssesdomainidentityverifications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
