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

	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/rest"

	kcpadmissionregistrationv1 "github.com/kcp-dev/client-go/kubernetes/typed/admissionregistration/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpadmissionregistrationv1.AdmissionregistrationV1ClusterInterface = (*AdmissionregistrationV1ClusterClient)(nil)

type AdmissionregistrationV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *AdmissionregistrationV1ClusterClient) Cluster(clusterPath logicalcluster.Path) admissionregistrationv1.AdmissionregistrationV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AdmissionregistrationV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *AdmissionregistrationV1ClusterClient) ValidatingAdmissionPolicies() kcpadmissionregistrationv1.ValidatingAdmissionPolicyClusterInterface {
	return &validatingAdmissionPoliciesClusterClient{Fake: c.Fake}
}

func (c *AdmissionregistrationV1ClusterClient) ValidatingAdmissionPolicyBindings() kcpadmissionregistrationv1.ValidatingAdmissionPolicyBindingClusterInterface {
	return &validatingAdmissionPolicyBindingsClusterClient{Fake: c.Fake}
}

func (c *AdmissionregistrationV1ClusterClient) ValidatingWebhookConfigurations() kcpadmissionregistrationv1.ValidatingWebhookConfigurationClusterInterface {
	return &validatingWebhookConfigurationsClusterClient{Fake: c.Fake}
}

func (c *AdmissionregistrationV1ClusterClient) MutatingWebhookConfigurations() kcpadmissionregistrationv1.MutatingWebhookConfigurationClusterInterface {
	return &mutatingWebhookConfigurationsClusterClient{Fake: c.Fake}
}

var _ admissionregistrationv1.AdmissionregistrationV1Interface = (*AdmissionregistrationV1Client)(nil)

type AdmissionregistrationV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *AdmissionregistrationV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AdmissionregistrationV1Client) ValidatingAdmissionPolicies() admissionregistrationv1.ValidatingAdmissionPolicyInterface {
	return &validatingAdmissionPoliciesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *AdmissionregistrationV1Client) ValidatingAdmissionPolicyBindings() admissionregistrationv1.ValidatingAdmissionPolicyBindingInterface {
	return &validatingAdmissionPolicyBindingsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *AdmissionregistrationV1Client) ValidatingWebhookConfigurations() admissionregistrationv1.ValidatingWebhookConfigurationInterface {
	return &validatingWebhookConfigurationsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *AdmissionregistrationV1Client) MutatingWebhookConfigurations() admissionregistrationv1.MutatingWebhookConfigurationInterface {
	return &mutatingWebhookConfigurationsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
