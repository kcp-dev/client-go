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

package v1alpha1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamadmissionregistrationv1alpha1informers "k8s.io/client-go/informers/admissionregistration/v1alpha1"
	upstreamadmissionregistrationv1alpha1listers "k8s.io/client-go/listers/admissionregistration/v1alpha1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	admissionregistrationv1alpha1listers "github.com/kcp-dev/client-go/listers/admissionregistration/v1alpha1"
)

// MutatingAdmissionPolicyClusterInformer provides access to a shared informer and lister for
// MutatingAdmissionPolicies.
type MutatingAdmissionPolicyClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamadmissionregistrationv1alpha1informers.MutatingAdmissionPolicyInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() admissionregistrationv1alpha1listers.MutatingAdmissionPolicyClusterLister
}

type mutatingAdmissionPolicyClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMutatingAdmissionPolicyClusterInformer constructs a new informer for MutatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMutatingAdmissionPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredMutatingAdmissionPolicyClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMutatingAdmissionPolicyClusterInformer constructs a new informer for MutatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMutatingAdmissionPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1alpha1().MutatingAdmissionPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1alpha1().MutatingAdmissionPolicies().Watch(context.TODO(), options)
			},
		},
		&admissionregistrationv1alpha1.MutatingAdmissionPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *mutatingAdmissionPolicyClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredMutatingAdmissionPolicyClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *mutatingAdmissionPolicyClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&admissionregistrationv1alpha1.MutatingAdmissionPolicy{}, f.defaultInformer)
}

func (f *mutatingAdmissionPolicyClusterInformer) Lister() admissionregistrationv1alpha1listers.MutatingAdmissionPolicyClusterLister {
	return admissionregistrationv1alpha1listers.NewMutatingAdmissionPolicyClusterLister(f.Informer().GetIndexer())
}

func (f *mutatingAdmissionPolicyClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamadmissionregistrationv1alpha1informers.MutatingAdmissionPolicyInformer {
	return &mutatingAdmissionPolicyInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type mutatingAdmissionPolicyInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamadmissionregistrationv1alpha1listers.MutatingAdmissionPolicyLister
}

func (f *mutatingAdmissionPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *mutatingAdmissionPolicyInformer) Lister() upstreamadmissionregistrationv1alpha1listers.MutatingAdmissionPolicyLister {
	return f.lister
}
