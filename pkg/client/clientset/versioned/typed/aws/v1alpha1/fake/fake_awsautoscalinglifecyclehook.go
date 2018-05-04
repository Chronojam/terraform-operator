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

// FakeAwsAutoscalingLifecycleHooks implements AwsAutoscalingLifecycleHookInterface
type FakeAwsAutoscalingLifecycleHooks struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsautoscalinglifecyclehooksResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsautoscalinglifecyclehooks"}

var awsautoscalinglifecyclehooksKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsAutoscalingLifecycleHook"}

// Get takes name of the awsAutoscalingLifecycleHook, and returns the corresponding awsAutoscalingLifecycleHook object, and an error if there is any.
func (c *FakeAwsAutoscalingLifecycleHooks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsautoscalinglifecyclehooksResource, c.ns, name), &v1alpha1.AwsAutoscalingLifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingLifecycleHook), err
}

// List takes label and field selectors, and returns the list of AwsAutoscalingLifecycleHooks that match those selectors.
func (c *FakeAwsAutoscalingLifecycleHooks) List(opts v1.ListOptions) (result *v1alpha1.AwsAutoscalingLifecycleHookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsautoscalinglifecyclehooksResource, awsautoscalinglifecyclehooksKind, c.ns, opts), &v1alpha1.AwsAutoscalingLifecycleHookList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAutoscalingLifecycleHookList{}
	for _, item := range obj.(*v1alpha1.AwsAutoscalingLifecycleHookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingLifecycleHooks.
func (c *FakeAwsAutoscalingLifecycleHooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsautoscalinglifecyclehooksResource, c.ns, opts))

}

// Create takes the representation of a awsAutoscalingLifecycleHook and creates it.  Returns the server's representation of the awsAutoscalingLifecycleHook, and an error, if there is any.
func (c *FakeAwsAutoscalingLifecycleHooks) Create(awsAutoscalingLifecycleHook *v1alpha1.AwsAutoscalingLifecycleHook) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsautoscalinglifecyclehooksResource, c.ns, awsAutoscalingLifecycleHook), &v1alpha1.AwsAutoscalingLifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingLifecycleHook), err
}

// Update takes the representation of a awsAutoscalingLifecycleHook and updates it. Returns the server's representation of the awsAutoscalingLifecycleHook, and an error, if there is any.
func (c *FakeAwsAutoscalingLifecycleHooks) Update(awsAutoscalingLifecycleHook *v1alpha1.AwsAutoscalingLifecycleHook) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsautoscalinglifecyclehooksResource, c.ns, awsAutoscalingLifecycleHook), &v1alpha1.AwsAutoscalingLifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingLifecycleHook), err
}

// Delete takes name of the awsAutoscalingLifecycleHook and deletes it. Returns an error if one occurs.
func (c *FakeAwsAutoscalingLifecycleHooks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsautoscalinglifecyclehooksResource, c.ns, name), &v1alpha1.AwsAutoscalingLifecycleHook{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAutoscalingLifecycleHooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsautoscalinglifecyclehooksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAutoscalingLifecycleHookList{})
	return err
}

// Patch applies the patch and returns the patched awsAutoscalingLifecycleHook.
func (c *FakeAwsAutoscalingLifecycleHooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAutoscalingLifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsautoscalinglifecyclehooksResource, c.ns, name, data, subresources...), &v1alpha1.AwsAutoscalingLifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAutoscalingLifecycleHook), err
}
