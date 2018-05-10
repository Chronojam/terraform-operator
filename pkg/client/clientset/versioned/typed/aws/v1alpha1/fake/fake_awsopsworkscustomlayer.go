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

// FakeAwsOpsworksCustomLayers implements AwsOpsworksCustomLayerInterface
type FakeAwsOpsworksCustomLayers struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsopsworkscustomlayersResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsopsworkscustomlayers"}

var awsopsworkscustomlayersKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsOpsworksCustomLayer"}

// Get takes name of the awsOpsworksCustomLayer, and returns the corresponding awsOpsworksCustomLayer object, and an error if there is any.
func (c *FakeAwsOpsworksCustomLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksCustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsopsworkscustomlayersResource, c.ns, name), &v1alpha1.AwsOpsworksCustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksCustomLayer), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksCustomLayers that match those selectors.
func (c *FakeAwsOpsworksCustomLayers) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksCustomLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsopsworkscustomlayersResource, awsopsworkscustomlayersKind, c.ns, opts), &v1alpha1.AwsOpsworksCustomLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOpsworksCustomLayerList{}
	for _, item := range obj.(*v1alpha1.AwsOpsworksCustomLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksCustomLayers.
func (c *FakeAwsOpsworksCustomLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsopsworkscustomlayersResource, c.ns, opts))

}

// Create takes the representation of a awsOpsworksCustomLayer and creates it.  Returns the server's representation of the awsOpsworksCustomLayer, and an error, if there is any.
func (c *FakeAwsOpsworksCustomLayers) Create(awsOpsworksCustomLayer *v1alpha1.AwsOpsworksCustomLayer) (result *v1alpha1.AwsOpsworksCustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsopsworkscustomlayersResource, c.ns, awsOpsworksCustomLayer), &v1alpha1.AwsOpsworksCustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksCustomLayer), err
}

// Update takes the representation of a awsOpsworksCustomLayer and updates it. Returns the server's representation of the awsOpsworksCustomLayer, and an error, if there is any.
func (c *FakeAwsOpsworksCustomLayers) Update(awsOpsworksCustomLayer *v1alpha1.AwsOpsworksCustomLayer) (result *v1alpha1.AwsOpsworksCustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsopsworkscustomlayersResource, c.ns, awsOpsworksCustomLayer), &v1alpha1.AwsOpsworksCustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksCustomLayer), err
}

// Delete takes name of the awsOpsworksCustomLayer and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksCustomLayers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsopsworkscustomlayersResource, c.ns, name), &v1alpha1.AwsOpsworksCustomLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksCustomLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsopsworkscustomlayersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOpsworksCustomLayerList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksCustomLayer.
func (c *FakeAwsOpsworksCustomLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksCustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsopsworkscustomlayersResource, c.ns, name, data, subresources...), &v1alpha1.AwsOpsworksCustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksCustomLayer), err
}
