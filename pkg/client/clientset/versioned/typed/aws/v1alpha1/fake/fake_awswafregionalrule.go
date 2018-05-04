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

// FakeAwsWafregionalRules implements AwsWafregionalRuleInterface
type FakeAwsWafregionalRules struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awswafregionalrulesResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awswafregionalrules"}

var awswafregionalrulesKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsWafregionalRule"}

// Get takes name of the awsWafregionalRule, and returns the corresponding awsWafregionalRule object, and an error if there is any.
func (c *FakeAwsWafregionalRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafregionalrulesResource, c.ns, name), &v1alpha1.AwsWafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRule), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalRules that match those selectors.
func (c *FakeAwsWafregionalRules) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafregionalrulesResource, awswafregionalrulesKind, c.ns, opts), &v1alpha1.AwsWafregionalRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafregionalRuleList{}
	for _, item := range obj.(*v1alpha1.AwsWafregionalRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalRules.
func (c *FakeAwsWafregionalRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafregionalrulesResource, c.ns, opts))

}

// Create takes the representation of a awsWafregionalRule and creates it.  Returns the server's representation of the awsWafregionalRule, and an error, if there is any.
func (c *FakeAwsWafregionalRules) Create(awsWafregionalRule *v1alpha1.AwsWafregionalRule) (result *v1alpha1.AwsWafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafregionalrulesResource, c.ns, awsWafregionalRule), &v1alpha1.AwsWafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRule), err
}

// Update takes the representation of a awsWafregionalRule and updates it. Returns the server's representation of the awsWafregionalRule, and an error, if there is any.
func (c *FakeAwsWafregionalRules) Update(awsWafregionalRule *v1alpha1.AwsWafregionalRule) (result *v1alpha1.AwsWafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafregionalrulesResource, c.ns, awsWafregionalRule), &v1alpha1.AwsWafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRule), err
}

// Delete takes name of the awsWafregionalRule and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafregionalrulesResource, c.ns, name), &v1alpha1.AwsWafregionalRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafregionalrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafregionalRuleList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalRule.
func (c *FakeAwsWafregionalRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafregionalrulesResource, c.ns, name, data, subresources...), &v1alpha1.AwsWafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafregionalRule), err
}
