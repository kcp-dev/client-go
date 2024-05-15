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

package v1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamadmissionregistrationv1informers "k8s.io/client-go/informers/admissionregistration/v1"
	upstreamadmissionregistrationv1listers "k8s.io/client-go/listers/admissionregistration/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	admissionregistrationv1listers "github.com/kcp-dev/client-go/listers/admissionregistration/v1"
)

// ValidatingAdmissionPolicyClusterInformer provides access to a shared informer and lister for
// ValidatingAdmissionPolicies.
type ValidatingAdmissionPolicyClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamadmissionregistrationv1informers.ValidatingAdmissionPolicyInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() admissionregistrationv1listers.ValidatingAdmissionPolicyClusterLister
}

type validatingAdmissionPolicyClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewValidatingAdmissionPolicyClusterInformer constructs a new informer for ValidatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewValidatingAdmissionPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredValidatingAdmissionPolicyClusterInformer constructs a new informer for ValidatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredValidatingAdmissionPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1().ValidatingAdmissionPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1().ValidatingAdmissionPolicies().Watch(context.TODO(), options)
			},
		},
		&admissionregistrationv1.ValidatingAdmissionPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *validatingAdmissionPolicyClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *validatingAdmissionPolicyClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&admissionregistrationv1.ValidatingAdmissionPolicy{}, f.defaultInformer)
}

func (f *validatingAdmissionPolicyClusterInformer) Lister() admissionregistrationv1listers.ValidatingAdmissionPolicyClusterLister {
	return admissionregistrationv1listers.NewValidatingAdmissionPolicyClusterLister(f.Informer().GetIndexer())
}

func (f *validatingAdmissionPolicyClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamadmissionregistrationv1informers.ValidatingAdmissionPolicyInformer {
	return &validatingAdmissionPolicyInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type validatingAdmissionPolicyInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamadmissionregistrationv1listers.ValidatingAdmissionPolicyLister
}

func (f *validatingAdmissionPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *validatingAdmissionPolicyInformer) Lister() upstreamadmissionregistrationv1listers.ValidatingAdmissionPolicyLister {
	return f.lister
}
