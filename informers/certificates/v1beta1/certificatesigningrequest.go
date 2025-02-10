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

	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamcertificatesv1beta1informers "k8s.io/client-go/informers/certificates/v1beta1"
	upstreamcertificatesv1beta1listers "k8s.io/client-go/listers/certificates/v1beta1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	certificatesv1beta1listers "github.com/kcp-dev/client-go/listers/certificates/v1beta1"
)

// CertificateSigningRequestClusterInformer provides access to a shared informer and lister for
// CertificateSigningRequests.
type CertificateSigningRequestClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamcertificatesv1beta1informers.CertificateSigningRequestInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() certificatesv1beta1listers.CertificateSigningRequestClusterLister
}

type certificateSigningRequestClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCertificateSigningRequestClusterInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateSigningRequestClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredCertificateSigningRequestClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateSigningRequestClusterInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateSigningRequestClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1beta1().CertificateSigningRequests().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1beta1().CertificateSigningRequests().Watch(context.TODO(), options)
			},
		},
		&certificatesv1beta1.CertificateSigningRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateSigningRequestClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredCertificateSigningRequestClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *certificateSigningRequestClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&certificatesv1beta1.CertificateSigningRequest{}, f.defaultInformer)
}

func (f *certificateSigningRequestClusterInformer) Lister() certificatesv1beta1listers.CertificateSigningRequestClusterLister {
	return certificatesv1beta1listers.NewCertificateSigningRequestClusterLister(f.Informer().GetIndexer())
}

func (f *certificateSigningRequestClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamcertificatesv1beta1informers.CertificateSigningRequestInformer {
	return &certificateSigningRequestInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type certificateSigningRequestInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamcertificatesv1beta1listers.CertificateSigningRequestLister
}

func (f *certificateSigningRequestInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *certificateSigningRequestInformer) Lister() upstreamcertificatesv1beta1listers.CertificateSigningRequestLister {
	return f.lister
}
