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

package v1beta1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	storagev1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	storagev1beta1client "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var cSIDriversResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "V1beta1", Resource: "csidrivers"}
var cSIDriversKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "V1beta1", Kind: "CSIDriver"}

type cSIDriversClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cSIDriversClusterClient) Cluster(cluster logicalcluster.Name) storagev1beta1client.CSIDriverInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cSIDriversClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of CSIDrivers that match those selectors across all clusters.
func (c *cSIDriversClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.CSIDriverList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(cSIDriversResource, cSIDriversKind, logicalcluster.Wildcard, opts), &storagev1beta1.CSIDriverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.CSIDriverList{ListMeta: obj.(*storagev1beta1.CSIDriverList).ListMeta}
	for _, item := range obj.(*storagev1beta1.CSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested CSIDrivers across all clusters.
func (c *cSIDriversClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(cSIDriversResource, logicalcluster.Wildcard, opts))
}

type cSIDriversClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *cSIDriversClient) Create(ctx context.Context, cSIDriver *storagev1beta1.CSIDriver, opts metav1.CreateOptions) (*storagev1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(cSIDriversResource, c.Cluster, cSIDriver), &storagev1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSIDriver), err
}

func (c *cSIDriversClient) Update(ctx context.Context, cSIDriver *storagev1beta1.CSIDriver, opts metav1.UpdateOptions) (*storagev1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(cSIDriversResource, c.Cluster, cSIDriver), &storagev1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSIDriver), err
}

func (c *cSIDriversClient) UpdateStatus(ctx context.Context, cSIDriver *storagev1beta1.CSIDriver, opts metav1.UpdateOptions) (*storagev1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(cSIDriversResource, c.Cluster, "status", cSIDriver), &storagev1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSIDriver), err
}

func (c *cSIDriversClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(cSIDriversResource, c.Cluster, name, opts), &storagev1beta1.CSIDriver{})
	return err
}

func (c *cSIDriversClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(cSIDriversResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1beta1.CSIDriverList{})
	return err
}

func (c *cSIDriversClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*storagev1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(cSIDriversResource, c.Cluster, name), &storagev1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSIDriver), err
}

// List takes label and field selectors, and returns the list of CSIDrivers that match those selectors.
func (c *cSIDriversClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.CSIDriverList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(cSIDriversResource, cSIDriversKind, c.Cluster, opts), &storagev1beta1.CSIDriverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.CSIDriverList{ListMeta: obj.(*storagev1beta1.CSIDriverList).ListMeta}
	for _, item := range obj.(*storagev1beta1.CSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *cSIDriversClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(cSIDriversResource, c.Cluster, opts))
}

func (c *cSIDriversClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagev1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(cSIDriversResource, c.Cluster, name, pt, data, subresources...), &storagev1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSIDriver), err
}

func (c *cSIDriversClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.CSIDriverApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.CSIDriver, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(cSIDriversResource, c.Cluster, *name, types.ApplyPatchType, data), &storagev1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSIDriver), err
}

func (c *cSIDriversClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.CSIDriverApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.CSIDriver, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(cSIDriversResource, c.Cluster, *name, types.ApplyPatchType, data, "status"), &storagev1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSIDriver), err
}