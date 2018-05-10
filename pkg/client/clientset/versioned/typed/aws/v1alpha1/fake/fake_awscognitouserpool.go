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

// FakeAwsCognitoUserPools implements AwsCognitoUserPoolInterface
type FakeAwsCognitoUserPools struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awscognitouserpoolsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awscognitouserpools"}

var awscognitouserpoolsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsCognitoUserPool"}

// Get takes name of the awsCognitoUserPool, and returns the corresponding awsCognitoUserPool object, and an error if there is any.
func (c *FakeAwsCognitoUserPools) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCognitoUserPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscognitouserpoolsResource, c.ns, name), &v1alpha1.AwsCognitoUserPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoUserPool), err
}

// List takes label and field selectors, and returns the list of AwsCognitoUserPools that match those selectors.
func (c *FakeAwsCognitoUserPools) List(opts v1.ListOptions) (result *v1alpha1.AwsCognitoUserPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscognitouserpoolsResource, awscognitouserpoolsKind, c.ns, opts), &v1alpha1.AwsCognitoUserPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCognitoUserPoolList{}
	for _, item := range obj.(*v1alpha1.AwsCognitoUserPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCognitoUserPools.
func (c *FakeAwsCognitoUserPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscognitouserpoolsResource, c.ns, opts))

}

// Create takes the representation of a awsCognitoUserPool and creates it.  Returns the server's representation of the awsCognitoUserPool, and an error, if there is any.
func (c *FakeAwsCognitoUserPools) Create(awsCognitoUserPool *v1alpha1.AwsCognitoUserPool) (result *v1alpha1.AwsCognitoUserPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscognitouserpoolsResource, c.ns, awsCognitoUserPool), &v1alpha1.AwsCognitoUserPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoUserPool), err
}

// Update takes the representation of a awsCognitoUserPool and updates it. Returns the server's representation of the awsCognitoUserPool, and an error, if there is any.
func (c *FakeAwsCognitoUserPools) Update(awsCognitoUserPool *v1alpha1.AwsCognitoUserPool) (result *v1alpha1.AwsCognitoUserPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscognitouserpoolsResource, c.ns, awsCognitoUserPool), &v1alpha1.AwsCognitoUserPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoUserPool), err
}

// Delete takes name of the awsCognitoUserPool and deletes it. Returns an error if one occurs.
func (c *FakeAwsCognitoUserPools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscognitouserpoolsResource, c.ns, name), &v1alpha1.AwsCognitoUserPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCognitoUserPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscognitouserpoolsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCognitoUserPoolList{})
	return err
}

// Patch applies the patch and returns the patched awsCognitoUserPool.
func (c *FakeAwsCognitoUserPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoUserPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscognitouserpoolsResource, c.ns, name, data, subresources...), &v1alpha1.AwsCognitoUserPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoUserPool), err
}
