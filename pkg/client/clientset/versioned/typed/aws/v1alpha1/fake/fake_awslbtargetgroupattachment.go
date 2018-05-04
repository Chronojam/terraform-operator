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

// FakeAwsLbTargetGroupAttachments implements AwsLbTargetGroupAttachmentInterface
type FakeAwsLbTargetGroupAttachments struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awslbtargetgroupattachmentsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awslbtargetgroupattachments"}

var awslbtargetgroupattachmentsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsLbTargetGroupAttachment"}

// Get takes name of the awsLbTargetGroupAttachment, and returns the corresponding awsLbTargetGroupAttachment object, and an error if there is any.
func (c *FakeAwsLbTargetGroupAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslbtargetgroupattachmentsResource, c.ns, name), &v1alpha1.AwsLbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroupAttachment), err
}

// List takes label and field selectors, and returns the list of AwsLbTargetGroupAttachments that match those selectors.
func (c *FakeAwsLbTargetGroupAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsLbTargetGroupAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslbtargetgroupattachmentsResource, awslbtargetgroupattachmentsKind, c.ns, opts), &v1alpha1.AwsLbTargetGroupAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLbTargetGroupAttachmentList{}
	for _, item := range obj.(*v1alpha1.AwsLbTargetGroupAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLbTargetGroupAttachments.
func (c *FakeAwsLbTargetGroupAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslbtargetgroupattachmentsResource, c.ns, opts))

}

// Create takes the representation of a awsLbTargetGroupAttachment and creates it.  Returns the server's representation of the awsLbTargetGroupAttachment, and an error, if there is any.
func (c *FakeAwsLbTargetGroupAttachments) Create(awsLbTargetGroupAttachment *v1alpha1.AwsLbTargetGroupAttachment) (result *v1alpha1.AwsLbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslbtargetgroupattachmentsResource, c.ns, awsLbTargetGroupAttachment), &v1alpha1.AwsLbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroupAttachment), err
}

// Update takes the representation of a awsLbTargetGroupAttachment and updates it. Returns the server's representation of the awsLbTargetGroupAttachment, and an error, if there is any.
func (c *FakeAwsLbTargetGroupAttachments) Update(awsLbTargetGroupAttachment *v1alpha1.AwsLbTargetGroupAttachment) (result *v1alpha1.AwsLbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslbtargetgroupattachmentsResource, c.ns, awsLbTargetGroupAttachment), &v1alpha1.AwsLbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroupAttachment), err
}

// Delete takes name of the awsLbTargetGroupAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAwsLbTargetGroupAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslbtargetgroupattachmentsResource, c.ns, name), &v1alpha1.AwsLbTargetGroupAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLbTargetGroupAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslbtargetgroupattachmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLbTargetGroupAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched awsLbTargetGroupAttachment.
func (c *FakeAwsLbTargetGroupAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslbtargetgroupattachmentsResource, c.ns, name, data, subresources...), &v1alpha1.AwsLbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbTargetGroupAttachment), err
}
