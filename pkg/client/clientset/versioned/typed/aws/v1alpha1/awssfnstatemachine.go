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

// AwsSfnStateMachinesGetter has a method to return a AwsSfnStateMachineInterface.
// A group's client should implement this interface.
type AwsSfnStateMachinesGetter interface {
	AwsSfnStateMachines(namespace string) AwsSfnStateMachineInterface
}

// AwsSfnStateMachineInterface has methods to work with AwsSfnStateMachine resources.
type AwsSfnStateMachineInterface interface {
	Create(*v1alpha1.AwsSfnStateMachine) (*v1alpha1.AwsSfnStateMachine, error)
	Update(*v1alpha1.AwsSfnStateMachine) (*v1alpha1.AwsSfnStateMachine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSfnStateMachine, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSfnStateMachineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSfnStateMachine, err error)
	AwsSfnStateMachineExpansion
}

// awsSfnStateMachines implements AwsSfnStateMachineInterface
type awsSfnStateMachines struct {
	client rest.Interface
	ns     string
}

// newAwsSfnStateMachines returns a AwsSfnStateMachines
func newAwsSfnStateMachines(c *ChronojamV1alpha1Client, namespace string) *awsSfnStateMachines {
	return &awsSfnStateMachines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSfnStateMachine, and returns the corresponding awsSfnStateMachine object, and an error if there is any.
func (c *awsSfnStateMachines) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSfnStateMachine, err error) {
	result = &v1alpha1.AwsSfnStateMachine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSfnStateMachines that match those selectors.
func (c *awsSfnStateMachines) List(opts v1.ListOptions) (result *v1alpha1.AwsSfnStateMachineList, err error) {
	result = &v1alpha1.AwsSfnStateMachineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSfnStateMachines.
func (c *awsSfnStateMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSfnStateMachine and creates it.  Returns the server's representation of the awsSfnStateMachine, and an error, if there is any.
func (c *awsSfnStateMachines) Create(awsSfnStateMachine *v1alpha1.AwsSfnStateMachine) (result *v1alpha1.AwsSfnStateMachine, err error) {
	result = &v1alpha1.AwsSfnStateMachine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		Body(awsSfnStateMachine).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSfnStateMachine and updates it. Returns the server's representation of the awsSfnStateMachine, and an error, if there is any.
func (c *awsSfnStateMachines) Update(awsSfnStateMachine *v1alpha1.AwsSfnStateMachine) (result *v1alpha1.AwsSfnStateMachine, err error) {
	result = &v1alpha1.AwsSfnStateMachine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		Name(awsSfnStateMachine.Name).
		Body(awsSfnStateMachine).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSfnStateMachine and deletes it. Returns an error if one occurs.
func (c *awsSfnStateMachines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSfnStateMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSfnStateMachine.
func (c *awsSfnStateMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSfnStateMachine, err error) {
	result = &v1alpha1.AwsSfnStateMachine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssfnstatemachines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
