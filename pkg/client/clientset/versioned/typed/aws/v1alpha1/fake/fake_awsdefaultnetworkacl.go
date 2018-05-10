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

// FakeAwsDefaultNetworkAcls implements AwsDefaultNetworkAclInterface
type FakeAwsDefaultNetworkAcls struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsdefaultnetworkaclsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsdefaultnetworkacls"}

var awsdefaultnetworkaclsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsDefaultNetworkAcl"}

// Get takes name of the awsDefaultNetworkAcl, and returns the corresponding awsDefaultNetworkAcl object, and an error if there is any.
func (c *FakeAwsDefaultNetworkAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdefaultnetworkaclsResource, c.ns, name), &v1alpha1.AwsDefaultNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDefaultNetworkAcl), err
}

// List takes label and field selectors, and returns the list of AwsDefaultNetworkAcls that match those selectors.
func (c *FakeAwsDefaultNetworkAcls) List(opts v1.ListOptions) (result *v1alpha1.AwsDefaultNetworkAclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdefaultnetworkaclsResource, awsdefaultnetworkaclsKind, c.ns, opts), &v1alpha1.AwsDefaultNetworkAclList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDefaultNetworkAclList{}
	for _, item := range obj.(*v1alpha1.AwsDefaultNetworkAclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDefaultNetworkAcls.
func (c *FakeAwsDefaultNetworkAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdefaultnetworkaclsResource, c.ns, opts))

}

// Create takes the representation of a awsDefaultNetworkAcl and creates it.  Returns the server's representation of the awsDefaultNetworkAcl, and an error, if there is any.
func (c *FakeAwsDefaultNetworkAcls) Create(awsDefaultNetworkAcl *v1alpha1.AwsDefaultNetworkAcl) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdefaultnetworkaclsResource, c.ns, awsDefaultNetworkAcl), &v1alpha1.AwsDefaultNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDefaultNetworkAcl), err
}

// Update takes the representation of a awsDefaultNetworkAcl and updates it. Returns the server's representation of the awsDefaultNetworkAcl, and an error, if there is any.
func (c *FakeAwsDefaultNetworkAcls) Update(awsDefaultNetworkAcl *v1alpha1.AwsDefaultNetworkAcl) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdefaultnetworkaclsResource, c.ns, awsDefaultNetworkAcl), &v1alpha1.AwsDefaultNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDefaultNetworkAcl), err
}

// Delete takes name of the awsDefaultNetworkAcl and deletes it. Returns an error if one occurs.
func (c *FakeAwsDefaultNetworkAcls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdefaultnetworkaclsResource, c.ns, name), &v1alpha1.AwsDefaultNetworkAcl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDefaultNetworkAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdefaultnetworkaclsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDefaultNetworkAclList{})
	return err
}

// Patch applies the patch and returns the patched awsDefaultNetworkAcl.
func (c *FakeAwsDefaultNetworkAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdefaultnetworkaclsResource, c.ns, name, data, subresources...), &v1alpha1.AwsDefaultNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDefaultNetworkAcl), err
}
