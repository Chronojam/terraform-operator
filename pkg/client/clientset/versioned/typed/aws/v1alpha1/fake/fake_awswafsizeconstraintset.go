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

// FakeAwsWafSizeConstraintSets implements AwsWafSizeConstraintSetInterface
type FakeAwsWafSizeConstraintSets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awswafsizeconstraintsetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awswafsizeconstraintsets"}

var awswafsizeconstraintsetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsWafSizeConstraintSet"}

// Get takes name of the awsWafSizeConstraintSet, and returns the corresponding awsWafSizeConstraintSet object, and an error if there is any.
func (c *FakeAwsWafSizeConstraintSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafSizeConstraintSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafsizeconstraintsetsResource, c.ns, name), &v1alpha1.AwsWafSizeConstraintSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafSizeConstraintSet), err
}

// List takes label and field selectors, and returns the list of AwsWafSizeConstraintSets that match those selectors.
func (c *FakeAwsWafSizeConstraintSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafSizeConstraintSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafsizeconstraintsetsResource, awswafsizeconstraintsetsKind, c.ns, opts), &v1alpha1.AwsWafSizeConstraintSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafSizeConstraintSetList{}
	for _, item := range obj.(*v1alpha1.AwsWafSizeConstraintSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafSizeConstraintSets.
func (c *FakeAwsWafSizeConstraintSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafsizeconstraintsetsResource, c.ns, opts))

}

// Create takes the representation of a awsWafSizeConstraintSet and creates it.  Returns the server's representation of the awsWafSizeConstraintSet, and an error, if there is any.
func (c *FakeAwsWafSizeConstraintSets) Create(awsWafSizeConstraintSet *v1alpha1.AwsWafSizeConstraintSet) (result *v1alpha1.AwsWafSizeConstraintSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafsizeconstraintsetsResource, c.ns, awsWafSizeConstraintSet), &v1alpha1.AwsWafSizeConstraintSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafSizeConstraintSet), err
}

// Update takes the representation of a awsWafSizeConstraintSet and updates it. Returns the server's representation of the awsWafSizeConstraintSet, and an error, if there is any.
func (c *FakeAwsWafSizeConstraintSets) Update(awsWafSizeConstraintSet *v1alpha1.AwsWafSizeConstraintSet) (result *v1alpha1.AwsWafSizeConstraintSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafsizeconstraintsetsResource, c.ns, awsWafSizeConstraintSet), &v1alpha1.AwsWafSizeConstraintSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafSizeConstraintSet), err
}

// Delete takes name of the awsWafSizeConstraintSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafSizeConstraintSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafsizeconstraintsetsResource, c.ns, name), &v1alpha1.AwsWafSizeConstraintSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafSizeConstraintSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafsizeconstraintsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafSizeConstraintSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafSizeConstraintSet.
func (c *FakeAwsWafSizeConstraintSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafSizeConstraintSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafsizeconstraintsetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsWafSizeConstraintSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafSizeConstraintSet), err
}
