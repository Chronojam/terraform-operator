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

package fake

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsSfnStateMachines implements AwsSfnStateMachineInterface
type FakeAwsSfnStateMachines struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awssfnstatemachinesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awssfnstatemachines"}

var awssfnstatemachinesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSfnStateMachine"}

// Get takes name of the awsSfnStateMachine, and returns the corresponding awsSfnStateMachine object, and an error if there is any.
func (c *FakeAwsSfnStateMachines) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssfnstatemachinesResource, c.ns, name), &v1alpha1.AwsSfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSfnStateMachine), err
}

// List takes label and field selectors, and returns the list of AwsSfnStateMachines that match those selectors.
func (c *FakeAwsSfnStateMachines) List(opts v1.ListOptions) (result *v1alpha1.AwsSfnStateMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssfnstatemachinesResource, awssfnstatemachinesKind, c.ns, opts), &v1alpha1.AwsSfnStateMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSfnStateMachineList{}
	for _, item := range obj.(*v1alpha1.AwsSfnStateMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSfnStateMachines.
func (c *FakeAwsSfnStateMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssfnstatemachinesResource, c.ns, opts))

}

// Create takes the representation of a awsSfnStateMachine and creates it.  Returns the server's representation of the awsSfnStateMachine, and an error, if there is any.
func (c *FakeAwsSfnStateMachines) Create(awsSfnStateMachine *v1alpha1.AwsSfnStateMachine) (result *v1alpha1.AwsSfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssfnstatemachinesResource, c.ns, awsSfnStateMachine), &v1alpha1.AwsSfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSfnStateMachine), err
}

// Update takes the representation of a awsSfnStateMachine and updates it. Returns the server's representation of the awsSfnStateMachine, and an error, if there is any.
func (c *FakeAwsSfnStateMachines) Update(awsSfnStateMachine *v1alpha1.AwsSfnStateMachine) (result *v1alpha1.AwsSfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssfnstatemachinesResource, c.ns, awsSfnStateMachine), &v1alpha1.AwsSfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSfnStateMachine), err
}

// Delete takes name of the awsSfnStateMachine and deletes it. Returns an error if one occurs.
func (c *FakeAwsSfnStateMachines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssfnstatemachinesResource, c.ns, name), &v1alpha1.AwsSfnStateMachine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSfnStateMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssfnstatemachinesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSfnStateMachineList{})
	return err
}

// Patch applies the patch and returns the patched awsSfnStateMachine.
func (c *FakeAwsSfnStateMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssfnstatemachinesResource, c.ns, name, data, subresources...), &v1alpha1.AwsSfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSfnStateMachine), err
}
