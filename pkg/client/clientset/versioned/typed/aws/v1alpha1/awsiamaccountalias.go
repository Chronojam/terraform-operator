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

// AwsIamAccountAliasesGetter has a method to return a AwsIamAccountAliasInterface.
// A group's client should implement this interface.
type AwsIamAccountAliasesGetter interface {
	AwsIamAccountAliases(namespace string) AwsIamAccountAliasInterface
}

// AwsIamAccountAliasInterface has methods to work with AwsIamAccountAlias resources.
type AwsIamAccountAliasInterface interface {
	Create(*v1alpha1.AwsIamAccountAlias) (*v1alpha1.AwsIamAccountAlias, error)
	Update(*v1alpha1.AwsIamAccountAlias) (*v1alpha1.AwsIamAccountAlias, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIamAccountAlias, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIamAccountAliasList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamAccountAlias, err error)
	AwsIamAccountAliasExpansion
}

// awsIamAccountAliases implements AwsIamAccountAliasInterface
type awsIamAccountAliases struct {
	client rest.Interface
	ns     string
}

// newAwsIamAccountAliases returns a AwsIamAccountAliases
func newAwsIamAccountAliases(c *ChronojamV1alpha1Client, namespace string) *awsIamAccountAliases {
	return &awsIamAccountAliases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamAccountAlias, and returns the corresponding awsIamAccountAlias object, and an error if there is any.
func (c *awsIamAccountAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamAccountAlias, err error) {
	result = &v1alpha1.AwsIamAccountAlias{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamAccountAliases that match those selectors.
func (c *awsIamAccountAliases) List(opts v1.ListOptions) (result *v1alpha1.AwsIamAccountAliasList, err error) {
	result = &v1alpha1.AwsIamAccountAliasList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamAccountAliases.
func (c *awsIamAccountAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamAccountAlias and creates it.  Returns the server's representation of the awsIamAccountAlias, and an error, if there is any.
func (c *awsIamAccountAliases) Create(awsIamAccountAlias *v1alpha1.AwsIamAccountAlias) (result *v1alpha1.AwsIamAccountAlias, err error) {
	result = &v1alpha1.AwsIamAccountAlias{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		Body(awsIamAccountAlias).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamAccountAlias and updates it. Returns the server's representation of the awsIamAccountAlias, and an error, if there is any.
func (c *awsIamAccountAliases) Update(awsIamAccountAlias *v1alpha1.AwsIamAccountAlias) (result *v1alpha1.AwsIamAccountAlias, err error) {
	result = &v1alpha1.AwsIamAccountAlias{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		Name(awsIamAccountAlias.Name).
		Body(awsIamAccountAlias).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamAccountAlias and deletes it. Returns an error if one occurs.
func (c *awsIamAccountAliases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamAccountAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamAccountAlias.
func (c *awsIamAccountAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamAccountAlias, err error) {
	result = &v1alpha1.AwsIamAccountAlias{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamaccountaliases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
