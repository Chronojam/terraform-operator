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

// AwsEmrClustersGetter has a method to return a AwsEmrClusterInterface.
// A group's client should implement this interface.
type AwsEmrClustersGetter interface {
	AwsEmrClusters(namespace string) AwsEmrClusterInterface
}

// AwsEmrClusterInterface has methods to work with AwsEmrCluster resources.
type AwsEmrClusterInterface interface {
	Create(*v1alpha1.AwsEmrCluster) (*v1alpha1.AwsEmrCluster, error)
	Update(*v1alpha1.AwsEmrCluster) (*v1alpha1.AwsEmrCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEmrCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEmrClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEmrCluster, err error)
	AwsEmrClusterExpansion
}

// awsEmrClusters implements AwsEmrClusterInterface
type awsEmrClusters struct {
	client rest.Interface
	ns     string
}

// newAwsEmrClusters returns a AwsEmrClusters
func newAwsEmrClusters(c *ChronojamV1alpha1Client, namespace string) *awsEmrClusters {
	return &awsEmrClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsEmrCluster, and returns the corresponding awsEmrCluster object, and an error if there is any.
func (c *awsEmrClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEmrCluster, err error) {
	result = &v1alpha1.AwsEmrCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsemrclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEmrClusters that match those selectors.
func (c *awsEmrClusters) List(opts v1.ListOptions) (result *v1alpha1.AwsEmrClusterList, err error) {
	result = &v1alpha1.AwsEmrClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsemrclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEmrClusters.
func (c *awsEmrClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsemrclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsEmrCluster and creates it.  Returns the server's representation of the awsEmrCluster, and an error, if there is any.
func (c *awsEmrClusters) Create(awsEmrCluster *v1alpha1.AwsEmrCluster) (result *v1alpha1.AwsEmrCluster, err error) {
	result = &v1alpha1.AwsEmrCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsemrclusters").
		Body(awsEmrCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEmrCluster and updates it. Returns the server's representation of the awsEmrCluster, and an error, if there is any.
func (c *awsEmrClusters) Update(awsEmrCluster *v1alpha1.AwsEmrCluster) (result *v1alpha1.AwsEmrCluster, err error) {
	result = &v1alpha1.AwsEmrCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsemrclusters").
		Name(awsEmrCluster.Name).
		Body(awsEmrCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEmrCluster and deletes it. Returns an error if one occurs.
func (c *awsEmrClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsemrclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEmrClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsemrclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEmrCluster.
func (c *awsEmrClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEmrCluster, err error) {
	result = &v1alpha1.AwsEmrCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsemrclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
