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

// FakeAwsNetworkAcls implements AwsNetworkAclInterface
type FakeAwsNetworkAcls struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsnetworkaclsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsnetworkacls"}

var awsnetworkaclsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsNetworkAcl"}

// Get takes name of the awsNetworkAcl, and returns the corresponding awsNetworkAcl object, and an error if there is any.
func (c *FakeAwsNetworkAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsnetworkaclsResource, c.ns, name), &v1alpha1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkAcl), err
}

// List takes label and field selectors, and returns the list of AwsNetworkAcls that match those selectors.
func (c *FakeAwsNetworkAcls) List(opts v1.ListOptions) (result *v1alpha1.AwsNetworkAclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsnetworkaclsResource, awsnetworkaclsKind, c.ns, opts), &v1alpha1.AwsNetworkAclList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsNetworkAclList{}
	for _, item := range obj.(*v1alpha1.AwsNetworkAclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsNetworkAcls.
func (c *FakeAwsNetworkAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsnetworkaclsResource, c.ns, opts))

}

// Create takes the representation of a awsNetworkAcl and creates it.  Returns the server's representation of the awsNetworkAcl, and an error, if there is any.
func (c *FakeAwsNetworkAcls) Create(awsNetworkAcl *v1alpha1.AwsNetworkAcl) (result *v1alpha1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsnetworkaclsResource, c.ns, awsNetworkAcl), &v1alpha1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkAcl), err
}

// Update takes the representation of a awsNetworkAcl and updates it. Returns the server's representation of the awsNetworkAcl, and an error, if there is any.
func (c *FakeAwsNetworkAcls) Update(awsNetworkAcl *v1alpha1.AwsNetworkAcl) (result *v1alpha1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsnetworkaclsResource, c.ns, awsNetworkAcl), &v1alpha1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkAcl), err
}

// Delete takes name of the awsNetworkAcl and deletes it. Returns an error if one occurs.
func (c *FakeAwsNetworkAcls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsnetworkaclsResource, c.ns, name), &v1alpha1.AwsNetworkAcl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsNetworkAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsnetworkaclsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsNetworkAclList{})
	return err
}

// Patch applies the patch and returns the patched awsNetworkAcl.
func (c *FakeAwsNetworkAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsnetworkaclsResource, c.ns, name, data, subresources...), &v1alpha1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkAcl), err
}
