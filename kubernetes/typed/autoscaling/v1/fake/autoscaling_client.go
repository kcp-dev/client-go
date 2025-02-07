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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	"k8s.io/client-go/rest"

	kcpautoscalingv1 "github.com/kcp-dev/client-go/kubernetes/typed/autoscaling/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpautoscalingv1.AutoscalingV1ClusterInterface = (*AutoscalingV1ClusterClient)(nil)

type AutoscalingV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *AutoscalingV1ClusterClient) Cluster(clusterPath logicalcluster.Path) autoscalingv1.AutoscalingV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AutoscalingV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *AutoscalingV1ClusterClient) HorizontalPodAutoscalers() kcpautoscalingv1.HorizontalPodAutoscalerClusterInterface {
	return &horizontalPodAutoscalersClusterClient{Fake: c.Fake}
}

var _ autoscalingv1.AutoscalingV1Interface = (*AutoscalingV1Client)(nil)

type AutoscalingV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *AutoscalingV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AutoscalingV1Client) HorizontalPodAutoscalers(namespace string) autoscalingv1.HorizontalPodAutoscalerInterface {
	return &horizontalPodAutoscalersClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
