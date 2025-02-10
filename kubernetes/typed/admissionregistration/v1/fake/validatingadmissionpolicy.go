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
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsadmissionregistrationv1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
	admissionregistrationv1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var validatingAdmissionPoliciesResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "v1", Resource: "validatingadmissionpolicies"}
var validatingAdmissionPoliciesKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "v1", Kind: "ValidatingAdmissionPolicy"}

type validatingAdmissionPoliciesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *validatingAdmissionPoliciesClusterClient) Cluster(clusterPath logicalcluster.Path) admissionregistrationv1client.ValidatingAdmissionPolicyInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &validatingAdmissionPoliciesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ValidatingAdmissionPolicies that match those selectors across all clusters.
func (c *validatingAdmissionPoliciesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1.ValidatingAdmissionPolicyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(validatingAdmissionPoliciesResource, validatingAdmissionPoliciesKind, logicalcluster.Wildcard, opts), &admissionregistrationv1.ValidatingAdmissionPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistrationv1.ValidatingAdmissionPolicyList{ListMeta: obj.(*admissionregistrationv1.ValidatingAdmissionPolicyList).ListMeta}
	for _, item := range obj.(*admissionregistrationv1.ValidatingAdmissionPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ValidatingAdmissionPolicies across all clusters.
func (c *validatingAdmissionPoliciesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(validatingAdmissionPoliciesResource, logicalcluster.Wildcard, opts))
}

type validatingAdmissionPoliciesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *validatingAdmissionPoliciesClient) Create(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1.ValidatingAdmissionPolicy, opts metav1.CreateOptions) (*admissionregistrationv1.ValidatingAdmissionPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(validatingAdmissionPoliciesResource, c.ClusterPath, validatingAdmissionPolicy), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.ValidatingAdmissionPolicy), err
}

func (c *validatingAdmissionPoliciesClient) Update(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1.ValidatingAdmissionPolicy, opts metav1.UpdateOptions) (*admissionregistrationv1.ValidatingAdmissionPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(validatingAdmissionPoliciesResource, c.ClusterPath, validatingAdmissionPolicy), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.ValidatingAdmissionPolicy), err
}

func (c *validatingAdmissionPoliciesClient) UpdateStatus(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1.ValidatingAdmissionPolicy, opts metav1.UpdateOptions) (*admissionregistrationv1.ValidatingAdmissionPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(validatingAdmissionPoliciesResource, c.ClusterPath, "status", validatingAdmissionPolicy), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.ValidatingAdmissionPolicy), err
}

func (c *validatingAdmissionPoliciesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(validatingAdmissionPoliciesResource, c.ClusterPath, name, opts), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	return err
}

func (c *validatingAdmissionPoliciesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(validatingAdmissionPoliciesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &admissionregistrationv1.ValidatingAdmissionPolicyList{})
	return err
}

func (c *validatingAdmissionPoliciesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*admissionregistrationv1.ValidatingAdmissionPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(validatingAdmissionPoliciesResource, c.ClusterPath, name), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.ValidatingAdmissionPolicy), err
}

// List takes label and field selectors, and returns the list of ValidatingAdmissionPolicies that match those selectors.
func (c *validatingAdmissionPoliciesClient) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1.ValidatingAdmissionPolicyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(validatingAdmissionPoliciesResource, validatingAdmissionPoliciesKind, c.ClusterPath, opts), &admissionregistrationv1.ValidatingAdmissionPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistrationv1.ValidatingAdmissionPolicyList{ListMeta: obj.(*admissionregistrationv1.ValidatingAdmissionPolicyList).ListMeta}
	for _, item := range obj.(*admissionregistrationv1.ValidatingAdmissionPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *validatingAdmissionPoliciesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(validatingAdmissionPoliciesResource, c.ClusterPath, opts))
}

func (c *validatingAdmissionPoliciesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*admissionregistrationv1.ValidatingAdmissionPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingAdmissionPoliciesResource, c.ClusterPath, name, pt, data, subresources...), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.ValidatingAdmissionPolicy), err
}

func (c *validatingAdmissionPoliciesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsadmissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration, opts metav1.ApplyOptions) (*admissionregistrationv1.ValidatingAdmissionPolicy, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingAdmissionPoliciesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.ValidatingAdmissionPolicy), err
}

func (c *validatingAdmissionPoliciesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsadmissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration, opts metav1.ApplyOptions) (*admissionregistrationv1.ValidatingAdmissionPolicy, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingAdmissionPoliciesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &admissionregistrationv1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1.ValidatingAdmissionPolicy), err
}
