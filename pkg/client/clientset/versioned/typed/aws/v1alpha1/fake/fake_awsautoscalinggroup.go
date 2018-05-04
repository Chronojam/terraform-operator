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

// FakeAwsAutoscalingGroups implements AwsAutoscalingGroupInterface
type FakeAwsAutoscalingGroups struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsautoscalinggroupsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsautoscalinggroups"}

var awsautoscalinggroupsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsAutoscalingGroup"}

// Get takes name of the awsAutoscalingGroup, and returns the corresponding awsAutoscalingGroup object, and an error if there is any.
func (c *FakeAwsAutoscalingGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsautoscalinggroupsResource, c.ns, name), &v1alpha1.AwsAutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingGroup), err
}

// List takes label and field selectors, and returns the list of AwsAutoscalingGroups that match those selectors.
func (c *FakeAwsAutoscalingGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsAutoscalingGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsautoscalinggroupsResource, awsautoscalinggroupsKind, c.ns, opts), &v1alpha1.AwsAutoscalingGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAutoscalingGroupList{}
	for _, item := range obj.(*v1alpha1.AwsAutoscalingGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingGroups.
func (c *FakeAwsAutoscalingGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsautoscalinggroupsResource, c.ns, opts))

}

// Create takes the representation of a awsAutoscalingGroup and creates it.  Returns the server's representation of the awsAutoscalingGroup, and an error, if there is any.
func (c *FakeAwsAutoscalingGroups) Create(awsAutoscalingGroup *v1alpha1.AwsAutoscalingGroup) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsautoscalinggroupsResource, c.ns, awsAutoscalingGroup), &v1alpha1.AwsAutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingGroup), err
}

// Update takes the representation of a awsAutoscalingGroup and updates it. Returns the server's representation of the awsAutoscalingGroup, and an error, if there is any.
func (c *FakeAwsAutoscalingGroups) Update(awsAutoscalingGroup *v1alpha1.AwsAutoscalingGroup) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsautoscalinggroupsResource, c.ns, awsAutoscalingGroup), &v1alpha1.AwsAutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingGroup), err
}

// Delete takes name of the awsAutoscalingGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsAutoscalingGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsautoscalinggroupsResource, c.ns, name), &v1alpha1.AwsAutoscalingGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAutoscalingGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsautoscalinggroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAutoscalingGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsAutoscalingGroup.
func (c *FakeAwsAutoscalingGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsautoscalinggroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsAutoscalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingGroup), err
}
