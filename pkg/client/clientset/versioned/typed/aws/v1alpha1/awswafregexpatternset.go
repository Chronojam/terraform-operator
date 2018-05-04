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

// AwsWafRegexPatternSetsGetter has a method to return a AwsWafRegexPatternSetInterface.
// A group's client should implement this interface.
type AwsWafRegexPatternSetsGetter interface {
	AwsWafRegexPatternSets(namespace string) AwsWafRegexPatternSetInterface
}

// AwsWafRegexPatternSetInterface has methods to work with AwsWafRegexPatternSet resources.
type AwsWafRegexPatternSetInterface interface {
	Create(*v1alpha1.AwsWafRegexPatternSet) (*v1alpha1.AwsWafRegexPatternSet, error)
	Update(*v1alpha1.AwsWafRegexPatternSet) (*v1alpha1.AwsWafRegexPatternSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafRegexPatternSet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafRegexPatternSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRegexPatternSet, err error)
	AwsWafRegexPatternSetExpansion
}

// awsWafRegexPatternSets implements AwsWafRegexPatternSetInterface
type awsWafRegexPatternSets struct {
	client rest.Interface
	ns     string
}

// newAwsWafRegexPatternSets returns a AwsWafRegexPatternSets
func newAwsWafRegexPatternSets(c *ChronojamV1alpha1Client, namespace string) *awsWafRegexPatternSets {
	return &awsWafRegexPatternSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafRegexPatternSet, and returns the corresponding awsWafRegexPatternSet object, and an error if there is any.
func (c *awsWafRegexPatternSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafRegexPatternSet, err error) {
	result = &v1alpha1.AwsWafRegexPatternSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafRegexPatternSets that match those selectors.
func (c *awsWafRegexPatternSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafRegexPatternSetList, err error) {
	result = &v1alpha1.AwsWafRegexPatternSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafRegexPatternSets.
func (c *awsWafRegexPatternSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafRegexPatternSet and creates it.  Returns the server's representation of the awsWafRegexPatternSet, and an error, if there is any.
func (c *awsWafRegexPatternSets) Create(awsWafRegexPatternSet *v1alpha1.AwsWafRegexPatternSet) (result *v1alpha1.AwsWafRegexPatternSet, err error) {
	result = &v1alpha1.AwsWafRegexPatternSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		Body(awsWafRegexPatternSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafRegexPatternSet and updates it. Returns the server's representation of the awsWafRegexPatternSet, and an error, if there is any.
func (c *awsWafRegexPatternSets) Update(awsWafRegexPatternSet *v1alpha1.AwsWafRegexPatternSet) (result *v1alpha1.AwsWafRegexPatternSet, err error) {
	result = &v1alpha1.AwsWafRegexPatternSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		Name(awsWafRegexPatternSet.Name).
		Body(awsWafRegexPatternSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafRegexPatternSet and deletes it. Returns an error if one occurs.
func (c *awsWafRegexPatternSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafRegexPatternSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafRegexPatternSet.
func (c *awsWafRegexPatternSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRegexPatternSet, err error) {
	result = &v1alpha1.AwsWafRegexPatternSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafregexpatternsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
