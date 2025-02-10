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

	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"k8s.io/client-go/rest"

	kcpextensionsv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/extensions/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpextensionsv1beta1.ExtensionsV1beta1ClusterInterface = (*ExtensionsV1beta1ClusterClient)(nil)

type ExtensionsV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ExtensionsV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) extensionsv1beta1.ExtensionsV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ExtensionsV1beta1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ExtensionsV1beta1ClusterClient) Deployments() kcpextensionsv1beta1.DeploymentClusterInterface {
	return &deploymentsClusterClient{Fake: c.Fake}
}

func (c *ExtensionsV1beta1ClusterClient) DaemonSets() kcpextensionsv1beta1.DaemonSetClusterInterface {
	return &daemonSetsClusterClient{Fake: c.Fake}
}

func (c *ExtensionsV1beta1ClusterClient) Ingresses() kcpextensionsv1beta1.IngressClusterInterface {
	return &ingressesClusterClient{Fake: c.Fake}
}

func (c *ExtensionsV1beta1ClusterClient) ReplicaSets() kcpextensionsv1beta1.ReplicaSetClusterInterface {
	return &replicaSetsClusterClient{Fake: c.Fake}
}

func (c *ExtensionsV1beta1ClusterClient) NetworkPolicies() kcpextensionsv1beta1.NetworkPolicyClusterInterface {
	return &networkPoliciesClusterClient{Fake: c.Fake}
}

var _ extensionsv1beta1.ExtensionsV1beta1Interface = (*ExtensionsV1beta1Client)(nil)

type ExtensionsV1beta1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ExtensionsV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ExtensionsV1beta1Client) Deployments(namespace string) extensionsv1beta1.DeploymentInterface {
	return &deploymentsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExtensionsV1beta1Client) DaemonSets(namespace string) extensionsv1beta1.DaemonSetInterface {
	return &daemonSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExtensionsV1beta1Client) Ingresses(namespace string) extensionsv1beta1.IngressInterface {
	return &ingressesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExtensionsV1beta1Client) ReplicaSets(namespace string) extensionsv1beta1.ReplicaSetInterface {
	return &replicaSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExtensionsV1beta1Client) NetworkPolicies(namespace string) extensionsv1beta1.NetworkPolicyInterface {
	return &networkPoliciesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
