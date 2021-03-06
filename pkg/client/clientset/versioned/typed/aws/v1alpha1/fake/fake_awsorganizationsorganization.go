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

// FakeAwsOrganizationsOrganizations implements AwsOrganizationsOrganizationInterface
type FakeAwsOrganizationsOrganizations struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsorganizationsorganizationsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsorganizationsorganizations"}

var awsorganizationsorganizationsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsOrganizationsOrganization"}

// Get takes name of the awsOrganizationsOrganization, and returns the corresponding awsOrganizationsOrganization object, and an error if there is any.
func (c *FakeAwsOrganizationsOrganizations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsorganizationsorganizationsResource, c.ns, name), &v1alpha1.AwsOrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOrganizationsOrganization), err
}

// List takes label and field selectors, and returns the list of AwsOrganizationsOrganizations that match those selectors.
func (c *FakeAwsOrganizationsOrganizations) List(opts v1.ListOptions) (result *v1alpha1.AwsOrganizationsOrganizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsorganizationsorganizationsResource, awsorganizationsorganizationsKind, c.ns, opts), &v1alpha1.AwsOrganizationsOrganizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOrganizationsOrganizationList{}
	for _, item := range obj.(*v1alpha1.AwsOrganizationsOrganizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOrganizationsOrganizations.
func (c *FakeAwsOrganizationsOrganizations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsorganizationsorganizationsResource, c.ns, opts))

}

// Create takes the representation of a awsOrganizationsOrganization and creates it.  Returns the server's representation of the awsOrganizationsOrganization, and an error, if there is any.
func (c *FakeAwsOrganizationsOrganizations) Create(awsOrganizationsOrganization *v1alpha1.AwsOrganizationsOrganization) (result *v1alpha1.AwsOrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsorganizationsorganizationsResource, c.ns, awsOrganizationsOrganization), &v1alpha1.AwsOrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOrganizationsOrganization), err
}

// Update takes the representation of a awsOrganizationsOrganization and updates it. Returns the server's representation of the awsOrganizationsOrganization, and an error, if there is any.
func (c *FakeAwsOrganizationsOrganizations) Update(awsOrganizationsOrganization *v1alpha1.AwsOrganizationsOrganization) (result *v1alpha1.AwsOrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsorganizationsorganizationsResource, c.ns, awsOrganizationsOrganization), &v1alpha1.AwsOrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOrganizationsOrganization), err
}

// Delete takes name of the awsOrganizationsOrganization and deletes it. Returns an error if one occurs.
func (c *FakeAwsOrganizationsOrganizations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsorganizationsorganizationsResource, c.ns, name), &v1alpha1.AwsOrganizationsOrganization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOrganizationsOrganizations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsorganizationsorganizationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOrganizationsOrganizationList{})
	return err
}

// Patch applies the patch and returns the patched awsOrganizationsOrganization.
func (c *FakeAwsOrganizationsOrganizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOrganizationsOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsorganizationsorganizationsResource, c.ns, name, data, subresources...), &v1alpha1.AwsOrganizationsOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOrganizationsOrganization), err
}
