// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImageContentSourcePoliciesGetter has a method to return a ImageContentSourcePolicyInterface.
// A group's client should implement this interface.
type ImageContentSourcePoliciesGetter interface {
	ImageContentSourcePolicies() ImageContentSourcePolicyInterface
}

// ImageContentSourcePolicyInterface has methods to work with ImageContentSourcePolicy resources.
type ImageContentSourcePolicyInterface interface {
	Create(*v1alpha1.ImageContentSourcePolicy) (*v1alpha1.ImageContentSourcePolicy, error)
	Update(*v1alpha1.ImageContentSourcePolicy) (*v1alpha1.ImageContentSourcePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ImageContentSourcePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.ImageContentSourcePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ImageContentSourcePolicy, err error)
	ImageContentSourcePolicyExpansion
}

// imageContentSourcePolicies implements ImageContentSourcePolicyInterface
type imageContentSourcePolicies struct {
	client rest.Interface
}

// newImageContentSourcePolicies returns a ImageContentSourcePolicies
func newImageContentSourcePolicies(c *OperatorV1alpha1Client) *imageContentSourcePolicies {
	return &imageContentSourcePolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the imageContentSourcePolicy, and returns the corresponding imageContentSourcePolicy object, and an error if there is any.
func (c *imageContentSourcePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	result = &v1alpha1.ImageContentSourcePolicy{}
	err = c.client.Get().
		Resource("imagecontentsourcepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ImageContentSourcePolicies that match those selectors.
func (c *imageContentSourcePolicies) List(opts v1.ListOptions) (result *v1alpha1.ImageContentSourcePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ImageContentSourcePolicyList{}
	err = c.client.Get().
		Resource("imagecontentsourcepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested imageContentSourcePolicies.
func (c *imageContentSourcePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("imagecontentsourcepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a imageContentSourcePolicy and creates it.  Returns the server's representation of the imageContentSourcePolicy, and an error, if there is any.
func (c *imageContentSourcePolicies) Create(imageContentSourcePolicy *v1alpha1.ImageContentSourcePolicy) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	result = &v1alpha1.ImageContentSourcePolicy{}
	err = c.client.Post().
		Resource("imagecontentsourcepolicies").
		Body(imageContentSourcePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a imageContentSourcePolicy and updates it. Returns the server's representation of the imageContentSourcePolicy, and an error, if there is any.
func (c *imageContentSourcePolicies) Update(imageContentSourcePolicy *v1alpha1.ImageContentSourcePolicy) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	result = &v1alpha1.ImageContentSourcePolicy{}
	err = c.client.Put().
		Resource("imagecontentsourcepolicies").
		Name(imageContentSourcePolicy.Name).
		Body(imageContentSourcePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the imageContentSourcePolicy and deletes it. Returns an error if one occurs.
func (c *imageContentSourcePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("imagecontentsourcepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *imageContentSourcePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("imagecontentsourcepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched imageContentSourcePolicy.
func (c *imageContentSourcePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ImageContentSourcePolicy, err error) {
	result = &v1alpha1.ImageContentSourcePolicy{}
	err = c.client.Patch(pt).
		Resource("imagecontentsourcepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
