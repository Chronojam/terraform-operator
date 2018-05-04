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

// FakeAwsEipAssociations implements AwsEipAssociationInterface
type FakeAwsEipAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awseipassociationsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awseipassociations"}

var awseipassociationsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsEipAssociation"}

// Get takes name of the awsEipAssociation, and returns the corresponding awsEipAssociation object, and an error if there is any.
func (c *FakeAwsEipAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEipAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awseipassociationsResource, c.ns, name), &v1alpha1.AwsEipAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEipAssociation), err
}

// List takes label and field selectors, and returns the list of AwsEipAssociations that match those selectors.
func (c *FakeAwsEipAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsEipAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awseipassociationsResource, awseipassociationsKind, c.ns, opts), &v1alpha1.AwsEipAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEipAssociationList{}
	for _, item := range obj.(*v1alpha1.AwsEipAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEipAssociations.
func (c *FakeAwsEipAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awseipassociationsResource, c.ns, opts))

}

// Create takes the representation of a awsEipAssociation and creates it.  Returns the server's representation of the awsEipAssociation, and an error, if there is any.
func (c *FakeAwsEipAssociations) Create(awsEipAssociation *v1alpha1.AwsEipAssociation) (result *v1alpha1.AwsEipAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awseipassociationsResource, c.ns, awsEipAssociation), &v1alpha1.AwsEipAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEipAssociation), err
}

// Update takes the representation of a awsEipAssociation and updates it. Returns the server's representation of the awsEipAssociation, and an error, if there is any.
func (c *FakeAwsEipAssociations) Update(awsEipAssociation *v1alpha1.AwsEipAssociation) (result *v1alpha1.AwsEipAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awseipassociationsResource, c.ns, awsEipAssociation), &v1alpha1.AwsEipAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEipAssociation), err
}

// Delete takes name of the awsEipAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsEipAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awseipassociationsResource, c.ns, name), &v1alpha1.AwsEipAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEipAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awseipassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEipAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsEipAssociation.
func (c *FakeAwsEipAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEipAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awseipassociationsResource, c.ns, name, data, subresources...), &v1alpha1.AwsEipAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEipAssociation), err
}
