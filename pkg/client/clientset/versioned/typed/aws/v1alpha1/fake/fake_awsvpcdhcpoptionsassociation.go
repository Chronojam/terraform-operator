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

// FakeAwsVpcDhcpOptionsAssociations implements AwsVpcDhcpOptionsAssociationInterface
type FakeAwsVpcDhcpOptionsAssociations struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsvpcdhcpoptionsassociationsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsvpcdhcpoptionsassociations"}

var awsvpcdhcpoptionsassociationsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsVpcDhcpOptionsAssociation"}

// Get takes name of the awsVpcDhcpOptionsAssociation, and returns the corresponding awsVpcDhcpOptionsAssociation object, and an error if there is any.
func (c *FakeAwsVpcDhcpOptionsAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcDhcpOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsvpcdhcpoptionsassociationsResource, c.ns, name), &v1alpha1.AwsVpcDhcpOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptionsAssociation), err
}

// List takes label and field selectors, and returns the list of AwsVpcDhcpOptionsAssociations that match those selectors.
func (c *FakeAwsVpcDhcpOptionsAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcDhcpOptionsAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsvpcdhcpoptionsassociationsResource, awsvpcdhcpoptionsassociationsKind, c.ns, opts), &v1alpha1.AwsVpcDhcpOptionsAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVpcDhcpOptionsAssociationList{}
	for _, item := range obj.(*v1alpha1.AwsVpcDhcpOptionsAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcDhcpOptionsAssociations.
func (c *FakeAwsVpcDhcpOptionsAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsvpcdhcpoptionsassociationsResource, c.ns, opts))

}

// Create takes the representation of a awsVpcDhcpOptionsAssociation and creates it.  Returns the server's representation of the awsVpcDhcpOptionsAssociation, and an error, if there is any.
func (c *FakeAwsVpcDhcpOptionsAssociations) Create(awsVpcDhcpOptionsAssociation *v1alpha1.AwsVpcDhcpOptionsAssociation) (result *v1alpha1.AwsVpcDhcpOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsvpcdhcpoptionsassociationsResource, c.ns, awsVpcDhcpOptionsAssociation), &v1alpha1.AwsVpcDhcpOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptionsAssociation), err
}

// Update takes the representation of a awsVpcDhcpOptionsAssociation and updates it. Returns the server's representation of the awsVpcDhcpOptionsAssociation, and an error, if there is any.
func (c *FakeAwsVpcDhcpOptionsAssociations) Update(awsVpcDhcpOptionsAssociation *v1alpha1.AwsVpcDhcpOptionsAssociation) (result *v1alpha1.AwsVpcDhcpOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsvpcdhcpoptionsassociationsResource, c.ns, awsVpcDhcpOptionsAssociation), &v1alpha1.AwsVpcDhcpOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptionsAssociation), err
}

// Delete takes name of the awsVpcDhcpOptionsAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcDhcpOptionsAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsvpcdhcpoptionsassociationsResource, c.ns, name), &v1alpha1.AwsVpcDhcpOptionsAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcDhcpOptionsAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsvpcdhcpoptionsassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVpcDhcpOptionsAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcDhcpOptionsAssociation.
func (c *FakeAwsVpcDhcpOptionsAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcDhcpOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsvpcdhcpoptionsassociationsResource, c.ns, name, data, subresources...), &v1alpha1.AwsVpcDhcpOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptionsAssociation), err
}
