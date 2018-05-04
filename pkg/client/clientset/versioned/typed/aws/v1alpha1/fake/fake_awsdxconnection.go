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

// FakeAwsDxConnections implements AwsDxConnectionInterface
type FakeAwsDxConnections struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsdxconnectionsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsdxconnections"}

var awsdxconnectionsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsDxConnection"}

// Get takes name of the awsDxConnection, and returns the corresponding awsDxConnection object, and an error if there is any.
func (c *FakeAwsDxConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdxconnectionsResource, c.ns, name), &v1alpha1.AwsDxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxConnection), err
}

// List takes label and field selectors, and returns the list of AwsDxConnections that match those selectors.
func (c *FakeAwsDxConnections) List(opts v1.ListOptions) (result *v1alpha1.AwsDxConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdxconnectionsResource, awsdxconnectionsKind, c.ns, opts), &v1alpha1.AwsDxConnectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDxConnectionList{}
	for _, item := range obj.(*v1alpha1.AwsDxConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDxConnections.
func (c *FakeAwsDxConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdxconnectionsResource, c.ns, opts))

}

// Create takes the representation of a awsDxConnection and creates it.  Returns the server's representation of the awsDxConnection, and an error, if there is any.
func (c *FakeAwsDxConnections) Create(awsDxConnection *v1alpha1.AwsDxConnection) (result *v1alpha1.AwsDxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdxconnectionsResource, c.ns, awsDxConnection), &v1alpha1.AwsDxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxConnection), err
}

// Update takes the representation of a awsDxConnection and updates it. Returns the server's representation of the awsDxConnection, and an error, if there is any.
func (c *FakeAwsDxConnections) Update(awsDxConnection *v1alpha1.AwsDxConnection) (result *v1alpha1.AwsDxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdxconnectionsResource, c.ns, awsDxConnection), &v1alpha1.AwsDxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxConnection), err
}

// Delete takes name of the awsDxConnection and deletes it. Returns an error if one occurs.
func (c *FakeAwsDxConnections) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdxconnectionsResource, c.ns, name), &v1alpha1.AwsDxConnection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDxConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdxconnectionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDxConnectionList{})
	return err
}

// Patch applies the patch and returns the patched awsDxConnection.
func (c *FakeAwsDxConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdxconnectionsResource, c.ns, name, data, subresources...), &v1alpha1.AwsDxConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxConnection), err
}
