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

// FakeAwsVpcPeeringConnectionOptionses implements AwsVpcPeeringConnectionOptionsInterface
type FakeAwsVpcPeeringConnectionOptionses struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsvpcpeeringconnectionoptionsesResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsvpcpeeringconnectionoptionses"}

var awsvpcpeeringconnectionoptionsesKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsVpcPeeringConnectionOptions"}

// Get takes name of the awsVpcPeeringConnectionOptions, and returns the corresponding awsVpcPeeringConnectionOptions object, and an error if there is any.
func (c *FakeAwsVpcPeeringConnectionOptionses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcPeeringConnectionOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsvpcpeeringconnectionoptionsesResource, c.ns, name), &v1alpha1.AwsVpcPeeringConnectionOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcPeeringConnectionOptions), err
}

// List takes label and field selectors, and returns the list of AwsVpcPeeringConnectionOptionses that match those selectors.
func (c *FakeAwsVpcPeeringConnectionOptionses) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcPeeringConnectionOptionsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsvpcpeeringconnectionoptionsesResource, awsvpcpeeringconnectionoptionsesKind, c.ns, opts), &v1alpha1.AwsVpcPeeringConnectionOptionsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVpcPeeringConnectionOptionsList{}
	for _, item := range obj.(*v1alpha1.AwsVpcPeeringConnectionOptionsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcPeeringConnectionOptionses.
func (c *FakeAwsVpcPeeringConnectionOptionses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsvpcpeeringconnectionoptionsesResource, c.ns, opts))

}

// Create takes the representation of a awsVpcPeeringConnectionOptions and creates it.  Returns the server's representation of the awsVpcPeeringConnectionOptions, and an error, if there is any.
func (c *FakeAwsVpcPeeringConnectionOptionses) Create(awsVpcPeeringConnectionOptions *v1alpha1.AwsVpcPeeringConnectionOptions) (result *v1alpha1.AwsVpcPeeringConnectionOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsvpcpeeringconnectionoptionsesResource, c.ns, awsVpcPeeringConnectionOptions), &v1alpha1.AwsVpcPeeringConnectionOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcPeeringConnectionOptions), err
}

// Update takes the representation of a awsVpcPeeringConnectionOptions and updates it. Returns the server's representation of the awsVpcPeeringConnectionOptions, and an error, if there is any.
func (c *FakeAwsVpcPeeringConnectionOptionses) Update(awsVpcPeeringConnectionOptions *v1alpha1.AwsVpcPeeringConnectionOptions) (result *v1alpha1.AwsVpcPeeringConnectionOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsvpcpeeringconnectionoptionsesResource, c.ns, awsVpcPeeringConnectionOptions), &v1alpha1.AwsVpcPeeringConnectionOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcPeeringConnectionOptions), err
}

// Delete takes name of the awsVpcPeeringConnectionOptions and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcPeeringConnectionOptionses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsvpcpeeringconnectionoptionsesResource, c.ns, name), &v1alpha1.AwsVpcPeeringConnectionOptions{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcPeeringConnectionOptionses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsvpcpeeringconnectionoptionsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVpcPeeringConnectionOptionsList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcPeeringConnectionOptions.
func (c *FakeAwsVpcPeeringConnectionOptionses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcPeeringConnectionOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsvpcpeeringconnectionoptionsesResource, c.ns, name, data, subresources...), &v1alpha1.AwsVpcPeeringConnectionOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcPeeringConnectionOptions), err
}
