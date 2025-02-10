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

package v1alpha3

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1alpha3 "k8s.io/api/resource/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamresourcev1alpha3informers "k8s.io/client-go/informers/resource/v1alpha3"
	upstreamresourcev1alpha3listers "k8s.io/client-go/listers/resource/v1alpha3"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	resourcev1alpha3listers "github.com/kcp-dev/client-go/listers/resource/v1alpha3"
)

// ResourceClaimClusterInformer provides access to a shared informer and lister for
// ResourceClaims.
type ResourceClaimClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamresourcev1alpha3informers.ResourceClaimInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() resourcev1alpha3listers.ResourceClaimClusterLister
}

type resourceClaimClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceClaimClusterInformer constructs a new informer for ResourceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceClaimClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResourceClaimClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceClaimClusterInformer constructs a new informer for ResourceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceClaimClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceClaims().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceClaims().Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha3.ResourceClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceClaimClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResourceClaimClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *resourceClaimClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha3.ResourceClaim{}, f.defaultInformer)
}

func (f *resourceClaimClusterInformer) Lister() resourcev1alpha3listers.ResourceClaimClusterLister {
	return resourcev1alpha3listers.NewResourceClaimClusterLister(f.Informer().GetIndexer())
}

func (f *resourceClaimClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamresourcev1alpha3informers.ResourceClaimInformer {
	return &resourceClaimInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type resourceClaimInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamresourcev1alpha3listers.ResourceClaimLister
}

func (f *resourceClaimInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *resourceClaimInformer) Lister() upstreamresourcev1alpha3listers.ResourceClaimLister {
	return f.lister
}
