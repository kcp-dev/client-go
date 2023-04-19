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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1alpha1 "k8s.io/client-go/kubernetes/typed/resource/v1alpha1"
	"k8s.io/client-go/rest"

	kcpresourcev1alpha1 "github.com/kcp-dev/client-go/kubernetes/typed/resource/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpresourcev1alpha1.ResourceV1alpha1ClusterInterface = (*ResourceV1alpha1ClusterClient)(nil)

type ResourceV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ResourceV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) resourcev1alpha1.ResourceV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ResourceV1alpha1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ResourceV1alpha1ClusterClient) ResourceClaims() kcpresourcev1alpha1.ResourceClaimClusterInterface {
	return &resourceClaimsClusterClient{Fake: c.Fake}
}

func (c *ResourceV1alpha1ClusterClient) PodSchedulings() kcpresourcev1alpha1.PodSchedulingClusterInterface {
	return &podSchedulingsClusterClient{Fake: c.Fake}
}

func (c *ResourceV1alpha1ClusterClient) ResourceClasses() kcpresourcev1alpha1.ResourceClassClusterInterface {
	return &resourceClassesClusterClient{Fake: c.Fake}
}

func (c *ResourceV1alpha1ClusterClient) ResourceClaimTemplates() kcpresourcev1alpha1.ResourceClaimTemplateClusterInterface {
	return &resourceClaimTemplatesClusterClient{Fake: c.Fake}
}

var _ resourcev1alpha1.ResourceV1alpha1Interface = (*ResourceV1alpha1Client)(nil)

type ResourceV1alpha1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ResourceV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ResourceV1alpha1Client) ResourceClaims(namespace string) resourcev1alpha1.ResourceClaimInterface {
	return &resourceClaimsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ResourceV1alpha1Client) PodSchedulings(namespace string) resourcev1alpha1.PodSchedulingInterface {
	return &podSchedulingsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ResourceV1alpha1Client) ResourceClasses() resourcev1alpha1.ResourceClassInterface {
	return &resourceClassesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *ResourceV1alpha1Client) ResourceClaimTemplates(namespace string) resourcev1alpha1.ResourceClaimTemplateInterface {
	return &resourceClaimTemplatesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}