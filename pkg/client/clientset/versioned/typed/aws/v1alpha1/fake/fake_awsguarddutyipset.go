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

// FakeAwsGuarddutyIpsets implements AwsGuarddutyIpsetInterface
type FakeAwsGuarddutyIpsets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsguarddutyipsetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsguarddutyipsets"}

var awsguarddutyipsetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsGuarddutyIpset"}

// Get takes name of the awsGuarddutyIpset, and returns the corresponding awsGuarddutyIpset object, and an error if there is any.
func (c *FakeAwsGuarddutyIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsguarddutyipsetsResource, c.ns, name), &v1alpha1.AwsGuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyIpset), err
}

// List takes label and field selectors, and returns the list of AwsGuarddutyIpsets that match those selectors.
func (c *FakeAwsGuarddutyIpsets) List(opts v1.ListOptions) (result *v1alpha1.AwsGuarddutyIpsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsguarddutyipsetsResource, awsguarddutyipsetsKind, c.ns, opts), &v1alpha1.AwsGuarddutyIpsetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGuarddutyIpsetList{}
	for _, item := range obj.(*v1alpha1.AwsGuarddutyIpsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGuarddutyIpsets.
func (c *FakeAwsGuarddutyIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsguarddutyipsetsResource, c.ns, opts))

}

// Create takes the representation of a awsGuarddutyIpset and creates it.  Returns the server's representation of the awsGuarddutyIpset, and an error, if there is any.
func (c *FakeAwsGuarddutyIpsets) Create(awsGuarddutyIpset *v1alpha1.AwsGuarddutyIpset) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsguarddutyipsetsResource, c.ns, awsGuarddutyIpset), &v1alpha1.AwsGuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyIpset), err
}

// Update takes the representation of a awsGuarddutyIpset and updates it. Returns the server's representation of the awsGuarddutyIpset, and an error, if there is any.
func (c *FakeAwsGuarddutyIpsets) Update(awsGuarddutyIpset *v1alpha1.AwsGuarddutyIpset) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsguarddutyipsetsResource, c.ns, awsGuarddutyIpset), &v1alpha1.AwsGuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyIpset), err
}

// Delete takes name of the awsGuarddutyIpset and deletes it. Returns an error if one occurs.
func (c *FakeAwsGuarddutyIpsets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsguarddutyipsetsResource, c.ns, name), &v1alpha1.AwsGuarddutyIpset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGuarddutyIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsguarddutyipsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGuarddutyIpsetList{})
	return err
}

// Patch applies the patch and returns the patched awsGuarddutyIpset.
func (c *FakeAwsGuarddutyIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsguarddutyipsetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsGuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGuarddutyIpset), err
}
