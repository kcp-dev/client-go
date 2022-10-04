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

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	coordinationv1 "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	coordinationv1client "k8s.io/client-go/kubernetes/typed/coordination/v1"
)

// LeasesClusterGetter has a method to return a LeasesClusterInterface.
// A group's cluster client should implement this interface.
type LeasesClusterGetter interface {
	Leases() LeasesClusterInterface
}

// LeasesClusterInterface can operate on Leases across all clusters,
// or scope down to one cluster and return a LeasesNamespacer.
type LeasesClusterInterface interface {
	Cluster(logicalcluster.Name) LeasesNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*coordinationv1.LeaseList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type leasesClusterInterface struct {
	clientCache kcpclient.Cache[*coordinationv1client.CoordinationV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *leasesClusterInterface) Cluster(name logicalcluster.Name) LeasesNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &leasesNamespacer{clientCache: c.clientCache, name: name}
}

// List returns the entire collection of all Leases across all clusters.
func (c *leasesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*coordinationv1.LeaseList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Leases(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Leases across all clusters.
func (c *leasesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Leases(metav1.NamespaceAll).Watch(ctx, opts)
}

// LeasesNamespacer can scope to objects within a namespace, returning a coordinationv1client.LeaseInterface.
type LeasesNamespacer interface {
	Namespace(string) coordinationv1client.LeaseInterface
}

type leasesNamespacer struct {
	clientCache kcpclient.Cache[*coordinationv1client.CoordinationV1Client]
	name        logicalcluster.Name
}

func (n *leasesNamespacer) Namespace(namespace string) coordinationv1client.LeaseInterface {
	return n.clientCache.ClusterOrDie(n.name).Leases(namespace)
}
