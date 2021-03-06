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

// FakeAwsIamUserGroupMemberships implements AwsIamUserGroupMembershipInterface
type FakeAwsIamUserGroupMemberships struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsiamusergroupmembershipsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsiamusergroupmemberships"}

var awsiamusergroupmembershipsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsIamUserGroupMembership"}

// Get takes name of the awsIamUserGroupMembership, and returns the corresponding awsIamUserGroupMembership object, and an error if there is any.
func (c *FakeAwsIamUserGroupMemberships) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiamusergroupmembershipsResource, c.ns, name), &v1alpha1.AwsIamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUserGroupMembership), err
}

// List takes label and field selectors, and returns the list of AwsIamUserGroupMemberships that match those selectors.
func (c *FakeAwsIamUserGroupMemberships) List(opts v1.ListOptions) (result *v1alpha1.AwsIamUserGroupMembershipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiamusergroupmembershipsResource, awsiamusergroupmembershipsKind, c.ns, opts), &v1alpha1.AwsIamUserGroupMembershipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamUserGroupMembershipList{}
	for _, item := range obj.(*v1alpha1.AwsIamUserGroupMembershipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamUserGroupMemberships.
func (c *FakeAwsIamUserGroupMemberships) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiamusergroupmembershipsResource, c.ns, opts))

}

// Create takes the representation of a awsIamUserGroupMembership and creates it.  Returns the server's representation of the awsIamUserGroupMembership, and an error, if there is any.
func (c *FakeAwsIamUserGroupMemberships) Create(awsIamUserGroupMembership *v1alpha1.AwsIamUserGroupMembership) (result *v1alpha1.AwsIamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiamusergroupmembershipsResource, c.ns, awsIamUserGroupMembership), &v1alpha1.AwsIamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUserGroupMembership), err
}

// Update takes the representation of a awsIamUserGroupMembership and updates it. Returns the server's representation of the awsIamUserGroupMembership, and an error, if there is any.
func (c *FakeAwsIamUserGroupMemberships) Update(awsIamUserGroupMembership *v1alpha1.AwsIamUserGroupMembership) (result *v1alpha1.AwsIamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiamusergroupmembershipsResource, c.ns, awsIamUserGroupMembership), &v1alpha1.AwsIamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUserGroupMembership), err
}

// Delete takes name of the awsIamUserGroupMembership and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamUserGroupMemberships) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiamusergroupmembershipsResource, c.ns, name), &v1alpha1.AwsIamUserGroupMembership{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamUserGroupMemberships) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiamusergroupmembershipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamUserGroupMembershipList{})
	return err
}

// Patch applies the patch and returns the patched awsIamUserGroupMembership.
func (c *FakeAwsIamUserGroupMemberships) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiamusergroupmembershipsResource, c.ns, name, data, subresources...), &v1alpha1.AwsIamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUserGroupMembership), err
}
