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

// AwsOrganizationsAccountsGetter has a method to return a AwsOrganizationsAccountInterface.
// A group's client should implement this interface.
type AwsOrganizationsAccountsGetter interface {
	AwsOrganizationsAccounts(namespace string) AwsOrganizationsAccountInterface
}

// AwsOrganizationsAccountInterface has methods to work with AwsOrganizationsAccount resources.
type AwsOrganizationsAccountInterface interface {
	Create(*v1alpha1.AwsOrganizationsAccount) (*v1alpha1.AwsOrganizationsAccount, error)
	Update(*v1alpha1.AwsOrganizationsAccount) (*v1alpha1.AwsOrganizationsAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOrganizationsAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOrganizationsAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOrganizationsAccount, err error)
	AwsOrganizationsAccountExpansion
}

// awsOrganizationsAccounts implements AwsOrganizationsAccountInterface
type awsOrganizationsAccounts struct {
	client rest.Interface
	ns     string
}

// newAwsOrganizationsAccounts returns a AwsOrganizationsAccounts
func newAwsOrganizationsAccounts(c *ChronojamV1alpha1Client, namespace string) *awsOrganizationsAccounts {
	return &awsOrganizationsAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOrganizationsAccount, and returns the corresponding awsOrganizationsAccount object, and an error if there is any.
func (c *awsOrganizationsAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOrganizationsAccount, err error) {
	result = &v1alpha1.AwsOrganizationsAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOrganizationsAccounts that match those selectors.
func (c *awsOrganizationsAccounts) List(opts v1.ListOptions) (result *v1alpha1.AwsOrganizationsAccountList, err error) {
	result = &v1alpha1.AwsOrganizationsAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOrganizationsAccounts.
func (c *awsOrganizationsAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOrganizationsAccount and creates it.  Returns the server's representation of the awsOrganizationsAccount, and an error, if there is any.
func (c *awsOrganizationsAccounts) Create(awsOrganizationsAccount *v1alpha1.AwsOrganizationsAccount) (result *v1alpha1.AwsOrganizationsAccount, err error) {
	result = &v1alpha1.AwsOrganizationsAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		Body(awsOrganizationsAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOrganizationsAccount and updates it. Returns the server's representation of the awsOrganizationsAccount, and an error, if there is any.
func (c *awsOrganizationsAccounts) Update(awsOrganizationsAccount *v1alpha1.AwsOrganizationsAccount) (result *v1alpha1.AwsOrganizationsAccount, err error) {
	result = &v1alpha1.AwsOrganizationsAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		Name(awsOrganizationsAccount.Name).
		Body(awsOrganizationsAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOrganizationsAccount and deletes it. Returns an error if one occurs.
func (c *awsOrganizationsAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOrganizationsAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOrganizationsAccount.
func (c *awsOrganizationsAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOrganizationsAccount, err error) {
	result = &v1alpha1.AwsOrganizationsAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsorganizationsaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
