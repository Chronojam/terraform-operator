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

// FakeAwsVpcDhcpOptionses implements AwsVpcDhcpOptionsInterface
type FakeAwsVpcDhcpOptionses struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsvpcdhcpoptionsesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsvpcdhcpoptionses"}

var awsvpcdhcpoptionsesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsVpcDhcpOptions"}

// Get takes name of the awsVpcDhcpOptions, and returns the corresponding awsVpcDhcpOptions object, and an error if there is any.
func (c *FakeAwsVpcDhcpOptionses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsvpcdhcpoptionsesResource, c.ns, name), &v1alpha1.AwsVpcDhcpOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptions), err
}

// List takes label and field selectors, and returns the list of AwsVpcDhcpOptionses that match those selectors.
func (c *FakeAwsVpcDhcpOptionses) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcDhcpOptionsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsvpcdhcpoptionsesResource, awsvpcdhcpoptionsesKind, c.ns, opts), &v1alpha1.AwsVpcDhcpOptionsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVpcDhcpOptionsList{}
	for _, item := range obj.(*v1alpha1.AwsVpcDhcpOptionsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcDhcpOptionses.
func (c *FakeAwsVpcDhcpOptionses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsvpcdhcpoptionsesResource, c.ns, opts))

}

// Create takes the representation of a awsVpcDhcpOptions and creates it.  Returns the server's representation of the awsVpcDhcpOptions, and an error, if there is any.
func (c *FakeAwsVpcDhcpOptionses) Create(awsVpcDhcpOptions *v1alpha1.AwsVpcDhcpOptions) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsvpcdhcpoptionsesResource, c.ns, awsVpcDhcpOptions), &v1alpha1.AwsVpcDhcpOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptions), err
}

// Update takes the representation of a awsVpcDhcpOptions and updates it. Returns the server's representation of the awsVpcDhcpOptions, and an error, if there is any.
func (c *FakeAwsVpcDhcpOptionses) Update(awsVpcDhcpOptions *v1alpha1.AwsVpcDhcpOptions) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsvpcdhcpoptionsesResource, c.ns, awsVpcDhcpOptions), &v1alpha1.AwsVpcDhcpOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptions), err
}

// Delete takes name of the awsVpcDhcpOptions and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcDhcpOptionses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsvpcdhcpoptionsesResource, c.ns, name), &v1alpha1.AwsVpcDhcpOptions{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcDhcpOptionses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsvpcdhcpoptionsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVpcDhcpOptionsList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcDhcpOptions.
func (c *FakeAwsVpcDhcpOptionses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcDhcpOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsvpcdhcpoptionsesResource, c.ns, name, data, subresources...), &v1alpha1.AwsVpcDhcpOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcDhcpOptions), err
}
