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

// FakeAwsEcsServices implements AwsEcsServiceInterface
type FakeAwsEcsServices struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsecsservicesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsecsservices"}

var awsecsservicesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsEcsService"}

// Get takes name of the awsEcsService, and returns the corresponding awsEcsService object, and an error if there is any.
func (c *FakeAwsEcsServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsecsservicesResource, c.ns, name), &v1alpha1.AwsEcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcsService), err
}

// List takes label and field selectors, and returns the list of AwsEcsServices that match those selectors.
func (c *FakeAwsEcsServices) List(opts v1.ListOptions) (result *v1alpha1.AwsEcsServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsecsservicesResource, awsecsservicesKind, c.ns, opts), &v1alpha1.AwsEcsServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEcsServiceList{}
	for _, item := range obj.(*v1alpha1.AwsEcsServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEcsServices.
func (c *FakeAwsEcsServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsecsservicesResource, c.ns, opts))

}

// Create takes the representation of a awsEcsService and creates it.  Returns the server's representation of the awsEcsService, and an error, if there is any.
func (c *FakeAwsEcsServices) Create(awsEcsService *v1alpha1.AwsEcsService) (result *v1alpha1.AwsEcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsecsservicesResource, c.ns, awsEcsService), &v1alpha1.AwsEcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcsService), err
}

// Update takes the representation of a awsEcsService and updates it. Returns the server's representation of the awsEcsService, and an error, if there is any.
func (c *FakeAwsEcsServices) Update(awsEcsService *v1alpha1.AwsEcsService) (result *v1alpha1.AwsEcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsecsservicesResource, c.ns, awsEcsService), &v1alpha1.AwsEcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcsService), err
}

// Delete takes name of the awsEcsService and deletes it. Returns an error if one occurs.
func (c *FakeAwsEcsServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsecsservicesResource, c.ns, name), &v1alpha1.AwsEcsService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEcsServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsecsservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEcsServiceList{})
	return err
}

// Patch applies the patch and returns the patched awsEcsService.
func (c *FakeAwsEcsServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsecsservicesResource, c.ns, name, data, subresources...), &v1alpha1.AwsEcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcsService), err
}
