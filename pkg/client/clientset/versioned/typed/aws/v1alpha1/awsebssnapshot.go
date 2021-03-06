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

package v1alpha1

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	scheme "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsEbsSnapshotsGetter has a method to return a AwsEbsSnapshotInterface.
// A group's client should implement this interface.
type AwsEbsSnapshotsGetter interface {
	AwsEbsSnapshots(namespace string) AwsEbsSnapshotInterface
}

// AwsEbsSnapshotInterface has methods to work with AwsEbsSnapshot resources.
type AwsEbsSnapshotInterface interface {
	Create(*v1alpha1.AwsEbsSnapshot) (*v1alpha1.AwsEbsSnapshot, error)
	Update(*v1alpha1.AwsEbsSnapshot) (*v1alpha1.AwsEbsSnapshot, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEbsSnapshot, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEbsSnapshotList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEbsSnapshot, err error)
	AwsEbsSnapshotExpansion
}

// awsEbsSnapshots implements AwsEbsSnapshotInterface
type awsEbsSnapshots struct {
	client rest.Interface
	ns     string
}

// newAwsEbsSnapshots returns a AwsEbsSnapshots
func newAwsEbsSnapshots(c *ChronojamV1alpha1Client, namespace string) *awsEbsSnapshots {
	return &awsEbsSnapshots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsEbsSnapshot, and returns the corresponding awsEbsSnapshot object, and an error if there is any.
func (c *awsEbsSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEbsSnapshot, err error) {
	result = &v1alpha1.AwsEbsSnapshot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsebssnapshots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEbsSnapshots that match those selectors.
func (c *awsEbsSnapshots) List(opts v1.ListOptions) (result *v1alpha1.AwsEbsSnapshotList, err error) {
	result = &v1alpha1.AwsEbsSnapshotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsebssnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEbsSnapshots.
func (c *awsEbsSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsebssnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsEbsSnapshot and creates it.  Returns the server's representation of the awsEbsSnapshot, and an error, if there is any.
func (c *awsEbsSnapshots) Create(awsEbsSnapshot *v1alpha1.AwsEbsSnapshot) (result *v1alpha1.AwsEbsSnapshot, err error) {
	result = &v1alpha1.AwsEbsSnapshot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsebssnapshots").
		Body(awsEbsSnapshot).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEbsSnapshot and updates it. Returns the server's representation of the awsEbsSnapshot, and an error, if there is any.
func (c *awsEbsSnapshots) Update(awsEbsSnapshot *v1alpha1.AwsEbsSnapshot) (result *v1alpha1.AwsEbsSnapshot, err error) {
	result = &v1alpha1.AwsEbsSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsebssnapshots").
		Name(awsEbsSnapshot.Name).
		Body(awsEbsSnapshot).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEbsSnapshot and deletes it. Returns an error if one occurs.
func (c *awsEbsSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsebssnapshots").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEbsSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsebssnapshots").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEbsSnapshot.
func (c *awsEbsSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEbsSnapshot, err error) {
	result = &v1alpha1.AwsEbsSnapshot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsebssnapshots").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
