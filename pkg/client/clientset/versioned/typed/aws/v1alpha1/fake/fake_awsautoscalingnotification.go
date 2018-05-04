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

// FakeAwsAutoscalingNotifications implements AwsAutoscalingNotificationInterface
type FakeAwsAutoscalingNotifications struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsautoscalingnotificationsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsautoscalingnotifications"}

var awsautoscalingnotificationsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsAutoscalingNotification"}

// Get takes name of the awsAutoscalingNotification, and returns the corresponding awsAutoscalingNotification object, and an error if there is any.
func (c *FakeAwsAutoscalingNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsautoscalingnotificationsResource, c.ns, name), &v1alpha1.AwsAutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingNotification), err
}

// List takes label and field selectors, and returns the list of AwsAutoscalingNotifications that match those selectors.
func (c *FakeAwsAutoscalingNotifications) List(opts v1.ListOptions) (result *v1alpha1.AwsAutoscalingNotificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsautoscalingnotificationsResource, awsautoscalingnotificationsKind, c.ns, opts), &v1alpha1.AwsAutoscalingNotificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAutoscalingNotificationList{}
	for _, item := range obj.(*v1alpha1.AwsAutoscalingNotificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingNotifications.
func (c *FakeAwsAutoscalingNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsautoscalingnotificationsResource, c.ns, opts))

}

// Create takes the representation of a awsAutoscalingNotification and creates it.  Returns the server's representation of the awsAutoscalingNotification, and an error, if there is any.
func (c *FakeAwsAutoscalingNotifications) Create(awsAutoscalingNotification *v1alpha1.AwsAutoscalingNotification) (result *v1alpha1.AwsAutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsautoscalingnotificationsResource, c.ns, awsAutoscalingNotification), &v1alpha1.AwsAutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingNotification), err
}

// Update takes the representation of a awsAutoscalingNotification and updates it. Returns the server's representation of the awsAutoscalingNotification, and an error, if there is any.
func (c *FakeAwsAutoscalingNotifications) Update(awsAutoscalingNotification *v1alpha1.AwsAutoscalingNotification) (result *v1alpha1.AwsAutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsautoscalingnotificationsResource, c.ns, awsAutoscalingNotification), &v1alpha1.AwsAutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingNotification), err
}

// Delete takes name of the awsAutoscalingNotification and deletes it. Returns an error if one occurs.
func (c *FakeAwsAutoscalingNotifications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsautoscalingnotificationsResource, c.ns, name), &v1alpha1.AwsAutoscalingNotification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAutoscalingNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsautoscalingnotificationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAutoscalingNotificationList{})
	return err
}

// Patch applies the patch and returns the patched awsAutoscalingNotification.
func (c *FakeAwsAutoscalingNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsautoscalingnotificationsResource, c.ns, name, data, subresources...), &v1alpha1.AwsAutoscalingNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingNotification), err
}
