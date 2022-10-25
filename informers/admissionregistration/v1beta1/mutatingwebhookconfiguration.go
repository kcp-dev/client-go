//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamadmissionregistrationv1beta1informers "k8s.io/client-go/informers/admissionregistration/v1beta1"
	upstreamadmissionregistrationv1beta1listers "k8s.io/client-go/listers/admissionregistration/v1beta1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	admissionregistrationv1beta1listers "github.com/kcp-dev/client-go/listers/admissionregistration/v1beta1"
)

// MutatingWebhookConfigurationClusterInformer provides access to a shared informer and lister for
// MutatingWebhookConfigurations.
type MutatingWebhookConfigurationClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamadmissionregistrationv1beta1informers.MutatingWebhookConfigurationInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() admissionregistrationv1beta1listers.MutatingWebhookConfigurationClusterLister
}

type mutatingWebhookConfigurationClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMutatingWebhookConfigurationClusterInformer constructs a new informer for MutatingWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMutatingWebhookConfigurationClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredMutatingWebhookConfigurationClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMutatingWebhookConfigurationClusterInformer constructs a new informer for MutatingWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMutatingWebhookConfigurationClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Watch(context.TODO(), options)
			},
		},
		&admissionregistrationv1beta1.MutatingWebhookConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *mutatingWebhookConfigurationClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredMutatingWebhookConfigurationClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *mutatingWebhookConfigurationClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&admissionregistrationv1beta1.MutatingWebhookConfiguration{}, f.defaultInformer)
}

func (f *mutatingWebhookConfigurationClusterInformer) Lister() admissionregistrationv1beta1listers.MutatingWebhookConfigurationClusterLister {
	return admissionregistrationv1beta1listers.NewMutatingWebhookConfigurationClusterLister(f.Informer().GetIndexer())
}

func (f *mutatingWebhookConfigurationClusterInformer) Cluster(cluster logicalcluster.Name) upstreamadmissionregistrationv1beta1informers.MutatingWebhookConfigurationInformer {
	return &mutatingWebhookConfigurationInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type mutatingWebhookConfigurationInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamadmissionregistrationv1beta1listers.MutatingWebhookConfigurationLister
}

func (f *mutatingWebhookConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *mutatingWebhookConfigurationInformer) Lister() upstreamadmissionregistrationv1beta1listers.MutatingWebhookConfigurationLister {
	return f.lister
}