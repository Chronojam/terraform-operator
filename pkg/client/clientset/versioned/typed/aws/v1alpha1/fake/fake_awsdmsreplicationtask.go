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

// FakeAwsDmsReplicationTasks implements AwsDmsReplicationTaskInterface
type FakeAwsDmsReplicationTasks struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsdmsreplicationtasksResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsdmsreplicationtasks"}

var awsdmsreplicationtasksKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsDmsReplicationTask"}

// Get takes name of the awsDmsReplicationTask, and returns the corresponding awsDmsReplicationTask object, and an error if there is any.
func (c *FakeAwsDmsReplicationTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdmsreplicationtasksResource, c.ns, name), &v1alpha1.AwsDmsReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDmsReplicationTask), err
}

// List takes label and field selectors, and returns the list of AwsDmsReplicationTasks that match those selectors.
func (c *FakeAwsDmsReplicationTasks) List(opts v1.ListOptions) (result *v1alpha1.AwsDmsReplicationTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdmsreplicationtasksResource, awsdmsreplicationtasksKind, c.ns, opts), &v1alpha1.AwsDmsReplicationTaskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDmsReplicationTaskList{}
	for _, item := range obj.(*v1alpha1.AwsDmsReplicationTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDmsReplicationTasks.
func (c *FakeAwsDmsReplicationTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdmsreplicationtasksResource, c.ns, opts))

}

// Create takes the representation of a awsDmsReplicationTask and creates it.  Returns the server's representation of the awsDmsReplicationTask, and an error, if there is any.
func (c *FakeAwsDmsReplicationTasks) Create(awsDmsReplicationTask *v1alpha1.AwsDmsReplicationTask) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdmsreplicationtasksResource, c.ns, awsDmsReplicationTask), &v1alpha1.AwsDmsReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDmsReplicationTask), err
}

// Update takes the representation of a awsDmsReplicationTask and updates it. Returns the server's representation of the awsDmsReplicationTask, and an error, if there is any.
func (c *FakeAwsDmsReplicationTasks) Update(awsDmsReplicationTask *v1alpha1.AwsDmsReplicationTask) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdmsreplicationtasksResource, c.ns, awsDmsReplicationTask), &v1alpha1.AwsDmsReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDmsReplicationTask), err
}

// Delete takes name of the awsDmsReplicationTask and deletes it. Returns an error if one occurs.
func (c *FakeAwsDmsReplicationTasks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdmsreplicationtasksResource, c.ns, name), &v1alpha1.AwsDmsReplicationTask{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDmsReplicationTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdmsreplicationtasksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDmsReplicationTaskList{})
	return err
}

// Patch applies the patch and returns the patched awsDmsReplicationTask.
func (c *FakeAwsDmsReplicationTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDmsReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdmsreplicationtasksResource, c.ns, name, data, subresources...), &v1alpha1.AwsDmsReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDmsReplicationTask), err
}
