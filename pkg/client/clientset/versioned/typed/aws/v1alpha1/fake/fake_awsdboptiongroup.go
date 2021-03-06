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

// FakeAwsDbOptionGroups implements AwsDbOptionGroupInterface
type FakeAwsDbOptionGroups struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsdboptiongroupsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsdboptiongroups"}

var awsdboptiongroupsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsDbOptionGroup"}

// Get takes name of the awsDbOptionGroup, and returns the corresponding awsDbOptionGroup object, and an error if there is any.
func (c *FakeAwsDbOptionGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdboptiongroupsResource, c.ns, name), &v1alpha1.AwsDbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbOptionGroup), err
}

// List takes label and field selectors, and returns the list of AwsDbOptionGroups that match those selectors.
func (c *FakeAwsDbOptionGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsDbOptionGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdboptiongroupsResource, awsdboptiongroupsKind, c.ns, opts), &v1alpha1.AwsDbOptionGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDbOptionGroupList{}
	for _, item := range obj.(*v1alpha1.AwsDbOptionGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDbOptionGroups.
func (c *FakeAwsDbOptionGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdboptiongroupsResource, c.ns, opts))

}

// Create takes the representation of a awsDbOptionGroup and creates it.  Returns the server's representation of the awsDbOptionGroup, and an error, if there is any.
func (c *FakeAwsDbOptionGroups) Create(awsDbOptionGroup *v1alpha1.AwsDbOptionGroup) (result *v1alpha1.AwsDbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdboptiongroupsResource, c.ns, awsDbOptionGroup), &v1alpha1.AwsDbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbOptionGroup), err
}

// Update takes the representation of a awsDbOptionGroup and updates it. Returns the server's representation of the awsDbOptionGroup, and an error, if there is any.
func (c *FakeAwsDbOptionGroups) Update(awsDbOptionGroup *v1alpha1.AwsDbOptionGroup) (result *v1alpha1.AwsDbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdboptiongroupsResource, c.ns, awsDbOptionGroup), &v1alpha1.AwsDbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbOptionGroup), err
}

// Delete takes name of the awsDbOptionGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsDbOptionGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdboptiongroupsResource, c.ns, name), &v1alpha1.AwsDbOptionGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDbOptionGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdboptiongroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDbOptionGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsDbOptionGroup.
func (c *FakeAwsDbOptionGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdboptiongroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsDbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbOptionGroup), err
}
