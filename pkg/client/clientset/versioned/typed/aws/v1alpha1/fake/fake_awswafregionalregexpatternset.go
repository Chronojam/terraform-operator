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

// FakeAwsWafregionalRegexPatternSets implements AwsWafregionalRegexPatternSetInterface
type FakeAwsWafregionalRegexPatternSets struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awswafregionalregexpatternsetsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awswafregionalregexpatternsets"}

var awswafregionalregexpatternsetsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsWafregionalRegexPatternSet"}

// Get takes name of the awsWafregionalRegexPatternSet, and returns the corresponding awsWafregionalRegexPatternSet object, and an error if there is any.
func (c *FakeAwsWafregionalRegexPatternSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalRegexPatternSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafregionalregexpatternsetsResource, c.ns, name), &v1alpha1.AwsWafregionalRegexPatternSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexPatternSet), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalRegexPatternSets that match those selectors.
func (c *FakeAwsWafregionalRegexPatternSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalRegexPatternSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafregionalregexpatternsetsResource, awswafregionalregexpatternsetsKind, c.ns, opts), &v1alpha1.AwsWafregionalRegexPatternSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafregionalRegexPatternSetList{}
	for _, item := range obj.(*v1alpha1.AwsWafregionalRegexPatternSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalRegexPatternSets.
func (c *FakeAwsWafregionalRegexPatternSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafregionalregexpatternsetsResource, c.ns, opts))

}

// Create takes the representation of a awsWafregionalRegexPatternSet and creates it.  Returns the server's representation of the awsWafregionalRegexPatternSet, and an error, if there is any.
func (c *FakeAwsWafregionalRegexPatternSets) Create(awsWafregionalRegexPatternSet *v1alpha1.AwsWafregionalRegexPatternSet) (result *v1alpha1.AwsWafregionalRegexPatternSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafregionalregexpatternsetsResource, c.ns, awsWafregionalRegexPatternSet), &v1alpha1.AwsWafregionalRegexPatternSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexPatternSet), err
}

// Update takes the representation of a awsWafregionalRegexPatternSet and updates it. Returns the server's representation of the awsWafregionalRegexPatternSet, and an error, if there is any.
func (c *FakeAwsWafregionalRegexPatternSets) Update(awsWafregionalRegexPatternSet *v1alpha1.AwsWafregionalRegexPatternSet) (result *v1alpha1.AwsWafregionalRegexPatternSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafregionalregexpatternsetsResource, c.ns, awsWafregionalRegexPatternSet), &v1alpha1.AwsWafregionalRegexPatternSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexPatternSet), err
}

// Delete takes name of the awsWafregionalRegexPatternSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalRegexPatternSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafregionalregexpatternsetsResource, c.ns, name), &v1alpha1.AwsWafregionalRegexPatternSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalRegexPatternSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafregionalregexpatternsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafregionalRegexPatternSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalRegexPatternSet.
func (c *FakeAwsWafregionalRegexPatternSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRegexPatternSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafregionalregexpatternsetsResource, c.ns, name, data, subresources...), &v1alpha1.AwsWafregionalRegexPatternSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRegexPatternSet), err
}
