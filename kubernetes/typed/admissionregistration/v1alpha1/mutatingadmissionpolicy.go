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

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1alpha1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
)

// MutatingAdmissionPoliciesClusterGetter has a method to return a MutatingAdmissionPolicyClusterInterface.
// A group's cluster client should implement this interface.
type MutatingAdmissionPoliciesClusterGetter interface {
	MutatingAdmissionPolicies() MutatingAdmissionPolicyClusterInterface
}

// MutatingAdmissionPolicyClusterInterface can operate on MutatingAdmissionPolicies across all clusters,
// or scope down to one cluster and return a admissionregistrationv1alpha1client.MutatingAdmissionPolicyInterface.
type MutatingAdmissionPolicyClusterInterface interface {
	Cluster(logicalcluster.Path) admissionregistrationv1alpha1client.MutatingAdmissionPolicyInterface
	List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1alpha1.MutatingAdmissionPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type mutatingAdmissionPoliciesClusterInterface struct {
	clientCache kcpclient.Cache[*admissionregistrationv1alpha1client.AdmissionregistrationV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *mutatingAdmissionPoliciesClusterInterface) Cluster(clusterPath logicalcluster.Path) admissionregistrationv1alpha1client.MutatingAdmissionPolicyInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).MutatingAdmissionPolicies()
}

// List returns the entire collection of all MutatingAdmissionPolicies across all clusters.
func (c *mutatingAdmissionPoliciesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1alpha1.MutatingAdmissionPolicyList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).MutatingAdmissionPolicies().List(ctx, opts)
}

// Watch begins to watch all MutatingAdmissionPolicies across all clusters.
func (c *mutatingAdmissionPoliciesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).MutatingAdmissionPolicies().Watch(ctx, opts)
}
