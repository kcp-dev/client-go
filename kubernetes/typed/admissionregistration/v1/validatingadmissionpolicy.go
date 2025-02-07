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

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
)

// ValidatingAdmissionPoliciesClusterGetter has a method to return a ValidatingAdmissionPolicyClusterInterface.
// A group's cluster client should implement this interface.
type ValidatingAdmissionPoliciesClusterGetter interface {
	ValidatingAdmissionPolicies() ValidatingAdmissionPolicyClusterInterface
}

// ValidatingAdmissionPolicyClusterInterface can operate on ValidatingAdmissionPolicies across all clusters,
// or scope down to one cluster and return a admissionregistrationv1client.ValidatingAdmissionPolicyInterface.
type ValidatingAdmissionPolicyClusterInterface interface {
	Cluster(logicalcluster.Path) admissionregistrationv1client.ValidatingAdmissionPolicyInterface
	List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1.ValidatingAdmissionPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type validatingAdmissionPoliciesClusterInterface struct {
	clientCache kcpclient.Cache[*admissionregistrationv1client.AdmissionregistrationV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *validatingAdmissionPoliciesClusterInterface) Cluster(clusterPath logicalcluster.Path) admissionregistrationv1client.ValidatingAdmissionPolicyInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).ValidatingAdmissionPolicies()
}

// List returns the entire collection of all ValidatingAdmissionPolicies across all clusters.
func (c *validatingAdmissionPoliciesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1.ValidatingAdmissionPolicyList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ValidatingAdmissionPolicies().List(ctx, opts)
}

// Watch begins to watch all ValidatingAdmissionPolicies across all clusters.
func (c *validatingAdmissionPoliciesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ValidatingAdmissionPolicies().Watch(ctx, opts)
}
