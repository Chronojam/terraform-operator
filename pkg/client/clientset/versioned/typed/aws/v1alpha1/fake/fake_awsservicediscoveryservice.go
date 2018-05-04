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

// FakeAwsServiceDiscoveryServices implements AwsServiceDiscoveryServiceInterface
type FakeAwsServiceDiscoveryServices struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsservicediscoveryservicesResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsservicediscoveryservices"}

var awsservicediscoveryservicesKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsServiceDiscoveryService"}

// Get takes name of the awsServiceDiscoveryService, and returns the corresponding awsServiceDiscoveryService object, and an error if there is any.
func (c *FakeAwsServiceDiscoveryServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsServiceDiscoveryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsservicediscoveryservicesResource, c.ns, name), &v1alpha1.AwsServiceDiscoveryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsServiceDiscoveryService), err
}

// List takes label and field selectors, and returns the list of AwsServiceDiscoveryServices that match those selectors.
func (c *FakeAwsServiceDiscoveryServices) List(opts v1.ListOptions) (result *v1alpha1.AwsServiceDiscoveryServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsservicediscoveryservicesResource, awsservicediscoveryservicesKind, c.ns, opts), &v1alpha1.AwsServiceDiscoveryServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsServiceDiscoveryServiceList{}
	for _, item := range obj.(*v1alpha1.AwsServiceDiscoveryServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsServiceDiscoveryServices.
func (c *FakeAwsServiceDiscoveryServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsservicediscoveryservicesResource, c.ns, opts))

}

// Create takes the representation of a awsServiceDiscoveryService and creates it.  Returns the server's representation of the awsServiceDiscoveryService, and an error, if there is any.
func (c *FakeAwsServiceDiscoveryServices) Create(awsServiceDiscoveryService *v1alpha1.AwsServiceDiscoveryService) (result *v1alpha1.AwsServiceDiscoveryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsservicediscoveryservicesResource, c.ns, awsServiceDiscoveryService), &v1alpha1.AwsServiceDiscoveryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsServiceDiscoveryService), err
}

// Update takes the representation of a awsServiceDiscoveryService and updates it. Returns the server's representation of the awsServiceDiscoveryService, and an error, if there is any.
func (c *FakeAwsServiceDiscoveryServices) Update(awsServiceDiscoveryService *v1alpha1.AwsServiceDiscoveryService) (result *v1alpha1.AwsServiceDiscoveryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsservicediscoveryservicesResource, c.ns, awsServiceDiscoveryService), &v1alpha1.AwsServiceDiscoveryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsServiceDiscoveryService), err
}

// Delete takes name of the awsServiceDiscoveryService and deletes it. Returns an error if one occurs.
func (c *FakeAwsServiceDiscoveryServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsservicediscoveryservicesResource, c.ns, name), &v1alpha1.AwsServiceDiscoveryService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsServiceDiscoveryServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsservicediscoveryservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsServiceDiscoveryServiceList{})
	return err
}

// Patch applies the patch and returns the patched awsServiceDiscoveryService.
func (c *FakeAwsServiceDiscoveryServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsServiceDiscoveryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsservicediscoveryservicesResource, c.ns, name, data, subresources...), &v1alpha1.AwsServiceDiscoveryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsServiceDiscoveryService), err
}
