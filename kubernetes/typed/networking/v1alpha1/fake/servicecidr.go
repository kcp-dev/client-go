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
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	networkingv1alpha1 "k8s.io/api/networking/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnetworkingv1alpha1 "k8s.io/client-go/applyconfigurations/networking/v1alpha1"
	networkingv1alpha1client "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var serviceCIDRsResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1alpha1", Resource: "servicecidrs"}
var serviceCIDRsKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1alpha1", Kind: "ServiceCIDR"}

type serviceCIDRsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *serviceCIDRsClusterClient) Cluster(clusterPath logicalcluster.Path) networkingv1alpha1client.ServiceCIDRInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &serviceCIDRsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ServiceCIDRs that match those selectors across all clusters.
func (c *serviceCIDRsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1alpha1.ServiceCIDRList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(serviceCIDRsResource, serviceCIDRsKind, logicalcluster.Wildcard, opts), &networkingv1alpha1.ServiceCIDRList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1alpha1.ServiceCIDRList{ListMeta: obj.(*networkingv1alpha1.ServiceCIDRList).ListMeta}
	for _, item := range obj.(*networkingv1alpha1.ServiceCIDRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ServiceCIDRs across all clusters.
func (c *serviceCIDRsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(serviceCIDRsResource, logicalcluster.Wildcard, opts))
}

type serviceCIDRsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *serviceCIDRsClient) Create(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDR, opts metav1.CreateOptions) (*networkingv1alpha1.ServiceCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(serviceCIDRsResource, c.ClusterPath, serviceCIDR), &networkingv1alpha1.ServiceCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), err
}

func (c *serviceCIDRsClient) Update(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDR, opts metav1.UpdateOptions) (*networkingv1alpha1.ServiceCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(serviceCIDRsResource, c.ClusterPath, serviceCIDR), &networkingv1alpha1.ServiceCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), err
}

func (c *serviceCIDRsClient) UpdateStatus(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDR, opts metav1.UpdateOptions) (*networkingv1alpha1.ServiceCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(serviceCIDRsResource, c.ClusterPath, "status", serviceCIDR), &networkingv1alpha1.ServiceCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), err
}

func (c *serviceCIDRsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(serviceCIDRsResource, c.ClusterPath, name, opts), &networkingv1alpha1.ServiceCIDR{})
	return err
}

func (c *serviceCIDRsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(serviceCIDRsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1alpha1.ServiceCIDRList{})
	return err
}

func (c *serviceCIDRsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*networkingv1alpha1.ServiceCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(serviceCIDRsResource, c.ClusterPath, name), &networkingv1alpha1.ServiceCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), err
}

// List takes label and field selectors, and returns the list of ServiceCIDRs that match those selectors.
func (c *serviceCIDRsClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1alpha1.ServiceCIDRList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(serviceCIDRsResource, serviceCIDRsKind, c.ClusterPath, opts), &networkingv1alpha1.ServiceCIDRList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1alpha1.ServiceCIDRList{ListMeta: obj.(*networkingv1alpha1.ServiceCIDRList).ListMeta}
	for _, item := range obj.(*networkingv1alpha1.ServiceCIDRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *serviceCIDRsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(serviceCIDRsResource, c.ClusterPath, opts))
}

func (c *serviceCIDRsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*networkingv1alpha1.ServiceCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(serviceCIDRsResource, c.ClusterPath, name, pt, data, subresources...), &networkingv1alpha1.ServiceCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), err
}

func (c *serviceCIDRsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1alpha1.ServiceCIDRApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1alpha1.ServiceCIDR, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(serviceCIDRsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &networkingv1alpha1.ServiceCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), err
}

func (c *serviceCIDRsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1alpha1.ServiceCIDRApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1alpha1.ServiceCIDR, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(serviceCIDRsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &networkingv1alpha1.ServiceCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), err
}
