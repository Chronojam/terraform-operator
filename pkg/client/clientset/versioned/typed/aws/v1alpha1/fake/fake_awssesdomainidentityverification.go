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

// FakeAwsSesDomainIdentityVerifications implements AwsSesDomainIdentityVerificationInterface
type FakeAwsSesDomainIdentityVerifications struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awssesdomainidentityverificationsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awssesdomainidentityverifications"}

var awssesdomainidentityverificationsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSesDomainIdentityVerification"}

// Get takes name of the awsSesDomainIdentityVerification, and returns the corresponding awsSesDomainIdentityVerification object, and an error if there is any.
func (c *FakeAwsSesDomainIdentityVerifications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssesdomainidentityverificationsResource, c.ns, name), &v1alpha1.AwsSesDomainIdentityVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesDomainIdentityVerification), err
}

// List takes label and field selectors, and returns the list of AwsSesDomainIdentityVerifications that match those selectors.
func (c *FakeAwsSesDomainIdentityVerifications) List(opts v1.ListOptions) (result *v1alpha1.AwsSesDomainIdentityVerificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssesdomainidentityverificationsResource, awssesdomainidentityverificationsKind, c.ns, opts), &v1alpha1.AwsSesDomainIdentityVerificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSesDomainIdentityVerificationList{}
	for _, item := range obj.(*v1alpha1.AwsSesDomainIdentityVerificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSesDomainIdentityVerifications.
func (c *FakeAwsSesDomainIdentityVerifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssesdomainidentityverificationsResource, c.ns, opts))

}

// Create takes the representation of a awsSesDomainIdentityVerification and creates it.  Returns the server's representation of the awsSesDomainIdentityVerification, and an error, if there is any.
func (c *FakeAwsSesDomainIdentityVerifications) Create(awsSesDomainIdentityVerification *v1alpha1.AwsSesDomainIdentityVerification) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssesdomainidentityverificationsResource, c.ns, awsSesDomainIdentityVerification), &v1alpha1.AwsSesDomainIdentityVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesDomainIdentityVerification), err
}

// Update takes the representation of a awsSesDomainIdentityVerification and updates it. Returns the server's representation of the awsSesDomainIdentityVerification, and an error, if there is any.
func (c *FakeAwsSesDomainIdentityVerifications) Update(awsSesDomainIdentityVerification *v1alpha1.AwsSesDomainIdentityVerification) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssesdomainidentityverificationsResource, c.ns, awsSesDomainIdentityVerification), &v1alpha1.AwsSesDomainIdentityVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesDomainIdentityVerification), err
}

// Delete takes name of the awsSesDomainIdentityVerification and deletes it. Returns an error if one occurs.
func (c *FakeAwsSesDomainIdentityVerifications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssesdomainidentityverificationsResource, c.ns, name), &v1alpha1.AwsSesDomainIdentityVerification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSesDomainIdentityVerifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssesdomainidentityverificationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSesDomainIdentityVerificationList{})
	return err
}

// Patch applies the patch and returns the patched awsSesDomainIdentityVerification.
func (c *FakeAwsSesDomainIdentityVerifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesDomainIdentityVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssesdomainidentityverificationsResource, c.ns, name, data, subresources...), &v1alpha1.AwsSesDomainIdentityVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesDomainIdentityVerification), err
}
