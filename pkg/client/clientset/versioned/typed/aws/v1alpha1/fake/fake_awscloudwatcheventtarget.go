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

// FakeAwsCloudwatchEventTargets implements AwsCloudwatchEventTargetInterface
type FakeAwsCloudwatchEventTargets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awscloudwatcheventtargetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awscloudwatcheventtargets"}

var awscloudwatcheventtargetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsCloudwatchEventTarget"}

// Get takes name of the awsCloudwatchEventTarget, and returns the corresponding awsCloudwatchEventTarget object, and an error if there is any.
func (c *FakeAwsCloudwatchEventTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscloudwatcheventtargetsResource, c.ns, name), &v1alpha1.AwsCloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventTarget), err
}

// List takes label and field selectors, and returns the list of AwsCloudwatchEventTargets that match those selectors.
func (c *FakeAwsCloudwatchEventTargets) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchEventTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscloudwatcheventtargetsResource, awscloudwatcheventtargetsKind, c.ns, opts), &v1alpha1.AwsCloudwatchEventTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCloudwatchEventTargetList{}
	for _, item := range obj.(*v1alpha1.AwsCloudwatchEventTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchEventTargets.
func (c *FakeAwsCloudwatchEventTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscloudwatcheventtargetsResource, c.ns, opts))

}

// Create takes the representation of a awsCloudwatchEventTarget and creates it.  Returns the server's representation of the awsCloudwatchEventTarget, and an error, if there is any.
func (c *FakeAwsCloudwatchEventTargets) Create(awsCloudwatchEventTarget *v1alpha1.AwsCloudwatchEventTarget) (result *v1alpha1.AwsCloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscloudwatcheventtargetsResource, c.ns, awsCloudwatchEventTarget), &v1alpha1.AwsCloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventTarget), err
}

// Update takes the representation of a awsCloudwatchEventTarget and updates it. Returns the server's representation of the awsCloudwatchEventTarget, and an error, if there is any.
func (c *FakeAwsCloudwatchEventTargets) Update(awsCloudwatchEventTarget *v1alpha1.AwsCloudwatchEventTarget) (result *v1alpha1.AwsCloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscloudwatcheventtargetsResource, c.ns, awsCloudwatchEventTarget), &v1alpha1.AwsCloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventTarget), err
}

// Delete takes name of the awsCloudwatchEventTarget and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudwatchEventTargets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscloudwatcheventtargetsResource, c.ns, name), &v1alpha1.AwsCloudwatchEventTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudwatchEventTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscloudwatcheventtargetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCloudwatchEventTargetList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudwatchEventTarget.
func (c *FakeAwsCloudwatchEventTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscloudwatcheventtargetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsCloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventTarget), err
}
