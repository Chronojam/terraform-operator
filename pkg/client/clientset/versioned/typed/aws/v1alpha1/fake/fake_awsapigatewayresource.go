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

// FakeAwsApiGatewayResources implements AwsApiGatewayResourceInterface
type FakeAwsApiGatewayResources struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsapigatewayresourcesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsapigatewayresources"}

var awsapigatewayresourcesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsApiGatewayResource"}

// Get takes name of the awsApiGatewayResource, and returns the corresponding awsApiGatewayResource object, and an error if there is any.
func (c *FakeAwsApiGatewayResources) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsapigatewayresourcesResource, c.ns, name), &v1alpha1.AwsApiGatewayResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayResource), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayResources that match those selectors.
func (c *FakeAwsApiGatewayResources) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsapigatewayresourcesResource, awsapigatewayresourcesKind, c.ns, opts), &v1alpha1.AwsApiGatewayResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsApiGatewayResourceList{}
	for _, item := range obj.(*v1alpha1.AwsApiGatewayResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayResources.
func (c *FakeAwsApiGatewayResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsapigatewayresourcesResource, c.ns, opts))

}

// Create takes the representation of a awsApiGatewayResource and creates it.  Returns the server's representation of the awsApiGatewayResource, and an error, if there is any.
func (c *FakeAwsApiGatewayResources) Create(awsApiGatewayResource *v1alpha1.AwsApiGatewayResource) (result *v1alpha1.AwsApiGatewayResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsapigatewayresourcesResource, c.ns, awsApiGatewayResource), &v1alpha1.AwsApiGatewayResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayResource), err
}

// Update takes the representation of a awsApiGatewayResource and updates it. Returns the server's representation of the awsApiGatewayResource, and an error, if there is any.
func (c *FakeAwsApiGatewayResources) Update(awsApiGatewayResource *v1alpha1.AwsApiGatewayResource) (result *v1alpha1.AwsApiGatewayResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsapigatewayresourcesResource, c.ns, awsApiGatewayResource), &v1alpha1.AwsApiGatewayResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayResource), err
}

// Delete takes name of the awsApiGatewayResource and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsapigatewayresourcesResource, c.ns, name), &v1alpha1.AwsApiGatewayResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsapigatewayresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsApiGatewayResourceList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayResource.
func (c *FakeAwsApiGatewayResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsapigatewayresourcesResource, c.ns, name, data, subresources...), &v1alpha1.AwsApiGatewayResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayResource), err
}
