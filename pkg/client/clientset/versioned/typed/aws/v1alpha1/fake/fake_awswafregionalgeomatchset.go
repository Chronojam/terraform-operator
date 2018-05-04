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

// FakeAwsWafregionalGeoMatchSets implements AwsWafregionalGeoMatchSetInterface
type FakeAwsWafregionalGeoMatchSets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awswafregionalgeomatchsetsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awswafregionalgeomatchsets"}

var awswafregionalgeomatchsetsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsWafregionalGeoMatchSet"}

// Get takes name of the awsWafregionalGeoMatchSet, and returns the corresponding awsWafregionalGeoMatchSet object, and an error if there is any.
func (c *FakeAwsWafregionalGeoMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafregionalgeomatchsetsResource, c.ns, name), &v1alpha1.AwsWafregionalGeoMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalGeoMatchSet), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalGeoMatchSets that match those selectors.
func (c *FakeAwsWafregionalGeoMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalGeoMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafregionalgeomatchsetsResource, awswafregionalgeomatchsetsKind, c.ns, opts), &v1alpha1.AwsWafregionalGeoMatchSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafregionalGeoMatchSetList{}
	for _, item := range obj.(*v1alpha1.AwsWafregionalGeoMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalGeoMatchSets.
func (c *FakeAwsWafregionalGeoMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafregionalgeomatchsetsResource, c.ns, opts))

}

// Create takes the representation of a awsWafregionalGeoMatchSet and creates it.  Returns the server's representation of the awsWafregionalGeoMatchSet, and an error, if there is any.
func (c *FakeAwsWafregionalGeoMatchSets) Create(awsWafregionalGeoMatchSet *v1alpha1.AwsWafregionalGeoMatchSet) (result *v1alpha1.AwsWafregionalGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafregionalgeomatchsetsResource, c.ns, awsWafregionalGeoMatchSet), &v1alpha1.AwsWafregionalGeoMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalGeoMatchSet), err
}

// Update takes the representation of a awsWafregionalGeoMatchSet and updates it. Returns the server's representation of the awsWafregionalGeoMatchSet, and an error, if there is any.
func (c *FakeAwsWafregionalGeoMatchSets) Update(awsWafregionalGeoMatchSet *v1alpha1.AwsWafregionalGeoMatchSet) (result *v1alpha1.AwsWafregionalGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafregionalgeomatchsetsResource, c.ns, awsWafregionalGeoMatchSet), &v1alpha1.AwsWafregionalGeoMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalGeoMatchSet), err
}

// Delete takes name of the awsWafregionalGeoMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalGeoMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafregionalgeomatchsetsResource, c.ns, name), &v1alpha1.AwsWafregionalGeoMatchSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalGeoMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafregionalgeomatchsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafregionalGeoMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalGeoMatchSet.
func (c *FakeAwsWafregionalGeoMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalGeoMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafregionalgeomatchsetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsWafregionalGeoMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalGeoMatchSet), err
}
