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

// FakeAwsWafregionalByteMatchSets implements AwsWafregionalByteMatchSetInterface
type FakeAwsWafregionalByteMatchSets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awswafregionalbytematchsetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awswafregionalbytematchsets"}

var awswafregionalbytematchsetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsWafregionalByteMatchSet"}

// Get takes name of the awsWafregionalByteMatchSet, and returns the corresponding awsWafregionalByteMatchSet object, and an error if there is any.
func (c *FakeAwsWafregionalByteMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafregionalbytematchsetsResource, c.ns, name), &v1alpha1.AwsWafregionalByteMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalByteMatchSet), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalByteMatchSets that match those selectors.
func (c *FakeAwsWafregionalByteMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalByteMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafregionalbytematchsetsResource, awswafregionalbytematchsetsKind, c.ns, opts), &v1alpha1.AwsWafregionalByteMatchSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafregionalByteMatchSetList{}
	for _, item := range obj.(*v1alpha1.AwsWafregionalByteMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalByteMatchSets.
func (c *FakeAwsWafregionalByteMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafregionalbytematchsetsResource, c.ns, opts))

}

// Create takes the representation of a awsWafregionalByteMatchSet and creates it.  Returns the server's representation of the awsWafregionalByteMatchSet, and an error, if there is any.
func (c *FakeAwsWafregionalByteMatchSets) Create(awsWafregionalByteMatchSet *v1alpha1.AwsWafregionalByteMatchSet) (result *v1alpha1.AwsWafregionalByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafregionalbytematchsetsResource, c.ns, awsWafregionalByteMatchSet), &v1alpha1.AwsWafregionalByteMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalByteMatchSet), err
}

// Update takes the representation of a awsWafregionalByteMatchSet and updates it. Returns the server's representation of the awsWafregionalByteMatchSet, and an error, if there is any.
func (c *FakeAwsWafregionalByteMatchSets) Update(awsWafregionalByteMatchSet *v1alpha1.AwsWafregionalByteMatchSet) (result *v1alpha1.AwsWafregionalByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafregionalbytematchsetsResource, c.ns, awsWafregionalByteMatchSet), &v1alpha1.AwsWafregionalByteMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalByteMatchSet), err
}

// Delete takes name of the awsWafregionalByteMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalByteMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafregionalbytematchsetsResource, c.ns, name), &v1alpha1.AwsWafregionalByteMatchSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalByteMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafregionalbytematchsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafregionalByteMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalByteMatchSet.
func (c *FakeAwsWafregionalByteMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalByteMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafregionalbytematchsetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsWafregionalByteMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalByteMatchSet), err
}
