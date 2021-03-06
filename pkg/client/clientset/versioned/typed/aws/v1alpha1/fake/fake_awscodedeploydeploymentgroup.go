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

// FakeAwsCodedeployDeploymentGroups implements AwsCodedeployDeploymentGroupInterface
type FakeAwsCodedeployDeploymentGroups struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awscodedeploydeploymentgroupsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awscodedeploydeploymentgroups"}

var awscodedeploydeploymentgroupsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsCodedeployDeploymentGroup"}

// Get takes name of the awsCodedeployDeploymentGroup, and returns the corresponding awsCodedeployDeploymentGroup object, and an error if there is any.
func (c *FakeAwsCodedeployDeploymentGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscodedeploydeploymentgroupsResource, c.ns, name), &v1alpha1.AwsCodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentGroup), err
}

// List takes label and field selectors, and returns the list of AwsCodedeployDeploymentGroups that match those selectors.
func (c *FakeAwsCodedeployDeploymentGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsCodedeployDeploymentGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscodedeploydeploymentgroupsResource, awscodedeploydeploymentgroupsKind, c.ns, opts), &v1alpha1.AwsCodedeployDeploymentGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCodedeployDeploymentGroupList{}
	for _, item := range obj.(*v1alpha1.AwsCodedeployDeploymentGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCodedeployDeploymentGroups.
func (c *FakeAwsCodedeployDeploymentGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscodedeploydeploymentgroupsResource, c.ns, opts))

}

// Create takes the representation of a awsCodedeployDeploymentGroup and creates it.  Returns the server's representation of the awsCodedeployDeploymentGroup, and an error, if there is any.
func (c *FakeAwsCodedeployDeploymentGroups) Create(awsCodedeployDeploymentGroup *v1alpha1.AwsCodedeployDeploymentGroup) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscodedeploydeploymentgroupsResource, c.ns, awsCodedeployDeploymentGroup), &v1alpha1.AwsCodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentGroup), err
}

// Update takes the representation of a awsCodedeployDeploymentGroup and updates it. Returns the server's representation of the awsCodedeployDeploymentGroup, and an error, if there is any.
func (c *FakeAwsCodedeployDeploymentGroups) Update(awsCodedeployDeploymentGroup *v1alpha1.AwsCodedeployDeploymentGroup) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscodedeploydeploymentgroupsResource, c.ns, awsCodedeployDeploymentGroup), &v1alpha1.AwsCodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentGroup), err
}

// Delete takes name of the awsCodedeployDeploymentGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsCodedeployDeploymentGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscodedeploydeploymentgroupsResource, c.ns, name), &v1alpha1.AwsCodedeployDeploymentGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCodedeployDeploymentGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscodedeploydeploymentgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCodedeployDeploymentGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsCodedeployDeploymentGroup.
func (c *FakeAwsCodedeployDeploymentGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscodedeploydeploymentgroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsCodedeployDeploymentGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployDeploymentGroup), err
}
