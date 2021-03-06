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

// FakeAwsRedshiftClusters implements AwsRedshiftClusterInterface
type FakeAwsRedshiftClusters struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsredshiftclustersResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsredshiftclusters"}

var awsredshiftclustersKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsRedshiftCluster"}

// Get takes name of the awsRedshiftCluster, and returns the corresponding awsRedshiftCluster object, and an error if there is any.
func (c *FakeAwsRedshiftClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRedshiftCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsredshiftclustersResource, c.ns, name), &v1alpha1.AwsRedshiftCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftCluster), err
}

// List takes label and field selectors, and returns the list of AwsRedshiftClusters that match those selectors.
func (c *FakeAwsRedshiftClusters) List(opts v1.ListOptions) (result *v1alpha1.AwsRedshiftClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsredshiftclustersResource, awsredshiftclustersKind, c.ns, opts), &v1alpha1.AwsRedshiftClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsRedshiftClusterList{}
	for _, item := range obj.(*v1alpha1.AwsRedshiftClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRedshiftClusters.
func (c *FakeAwsRedshiftClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsredshiftclustersResource, c.ns, opts))

}

// Create takes the representation of a awsRedshiftCluster and creates it.  Returns the server's representation of the awsRedshiftCluster, and an error, if there is any.
func (c *FakeAwsRedshiftClusters) Create(awsRedshiftCluster *v1alpha1.AwsRedshiftCluster) (result *v1alpha1.AwsRedshiftCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsredshiftclustersResource, c.ns, awsRedshiftCluster), &v1alpha1.AwsRedshiftCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftCluster), err
}

// Update takes the representation of a awsRedshiftCluster and updates it. Returns the server's representation of the awsRedshiftCluster, and an error, if there is any.
func (c *FakeAwsRedshiftClusters) Update(awsRedshiftCluster *v1alpha1.AwsRedshiftCluster) (result *v1alpha1.AwsRedshiftCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsredshiftclustersResource, c.ns, awsRedshiftCluster), &v1alpha1.AwsRedshiftCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftCluster), err
}

// Delete takes name of the awsRedshiftCluster and deletes it. Returns an error if one occurs.
func (c *FakeAwsRedshiftClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsredshiftclustersResource, c.ns, name), &v1alpha1.AwsRedshiftCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRedshiftClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsredshiftclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsRedshiftClusterList{})
	return err
}

// Patch applies the patch and returns the patched awsRedshiftCluster.
func (c *FakeAwsRedshiftClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRedshiftCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsredshiftclustersResource, c.ns, name, data, subresources...), &v1alpha1.AwsRedshiftCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftCluster), err
}
