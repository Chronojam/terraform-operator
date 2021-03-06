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

// AwsBudgetsBudgetsGetter has a method to return a AwsBudgetsBudgetInterface.
// A group's client should implement this interface.
type AwsBudgetsBudgetsGetter interface {
	AwsBudgetsBudgets(namespace string) AwsBudgetsBudgetInterface
}

// AwsBudgetsBudgetInterface has methods to work with AwsBudgetsBudget resources.
type AwsBudgetsBudgetInterface interface {
	Create(*v1alpha1.AwsBudgetsBudget) (*v1alpha1.AwsBudgetsBudget, error)
	Update(*v1alpha1.AwsBudgetsBudget) (*v1alpha1.AwsBudgetsBudget, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsBudgetsBudget, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsBudgetsBudgetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBudgetsBudget, err error)
	AwsBudgetsBudgetExpansion
}

// awsBudgetsBudgets implements AwsBudgetsBudgetInterface
type awsBudgetsBudgets struct {
	client rest.Interface
	ns     string
}

// newAwsBudgetsBudgets returns a AwsBudgetsBudgets
func newAwsBudgetsBudgets(c *ChronojamV1alpha1Client, namespace string) *awsBudgetsBudgets {
	return &awsBudgetsBudgets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsBudgetsBudget, and returns the corresponding awsBudgetsBudget object, and an error if there is any.
func (c *awsBudgetsBudgets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsBudgetsBudget, err error) {
	result = &v1alpha1.AwsBudgetsBudget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsBudgetsBudgets that match those selectors.
func (c *awsBudgetsBudgets) List(opts v1.ListOptions) (result *v1alpha1.AwsBudgetsBudgetList, err error) {
	result = &v1alpha1.AwsBudgetsBudgetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsBudgetsBudgets.
func (c *awsBudgetsBudgets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsBudgetsBudget and creates it.  Returns the server's representation of the awsBudgetsBudget, and an error, if there is any.
func (c *awsBudgetsBudgets) Create(awsBudgetsBudget *v1alpha1.AwsBudgetsBudget) (result *v1alpha1.AwsBudgetsBudget, err error) {
	result = &v1alpha1.AwsBudgetsBudget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		Body(awsBudgetsBudget).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsBudgetsBudget and updates it. Returns the server's representation of the awsBudgetsBudget, and an error, if there is any.
func (c *awsBudgetsBudgets) Update(awsBudgetsBudget *v1alpha1.AwsBudgetsBudget) (result *v1alpha1.AwsBudgetsBudget, err error) {
	result = &v1alpha1.AwsBudgetsBudget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		Name(awsBudgetsBudget.Name).
		Body(awsBudgetsBudget).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsBudgetsBudget and deletes it. Returns an error if one occurs.
func (c *awsBudgetsBudgets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsBudgetsBudgets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsBudgetsBudget.
func (c *awsBudgetsBudgets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBudgetsBudget, err error) {
	result = &v1alpha1.AwsBudgetsBudget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsbudgetsbudgets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
