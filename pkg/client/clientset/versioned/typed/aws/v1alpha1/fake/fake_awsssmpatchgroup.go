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

// FakeAwsSsmPatchGroups implements AwsSsmPatchGroupInterface
type FakeAwsSsmPatchGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsssmpatchgroupsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsssmpatchgroups"}

var awsssmpatchgroupsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsSsmPatchGroup"}

// Get takes name of the awsSsmPatchGroup, and returns the corresponding awsSsmPatchGroup object, and an error if there is any.
func (c *FakeAwsSsmPatchGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsssmpatchgroupsResource, c.ns, name), &v1alpha1.AwsSsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmPatchGroup), err
}

// List takes label and field selectors, and returns the list of AwsSsmPatchGroups that match those selectors.
func (c *FakeAwsSsmPatchGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsSsmPatchGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsssmpatchgroupsResource, awsssmpatchgroupsKind, c.ns, opts), &v1alpha1.AwsSsmPatchGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSsmPatchGroupList{}
	for _, item := range obj.(*v1alpha1.AwsSsmPatchGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSsmPatchGroups.
func (c *FakeAwsSsmPatchGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsssmpatchgroupsResource, c.ns, opts))

}

// Create takes the representation of a awsSsmPatchGroup and creates it.  Returns the server's representation of the awsSsmPatchGroup, and an error, if there is any.
func (c *FakeAwsSsmPatchGroups) Create(awsSsmPatchGroup *v1alpha1.AwsSsmPatchGroup) (result *v1alpha1.AwsSsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsssmpatchgroupsResource, c.ns, awsSsmPatchGroup), &v1alpha1.AwsSsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmPatchGroup), err
}

// Update takes the representation of a awsSsmPatchGroup and updates it. Returns the server's representation of the awsSsmPatchGroup, and an error, if there is any.
func (c *FakeAwsSsmPatchGroups) Update(awsSsmPatchGroup *v1alpha1.AwsSsmPatchGroup) (result *v1alpha1.AwsSsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsssmpatchgroupsResource, c.ns, awsSsmPatchGroup), &v1alpha1.AwsSsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmPatchGroup), err
}

// Delete takes name of the awsSsmPatchGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsSsmPatchGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsssmpatchgroupsResource, c.ns, name), &v1alpha1.AwsSsmPatchGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSsmPatchGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsssmpatchgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSsmPatchGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsSsmPatchGroup.
func (c *FakeAwsSsmPatchGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsssmpatchgroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsSsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmPatchGroup), err
}
