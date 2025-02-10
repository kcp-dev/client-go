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

	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	"k8s.io/client-go/rest"

	kcpcoordinationv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/coordination/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpcoordinationv1beta1.CoordinationV1beta1ClusterInterface = (*CoordinationV1beta1ClusterClient)(nil)

type CoordinationV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *CoordinationV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) coordinationv1beta1.CoordinationV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &CoordinationV1beta1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *CoordinationV1beta1ClusterClient) Leases() kcpcoordinationv1beta1.LeaseClusterInterface {
	return &leasesClusterClient{Fake: c.Fake}
}

var _ coordinationv1beta1.CoordinationV1beta1Interface = (*CoordinationV1beta1Client)(nil)

type CoordinationV1beta1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *CoordinationV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *CoordinationV1beta1Client) Leases(namespace string) coordinationv1beta1.LeaseInterface {
	return &leasesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
