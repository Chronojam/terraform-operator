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

// FakeAwsAlbs implements AwsAlbInterface
type FakeAwsAlbs struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsalbsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsalbs"}

var awsalbsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsAlb"}

// Get takes name of the awsAlb, and returns the corresponding awsAlb object, and an error if there is any.
func (c *FakeAwsAlbs) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsalbsResource, c.ns, name), &v1alpha1.AwsAlb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}

// List takes label and field selectors, and returns the list of AwsAlbs that match those selectors.
func (c *FakeAwsAlbs) List(opts v1.ListOptions) (result *v1alpha1.AwsAlbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsalbsResource, awsalbsKind, c.ns, opts), &v1alpha1.AwsAlbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAlbList{}
	for _, item := range obj.(*v1alpha1.AwsAlbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAlbs.
func (c *FakeAwsAlbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsalbsResource, c.ns, opts))

}

// Create takes the representation of a awsAlb and creates it.  Returns the server's representation of the awsAlb, and an error, if there is any.
func (c *FakeAwsAlbs) Create(awsAlb *v1alpha1.AwsAlb) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsalbsResource, c.ns, awsAlb), &v1alpha1.AwsAlb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}

// Update takes the representation of a awsAlb and updates it. Returns the server's representation of the awsAlb, and an error, if there is any.
func (c *FakeAwsAlbs) Update(awsAlb *v1alpha1.AwsAlb) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsalbsResource, c.ns, awsAlb), &v1alpha1.AwsAlb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}

// Delete takes name of the awsAlb and deletes it. Returns an error if one occurs.
func (c *FakeAwsAlbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsalbsResource, c.ns, name), &v1alpha1.AwsAlb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAlbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsalbsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAlbList{})
	return err
}

// Patch applies the patch and returns the patched awsAlb.
func (c *FakeAwsAlbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsalbsResource, c.ns, name, data, subresources...), &v1alpha1.AwsAlb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}
