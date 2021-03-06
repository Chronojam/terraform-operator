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

// FakeAwsSnsTopicPolicies implements AwsSnsTopicPolicyInterface
type FakeAwsSnsTopicPolicies struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awssnstopicpoliciesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awssnstopicpolicies"}

var awssnstopicpoliciesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSnsTopicPolicy"}

// Get takes name of the awsSnsTopicPolicy, and returns the corresponding awsSnsTopicPolicy object, and an error if there is any.
func (c *FakeAwsSnsTopicPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssnstopicpoliciesResource, c.ns, name), &v1alpha1.AwsSnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicPolicy), err
}

// List takes label and field selectors, and returns the list of AwsSnsTopicPolicies that match those selectors.
func (c *FakeAwsSnsTopicPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsSnsTopicPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssnstopicpoliciesResource, awssnstopicpoliciesKind, c.ns, opts), &v1alpha1.AwsSnsTopicPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSnsTopicPolicyList{}
	for _, item := range obj.(*v1alpha1.AwsSnsTopicPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSnsTopicPolicies.
func (c *FakeAwsSnsTopicPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssnstopicpoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsSnsTopicPolicy and creates it.  Returns the server's representation of the awsSnsTopicPolicy, and an error, if there is any.
func (c *FakeAwsSnsTopicPolicies) Create(awsSnsTopicPolicy *v1alpha1.AwsSnsTopicPolicy) (result *v1alpha1.AwsSnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssnstopicpoliciesResource, c.ns, awsSnsTopicPolicy), &v1alpha1.AwsSnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicPolicy), err
}

// Update takes the representation of a awsSnsTopicPolicy and updates it. Returns the server's representation of the awsSnsTopicPolicy, and an error, if there is any.
func (c *FakeAwsSnsTopicPolicies) Update(awsSnsTopicPolicy *v1alpha1.AwsSnsTopicPolicy) (result *v1alpha1.AwsSnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssnstopicpoliciesResource, c.ns, awsSnsTopicPolicy), &v1alpha1.AwsSnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicPolicy), err
}

// Delete takes name of the awsSnsTopicPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsSnsTopicPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssnstopicpoliciesResource, c.ns, name), &v1alpha1.AwsSnsTopicPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSnsTopicPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssnstopicpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSnsTopicPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsSnsTopicPolicy.
func (c *FakeAwsSnsTopicPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssnstopicpoliciesResource, c.ns, name, data, subresources...), &v1alpha1.AwsSnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSnsTopicPolicy), err
}
