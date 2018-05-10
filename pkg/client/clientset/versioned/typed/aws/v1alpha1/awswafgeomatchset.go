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

// AwsWafGeoMatchSetsGetter has a method to return a AwsWafGeoMatchSetInterface.
// A group's client should implement this interface.
type AwsWafGeoMatchSetsGetter interface {
	AwsWafGeoMatchSets(namespace string) AwsWafGeoMatchSetInterface
}

// AwsWafGeoMatchSetInterface has methods to work with AwsWafGeoMatchSet resources.
type AwsWafGeoMatchSetInterface interface {
	Create(*v1alpha1.AwsWafGeoMatchSet) (*v1alpha1.AwsWafGeoMatchSet, error)
	Update(*v1alpha1.AwsWafGeoMatchSet) (*v1alpha1.AwsWafGeoMatchSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafGeoMatchSet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafGeoMatchSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafGeoMatchSet, err error)
	AwsWafGeoMatchSetExpansion
}

// awsWafGeoMatchSets implements AwsWafGeoMatchSetInterface
type awsWafGeoMatchSets struct {
	client rest.Interface
	ns     string
}

// newAwsWafGeoMatchSets returns a AwsWafGeoMatchSets
func newAwsWafGeoMatchSets(c *ChronojamV1alpha1Client, namespace string) *awsWafGeoMatchSets {
	return &awsWafGeoMatchSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafGeoMatchSet, and returns the corresponding awsWafGeoMatchSet object, and an error if there is any.
func (c *awsWafGeoMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	result = &v1alpha1.AwsWafGeoMatchSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafGeoMatchSets that match those selectors.
func (c *awsWafGeoMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafGeoMatchSetList, err error) {
	result = &v1alpha1.AwsWafGeoMatchSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafGeoMatchSets.
func (c *awsWafGeoMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafGeoMatchSet and creates it.  Returns the server's representation of the awsWafGeoMatchSet, and an error, if there is any.
func (c *awsWafGeoMatchSets) Create(awsWafGeoMatchSet *v1alpha1.AwsWafGeoMatchSet) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	result = &v1alpha1.AwsWafGeoMatchSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		Body(awsWafGeoMatchSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafGeoMatchSet and updates it. Returns the server's representation of the awsWafGeoMatchSet, and an error, if there is any.
func (c *awsWafGeoMatchSets) Update(awsWafGeoMatchSet *v1alpha1.AwsWafGeoMatchSet) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	result = &v1alpha1.AwsWafGeoMatchSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		Name(awsWafGeoMatchSet.Name).
		Body(awsWafGeoMatchSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafGeoMatchSet and deletes it. Returns an error if one occurs.
func (c *awsWafGeoMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafGeoMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafGeoMatchSet.
func (c *awsWafGeoMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafGeoMatchSet, err error) {
	result = &v1alpha1.AwsWafGeoMatchSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafgeomatchsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
