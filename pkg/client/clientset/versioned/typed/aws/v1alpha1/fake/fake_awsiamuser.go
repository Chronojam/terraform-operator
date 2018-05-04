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

// FakeAwsIamUsers implements AwsIamUserInterface
type FakeAwsIamUsers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsiamusersResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsiamusers"}

var awsiamusersKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsIamUser"}

// Get takes name of the awsIamUser, and returns the corresponding awsIamUser object, and an error if there is any.
func (c *FakeAwsIamUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiamusersResource, c.ns, name), &v1alpha1.AwsIamUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUser), err
}

// List takes label and field selectors, and returns the list of AwsIamUsers that match those selectors.
func (c *FakeAwsIamUsers) List(opts v1.ListOptions) (result *v1alpha1.AwsIamUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiamusersResource, awsiamusersKind, c.ns, opts), &v1alpha1.AwsIamUserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamUserList{}
	for _, item := range obj.(*v1alpha1.AwsIamUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamUsers.
func (c *FakeAwsIamUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiamusersResource, c.ns, opts))

}

// Create takes the representation of a awsIamUser and creates it.  Returns the server's representation of the awsIamUser, and an error, if there is any.
func (c *FakeAwsIamUsers) Create(awsIamUser *v1alpha1.AwsIamUser) (result *v1alpha1.AwsIamUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiamusersResource, c.ns, awsIamUser), &v1alpha1.AwsIamUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUser), err
}

// Update takes the representation of a awsIamUser and updates it. Returns the server's representation of the awsIamUser, and an error, if there is any.
func (c *FakeAwsIamUsers) Update(awsIamUser *v1alpha1.AwsIamUser) (result *v1alpha1.AwsIamUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiamusersResource, c.ns, awsIamUser), &v1alpha1.AwsIamUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUser), err
}

// Delete takes name of the awsIamUser and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiamusersResource, c.ns, name), &v1alpha1.AwsIamUser{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiamusersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamUserList{})
	return err
}

// Patch applies the patch and returns the patched awsIamUser.
func (c *FakeAwsIamUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiamusersResource, c.ns, name, data, subresources...), &v1alpha1.AwsIamUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamUser), err
}
