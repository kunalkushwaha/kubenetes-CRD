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

package v1alpha

import (
	v1alpha "github.com/xieydd/kubenetes-crd/pkg/apis/unisound.org/v1alpha"
	scheme "github.com/xieydd/kubenetes-crd/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// QueueJobsGetter has a method to return a QueueJobInterface.
// A group's client should implement this interface.
type QueueJobsGetter interface {
	QueueJobs(namespace string) QueueJobInterface
}

// QueueJobInterface has methods to work with QueueJob resources.
type QueueJobInterface interface {
	Create(*v1alpha.QueueJob) (*v1alpha.QueueJob, error)
	Update(*v1alpha.QueueJob) (*v1alpha.QueueJob, error)
	UpdateStatus(*v1alpha.QueueJob) (*v1alpha.QueueJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha.QueueJob, error)
	List(opts v1.ListOptions) (*v1alpha.QueueJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.QueueJob, err error)
	QueueJobExpansion
}

// queueJobs implements QueueJobInterface
type queueJobs struct {
	client rest.Interface
	ns     string
}

// newQueueJobs returns a QueueJobs
func newQueueJobs(c *UnisoundV1alphaClient, namespace string) *queueJobs {
	return &queueJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the queueJob, and returns the corresponding queueJob object, and an error if there is any.
func (c *queueJobs) Get(name string, options v1.GetOptions) (result *v1alpha.QueueJob, err error) {
	result = &v1alpha.QueueJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("queuejobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of QueueJobs that match those selectors.
func (c *queueJobs) List(opts v1.ListOptions) (result *v1alpha.QueueJobList, err error) {
	result = &v1alpha.QueueJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("queuejobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested queueJobs.
func (c *queueJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("queuejobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a queueJob and creates it.  Returns the server's representation of the queueJob, and an error, if there is any.
func (c *queueJobs) Create(queueJob *v1alpha.QueueJob) (result *v1alpha.QueueJob, err error) {
	result = &v1alpha.QueueJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("queuejobs").
		Body(queueJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a queueJob and updates it. Returns the server's representation of the queueJob, and an error, if there is any.
func (c *queueJobs) Update(queueJob *v1alpha.QueueJob) (result *v1alpha.QueueJob, err error) {
	result = &v1alpha.QueueJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("queuejobs").
		Name(queueJob.Name).
		Body(queueJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *queueJobs) UpdateStatus(queueJob *v1alpha.QueueJob) (result *v1alpha.QueueJob, err error) {
	result = &v1alpha.QueueJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("queuejobs").
		Name(queueJob.Name).
		SubResource("status").
		Body(queueJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the queueJob and deletes it. Returns an error if one occurs.
func (c *queueJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("queuejobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *queueJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("queuejobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched queueJob.
func (c *queueJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.QueueJob, err error) {
	result = &v1alpha.QueueJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("queuejobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}