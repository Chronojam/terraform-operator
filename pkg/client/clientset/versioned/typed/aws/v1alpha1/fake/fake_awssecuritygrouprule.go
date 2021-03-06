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

// FakeAwsSecurityGroupRules implements AwsSecurityGroupRuleInterface
type FakeAwsSecurityGroupRules struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awssecuritygrouprulesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awssecuritygrouprules"}

var awssecuritygrouprulesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSecurityGroupRule"}

// Get takes name of the awsSecurityGroupRule, and returns the corresponding awsSecurityGroupRule object, and an error if there is any.
func (c *FakeAwsSecurityGroupRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssecuritygrouprulesResource, c.ns, name), &v1alpha1.AwsSecurityGroupRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}

// List takes label and field selectors, and returns the list of AwsSecurityGroupRules that match those selectors.
func (c *FakeAwsSecurityGroupRules) List(opts v1.ListOptions) (result *v1alpha1.AwsSecurityGroupRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssecuritygrouprulesResource, awssecuritygrouprulesKind, c.ns, opts), &v1alpha1.AwsSecurityGroupRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSecurityGroupRuleList{}
	for _, item := range obj.(*v1alpha1.AwsSecurityGroupRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSecurityGroupRules.
func (c *FakeAwsSecurityGroupRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssecuritygrouprulesResource, c.ns, opts))

}

// Create takes the representation of a awsSecurityGroupRule and creates it.  Returns the server's representation of the awsSecurityGroupRule, and an error, if there is any.
func (c *FakeAwsSecurityGroupRules) Create(awsSecurityGroupRule *v1alpha1.AwsSecurityGroupRule) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssecuritygrouprulesResource, c.ns, awsSecurityGroupRule), &v1alpha1.AwsSecurityGroupRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}

// Update takes the representation of a awsSecurityGroupRule and updates it. Returns the server's representation of the awsSecurityGroupRule, and an error, if there is any.
func (c *FakeAwsSecurityGroupRules) Update(awsSecurityGroupRule *v1alpha1.AwsSecurityGroupRule) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssecuritygrouprulesResource, c.ns, awsSecurityGroupRule), &v1alpha1.AwsSecurityGroupRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}

// Delete takes name of the awsSecurityGroupRule and deletes it. Returns an error if one occurs.
func (c *FakeAwsSecurityGroupRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssecuritygrouprulesResource, c.ns, name), &v1alpha1.AwsSecurityGroupRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSecurityGroupRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssecuritygrouprulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSecurityGroupRuleList{})
	return err
}

// Patch applies the patch and returns the patched awsSecurityGroupRule.
func (c *FakeAwsSecurityGroupRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssecuritygrouprulesResource, c.ns, name, data, subresources...), &v1alpha1.AwsSecurityGroupRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}
