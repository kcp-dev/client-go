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
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	policyv1beta1client "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
)

// EvictionsClusterGetter has a method to return a EvictionClusterInterface.
// A group's cluster client should implement this interface.
type EvictionsClusterGetter interface {
	Evictions() EvictionClusterInterface
}

// EvictionClusterInterface can scope down to one cluster and return a EvictionsNamespacer.
type EvictionClusterInterface interface {
	Cluster(logicalcluster.Path) EvictionsNamespacer
}

type evictionsClusterInterface struct {
	clientCache kcpclient.Cache[*policyv1beta1client.PolicyV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *evictionsClusterInterface) Cluster(clusterPath logicalcluster.Path) EvictionsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &evictionsNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// EvictionsNamespacer can scope to objects within a namespace, returning a policyv1beta1client.EvictionInterface.
type EvictionsNamespacer interface {
	Namespace(string) policyv1beta1client.EvictionInterface
}

type evictionsNamespacer struct {
	clientCache kcpclient.Cache[*policyv1beta1client.PolicyV1beta1Client]
	clusterPath logicalcluster.Path
}

func (n *evictionsNamespacer) Namespace(namespace string) policyv1beta1client.EvictionInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).Evictions(namespace)
}
