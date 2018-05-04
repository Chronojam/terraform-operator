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

package v1alpha1

import (
	v1alpha1 "github.com/chronojam/terraform-operator/pkg/apis/aws/v1alpha1"
	scheme "github.com/chronojam/terraform-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsCodecommitRepositoriesGetter has a method to return a AwsCodecommitRepositoryInterface.
// A group's client should implement this interface.
type AwsCodecommitRepositoriesGetter interface {
	AwsCodecommitRepositories(namespace string) AwsCodecommitRepositoryInterface
}

// AwsCodecommitRepositoryInterface has methods to work with AwsCodecommitRepository resources.
type AwsCodecommitRepositoryInterface interface {
	Create(*v1alpha1.AwsCodecommitRepository) (*v1alpha1.AwsCodecommitRepository, error)
	Update(*v1alpha1.AwsCodecommitRepository) (*v1alpha1.AwsCodecommitRepository, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCodecommitRepository, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCodecommitRepositoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodecommitRepository, err error)
	AwsCodecommitRepositoryExpansion
}

// awsCodecommitRepositories implements AwsCodecommitRepositoryInterface
type awsCodecommitRepositories struct {
	client rest.Interface
	ns     string
}

// newAwsCodecommitRepositories returns a AwsCodecommitRepositories
func newAwsCodecommitRepositories(c *ChronojamV1alpha1Client, namespace string) *awsCodecommitRepositories {
	return &awsCodecommitRepositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCodecommitRepository, and returns the corresponding awsCodecommitRepository object, and an error if there is any.
func (c *awsCodecommitRepositories) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodecommitRepository, err error) {
	result = &v1alpha1.AwsCodecommitRepository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCodecommitRepositories that match those selectors.
func (c *awsCodecommitRepositories) List(opts v1.ListOptions) (result *v1alpha1.AwsCodecommitRepositoryList, err error) {
	result = &v1alpha1.AwsCodecommitRepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCodecommitRepositories.
func (c *awsCodecommitRepositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCodecommitRepository and creates it.  Returns the server's representation of the awsCodecommitRepository, and an error, if there is any.
func (c *awsCodecommitRepositories) Create(awsCodecommitRepository *v1alpha1.AwsCodecommitRepository) (result *v1alpha1.AwsCodecommitRepository, err error) {
	result = &v1alpha1.AwsCodecommitRepository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		Body(awsCodecommitRepository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCodecommitRepository and updates it. Returns the server's representation of the awsCodecommitRepository, and an error, if there is any.
func (c *awsCodecommitRepositories) Update(awsCodecommitRepository *v1alpha1.AwsCodecommitRepository) (result *v1alpha1.AwsCodecommitRepository, err error) {
	result = &v1alpha1.AwsCodecommitRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		Name(awsCodecommitRepository.Name).
		Body(awsCodecommitRepository).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCodecommitRepository and deletes it. Returns an error if one occurs.
func (c *awsCodecommitRepositories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCodecommitRepositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCodecommitRepository.
func (c *awsCodecommitRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodecommitRepository, err error) {
	result = &v1alpha1.AwsCodecommitRepository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscodecommitrepositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
