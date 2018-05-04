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

// FakeAwsDxLags implements AwsDxLagInterface
type FakeAwsDxLags struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsdxlagsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsdxlags"}

var awsdxlagsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsDxLag"}

// Get takes name of the awsDxLag, and returns the corresponding awsDxLag object, and an error if there is any.
func (c *FakeAwsDxLags) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDxLag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdxlagsResource, c.ns, name), &v1alpha1.AwsDxLag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxLag), err
}

// List takes label and field selectors, and returns the list of AwsDxLags that match those selectors.
func (c *FakeAwsDxLags) List(opts v1.ListOptions) (result *v1alpha1.AwsDxLagList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdxlagsResource, awsdxlagsKind, c.ns, opts), &v1alpha1.AwsDxLagList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDxLagList{}
	for _, item := range obj.(*v1alpha1.AwsDxLagList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDxLags.
func (c *FakeAwsDxLags) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdxlagsResource, c.ns, opts))

}

// Create takes the representation of a awsDxLag and creates it.  Returns the server's representation of the awsDxLag, and an error, if there is any.
func (c *FakeAwsDxLags) Create(awsDxLag *v1alpha1.AwsDxLag) (result *v1alpha1.AwsDxLag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdxlagsResource, c.ns, awsDxLag), &v1alpha1.AwsDxLag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxLag), err
}

// Update takes the representation of a awsDxLag and updates it. Returns the server's representation of the awsDxLag, and an error, if there is any.
func (c *FakeAwsDxLags) Update(awsDxLag *v1alpha1.AwsDxLag) (result *v1alpha1.AwsDxLag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdxlagsResource, c.ns, awsDxLag), &v1alpha1.AwsDxLag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxLag), err
}

// Delete takes name of the awsDxLag and deletes it. Returns an error if one occurs.
func (c *FakeAwsDxLags) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdxlagsResource, c.ns, name), &v1alpha1.AwsDxLag{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDxLags) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdxlagsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDxLagList{})
	return err
}

// Patch applies the patch and returns the patched awsDxLag.
func (c *FakeAwsDxLags) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDxLag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdxlagsResource, c.ns, name, data, subresources...), &v1alpha1.AwsDxLag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxLag), err
}
