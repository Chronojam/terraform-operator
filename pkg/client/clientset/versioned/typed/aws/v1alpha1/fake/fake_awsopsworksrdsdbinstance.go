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

// FakeAwsOpsworksRdsDbInstances implements AwsOpsworksRdsDbInstanceInterface
type FakeAwsOpsworksRdsDbInstances struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var awsopsworksrdsdbinstancesResource = schema.GroupVersionResource{Group: "aws", Version: "v1alpha1", Resource: "awsopsworksrdsdbinstances"}

var awsopsworksrdsdbinstancesKind = schema.GroupVersionKind{Group: "aws", Version: "v1alpha1", Kind: "AwsOpsworksRdsDbInstance"}

// Get takes name of the awsOpsworksRdsDbInstance, and returns the corresponding awsOpsworksRdsDbInstance object, and an error if there is any.
func (c *FakeAwsOpsworksRdsDbInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsopsworksrdsdbinstancesResource, c.ns, name), &v1alpha1.AwsOpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksRdsDbInstances that match those selectors.
func (c *FakeAwsOpsworksRdsDbInstances) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksRdsDbInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsopsworksrdsdbinstancesResource, awsopsworksrdsdbinstancesKind, c.ns, opts), &v1alpha1.AwsOpsworksRdsDbInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOpsworksRdsDbInstanceList{}
	for _, item := range obj.(*v1alpha1.AwsOpsworksRdsDbInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksRdsDbInstances.
func (c *FakeAwsOpsworksRdsDbInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsopsworksrdsdbinstancesResource, c.ns, opts))

}

// Create takes the representation of a awsOpsworksRdsDbInstance and creates it.  Returns the server's representation of the awsOpsworksRdsDbInstance, and an error, if there is any.
func (c *FakeAwsOpsworksRdsDbInstances) Create(awsOpsworksRdsDbInstance *v1alpha1.AwsOpsworksRdsDbInstance) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsopsworksrdsdbinstancesResource, c.ns, awsOpsworksRdsDbInstance), &v1alpha1.AwsOpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}

// Update takes the representation of a awsOpsworksRdsDbInstance and updates it. Returns the server's representation of the awsOpsworksRdsDbInstance, and an error, if there is any.
func (c *FakeAwsOpsworksRdsDbInstances) Update(awsOpsworksRdsDbInstance *v1alpha1.AwsOpsworksRdsDbInstance) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsopsworksrdsdbinstancesResource, c.ns, awsOpsworksRdsDbInstance), &v1alpha1.AwsOpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}

// Delete takes name of the awsOpsworksRdsDbInstance and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksRdsDbInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsopsworksrdsdbinstancesResource, c.ns, name), &v1alpha1.AwsOpsworksRdsDbInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksRdsDbInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsopsworksrdsdbinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOpsworksRdsDbInstanceList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksRdsDbInstance.
func (c *FakeAwsOpsworksRdsDbInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsopsworksrdsdbinstancesResource, c.ns, name, data, subresources...), &v1alpha1.AwsOpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksRdsDbInstance), err
}
