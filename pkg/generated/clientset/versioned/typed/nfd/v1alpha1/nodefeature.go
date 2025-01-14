/*
Copyright 2022 The Kubernetes Authors.

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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "github.com/openshift/node-feature-discovery/pkg/apis/nfd/v1alpha1"
	scheme "github.com/openshift/node-feature-discovery/pkg/generated/clientset/versioned/scheme"
)

// NodeFeaturesGetter has a method to return a NodeFeatureInterface.
// A group's client should implement this interface.
type NodeFeaturesGetter interface {
	NodeFeatures(namespace string) NodeFeatureInterface
}

// NodeFeatureInterface has methods to work with NodeFeature resources.
type NodeFeatureInterface interface {
	Create(ctx context.Context, nodeFeature *v1alpha1.NodeFeature, opts v1.CreateOptions) (*v1alpha1.NodeFeature, error)
	Update(ctx context.Context, nodeFeature *v1alpha1.NodeFeature, opts v1.UpdateOptions) (*v1alpha1.NodeFeature, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NodeFeature, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodeFeatureList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeFeature, err error)
	NodeFeatureExpansion
}

// nodeFeatures implements NodeFeatureInterface
type nodeFeatures struct {
	client rest.Interface
	ns     string
}

// newNodeFeatures returns a NodeFeatures
func newNodeFeatures(c *NfdV1alpha1Client, namespace string) *nodeFeatures {
	return &nodeFeatures{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodeFeature, and returns the corresponding nodeFeature object, and an error if there is any.
func (c *nodeFeatures) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeFeature, err error) {
	result = &v1alpha1.NodeFeature{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodefeatures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeFeatures that match those selectors.
func (c *nodeFeatures) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeFeatureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodeFeatureList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodefeatures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeFeatures.
func (c *nodeFeatures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodefeatures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodeFeature and creates it.  Returns the server's representation of the nodeFeature, and an error, if there is any.
func (c *nodeFeatures) Create(ctx context.Context, nodeFeature *v1alpha1.NodeFeature, opts v1.CreateOptions) (result *v1alpha1.NodeFeature, err error) {
	result = &v1alpha1.NodeFeature{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodefeatures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeFeature).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodeFeature and updates it. Returns the server's representation of the nodeFeature, and an error, if there is any.
func (c *nodeFeatures) Update(ctx context.Context, nodeFeature *v1alpha1.NodeFeature, opts v1.UpdateOptions) (result *v1alpha1.NodeFeature, err error) {
	result = &v1alpha1.NodeFeature{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodefeatures").
		Name(nodeFeature.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeFeature).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodeFeature and deletes it. Returns an error if one occurs.
func (c *nodeFeatures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodefeatures").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeFeatures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodefeatures").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodeFeature.
func (c *nodeFeatures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeFeature, err error) {
	result = &v1alpha1.NodeFeature{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodefeatures").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
