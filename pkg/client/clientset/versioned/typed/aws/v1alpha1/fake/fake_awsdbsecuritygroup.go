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

// FakeAwsDbSecurityGroups implements AwsDbSecurityGroupInterface
type FakeAwsDbSecurityGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsdbsecuritygroupsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsdbsecuritygroups"}

var awsdbsecuritygroupsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsDbSecurityGroup"}

// Get takes name of the awsDbSecurityGroup, and returns the corresponding awsDbSecurityGroup object, and an error if there is any.
func (c *FakeAwsDbSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDbSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdbsecuritygroupsResource, c.ns, name), &v1alpha1.AwsDbSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSecurityGroup), err
}

// List takes label and field selectors, and returns the list of AwsDbSecurityGroups that match those selectors.
func (c *FakeAwsDbSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsDbSecurityGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdbsecuritygroupsResource, awsdbsecuritygroupsKind, c.ns, opts), &v1alpha1.AwsDbSecurityGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDbSecurityGroupList{}
	for _, item := range obj.(*v1alpha1.AwsDbSecurityGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDbSecurityGroups.
func (c *FakeAwsDbSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdbsecuritygroupsResource, c.ns, opts))

}

// Create takes the representation of a awsDbSecurityGroup and creates it.  Returns the server's representation of the awsDbSecurityGroup, and an error, if there is any.
func (c *FakeAwsDbSecurityGroups) Create(awsDbSecurityGroup *v1alpha1.AwsDbSecurityGroup) (result *v1alpha1.AwsDbSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdbsecuritygroupsResource, c.ns, awsDbSecurityGroup), &v1alpha1.AwsDbSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSecurityGroup), err
}

// Update takes the representation of a awsDbSecurityGroup and updates it. Returns the server's representation of the awsDbSecurityGroup, and an error, if there is any.
func (c *FakeAwsDbSecurityGroups) Update(awsDbSecurityGroup *v1alpha1.AwsDbSecurityGroup) (result *v1alpha1.AwsDbSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdbsecuritygroupsResource, c.ns, awsDbSecurityGroup), &v1alpha1.AwsDbSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSecurityGroup), err
}

// Delete takes name of the awsDbSecurityGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsDbSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdbsecuritygroupsResource, c.ns, name), &v1alpha1.AwsDbSecurityGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDbSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdbsecuritygroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDbSecurityGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsDbSecurityGroup.
func (c *FakeAwsDbSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDbSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdbsecuritygroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsDbSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSecurityGroup), err
}
