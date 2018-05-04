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

// FakeAwsAmis implements AwsAmiInterface
type FakeAwsAmis struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsamisResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsamis"}

var awsamisKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsAmi"}

// Get takes name of the awsAmi, and returns the corresponding awsAmi object, and an error if there is any.
func (c *FakeAwsAmis) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAmi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsamisResource, c.ns, name), &v1alpha1.AwsAmi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmi), err
}

// List takes label and field selectors, and returns the list of AwsAmis that match those selectors.
func (c *FakeAwsAmis) List(opts v1.ListOptions) (result *v1alpha1.AwsAmiList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsamisResource, awsamisKind, c.ns, opts), &v1alpha1.AwsAmiList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAmiList{}
	for _, item := range obj.(*v1alpha1.AwsAmiList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAmis.
func (c *FakeAwsAmis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsamisResource, c.ns, opts))

}

// Create takes the representation of a awsAmi and creates it.  Returns the server's representation of the awsAmi, and an error, if there is any.
func (c *FakeAwsAmis) Create(awsAmi *v1alpha1.AwsAmi) (result *v1alpha1.AwsAmi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsamisResource, c.ns, awsAmi), &v1alpha1.AwsAmi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmi), err
}

// Update takes the representation of a awsAmi and updates it. Returns the server's representation of the awsAmi, and an error, if there is any.
func (c *FakeAwsAmis) Update(awsAmi *v1alpha1.AwsAmi) (result *v1alpha1.AwsAmi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsamisResource, c.ns, awsAmi), &v1alpha1.AwsAmi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmi), err
}

// Delete takes name of the awsAmi and deletes it. Returns an error if one occurs.
func (c *FakeAwsAmis) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsamisResource, c.ns, name), &v1alpha1.AwsAmi{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAmis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsamisResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAmiList{})
	return err
}

// Patch applies the patch and returns the patched awsAmi.
func (c *FakeAwsAmis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsamisResource, c.ns, name, data, subresources...), &v1alpha1.AwsAmi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAmi), err
}
