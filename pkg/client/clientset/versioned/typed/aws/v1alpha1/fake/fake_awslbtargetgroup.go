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

// FakeAwsLbTargetGroups implements AwsLbTargetGroupInterface
type FakeAwsLbTargetGroups struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awslbtargetgroupsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awslbtargetgroups"}

var awslbtargetgroupsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsLbTargetGroup"}

// Get takes name of the awsLbTargetGroup, and returns the corresponding awsLbTargetGroup object, and an error if there is any.
func (c *FakeAwsLbTargetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslbtargetgroupsResource, c.ns, name), &v1alpha1.AwsLbTargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroup), err
}

// List takes label and field selectors, and returns the list of AwsLbTargetGroups that match those selectors.
func (c *FakeAwsLbTargetGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsLbTargetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslbtargetgroupsResource, awslbtargetgroupsKind, c.ns, opts), &v1alpha1.AwsLbTargetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLbTargetGroupList{}
	for _, item := range obj.(*v1alpha1.AwsLbTargetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLbTargetGroups.
func (c *FakeAwsLbTargetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslbtargetgroupsResource, c.ns, opts))

}

// Create takes the representation of a awsLbTargetGroup and creates it.  Returns the server's representation of the awsLbTargetGroup, and an error, if there is any.
func (c *FakeAwsLbTargetGroups) Create(awsLbTargetGroup *v1alpha1.AwsLbTargetGroup) (result *v1alpha1.AwsLbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslbtargetgroupsResource, c.ns, awsLbTargetGroup), &v1alpha1.AwsLbTargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroup), err
}

// Update takes the representation of a awsLbTargetGroup and updates it. Returns the server's representation of the awsLbTargetGroup, and an error, if there is any.
func (c *FakeAwsLbTargetGroups) Update(awsLbTargetGroup *v1alpha1.AwsLbTargetGroup) (result *v1alpha1.AwsLbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslbtargetgroupsResource, c.ns, awsLbTargetGroup), &v1alpha1.AwsLbTargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroup), err
}

// Delete takes name of the awsLbTargetGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsLbTargetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslbtargetgroupsResource, c.ns, name), &v1alpha1.AwsLbTargetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLbTargetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslbtargetgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLbTargetGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsLbTargetGroup.
func (c *FakeAwsLbTargetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslbtargetgroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsLbTargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroup), err
}
