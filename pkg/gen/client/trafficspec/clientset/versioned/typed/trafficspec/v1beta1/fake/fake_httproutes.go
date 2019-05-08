/*
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
	v1beta1 "github.com/deislabs/smi-sdk-go/pkg/apis/trafficspec/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHTTPRouteses implements HTTPRoutesInterface
type FakeHTTPRouteses struct {
	Fake *FakeSmispecV1beta1
	ns   string
}

var httproutesesResource = schema.GroupVersionResource{Group: "smispec.io", Version: "v1beta1", Resource: "httprouteses"}

var httproutesesKind = schema.GroupVersionKind{Group: "smispec.io", Version: "v1beta1", Kind: "HTTPRoutes"}

// Get takes name of the hTTPRoutes, and returns the corresponding hTTPRoutes object, and an error if there is any.
func (c *FakeHTTPRouteses) Get(name string, options v1.GetOptions) (result *v1beta1.HTTPRoutes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httproutesesResource, c.ns, name), &v1beta1.HTTPRoutes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HTTPRoutes), err
}

// List takes label and field selectors, and returns the list of HTTPRouteses that match those selectors.
func (c *FakeHTTPRouteses) List(opts v1.ListOptions) (result *v1beta1.HTTPRoutesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httproutesesResource, httproutesesKind, c.ns, opts), &v1beta1.HTTPRoutesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.HTTPRoutesList{ListMeta: obj.(*v1beta1.HTTPRoutesList).ListMeta}
	for _, item := range obj.(*v1beta1.HTTPRoutesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPRouteses.
func (c *FakeHTTPRouteses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httproutesesResource, c.ns, opts))

}

// Create takes the representation of a hTTPRoutes and creates it.  Returns the server's representation of the hTTPRoutes, and an error, if there is any.
func (c *FakeHTTPRouteses) Create(hTTPRoutes *v1beta1.HTTPRoutes) (result *v1beta1.HTTPRoutes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httproutesesResource, c.ns, hTTPRoutes), &v1beta1.HTTPRoutes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HTTPRoutes), err
}

// Update takes the representation of a hTTPRoutes and updates it. Returns the server's representation of the hTTPRoutes, and an error, if there is any.
func (c *FakeHTTPRouteses) Update(hTTPRoutes *v1beta1.HTTPRoutes) (result *v1beta1.HTTPRoutes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httproutesesResource, c.ns, hTTPRoutes), &v1beta1.HTTPRoutes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HTTPRoutes), err
}

// Delete takes name of the hTTPRoutes and deletes it. Returns an error if one occurs.
func (c *FakeHTTPRouteses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(httproutesesResource, c.ns, name), &v1beta1.HTTPRoutes{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPRouteses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httproutesesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.HTTPRoutesList{})
	return err
}

// Patch applies the patch and returns the patched hTTPRoutes.
func (c *FakeHTTPRouteses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.HTTPRoutes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httproutesesResource, c.ns, name, pt, data, subresources...), &v1beta1.HTTPRoutes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HTTPRoutes), err
}