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

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	certificatesv1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	certificatesv1client "k8s.io/client-go/kubernetes/typed/certificates/v1"
)

// CertificateSigningRequestsClusterGetter has a method to return a CertificateSigningRequestClusterInterface.
// A group's cluster client should implement this interface.
type CertificateSigningRequestsClusterGetter interface {
	CertificateSigningRequests() CertificateSigningRequestClusterInterface
}

// CertificateSigningRequestClusterInterface can operate on CertificateSigningRequests across all clusters,
// or scope down to one cluster and return a certificatesv1client.CertificateSigningRequestInterface.
type CertificateSigningRequestClusterInterface interface {
	Cluster(logicalcluster.Path) certificatesv1client.CertificateSigningRequestInterface
	List(ctx context.Context, opts metav1.ListOptions) (*certificatesv1.CertificateSigningRequestList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type certificateSigningRequestsClusterInterface struct {
	clientCache kcpclient.Cache[*certificatesv1client.CertificatesV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *certificateSigningRequestsClusterInterface) Cluster(clusterPath logicalcluster.Path) certificatesv1client.CertificateSigningRequestInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).CertificateSigningRequests()
}

// List returns the entire collection of all CertificateSigningRequests across all clusters.
func (c *certificateSigningRequestsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*certificatesv1.CertificateSigningRequestList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).CertificateSigningRequests().List(ctx, opts)
}

// Watch begins to watch all CertificateSigningRequests across all clusters.
func (c *certificateSigningRequestsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).CertificateSigningRequests().Watch(ctx, opts)
}
