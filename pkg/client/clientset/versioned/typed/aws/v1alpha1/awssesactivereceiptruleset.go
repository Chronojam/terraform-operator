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

// AwsSesActiveReceiptRuleSetsGetter has a method to return a AwsSesActiveReceiptRuleSetInterface.
// A group's client should implement this interface.
type AwsSesActiveReceiptRuleSetsGetter interface {
	AwsSesActiveReceiptRuleSets(namespace string) AwsSesActiveReceiptRuleSetInterface
}

// AwsSesActiveReceiptRuleSetInterface has methods to work with AwsSesActiveReceiptRuleSet resources.
type AwsSesActiveReceiptRuleSetInterface interface {
	Create(*v1alpha1.AwsSesActiveReceiptRuleSet) (*v1alpha1.AwsSesActiveReceiptRuleSet, error)
	Update(*v1alpha1.AwsSesActiveReceiptRuleSet) (*v1alpha1.AwsSesActiveReceiptRuleSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSesActiveReceiptRuleSet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSesActiveReceiptRuleSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesActiveReceiptRuleSet, err error)
	AwsSesActiveReceiptRuleSetExpansion
}

// awsSesActiveReceiptRuleSets implements AwsSesActiveReceiptRuleSetInterface
type awsSesActiveReceiptRuleSets struct {
	client rest.Interface
	ns     string
}

// newAwsSesActiveReceiptRuleSets returns a AwsSesActiveReceiptRuleSets
func newAwsSesActiveReceiptRuleSets(c *ChronojamV1alpha1Client, namespace string) *awsSesActiveReceiptRuleSets {
	return &awsSesActiveReceiptRuleSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSesActiveReceiptRuleSet, and returns the corresponding awsSesActiveReceiptRuleSet object, and an error if there is any.
func (c *awsSesActiveReceiptRuleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesActiveReceiptRuleSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSesActiveReceiptRuleSets that match those selectors.
func (c *awsSesActiveReceiptRuleSets) List(opts v1.ListOptions) (result *v1alpha1.AwsSesActiveReceiptRuleSetList, err error) {
	result = &v1alpha1.AwsSesActiveReceiptRuleSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSesActiveReceiptRuleSets.
func (c *awsSesActiveReceiptRuleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSesActiveReceiptRuleSet and creates it.  Returns the server's representation of the awsSesActiveReceiptRuleSet, and an error, if there is any.
func (c *awsSesActiveReceiptRuleSets) Create(awsSesActiveReceiptRuleSet *v1alpha1.AwsSesActiveReceiptRuleSet) (result *v1alpha1.AwsSesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesActiveReceiptRuleSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		Body(awsSesActiveReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSesActiveReceiptRuleSet and updates it. Returns the server's representation of the awsSesActiveReceiptRuleSet, and an error, if there is any.
func (c *awsSesActiveReceiptRuleSets) Update(awsSesActiveReceiptRuleSet *v1alpha1.AwsSesActiveReceiptRuleSet) (result *v1alpha1.AwsSesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesActiveReceiptRuleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		Name(awsSesActiveReceiptRuleSet.Name).
		Body(awsSesActiveReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSesActiveReceiptRuleSet and deletes it. Returns an error if one occurs.
func (c *awsSesActiveReceiptRuleSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSesActiveReceiptRuleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSesActiveReceiptRuleSet.
func (c *awsSesActiveReceiptRuleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesActiveReceiptRuleSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssesactivereceiptrulesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
