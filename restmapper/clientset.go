/*
Copyright 2022 The KCP Authors.

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

package restmapper

import (
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/restmapper"

	"github.com/kcp-dev/client-go/cache"
	clusterdiscovery "github.com/kcp-dev/client-go/discovery"
)

type DeferredDiscoveryRESTMapper struct {
	clientCache cache.Cache[meta.RESTMapper]
}

func (c *DeferredDiscoveryRESTMapper) Cluster(clusterPath logicalcluster.Path) meta.RESTMapper {
	return c.clientCache.ClusterOrDie(clusterPath)
}

var _ ClusterInterface = (*DeferredDiscoveryRESTMapper)(nil)

func NewDeferredDiscoveryRESTMapper(cl clusterdiscovery.DiscoveryClusterInterface) ClusterInterface {
	return &DeferredDiscoveryRESTMapper{
		clientCache: cache.NewCache(&cache.Constructor[meta.RESTMapper]{Provider: func(clusterPath logicalcluster.Path) (meta.RESTMapper, error) {
			return restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(cl.Cluster(clusterPath))), nil
		}}),
	}
}
