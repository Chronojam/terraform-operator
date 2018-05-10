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

// FakeAwsSqsQueues implements AwsSqsQueueInterface
type FakeAwsSqsQueues struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awssqsqueuesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awssqsqueues"}

var awssqsqueuesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSqsQueue"}

// Get takes name of the awsSqsQueue, and returns the corresponding awsSqsQueue object, and an error if there is any.
func (c *FakeAwsSqsQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssqsqueuesResource, c.ns, name), &v1alpha1.AwsSqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}

// List takes label and field selectors, and returns the list of AwsSqsQueues that match those selectors.
func (c *FakeAwsSqsQueues) List(opts v1.ListOptions) (result *v1alpha1.AwsSqsQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssqsqueuesResource, awssqsqueuesKind, c.ns, opts), &v1alpha1.AwsSqsQueueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSqsQueueList{}
	for _, item := range obj.(*v1alpha1.AwsSqsQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSqsQueues.
func (c *FakeAwsSqsQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssqsqueuesResource, c.ns, opts))

}

// Create takes the representation of a awsSqsQueue and creates it.  Returns the server's representation of the awsSqsQueue, and an error, if there is any.
func (c *FakeAwsSqsQueues) Create(awsSqsQueue *v1alpha1.AwsSqsQueue) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssqsqueuesResource, c.ns, awsSqsQueue), &v1alpha1.AwsSqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}

// Update takes the representation of a awsSqsQueue and updates it. Returns the server's representation of the awsSqsQueue, and an error, if there is any.
func (c *FakeAwsSqsQueues) Update(awsSqsQueue *v1alpha1.AwsSqsQueue) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssqsqueuesResource, c.ns, awsSqsQueue), &v1alpha1.AwsSqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}

// Delete takes name of the awsSqsQueue and deletes it. Returns an error if one occurs.
func (c *FakeAwsSqsQueues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssqsqueuesResource, c.ns, name), &v1alpha1.AwsSqsQueue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSqsQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssqsqueuesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSqsQueueList{})
	return err
}

// Patch applies the patch and returns the patched awsSqsQueue.
func (c *FakeAwsSqsQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssqsqueuesResource, c.ns, name, data, subresources...), &v1alpha1.AwsSqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsQueue), err
}
