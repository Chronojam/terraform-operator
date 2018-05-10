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

// FakeAwsElbs implements AwsElbInterface
type FakeAwsElbs struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awselbsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awselbs"}

var awselbsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsElb"}

// Get takes name of the awsElb, and returns the corresponding awsElb object, and an error if there is any.
func (c *FakeAwsElbs) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awselbsResource, c.ns, name), &v1alpha1.AwsElb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElb), err
}

// List takes label and field selectors, and returns the list of AwsElbs that match those selectors.
func (c *FakeAwsElbs) List(opts v1.ListOptions) (result *v1alpha1.AwsElbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awselbsResource, awselbsKind, c.ns, opts), &v1alpha1.AwsElbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElbList{}
	for _, item := range obj.(*v1alpha1.AwsElbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElbs.
func (c *FakeAwsElbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awselbsResource, c.ns, opts))

}

// Create takes the representation of a awsElb and creates it.  Returns the server's representation of the awsElb, and an error, if there is any.
func (c *FakeAwsElbs) Create(awsElb *v1alpha1.AwsElb) (result *v1alpha1.AwsElb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awselbsResource, c.ns, awsElb), &v1alpha1.AwsElb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElb), err
}

// Update takes the representation of a awsElb and updates it. Returns the server's representation of the awsElb, and an error, if there is any.
func (c *FakeAwsElbs) Update(awsElb *v1alpha1.AwsElb) (result *v1alpha1.AwsElb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awselbsResource, c.ns, awsElb), &v1alpha1.AwsElb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElb), err
}

// Delete takes name of the awsElb and deletes it. Returns an error if one occurs.
func (c *FakeAwsElbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awselbsResource, c.ns, name), &v1alpha1.AwsElb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awselbsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElbList{})
	return err
}

// Patch applies the patch and returns the patched awsElb.
func (c *FakeAwsElbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awselbsResource, c.ns, name, data, subresources...), &v1alpha1.AwsElb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElb), err
}
