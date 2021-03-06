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

// FakeAwsSsmMaintenanceWindowTargets implements AwsSsmMaintenanceWindowTargetInterface
type FakeAwsSsmMaintenanceWindowTargets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsssmmaintenancewindowtargetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsssmmaintenancewindowtargets"}

var awsssmmaintenancewindowtargetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSsmMaintenanceWindowTarget"}

// Get takes name of the awsSsmMaintenanceWindowTarget, and returns the corresponding awsSsmMaintenanceWindowTarget object, and an error if there is any.
func (c *FakeAwsSsmMaintenanceWindowTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSsmMaintenanceWindowTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsssmmaintenancewindowtargetsResource, c.ns, name), &v1alpha1.AwsSsmMaintenanceWindowTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTarget), err
}

// List takes label and field selectors, and returns the list of AwsSsmMaintenanceWindowTargets that match those selectors.
func (c *FakeAwsSsmMaintenanceWindowTargets) List(opts v1.ListOptions) (result *v1alpha1.AwsSsmMaintenanceWindowTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsssmmaintenancewindowtargetsResource, awsssmmaintenancewindowtargetsKind, c.ns, opts), &v1alpha1.AwsSsmMaintenanceWindowTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSsmMaintenanceWindowTargetList{}
	for _, item := range obj.(*v1alpha1.AwsSsmMaintenanceWindowTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSsmMaintenanceWindowTargets.
func (c *FakeAwsSsmMaintenanceWindowTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsssmmaintenancewindowtargetsResource, c.ns, opts))

}

// Create takes the representation of a awsSsmMaintenanceWindowTarget and creates it.  Returns the server's representation of the awsSsmMaintenanceWindowTarget, and an error, if there is any.
func (c *FakeAwsSsmMaintenanceWindowTargets) Create(awsSsmMaintenanceWindowTarget *v1alpha1.AwsSsmMaintenanceWindowTarget) (result *v1alpha1.AwsSsmMaintenanceWindowTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsssmmaintenancewindowtargetsResource, c.ns, awsSsmMaintenanceWindowTarget), &v1alpha1.AwsSsmMaintenanceWindowTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTarget), err
}

// Update takes the representation of a awsSsmMaintenanceWindowTarget and updates it. Returns the server's representation of the awsSsmMaintenanceWindowTarget, and an error, if there is any.
func (c *FakeAwsSsmMaintenanceWindowTargets) Update(awsSsmMaintenanceWindowTarget *v1alpha1.AwsSsmMaintenanceWindowTarget) (result *v1alpha1.AwsSsmMaintenanceWindowTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsssmmaintenancewindowtargetsResource, c.ns, awsSsmMaintenanceWindowTarget), &v1alpha1.AwsSsmMaintenanceWindowTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTarget), err
}

// Delete takes name of the awsSsmMaintenanceWindowTarget and deletes it. Returns an error if one occurs.
func (c *FakeAwsSsmMaintenanceWindowTargets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsssmmaintenancewindowtargetsResource, c.ns, name), &v1alpha1.AwsSsmMaintenanceWindowTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSsmMaintenanceWindowTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsssmmaintenancewindowtargetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSsmMaintenanceWindowTargetList{})
	return err
}

// Patch applies the patch and returns the patched awsSsmMaintenanceWindowTarget.
func (c *FakeAwsSsmMaintenanceWindowTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSsmMaintenanceWindowTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsssmmaintenancewindowtargetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsSsmMaintenanceWindowTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTarget), err
}
