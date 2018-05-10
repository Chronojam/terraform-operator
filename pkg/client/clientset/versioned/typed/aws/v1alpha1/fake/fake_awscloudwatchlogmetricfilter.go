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

// FakeAwsCloudwatchLogMetricFilters implements AwsCloudwatchLogMetricFilterInterface
type FakeAwsCloudwatchLogMetricFilters struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awscloudwatchlogmetricfiltersResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awscloudwatchlogmetricfilters"}

var awscloudwatchlogmetricfiltersKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsCloudwatchLogMetricFilter"}

// Get takes name of the awsCloudwatchLogMetricFilter, and returns the corresponding awsCloudwatchLogMetricFilter object, and an error if there is any.
func (c *FakeAwsCloudwatchLogMetricFilters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscloudwatchlogmetricfiltersResource, c.ns, name), &v1alpha1.AwsCloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchLogMetricFilter), err
}

// List takes label and field selectors, and returns the list of AwsCloudwatchLogMetricFilters that match those selectors.
func (c *FakeAwsCloudwatchLogMetricFilters) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchLogMetricFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscloudwatchlogmetricfiltersResource, awscloudwatchlogmetricfiltersKind, c.ns, opts), &v1alpha1.AwsCloudwatchLogMetricFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCloudwatchLogMetricFilterList{}
	for _, item := range obj.(*v1alpha1.AwsCloudwatchLogMetricFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchLogMetricFilters.
func (c *FakeAwsCloudwatchLogMetricFilters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscloudwatchlogmetricfiltersResource, c.ns, opts))

}

// Create takes the representation of a awsCloudwatchLogMetricFilter and creates it.  Returns the server's representation of the awsCloudwatchLogMetricFilter, and an error, if there is any.
func (c *FakeAwsCloudwatchLogMetricFilters) Create(awsCloudwatchLogMetricFilter *v1alpha1.AwsCloudwatchLogMetricFilter) (result *v1alpha1.AwsCloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscloudwatchlogmetricfiltersResource, c.ns, awsCloudwatchLogMetricFilter), &v1alpha1.AwsCloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchLogMetricFilter), err
}

// Update takes the representation of a awsCloudwatchLogMetricFilter and updates it. Returns the server's representation of the awsCloudwatchLogMetricFilter, and an error, if there is any.
func (c *FakeAwsCloudwatchLogMetricFilters) Update(awsCloudwatchLogMetricFilter *v1alpha1.AwsCloudwatchLogMetricFilter) (result *v1alpha1.AwsCloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscloudwatchlogmetricfiltersResource, c.ns, awsCloudwatchLogMetricFilter), &v1alpha1.AwsCloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchLogMetricFilter), err
}

// Delete takes name of the awsCloudwatchLogMetricFilter and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudwatchLogMetricFilters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscloudwatchlogmetricfiltersResource, c.ns, name), &v1alpha1.AwsCloudwatchLogMetricFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudwatchLogMetricFilters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscloudwatchlogmetricfiltersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCloudwatchLogMetricFilterList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudwatchLogMetricFilter.
func (c *FakeAwsCloudwatchLogMetricFilters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscloudwatchlogmetricfiltersResource, c.ns, name, data, subresources...), &v1alpha1.AwsCloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchLogMetricFilter), err
}
