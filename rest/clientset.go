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

package rest

import (
	"fmt"
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type ClusterClientset struct {
	clientCache kcpclient.Cache[rest.Interface]
}

func (c *ClusterClientset) Cluster(path logicalcluster.Path) rest.Interface {
	return c.clientCache.ClusterOrDie(path)
}

var _ ClusterInterface = (*ClusterClientset)(nil)

func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*ClusterClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	cache := kcpclient.NewCache(c, httpClient, &kcpclient.Constructor[rest.Interface]{
		NewForConfigAndClient: func(cfg *rest.Config, client *http.Client) (rest.Interface, error) {
			return rest.RESTClientForConfigAndClient(cfg, client)
		},
	})

	return &ClusterClientset{clientCache: cache}, nil
}
