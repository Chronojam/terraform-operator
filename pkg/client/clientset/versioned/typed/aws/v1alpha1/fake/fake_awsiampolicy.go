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

// FakeAwsIamPolicies implements AwsIamPolicyInterface
type FakeAwsIamPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsiampoliciesResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsiampolicies"}

var awsiampoliciesKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsIamPolicy"}

// Get takes name of the awsIamPolicy, and returns the corresponding awsIamPolicy object, and an error if there is any.
func (c *FakeAwsIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiampoliciesResource, c.ns, name), &v1alpha1.AwsIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamPolicy), err
}

// List takes label and field selectors, and returns the list of AwsIamPolicies that match those selectors.
func (c *FakeAwsIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiampoliciesResource, awsiampoliciesKind, c.ns, opts), &v1alpha1.AwsIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamPolicyList{}
	for _, item := range obj.(*v1alpha1.AwsIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamPolicies.
func (c *FakeAwsIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiampoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsIamPolicy and creates it.  Returns the server's representation of the awsIamPolicy, and an error, if there is any.
func (c *FakeAwsIamPolicies) Create(awsIamPolicy *v1alpha1.AwsIamPolicy) (result *v1alpha1.AwsIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiampoliciesResource, c.ns, awsIamPolicy), &v1alpha1.AwsIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamPolicy), err
}

// Update takes the representation of a awsIamPolicy and updates it. Returns the server's representation of the awsIamPolicy, and an error, if there is any.
func (c *FakeAwsIamPolicies) Update(awsIamPolicy *v1alpha1.AwsIamPolicy) (result *v1alpha1.AwsIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiampoliciesResource, c.ns, awsIamPolicy), &v1alpha1.AwsIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamPolicy), err
}

// Delete takes name of the awsIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiampoliciesResource, c.ns, name), &v1alpha1.AwsIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiampoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsIamPolicy.
func (c *FakeAwsIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiampoliciesResource, c.ns, name, data, subresources...), &v1alpha1.AwsIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamPolicy), err
}
