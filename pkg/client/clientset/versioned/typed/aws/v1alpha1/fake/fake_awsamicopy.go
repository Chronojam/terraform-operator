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

// FakeAwsAmiCopies implements AwsAmiCopyInterface
type FakeAwsAmiCopies struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsamicopiesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsamicopies"}

var awsamicopiesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsAmiCopy"}

// Get takes name of the awsAmiCopy, and returns the corresponding awsAmiCopy object, and an error if there is any.
func (c *FakeAwsAmiCopies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsamicopiesResource, c.ns, name), &v1alpha1.AwsAmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmiCopy), err
}

// List takes label and field selectors, and returns the list of AwsAmiCopies that match those selectors.
func (c *FakeAwsAmiCopies) List(opts v1.ListOptions) (result *v1alpha1.AwsAmiCopyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsamicopiesResource, awsamicopiesKind, c.ns, opts), &v1alpha1.AwsAmiCopyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAmiCopyList{}
	for _, item := range obj.(*v1alpha1.AwsAmiCopyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAmiCopies.
func (c *FakeAwsAmiCopies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsamicopiesResource, c.ns, opts))

}

// Create takes the representation of a awsAmiCopy and creates it.  Returns the server's representation of the awsAmiCopy, and an error, if there is any.
func (c *FakeAwsAmiCopies) Create(awsAmiCopy *v1alpha1.AwsAmiCopy) (result *v1alpha1.AwsAmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsamicopiesResource, c.ns, awsAmiCopy), &v1alpha1.AwsAmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmiCopy), err
}

// Update takes the representation of a awsAmiCopy and updates it. Returns the server's representation of the awsAmiCopy, and an error, if there is any.
func (c *FakeAwsAmiCopies) Update(awsAmiCopy *v1alpha1.AwsAmiCopy) (result *v1alpha1.AwsAmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsamicopiesResource, c.ns, awsAmiCopy), &v1alpha1.AwsAmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmiCopy), err
}

// Delete takes name of the awsAmiCopy and deletes it. Returns an error if one occurs.
func (c *FakeAwsAmiCopies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsamicopiesResource, c.ns, name), &v1alpha1.AwsAmiCopy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAmiCopies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsamicopiesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAmiCopyList{})
	return err
}

// Patch applies the patch and returns the patched awsAmiCopy.
func (c *FakeAwsAmiCopies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsamicopiesResource, c.ns, name, data, subresources...), &v1alpha1.AwsAmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmiCopy), err
}
