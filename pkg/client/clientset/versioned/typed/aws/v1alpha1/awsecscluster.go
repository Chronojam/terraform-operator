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

// AwsEcsClustersGetter has a method to return a AwsEcsClusterInterface.
// A group's client should implement this interface.
type AwsEcsClustersGetter interface {
	AwsEcsClusters(namespace string) AwsEcsClusterInterface
}

// AwsEcsClusterInterface has methods to work with AwsEcsCluster resources.
type AwsEcsClusterInterface interface {
	Create(*v1alpha1.AwsEcsCluster) (*v1alpha1.AwsEcsCluster, error)
	Update(*v1alpha1.AwsEcsCluster) (*v1alpha1.AwsEcsCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEcsCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEcsClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcsCluster, err error)
	AwsEcsClusterExpansion
}

// awsEcsClusters implements AwsEcsClusterInterface
type awsEcsClusters struct {
	client rest.Interface
	ns     string
}

// newAwsEcsClusters returns a AwsEcsClusters
func newAwsEcsClusters(c *ChronojamV1alpha1Client, namespace string) *awsEcsClusters {
	return &awsEcsClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsEcsCluster, and returns the corresponding awsEcsCluster object, and an error if there is any.
func (c *awsEcsClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEcsCluster, err error) {
	result = &v1alpha1.AwsEcsCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecsclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEcsClusters that match those selectors.
func (c *awsEcsClusters) List(opts v1.ListOptions) (result *v1alpha1.AwsEcsClusterList, err error) {
	result = &v1alpha1.AwsEcsClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEcsClusters.
func (c *awsEcsClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsecsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsEcsCluster and creates it.  Returns the server's representation of the awsEcsCluster, and an error, if there is any.
func (c *awsEcsClusters) Create(awsEcsCluster *v1alpha1.AwsEcsCluster) (result *v1alpha1.AwsEcsCluster, err error) {
	result = &v1alpha1.AwsEcsCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsecsclusters").
		Body(awsEcsCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEcsCluster and updates it. Returns the server's representation of the awsEcsCluster, and an error, if there is any.
func (c *awsEcsClusters) Update(awsEcsCluster *v1alpha1.AwsEcsCluster) (result *v1alpha1.AwsEcsCluster, err error) {
	result = &v1alpha1.AwsEcsCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsecsclusters").
		Name(awsEcsCluster.Name).
		Body(awsEcsCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEcsCluster and deletes it. Returns an error if one occurs.
func (c *awsEcsClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecsclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEcsClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecsclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEcsCluster.
func (c *awsEcsClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcsCluster, err error) {
	result = &v1alpha1.AwsEcsCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsecsclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
