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

// FakeAwsSpotFleetRequests implements AwsSpotFleetRequestInterface
type FakeAwsSpotFleetRequests struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsspotfleetrequestsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsspotfleetrequests"}

var awsspotfleetrequestsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSpotFleetRequest"}

// Get takes name of the awsSpotFleetRequest, and returns the corresponding awsSpotFleetRequest object, and an error if there is any.
func (c *FakeAwsSpotFleetRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsspotfleetrequestsResource, c.ns, name), &v1alpha1.AwsSpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSpotFleetRequest), err
}

// List takes label and field selectors, and returns the list of AwsSpotFleetRequests that match those selectors.
func (c *FakeAwsSpotFleetRequests) List(opts v1.ListOptions) (result *v1alpha1.AwsSpotFleetRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsspotfleetrequestsResource, awsspotfleetrequestsKind, c.ns, opts), &v1alpha1.AwsSpotFleetRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSpotFleetRequestList{}
	for _, item := range obj.(*v1alpha1.AwsSpotFleetRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSpotFleetRequests.
func (c *FakeAwsSpotFleetRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsspotfleetrequestsResource, c.ns, opts))

}

// Create takes the representation of a awsSpotFleetRequest and creates it.  Returns the server's representation of the awsSpotFleetRequest, and an error, if there is any.
func (c *FakeAwsSpotFleetRequests) Create(awsSpotFleetRequest *v1alpha1.AwsSpotFleetRequest) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsspotfleetrequestsResource, c.ns, awsSpotFleetRequest), &v1alpha1.AwsSpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSpotFleetRequest), err
}

// Update takes the representation of a awsSpotFleetRequest and updates it. Returns the server's representation of the awsSpotFleetRequest, and an error, if there is any.
func (c *FakeAwsSpotFleetRequests) Update(awsSpotFleetRequest *v1alpha1.AwsSpotFleetRequest) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsspotfleetrequestsResource, c.ns, awsSpotFleetRequest), &v1alpha1.AwsSpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSpotFleetRequest), err
}

// Delete takes name of the awsSpotFleetRequest and deletes it. Returns an error if one occurs.
func (c *FakeAwsSpotFleetRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsspotfleetrequestsResource, c.ns, name), &v1alpha1.AwsSpotFleetRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSpotFleetRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsspotfleetrequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSpotFleetRequestList{})
	return err
}

// Patch applies the patch and returns the patched awsSpotFleetRequest.
func (c *FakeAwsSpotFleetRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsspotfleetrequestsResource, c.ns, name, data, subresources...), &v1alpha1.AwsSpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSpotFleetRequest), err
}
