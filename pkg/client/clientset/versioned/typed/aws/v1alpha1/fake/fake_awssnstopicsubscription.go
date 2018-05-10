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

// FakeAwsSnsTopicSubscriptions implements AwsSnsTopicSubscriptionInterface
type FakeAwsSnsTopicSubscriptions struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awssnstopicsubscriptionsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awssnstopicsubscriptions"}

var awssnstopicsubscriptionsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSnsTopicSubscription"}

// Get takes name of the awsSnsTopicSubscription, and returns the corresponding awsSnsTopicSubscription object, and an error if there is any.
func (c *FakeAwsSnsTopicSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssnstopicsubscriptionsResource, c.ns, name), &v1alpha1.AwsSnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicSubscription), err
}

// List takes label and field selectors, and returns the list of AwsSnsTopicSubscriptions that match those selectors.
func (c *FakeAwsSnsTopicSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.AwsSnsTopicSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssnstopicsubscriptionsResource, awssnstopicsubscriptionsKind, c.ns, opts), &v1alpha1.AwsSnsTopicSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSnsTopicSubscriptionList{}
	for _, item := range obj.(*v1alpha1.AwsSnsTopicSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSnsTopicSubscriptions.
func (c *FakeAwsSnsTopicSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssnstopicsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a awsSnsTopicSubscription and creates it.  Returns the server's representation of the awsSnsTopicSubscription, and an error, if there is any.
func (c *FakeAwsSnsTopicSubscriptions) Create(awsSnsTopicSubscription *v1alpha1.AwsSnsTopicSubscription) (result *v1alpha1.AwsSnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssnstopicsubscriptionsResource, c.ns, awsSnsTopicSubscription), &v1alpha1.AwsSnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicSubscription), err
}

// Update takes the representation of a awsSnsTopicSubscription and updates it. Returns the server's representation of the awsSnsTopicSubscription, and an error, if there is any.
func (c *FakeAwsSnsTopicSubscriptions) Update(awsSnsTopicSubscription *v1alpha1.AwsSnsTopicSubscription) (result *v1alpha1.AwsSnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssnstopicsubscriptionsResource, c.ns, awsSnsTopicSubscription), &v1alpha1.AwsSnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicSubscription), err
}

// Delete takes name of the awsSnsTopicSubscription and deletes it. Returns an error if one occurs.
func (c *FakeAwsSnsTopicSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssnstopicsubscriptionsResource, c.ns, name), &v1alpha1.AwsSnsTopicSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSnsTopicSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssnstopicsubscriptionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSnsTopicSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched awsSnsTopicSubscription.
func (c *FakeAwsSnsTopicSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssnstopicsubscriptionsResource, c.ns, name, data, subresources...), &v1alpha1.AwsSnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicSubscription), err
}
