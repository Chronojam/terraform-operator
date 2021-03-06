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

// FakeAwsSimpledbDomains implements AwsSimpledbDomainInterface
type FakeAwsSimpledbDomains struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awssimpledbdomainsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awssimpledbdomains"}

var awssimpledbdomainsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsSimpledbDomain"}

// Get takes name of the awsSimpledbDomain, and returns the corresponding awsSimpledbDomain object, and an error if there is any.
func (c *FakeAwsSimpledbDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssimpledbdomainsResource, c.ns, name), &v1alpha1.AwsSimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSimpledbDomain), err
}

// List takes label and field selectors, and returns the list of AwsSimpledbDomains that match those selectors.
func (c *FakeAwsSimpledbDomains) List(opts v1.ListOptions) (result *v1alpha1.AwsSimpledbDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssimpledbdomainsResource, awssimpledbdomainsKind, c.ns, opts), &v1alpha1.AwsSimpledbDomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSimpledbDomainList{}
	for _, item := range obj.(*v1alpha1.AwsSimpledbDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSimpledbDomains.
func (c *FakeAwsSimpledbDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssimpledbdomainsResource, c.ns, opts))

}

// Create takes the representation of a awsSimpledbDomain and creates it.  Returns the server's representation of the awsSimpledbDomain, and an error, if there is any.
func (c *FakeAwsSimpledbDomains) Create(awsSimpledbDomain *v1alpha1.AwsSimpledbDomain) (result *v1alpha1.AwsSimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssimpledbdomainsResource, c.ns, awsSimpledbDomain), &v1alpha1.AwsSimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSimpledbDomain), err
}

// Update takes the representation of a awsSimpledbDomain and updates it. Returns the server's representation of the awsSimpledbDomain, and an error, if there is any.
func (c *FakeAwsSimpledbDomains) Update(awsSimpledbDomain *v1alpha1.AwsSimpledbDomain) (result *v1alpha1.AwsSimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssimpledbdomainsResource, c.ns, awsSimpledbDomain), &v1alpha1.AwsSimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSimpledbDomain), err
}

// Delete takes name of the awsSimpledbDomain and deletes it. Returns an error if one occurs.
func (c *FakeAwsSimpledbDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssimpledbdomainsResource, c.ns, name), &v1alpha1.AwsSimpledbDomain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSimpledbDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssimpledbdomainsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSimpledbDomainList{})
	return err
}

// Patch applies the patch and returns the patched awsSimpledbDomain.
func (c *FakeAwsSimpledbDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSimpledbDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssimpledbdomainsResource, c.ns, name, data, subresources...), &v1alpha1.AwsSimpledbDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSimpledbDomain), err
}
