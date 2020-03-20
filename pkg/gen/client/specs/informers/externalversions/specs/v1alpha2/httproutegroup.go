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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	specsv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha2"
	versioned "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/specs/clientset/versioned"
	internalinterfaces "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/specs/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/specs/listers/specs/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HTTPRouteGroupInformer provides access to a shared informer and lister for
// HTTPRouteGroups.
type HTTPRouteGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.HTTPRouteGroupLister
}

type hTTPRouteGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHTTPRouteGroupInformer constructs a new informer for HTTPRouteGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHTTPRouteGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHTTPRouteGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHTTPRouteGroupInformer constructs a new informer for HTTPRouteGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHTTPRouteGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpecsV1alpha2().HTTPRouteGroups(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpecsV1alpha2().HTTPRouteGroups(namespace).Watch(options)
			},
		},
		&specsv1alpha2.HTTPRouteGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *hTTPRouteGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHTTPRouteGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hTTPRouteGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&specsv1alpha2.HTTPRouteGroup{}, f.defaultInformer)
}

func (f *hTTPRouteGroupInformer) Lister() v1alpha2.HTTPRouteGroupLister {
	return v1alpha2.NewHTTPRouteGroupLister(f.Informer().GetIndexer())
}
