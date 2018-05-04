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

// FakeAwsOpsworksApplications implements AwsOpsworksApplicationInterface
type FakeAwsOpsworksApplications struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsopsworksapplicationsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsopsworksapplications"}

var awsopsworksapplicationsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsOpsworksApplication"}

// Get takes name of the awsOpsworksApplication, and returns the corresponding awsOpsworksApplication object, and an error if there is any.
func (c *FakeAwsOpsworksApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsopsworksapplicationsResource, c.ns, name), &v1alpha1.AwsOpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksApplication), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksApplications that match those selectors.
func (c *FakeAwsOpsworksApplications) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsopsworksapplicationsResource, awsopsworksapplicationsKind, c.ns, opts), &v1alpha1.AwsOpsworksApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOpsworksApplicationList{}
	for _, item := range obj.(*v1alpha1.AwsOpsworksApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksApplications.
func (c *FakeAwsOpsworksApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsopsworksapplicationsResource, c.ns, opts))

}

// Create takes the representation of a awsOpsworksApplication and creates it.  Returns the server's representation of the awsOpsworksApplication, and an error, if there is any.
func (c *FakeAwsOpsworksApplications) Create(awsOpsworksApplication *v1alpha1.AwsOpsworksApplication) (result *v1alpha1.AwsOpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsopsworksapplicationsResource, c.ns, awsOpsworksApplication), &v1alpha1.AwsOpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksApplication), err
}

// Update takes the representation of a awsOpsworksApplication and updates it. Returns the server's representation of the awsOpsworksApplication, and an error, if there is any.
func (c *FakeAwsOpsworksApplications) Update(awsOpsworksApplication *v1alpha1.AwsOpsworksApplication) (result *v1alpha1.AwsOpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsopsworksapplicationsResource, c.ns, awsOpsworksApplication), &v1alpha1.AwsOpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksApplication), err
}

// Delete takes name of the awsOpsworksApplication and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsopsworksapplicationsResource, c.ns, name), &v1alpha1.AwsOpsworksApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsopsworksapplicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOpsworksApplicationList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksApplication.
func (c *FakeAwsOpsworksApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsopsworksapplicationsResource, c.ns, name, data, subresources...), &v1alpha1.AwsOpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksApplication), err
}
