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

// FakeAwsIamAccountPasswordPolicies implements AwsIamAccountPasswordPolicyInterface
type FakeAwsIamAccountPasswordPolicies struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsiamaccountpasswordpoliciesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsiamaccountpasswordpolicies"}

var awsiamaccountpasswordpoliciesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsIamAccountPasswordPolicy"}

// Get takes name of the awsIamAccountPasswordPolicy, and returns the corresponding awsIamAccountPasswordPolicy object, and an error if there is any.
func (c *FakeAwsIamAccountPasswordPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamAccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiamaccountpasswordpoliciesResource, c.ns, name), &v1alpha1.AwsIamAccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamAccountPasswordPolicy), err
}

// List takes label and field selectors, and returns the list of AwsIamAccountPasswordPolicies that match those selectors.
func (c *FakeAwsIamAccountPasswordPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsIamAccountPasswordPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiamaccountpasswordpoliciesResource, awsiamaccountpasswordpoliciesKind, c.ns, opts), &v1alpha1.AwsIamAccountPasswordPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamAccountPasswordPolicyList{}
	for _, item := range obj.(*v1alpha1.AwsIamAccountPasswordPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamAccountPasswordPolicies.
func (c *FakeAwsIamAccountPasswordPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiamaccountpasswordpoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsIamAccountPasswordPolicy and creates it.  Returns the server's representation of the awsIamAccountPasswordPolicy, and an error, if there is any.
func (c *FakeAwsIamAccountPasswordPolicies) Create(awsIamAccountPasswordPolicy *v1alpha1.AwsIamAccountPasswordPolicy) (result *v1alpha1.AwsIamAccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiamaccountpasswordpoliciesResource, c.ns, awsIamAccountPasswordPolicy), &v1alpha1.AwsIamAccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamAccountPasswordPolicy), err
}

// Update takes the representation of a awsIamAccountPasswordPolicy and updates it. Returns the server's representation of the awsIamAccountPasswordPolicy, and an error, if there is any.
func (c *FakeAwsIamAccountPasswordPolicies) Update(awsIamAccountPasswordPolicy *v1alpha1.AwsIamAccountPasswordPolicy) (result *v1alpha1.AwsIamAccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiamaccountpasswordpoliciesResource, c.ns, awsIamAccountPasswordPolicy), &v1alpha1.AwsIamAccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamAccountPasswordPolicy), err
}

// Delete takes name of the awsIamAccountPasswordPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamAccountPasswordPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiamaccountpasswordpoliciesResource, c.ns, name), &v1alpha1.AwsIamAccountPasswordPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamAccountPasswordPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiamaccountpasswordpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamAccountPasswordPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsIamAccountPasswordPolicy.
func (c *FakeAwsIamAccountPasswordPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamAccountPasswordPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiamaccountpasswordpoliciesResource, c.ns, name, data, subresources...), &v1alpha1.AwsIamAccountPasswordPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamAccountPasswordPolicy), err
}
