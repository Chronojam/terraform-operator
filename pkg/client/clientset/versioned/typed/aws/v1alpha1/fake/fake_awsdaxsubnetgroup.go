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

// FakeAwsDaxSubnetGroups implements AwsDaxSubnetGroupInterface
type FakeAwsDaxSubnetGroups struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsdaxsubnetgroupsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsdaxsubnetgroups"}

var awsdaxsubnetgroupsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsDaxSubnetGroup"}

// Get takes name of the awsDaxSubnetGroup, and returns the corresponding awsDaxSubnetGroup object, and an error if there is any.
func (c *FakeAwsDaxSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdaxsubnetgroupsResource, c.ns, name), &v1alpha1.AwsDaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxSubnetGroup), err
}

// List takes label and field selectors, and returns the list of AwsDaxSubnetGroups that match those selectors.
func (c *FakeAwsDaxSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsDaxSubnetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdaxsubnetgroupsResource, awsdaxsubnetgroupsKind, c.ns, opts), &v1alpha1.AwsDaxSubnetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDaxSubnetGroupList{}
	for _, item := range obj.(*v1alpha1.AwsDaxSubnetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDaxSubnetGroups.
func (c *FakeAwsDaxSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdaxsubnetgroupsResource, c.ns, opts))

}

// Create takes the representation of a awsDaxSubnetGroup and creates it.  Returns the server's representation of the awsDaxSubnetGroup, and an error, if there is any.
func (c *FakeAwsDaxSubnetGroups) Create(awsDaxSubnetGroup *v1alpha1.AwsDaxSubnetGroup) (result *v1alpha1.AwsDaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdaxsubnetgroupsResource, c.ns, awsDaxSubnetGroup), &v1alpha1.AwsDaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxSubnetGroup), err
}

// Update takes the representation of a awsDaxSubnetGroup and updates it. Returns the server's representation of the awsDaxSubnetGroup, and an error, if there is any.
func (c *FakeAwsDaxSubnetGroups) Update(awsDaxSubnetGroup *v1alpha1.AwsDaxSubnetGroup) (result *v1alpha1.AwsDaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdaxsubnetgroupsResource, c.ns, awsDaxSubnetGroup), &v1alpha1.AwsDaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxSubnetGroup), err
}

// Delete takes name of the awsDaxSubnetGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsDaxSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdaxsubnetgroupsResource, c.ns, name), &v1alpha1.AwsDaxSubnetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDaxSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdaxsubnetgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDaxSubnetGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsDaxSubnetGroup.
func (c *FakeAwsDaxSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdaxsubnetgroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsDaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxSubnetGroup), err
}
