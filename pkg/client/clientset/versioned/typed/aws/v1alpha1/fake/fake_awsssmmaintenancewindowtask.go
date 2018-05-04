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

// FakeAwsSsmMaintenanceWindowTasks implements AwsSsmMaintenanceWindowTaskInterface
type FakeAwsSsmMaintenanceWindowTasks struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsssmmaintenancewindowtasksResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsssmmaintenancewindowtasks"}

var awsssmmaintenancewindowtasksKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsSsmMaintenanceWindowTask"}

// Get takes name of the awsSsmMaintenanceWindowTask, and returns the corresponding awsSsmMaintenanceWindowTask object, and an error if there is any.
func (c *FakeAwsSsmMaintenanceWindowTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsssmmaintenancewindowtasksResource, c.ns, name), &v1alpha1.AwsSsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTask), err
}

// List takes label and field selectors, and returns the list of AwsSsmMaintenanceWindowTasks that match those selectors.
func (c *FakeAwsSsmMaintenanceWindowTasks) List(opts v1.ListOptions) (result *v1alpha1.AwsSsmMaintenanceWindowTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsssmmaintenancewindowtasksResource, awsssmmaintenancewindowtasksKind, c.ns, opts), &v1alpha1.AwsSsmMaintenanceWindowTaskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSsmMaintenanceWindowTaskList{}
	for _, item := range obj.(*v1alpha1.AwsSsmMaintenanceWindowTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSsmMaintenanceWindowTasks.
func (c *FakeAwsSsmMaintenanceWindowTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsssmmaintenancewindowtasksResource, c.ns, opts))

}

// Create takes the representation of a awsSsmMaintenanceWindowTask and creates it.  Returns the server's representation of the awsSsmMaintenanceWindowTask, and an error, if there is any.
func (c *FakeAwsSsmMaintenanceWindowTasks) Create(awsSsmMaintenanceWindowTask *v1alpha1.AwsSsmMaintenanceWindowTask) (result *v1alpha1.AwsSsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsssmmaintenancewindowtasksResource, c.ns, awsSsmMaintenanceWindowTask), &v1alpha1.AwsSsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTask), err
}

// Update takes the representation of a awsSsmMaintenanceWindowTask and updates it. Returns the server's representation of the awsSsmMaintenanceWindowTask, and an error, if there is any.
func (c *FakeAwsSsmMaintenanceWindowTasks) Update(awsSsmMaintenanceWindowTask *v1alpha1.AwsSsmMaintenanceWindowTask) (result *v1alpha1.AwsSsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsssmmaintenancewindowtasksResource, c.ns, awsSsmMaintenanceWindowTask), &v1alpha1.AwsSsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTask), err
}

// Delete takes name of the awsSsmMaintenanceWindowTask and deletes it. Returns an error if one occurs.
func (c *FakeAwsSsmMaintenanceWindowTasks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsssmmaintenancewindowtasksResource, c.ns, name), &v1alpha1.AwsSsmMaintenanceWindowTask{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSsmMaintenanceWindowTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsssmmaintenancewindowtasksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSsmMaintenanceWindowTaskList{})
	return err
}

// Patch applies the patch and returns the patched awsSsmMaintenanceWindowTask.
func (c *FakeAwsSsmMaintenanceWindowTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsssmmaintenancewindowtasksResource, c.ns, name, data, subresources...), &v1alpha1.AwsSsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSsmMaintenanceWindowTask), err
}
