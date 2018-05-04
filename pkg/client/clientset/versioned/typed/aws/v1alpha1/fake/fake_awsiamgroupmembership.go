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

// FakeAwsIamGroupMemberships implements AwsIamGroupMembershipInterface
type FakeAwsIamGroupMemberships struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsiamgroupmembershipsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsiamgroupmemberships"}

var awsiamgroupmembershipsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsIamGroupMembership"}

// Get takes name of the awsIamGroupMembership, and returns the corresponding awsIamGroupMembership object, and an error if there is any.
func (c *FakeAwsIamGroupMemberships) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiamgroupmembershipsResource, c.ns, name), &v1alpha1.AwsIamGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupMembership), err
}

// List takes label and field selectors, and returns the list of AwsIamGroupMemberships that match those selectors.
func (c *FakeAwsIamGroupMemberships) List(opts v1.ListOptions) (result *v1alpha1.AwsIamGroupMembershipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiamgroupmembershipsResource, awsiamgroupmembershipsKind, c.ns, opts), &v1alpha1.AwsIamGroupMembershipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamGroupMembershipList{}
	for _, item := range obj.(*v1alpha1.AwsIamGroupMembershipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamGroupMemberships.
func (c *FakeAwsIamGroupMemberships) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiamgroupmembershipsResource, c.ns, opts))

}

// Create takes the representation of a awsIamGroupMembership and creates it.  Returns the server's representation of the awsIamGroupMembership, and an error, if there is any.
func (c *FakeAwsIamGroupMemberships) Create(awsIamGroupMembership *v1alpha1.AwsIamGroupMembership) (result *v1alpha1.AwsIamGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiamgroupmembershipsResource, c.ns, awsIamGroupMembership), &v1alpha1.AwsIamGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupMembership), err
}

// Update takes the representation of a awsIamGroupMembership and updates it. Returns the server's representation of the awsIamGroupMembership, and an error, if there is any.
func (c *FakeAwsIamGroupMemberships) Update(awsIamGroupMembership *v1alpha1.AwsIamGroupMembership) (result *v1alpha1.AwsIamGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiamgroupmembershipsResource, c.ns, awsIamGroupMembership), &v1alpha1.AwsIamGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupMembership), err
}

// Delete takes name of the awsIamGroupMembership and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamGroupMemberships) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiamgroupmembershipsResource, c.ns, name), &v1alpha1.AwsIamGroupMembership{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamGroupMemberships) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiamgroupmembershipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamGroupMembershipList{})
	return err
}

// Patch applies the patch and returns the patched awsIamGroupMembership.
func (c *FakeAwsIamGroupMemberships) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiamgroupmembershipsResource, c.ns, name, data, subresources...), &v1alpha1.AwsIamGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupMembership), err
}
