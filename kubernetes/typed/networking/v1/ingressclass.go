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

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	networkingv1client "k8s.io/client-go/kubernetes/typed/networking/v1"
)

// IngressClassesClusterGetter has a method to return a IngressClassClusterInterface.
// A group's cluster client should implement this interface.
type IngressClassesClusterGetter interface {
	IngressClasses() IngressClassClusterInterface
}

// IngressClassClusterInterface can operate on IngressClasses across all clusters,
// or scope down to one cluster and return a networkingv1client.IngressClassInterface.
type IngressClassClusterInterface interface {
	Cluster(logicalcluster.Path) networkingv1client.IngressClassInterface
	List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.IngressClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type ingressClassesClusterInterface struct {
	clientCache kcpclient.Cache[*networkingv1client.NetworkingV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *ingressClassesClusterInterface) Cluster(clusterPath logicalcluster.Path) networkingv1client.IngressClassInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).IngressClasses()
}

// List returns the entire collection of all IngressClasses across all clusters.
func (c *ingressClassesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.IngressClassList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).IngressClasses().List(ctx, opts)
}

// Watch begins to watch all IngressClasses across all clusters.
func (c *ingressClassesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).IngressClasses().Watch(ctx, opts)
}
