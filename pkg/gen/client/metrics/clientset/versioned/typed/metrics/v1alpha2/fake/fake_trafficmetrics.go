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
	v1alpha2 "github.com/deislabs/smi-sdk-go/pkg/apis/metrics/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTrafficMetricses implements TrafficMetricsInterface
type FakeTrafficMetricses struct {
	Fake *FakeMetricsV1alpha2
	ns   string
}

var trafficmetricsesResource = schema.GroupVersionResource{Group: "metrics.smi-spec.io", Version: "v1alpha2", Resource: "trafficmetricses"}

var trafficmetricsesKind = schema.GroupVersionKind{Group: "metrics.smi-spec.io", Version: "v1alpha2", Kind: "TrafficMetrics"}

// Get takes name of the trafficMetrics, and returns the corresponding trafficMetrics object, and an error if there is any.
func (c *FakeTrafficMetricses) Get(name string, options v1.GetOptions) (result *v1alpha2.TrafficMetrics, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(trafficmetricsesResource, c.ns, name), &v1alpha2.TrafficMetrics{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficMetrics), err
}

// List takes label and field selectors, and returns the list of TrafficMetricses that match those selectors.
func (c *FakeTrafficMetricses) List(opts v1.ListOptions) (result *v1alpha2.TrafficMetricsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(trafficmetricsesResource, trafficmetricsesKind, c.ns, opts), &v1alpha2.TrafficMetricsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.TrafficMetricsList{ListMeta: obj.(*v1alpha2.TrafficMetricsList).ListMeta}
	for _, item := range obj.(*v1alpha2.TrafficMetricsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trafficMetricses.
func (c *FakeTrafficMetricses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(trafficmetricsesResource, c.ns, opts))

}

// Create takes the representation of a trafficMetrics and creates it.  Returns the server's representation of the trafficMetrics, and an error, if there is any.
func (c *FakeTrafficMetricses) Create(trafficMetrics *v1alpha2.TrafficMetrics) (result *v1alpha2.TrafficMetrics, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(trafficmetricsesResource, c.ns, trafficMetrics), &v1alpha2.TrafficMetrics{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficMetrics), err
}

// Update takes the representation of a trafficMetrics and updates it. Returns the server's representation of the trafficMetrics, and an error, if there is any.
func (c *FakeTrafficMetricses) Update(trafficMetrics *v1alpha2.TrafficMetrics) (result *v1alpha2.TrafficMetrics, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(trafficmetricsesResource, c.ns, trafficMetrics), &v1alpha2.TrafficMetrics{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficMetrics), err
}

// Delete takes name of the trafficMetrics and deletes it. Returns an error if one occurs.
func (c *FakeTrafficMetricses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(trafficmetricsesResource, c.ns, name), &v1alpha2.TrafficMetrics{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrafficMetricses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(trafficmetricsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.TrafficMetricsList{})
	return err
}

// Patch applies the patch and returns the patched trafficMetrics.
func (c *FakeTrafficMetricses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.TrafficMetrics, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficmetricsesResource, c.ns, name, pt, data, subresources...), &v1alpha2.TrafficMetrics{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficMetrics), err
}
