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

// FakeAwsElasticacheSecurityGroups implements AwsElasticacheSecurityGroupInterface
type FakeAwsElasticacheSecurityGroups struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awselasticachesecuritygroupsResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awselasticachesecuritygroups"}

var awselasticachesecuritygroupsKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsElasticacheSecurityGroup"}

// Get takes name of the awsElasticacheSecurityGroup, and returns the corresponding awsElasticacheSecurityGroup object, and an error if there is any.
func (c *FakeAwsElasticacheSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticacheSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awselasticachesecuritygroupsResource, c.ns, name), &v1alpha1.AwsElasticacheSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheSecurityGroup), err
}

// List takes label and field selectors, and returns the list of AwsElasticacheSecurityGroups that match those selectors.
func (c *FakeAwsElasticacheSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticacheSecurityGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awselasticachesecuritygroupsResource, awselasticachesecuritygroupsKind, c.ns, opts), &v1alpha1.AwsElasticacheSecurityGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElasticacheSecurityGroupList{}
	for _, item := range obj.(*v1alpha1.AwsElasticacheSecurityGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticacheSecurityGroups.
func (c *FakeAwsElasticacheSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awselasticachesecuritygroupsResource, c.ns, opts))

}

// Create takes the representation of a awsElasticacheSecurityGroup and creates it.  Returns the server's representation of the awsElasticacheSecurityGroup, and an error, if there is any.
func (c *FakeAwsElasticacheSecurityGroups) Create(awsElasticacheSecurityGroup *v1alpha1.AwsElasticacheSecurityGroup) (result *v1alpha1.AwsElasticacheSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awselasticachesecuritygroupsResource, c.ns, awsElasticacheSecurityGroup), &v1alpha1.AwsElasticacheSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheSecurityGroup), err
}

// Update takes the representation of a awsElasticacheSecurityGroup and updates it. Returns the server's representation of the awsElasticacheSecurityGroup, and an error, if there is any.
func (c *FakeAwsElasticacheSecurityGroups) Update(awsElasticacheSecurityGroup *v1alpha1.AwsElasticacheSecurityGroup) (result *v1alpha1.AwsElasticacheSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awselasticachesecuritygroupsResource, c.ns, awsElasticacheSecurityGroup), &v1alpha1.AwsElasticacheSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheSecurityGroup), err
}

// Delete takes name of the awsElasticacheSecurityGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticacheSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awselasticachesecuritygroupsResource, c.ns, name), &v1alpha1.AwsElasticacheSecurityGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticacheSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awselasticachesecuritygroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElasticacheSecurityGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticacheSecurityGroup.
func (c *FakeAwsElasticacheSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticacheSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awselasticachesecuritygroupsResource, c.ns, name, data, subresources...), &v1alpha1.AwsElasticacheSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheSecurityGroup), err
}
