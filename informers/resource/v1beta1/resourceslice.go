/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1beta1 "k8s.io/api/resource/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamresourcev1beta1informers "k8s.io/client-go/informers/resource/v1beta1"
	upstreamresourcev1beta1listers "k8s.io/client-go/listers/resource/v1beta1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	resourcev1beta1listers "github.com/kcp-dev/client-go/listers/resource/v1beta1"
)

// ResourceSliceClusterInformer provides access to a shared informer and lister for
// ResourceSlices.
type ResourceSliceClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamresourcev1beta1informers.ResourceSliceInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() resourcev1beta1listers.ResourceSliceClusterLister
}

type resourceSliceClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceSliceClusterInformer constructs a new informer for ResourceSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceSliceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResourceSliceClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceSliceClusterInformer constructs a new informer for ResourceSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceSliceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1beta1().ResourceSlices().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1beta1().ResourceSlices().Watch(context.TODO(), options)
			},
		},
		&resourcev1beta1.ResourceSlice{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceSliceClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResourceSliceClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *resourceSliceClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&resourcev1beta1.ResourceSlice{}, f.defaultInformer)
}

func (f *resourceSliceClusterInformer) Lister() resourcev1beta1listers.ResourceSliceClusterLister {
	return resourcev1beta1listers.NewResourceSliceClusterLister(f.Informer().GetIndexer())
}

func (f *resourceSliceClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamresourcev1beta1informers.ResourceSliceInformer {
	return &resourceSliceInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type resourceSliceInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamresourcev1beta1listers.ResourceSliceLister
}

func (f *resourceSliceInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *resourceSliceInformer) Lister() upstreamresourcev1beta1listers.ResourceSliceLister {
	return f.lister
}
