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

// FakeAwsEfsFileSystems implements AwsEfsFileSystemInterface
type FakeAwsEfsFileSystems struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsefsfilesystemsResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsefsfilesystems"}

var awsefsfilesystemsKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsEfsFileSystem"}

// Get takes name of the awsEfsFileSystem, and returns the corresponding awsEfsFileSystem object, and an error if there is any.
func (c *FakeAwsEfsFileSystems) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsefsfilesystemsResource, c.ns, name), &v1alpha1.AwsEfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEfsFileSystem), err
}

// List takes label and field selectors, and returns the list of AwsEfsFileSystems that match those selectors.
func (c *FakeAwsEfsFileSystems) List(opts v1.ListOptions) (result *v1alpha1.AwsEfsFileSystemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsefsfilesystemsResource, awsefsfilesystemsKind, c.ns, opts), &v1alpha1.AwsEfsFileSystemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEfsFileSystemList{}
	for _, item := range obj.(*v1alpha1.AwsEfsFileSystemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEfsFileSystems.
func (c *FakeAwsEfsFileSystems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsefsfilesystemsResource, c.ns, opts))

}

// Create takes the representation of a awsEfsFileSystem and creates it.  Returns the server's representation of the awsEfsFileSystem, and an error, if there is any.
func (c *FakeAwsEfsFileSystems) Create(awsEfsFileSystem *v1alpha1.AwsEfsFileSystem) (result *v1alpha1.AwsEfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsefsfilesystemsResource, c.ns, awsEfsFileSystem), &v1alpha1.AwsEfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEfsFileSystem), err
}

// Update takes the representation of a awsEfsFileSystem and updates it. Returns the server's representation of the awsEfsFileSystem, and an error, if there is any.
func (c *FakeAwsEfsFileSystems) Update(awsEfsFileSystem *v1alpha1.AwsEfsFileSystem) (result *v1alpha1.AwsEfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsefsfilesystemsResource, c.ns, awsEfsFileSystem), &v1alpha1.AwsEfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEfsFileSystem), err
}

// Delete takes name of the awsEfsFileSystem and deletes it. Returns an error if one occurs.
func (c *FakeAwsEfsFileSystems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsefsfilesystemsResource, c.ns, name), &v1alpha1.AwsEfsFileSystem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEfsFileSystems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsefsfilesystemsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEfsFileSystemList{})
	return err
}

// Patch applies the patch and returns the patched awsEfsFileSystem.
func (c *FakeAwsEfsFileSystems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEfsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsefsfilesystemsResource, c.ns, name, data, subresources...), &v1alpha1.AwsEfsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEfsFileSystem), err
}
