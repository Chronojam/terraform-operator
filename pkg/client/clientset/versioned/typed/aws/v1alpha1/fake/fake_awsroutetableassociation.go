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

// FakeAwsRouteTableAssociations implements AwsRouteTableAssociationInterface
type FakeAwsRouteTableAssociations struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsroutetableassociationsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsroutetableassociations"}

var awsroutetableassociationsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsRouteTableAssociation"}

// Get takes name of the awsRouteTableAssociation, and returns the corresponding awsRouteTableAssociation object, and an error if there is any.
func (c *FakeAwsRouteTableAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsroutetableassociationsResource, c.ns, name), &v1alpha1.AwsRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRouteTableAssociation), err
}

// List takes label and field selectors, and returns the list of AwsRouteTableAssociations that match those selectors.
func (c *FakeAwsRouteTableAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsRouteTableAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsroutetableassociationsResource, awsroutetableassociationsKind, c.ns, opts), &v1alpha1.AwsRouteTableAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsRouteTableAssociationList{}
	for _, item := range obj.(*v1alpha1.AwsRouteTableAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRouteTableAssociations.
func (c *FakeAwsRouteTableAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsroutetableassociationsResource, c.ns, opts))

}

// Create takes the representation of a awsRouteTableAssociation and creates it.  Returns the server's representation of the awsRouteTableAssociation, and an error, if there is any.
func (c *FakeAwsRouteTableAssociations) Create(awsRouteTableAssociation *v1alpha1.AwsRouteTableAssociation) (result *v1alpha1.AwsRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsroutetableassociationsResource, c.ns, awsRouteTableAssociation), &v1alpha1.AwsRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRouteTableAssociation), err
}

// Update takes the representation of a awsRouteTableAssociation and updates it. Returns the server's representation of the awsRouteTableAssociation, and an error, if there is any.
func (c *FakeAwsRouteTableAssociations) Update(awsRouteTableAssociation *v1alpha1.AwsRouteTableAssociation) (result *v1alpha1.AwsRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsroutetableassociationsResource, c.ns, awsRouteTableAssociation), &v1alpha1.AwsRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRouteTableAssociation), err
}

// Delete takes name of the awsRouteTableAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsRouteTableAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsroutetableassociationsResource, c.ns, name), &v1alpha1.AwsRouteTableAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRouteTableAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsroutetableassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsRouteTableAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsRouteTableAssociation.
func (c *FakeAwsRouteTableAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsroutetableassociationsResource, c.ns, name, data, subresources...), &v1alpha1.AwsRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRouteTableAssociation), err
}
