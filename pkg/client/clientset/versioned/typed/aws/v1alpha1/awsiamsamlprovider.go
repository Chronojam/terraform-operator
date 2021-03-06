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

// AwsIamSamlProvidersGetter has a method to return a AwsIamSamlProviderInterface.
// A group's client should implement this interface.
type AwsIamSamlProvidersGetter interface {
	AwsIamSamlProviders(namespace string) AwsIamSamlProviderInterface
}

// AwsIamSamlProviderInterface has methods to work with AwsIamSamlProvider resources.
type AwsIamSamlProviderInterface interface {
	Create(*v1alpha1.AwsIamSamlProvider) (*v1alpha1.AwsIamSamlProvider, error)
	Update(*v1alpha1.AwsIamSamlProvider) (*v1alpha1.AwsIamSamlProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIamSamlProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIamSamlProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamSamlProvider, err error)
	AwsIamSamlProviderExpansion
}

// awsIamSamlProviders implements AwsIamSamlProviderInterface
type awsIamSamlProviders struct {
	client rest.Interface
	ns     string
}

// newAwsIamSamlProviders returns a AwsIamSamlProviders
func newAwsIamSamlProviders(c *ChronojamV1alpha1Client, namespace string) *awsIamSamlProviders {
	return &awsIamSamlProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamSamlProvider, and returns the corresponding awsIamSamlProvider object, and an error if there is any.
func (c *awsIamSamlProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamSamlProvider, err error) {
	result = &v1alpha1.AwsIamSamlProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamSamlProviders that match those selectors.
func (c *awsIamSamlProviders) List(opts v1.ListOptions) (result *v1alpha1.AwsIamSamlProviderList, err error) {
	result = &v1alpha1.AwsIamSamlProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamSamlProviders.
func (c *awsIamSamlProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamSamlProvider and creates it.  Returns the server's representation of the awsIamSamlProvider, and an error, if there is any.
func (c *awsIamSamlProviders) Create(awsIamSamlProvider *v1alpha1.AwsIamSamlProvider) (result *v1alpha1.AwsIamSamlProvider, err error) {
	result = &v1alpha1.AwsIamSamlProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		Body(awsIamSamlProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamSamlProvider and updates it. Returns the server's representation of the awsIamSamlProvider, and an error, if there is any.
func (c *awsIamSamlProviders) Update(awsIamSamlProvider *v1alpha1.AwsIamSamlProvider) (result *v1alpha1.AwsIamSamlProvider, err error) {
	result = &v1alpha1.AwsIamSamlProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		Name(awsIamSamlProvider.Name).
		Body(awsIamSamlProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamSamlProvider and deletes it. Returns an error if one occurs.
func (c *awsIamSamlProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamSamlProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamSamlProvider.
func (c *awsIamSamlProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamSamlProvider, err error) {
	result = &v1alpha1.AwsIamSamlProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamsamlproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
