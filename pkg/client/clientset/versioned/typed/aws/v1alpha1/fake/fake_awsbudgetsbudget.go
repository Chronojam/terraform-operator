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

// FakeAwsBudgetsBudgets implements AwsBudgetsBudgetInterface
type FakeAwsBudgetsBudgets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsbudgetsbudgetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsbudgetsbudgets"}

var awsbudgetsbudgetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsBudgetsBudget"}

// Get takes name of the awsBudgetsBudget, and returns the corresponding awsBudgetsBudget object, and an error if there is any.
func (c *FakeAwsBudgetsBudgets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsBudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsbudgetsbudgetsResource, c.ns, name), &v1alpha1.AwsBudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBudgetsBudget), err
}

// List takes label and field selectors, and returns the list of AwsBudgetsBudgets that match those selectors.
func (c *FakeAwsBudgetsBudgets) List(opts v1.ListOptions) (result *v1alpha1.AwsBudgetsBudgetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsbudgetsbudgetsResource, awsbudgetsbudgetsKind, c.ns, opts), &v1alpha1.AwsBudgetsBudgetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsBudgetsBudgetList{}
	for _, item := range obj.(*v1alpha1.AwsBudgetsBudgetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsBudgetsBudgets.
func (c *FakeAwsBudgetsBudgets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsbudgetsbudgetsResource, c.ns, opts))

}

// Create takes the representation of a awsBudgetsBudget and creates it.  Returns the server's representation of the awsBudgetsBudget, and an error, if there is any.
func (c *FakeAwsBudgetsBudgets) Create(awsBudgetsBudget *v1alpha1.AwsBudgetsBudget) (result *v1alpha1.AwsBudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsbudgetsbudgetsResource, c.ns, awsBudgetsBudget), &v1alpha1.AwsBudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBudgetsBudget), err
}

// Update takes the representation of a awsBudgetsBudget and updates it. Returns the server's representation of the awsBudgetsBudget, and an error, if there is any.
func (c *FakeAwsBudgetsBudgets) Update(awsBudgetsBudget *v1alpha1.AwsBudgetsBudget) (result *v1alpha1.AwsBudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsbudgetsbudgetsResource, c.ns, awsBudgetsBudget), &v1alpha1.AwsBudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBudgetsBudget), err
}

// Delete takes name of the awsBudgetsBudget and deletes it. Returns an error if one occurs.
func (c *FakeAwsBudgetsBudgets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsbudgetsbudgetsResource, c.ns, name), &v1alpha1.AwsBudgetsBudget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsBudgetsBudgets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsbudgetsbudgetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsBudgetsBudgetList{})
	return err
}

// Patch applies the patch and returns the patched awsBudgetsBudget.
func (c *FakeAwsBudgetsBudgets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsbudgetsbudgetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsBudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBudgetsBudget), err
}
