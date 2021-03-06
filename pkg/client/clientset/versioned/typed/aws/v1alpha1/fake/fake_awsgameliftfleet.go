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

// FakeAwsGameliftFleets implements AwsGameliftFleetInterface
type FakeAwsGameliftFleets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsgameliftfleetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsgameliftfleets"}

var awsgameliftfleetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsGameliftFleet"}

// Get takes name of the awsGameliftFleet, and returns the corresponding awsGameliftFleet object, and an error if there is any.
func (c *FakeAwsGameliftFleets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGameliftFleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsgameliftfleetsResource, c.ns, name), &v1alpha1.AwsGameliftFleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftFleet), err
}

// List takes label and field selectors, and returns the list of AwsGameliftFleets that match those selectors.
func (c *FakeAwsGameliftFleets) List(opts v1.ListOptions) (result *v1alpha1.AwsGameliftFleetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsgameliftfleetsResource, awsgameliftfleetsKind, c.ns, opts), &v1alpha1.AwsGameliftFleetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGameliftFleetList{}
	for _, item := range obj.(*v1alpha1.AwsGameliftFleetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGameliftFleets.
func (c *FakeAwsGameliftFleets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsgameliftfleetsResource, c.ns, opts))

}

// Create takes the representation of a awsGameliftFleet and creates it.  Returns the server's representation of the awsGameliftFleet, and an error, if there is any.
func (c *FakeAwsGameliftFleets) Create(awsGameliftFleet *v1alpha1.AwsGameliftFleet) (result *v1alpha1.AwsGameliftFleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsgameliftfleetsResource, c.ns, awsGameliftFleet), &v1alpha1.AwsGameliftFleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftFleet), err
}

// Update takes the representation of a awsGameliftFleet and updates it. Returns the server's representation of the awsGameliftFleet, and an error, if there is any.
func (c *FakeAwsGameliftFleets) Update(awsGameliftFleet *v1alpha1.AwsGameliftFleet) (result *v1alpha1.AwsGameliftFleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsgameliftfleetsResource, c.ns, awsGameliftFleet), &v1alpha1.AwsGameliftFleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftFleet), err
}

// Delete takes name of the awsGameliftFleet and deletes it. Returns an error if one occurs.
func (c *FakeAwsGameliftFleets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsgameliftfleetsResource, c.ns, name), &v1alpha1.AwsGameliftFleet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGameliftFleets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsgameliftfleetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGameliftFleetList{})
	return err
}

// Patch applies the patch and returns the patched awsGameliftFleet.
func (c *FakeAwsGameliftFleets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGameliftFleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsgameliftfleetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsGameliftFleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftFleet), err
}
