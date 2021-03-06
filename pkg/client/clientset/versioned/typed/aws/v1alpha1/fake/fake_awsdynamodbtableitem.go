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

// FakeAwsDynamodbTableItems implements AwsDynamodbTableItemInterface
type FakeAwsDynamodbTableItems struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsdynamodbtableitemsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsdynamodbtableitems"}

var awsdynamodbtableitemsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsDynamodbTableItem"}

// Get takes name of the awsDynamodbTableItem, and returns the corresponding awsDynamodbTableItem object, and an error if there is any.
func (c *FakeAwsDynamodbTableItems) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdynamodbtableitemsResource, c.ns, name), &v1alpha1.AwsDynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbTableItem), err
}

// List takes label and field selectors, and returns the list of AwsDynamodbTableItems that match those selectors.
func (c *FakeAwsDynamodbTableItems) List(opts v1.ListOptions) (result *v1alpha1.AwsDynamodbTableItemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdynamodbtableitemsResource, awsdynamodbtableitemsKind, c.ns, opts), &v1alpha1.AwsDynamodbTableItemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDynamodbTableItemList{}
	for _, item := range obj.(*v1alpha1.AwsDynamodbTableItemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDynamodbTableItems.
func (c *FakeAwsDynamodbTableItems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdynamodbtableitemsResource, c.ns, opts))

}

// Create takes the representation of a awsDynamodbTableItem and creates it.  Returns the server's representation of the awsDynamodbTableItem, and an error, if there is any.
func (c *FakeAwsDynamodbTableItems) Create(awsDynamodbTableItem *v1alpha1.AwsDynamodbTableItem) (result *v1alpha1.AwsDynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdynamodbtableitemsResource, c.ns, awsDynamodbTableItem), &v1alpha1.AwsDynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbTableItem), err
}

// Update takes the representation of a awsDynamodbTableItem and updates it. Returns the server's representation of the awsDynamodbTableItem, and an error, if there is any.
func (c *FakeAwsDynamodbTableItems) Update(awsDynamodbTableItem *v1alpha1.AwsDynamodbTableItem) (result *v1alpha1.AwsDynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdynamodbtableitemsResource, c.ns, awsDynamodbTableItem), &v1alpha1.AwsDynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbTableItem), err
}

// Delete takes name of the awsDynamodbTableItem and deletes it. Returns an error if one occurs.
func (c *FakeAwsDynamodbTableItems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdynamodbtableitemsResource, c.ns, name), &v1alpha1.AwsDynamodbTableItem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDynamodbTableItems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdynamodbtableitemsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDynamodbTableItemList{})
	return err
}

// Patch applies the patch and returns the patched awsDynamodbTableItem.
func (c *FakeAwsDynamodbTableItems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdynamodbtableitemsResource, c.ns, name, data, subresources...), &v1alpha1.AwsDynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbTableItem), err
}
