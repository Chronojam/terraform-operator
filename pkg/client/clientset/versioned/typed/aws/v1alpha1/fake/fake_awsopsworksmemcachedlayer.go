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

// FakeAwsOpsworksMemcachedLayers implements AwsOpsworksMemcachedLayerInterface
type FakeAwsOpsworksMemcachedLayers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsopsworksmemcachedlayersResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsopsworksmemcachedlayers"}

var awsopsworksmemcachedlayersKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsOpsworksMemcachedLayer"}

// Get takes name of the awsOpsworksMemcachedLayer, and returns the corresponding awsOpsworksMemcachedLayer object, and an error if there is any.
func (c *FakeAwsOpsworksMemcachedLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksMemcachedLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsopsworksmemcachedlayersResource, c.ns, name), &v1alpha1.AwsOpsworksMemcachedLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksMemcachedLayer), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksMemcachedLayers that match those selectors.
func (c *FakeAwsOpsworksMemcachedLayers) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksMemcachedLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsopsworksmemcachedlayersResource, awsopsworksmemcachedlayersKind, c.ns, opts), &v1alpha1.AwsOpsworksMemcachedLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOpsworksMemcachedLayerList{}
	for _, item := range obj.(*v1alpha1.AwsOpsworksMemcachedLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksMemcachedLayers.
func (c *FakeAwsOpsworksMemcachedLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsopsworksmemcachedlayersResource, c.ns, opts))

}

// Create takes the representation of a awsOpsworksMemcachedLayer and creates it.  Returns the server's representation of the awsOpsworksMemcachedLayer, and an error, if there is any.
func (c *FakeAwsOpsworksMemcachedLayers) Create(awsOpsworksMemcachedLayer *v1alpha1.AwsOpsworksMemcachedLayer) (result *v1alpha1.AwsOpsworksMemcachedLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsopsworksmemcachedlayersResource, c.ns, awsOpsworksMemcachedLayer), &v1alpha1.AwsOpsworksMemcachedLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksMemcachedLayer), err
}

// Update takes the representation of a awsOpsworksMemcachedLayer and updates it. Returns the server's representation of the awsOpsworksMemcachedLayer, and an error, if there is any.
func (c *FakeAwsOpsworksMemcachedLayers) Update(awsOpsworksMemcachedLayer *v1alpha1.AwsOpsworksMemcachedLayer) (result *v1alpha1.AwsOpsworksMemcachedLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsopsworksmemcachedlayersResource, c.ns, awsOpsworksMemcachedLayer), &v1alpha1.AwsOpsworksMemcachedLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksMemcachedLayer), err
}

// Delete takes name of the awsOpsworksMemcachedLayer and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksMemcachedLayers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsopsworksmemcachedlayersResource, c.ns, name), &v1alpha1.AwsOpsworksMemcachedLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksMemcachedLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsopsworksmemcachedlayersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOpsworksMemcachedLayerList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksMemcachedLayer.
func (c *FakeAwsOpsworksMemcachedLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksMemcachedLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsopsworksmemcachedlayersResource, c.ns, name, data, subresources...), &v1alpha1.AwsOpsworksMemcachedLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksMemcachedLayer), err
}
