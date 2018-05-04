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

// FakeAwsApiGatewayRestApis implements AwsApiGatewayRestApiInterface
type FakeAwsApiGatewayRestApis struct {
	Fake *FakeChronojamV1alpha1
	ns   string
}

var awsapigatewayrestapisResource = schema.GroupVersionResource{Group: "chronojam.co.uk", Version: "v1alpha1", Resource: "awsapigatewayrestapis"}

var awsapigatewayrestapisKind = schema.GroupVersionKind{Group: "chronojam.co.uk", Version: "v1alpha1", Kind: "AwsApiGatewayRestApi"}

// Get takes name of the awsApiGatewayRestApi, and returns the corresponding awsApiGatewayRestApi object, and an error if there is any.
func (c *FakeAwsApiGatewayRestApis) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsapigatewayrestapisResource, c.ns, name), &v1alpha1.AwsApiGatewayRestApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayRestApi), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayRestApis that match those selectors.
func (c *FakeAwsApiGatewayRestApis) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayRestApiList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsapigatewayrestapisResource, awsapigatewayrestapisKind, c.ns, opts), &v1alpha1.AwsApiGatewayRestApiList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsApiGatewayRestApiList{}
	for _, item := range obj.(*v1alpha1.AwsApiGatewayRestApiList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayRestApis.
func (c *FakeAwsApiGatewayRestApis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsapigatewayrestapisResource, c.ns, opts))

}

// Create takes the representation of a awsApiGatewayRestApi and creates it.  Returns the server's representation of the awsApiGatewayRestApi, and an error, if there is any.
func (c *FakeAwsApiGatewayRestApis) Create(awsApiGatewayRestApi *v1alpha1.AwsApiGatewayRestApi) (result *v1alpha1.AwsApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsapigatewayrestapisResource, c.ns, awsApiGatewayRestApi), &v1alpha1.AwsApiGatewayRestApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayRestApi), err
}

// Update takes the representation of a awsApiGatewayRestApi and updates it. Returns the server's representation of the awsApiGatewayRestApi, and an error, if there is any.
func (c *FakeAwsApiGatewayRestApis) Update(awsApiGatewayRestApi *v1alpha1.AwsApiGatewayRestApi) (result *v1alpha1.AwsApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsapigatewayrestapisResource, c.ns, awsApiGatewayRestApi), &v1alpha1.AwsApiGatewayRestApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayRestApi), err
}

// Delete takes name of the awsApiGatewayRestApi and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayRestApis) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsapigatewayrestapisResource, c.ns, name), &v1alpha1.AwsApiGatewayRestApi{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayRestApis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsapigatewayrestapisResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsApiGatewayRestApiList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayRestApi.
func (c *FakeAwsApiGatewayRestApis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsapigatewayrestapisResource, c.ns, name, data, subresources...), &v1alpha1.AwsApiGatewayRestApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayRestApi), err
}
