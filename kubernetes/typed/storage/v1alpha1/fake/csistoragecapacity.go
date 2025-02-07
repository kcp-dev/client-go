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

	storagev1alpha1 "k8s.io/api/storage/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1alpha1 "k8s.io/client-go/applyconfigurations/storage/v1alpha1"
	storagev1alpha1client "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	"k8s.io/client-go/testing"

	kcpstoragev1alpha1 "github.com/kcp-dev/client-go/kubernetes/typed/storage/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var cSIStorageCapacitiesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1alpha1", Resource: "csistoragecapacities"}
var cSIStorageCapacitiesKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1alpha1", Kind: "CSIStorageCapacity"}

type cSIStorageCapacitiesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cSIStorageCapacitiesClusterClient) Cluster(clusterPath logicalcluster.Path) kcpstoragev1alpha1.CSIStorageCapacitiesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cSIStorageCapacitiesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of CSIStorageCapacities that match those selectors across all clusters.
func (c *cSIStorageCapacitiesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1alpha1.CSIStorageCapacityList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cSIStorageCapacitiesResource, cSIStorageCapacitiesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &storagev1alpha1.CSIStorageCapacityList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1alpha1.CSIStorageCapacityList{ListMeta: obj.(*storagev1alpha1.CSIStorageCapacityList).ListMeta}
	for _, item := range obj.(*storagev1alpha1.CSIStorageCapacityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested CSIStorageCapacities across all clusters.
func (c *cSIStorageCapacitiesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cSIStorageCapacitiesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type cSIStorageCapacitiesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *cSIStorageCapacitiesNamespacer) Namespace(namespace string) storagev1alpha1client.CSIStorageCapacityInterface {
	return &cSIStorageCapacitiesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type cSIStorageCapacitiesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *cSIStorageCapacitiesClient) Create(ctx context.Context, cSIStorageCapacity *storagev1alpha1.CSIStorageCapacity, opts metav1.CreateOptions) (*storagev1alpha1.CSIStorageCapacity, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, cSIStorageCapacity), &storagev1alpha1.CSIStorageCapacity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), err
}

func (c *cSIStorageCapacitiesClient) Update(ctx context.Context, cSIStorageCapacity *storagev1alpha1.CSIStorageCapacity, opts metav1.UpdateOptions) (*storagev1alpha1.CSIStorageCapacity, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, cSIStorageCapacity), &storagev1alpha1.CSIStorageCapacity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), err
}

func (c *cSIStorageCapacitiesClient) UpdateStatus(ctx context.Context, cSIStorageCapacity *storagev1alpha1.CSIStorageCapacity, opts metav1.UpdateOptions) (*storagev1alpha1.CSIStorageCapacity, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(cSIStorageCapacitiesResource, c.ClusterPath, "status", c.Namespace, cSIStorageCapacity), &storagev1alpha1.CSIStorageCapacity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), err
}

func (c *cSIStorageCapacitiesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, name, opts), &storagev1alpha1.CSIStorageCapacity{})
	return err
}

func (c *cSIStorageCapacitiesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1alpha1.CSIStorageCapacityList{})
	return err
}

func (c *cSIStorageCapacitiesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*storagev1alpha1.CSIStorageCapacity, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, name), &storagev1alpha1.CSIStorageCapacity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), err
}

// List takes label and field selectors, and returns the list of CSIStorageCapacities that match those selectors.
func (c *cSIStorageCapacitiesClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1alpha1.CSIStorageCapacityList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cSIStorageCapacitiesResource, cSIStorageCapacitiesKind, c.ClusterPath, c.Namespace, opts), &storagev1alpha1.CSIStorageCapacityList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1alpha1.CSIStorageCapacityList{ListMeta: obj.(*storagev1alpha1.CSIStorageCapacityList).ListMeta}
	for _, item := range obj.(*storagev1alpha1.CSIStorageCapacityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *cSIStorageCapacitiesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *cSIStorageCapacitiesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagev1alpha1.CSIStorageCapacity, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &storagev1alpha1.CSIStorageCapacity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), err
}

func (c *cSIStorageCapacitiesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1alpha1.CSIStorageCapacityApplyConfiguration, opts metav1.ApplyOptions) (*storagev1alpha1.CSIStorageCapacity, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &storagev1alpha1.CSIStorageCapacity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), err
}

func (c *cSIStorageCapacitiesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1alpha1.CSIStorageCapacityApplyConfiguration, opts metav1.ApplyOptions) (*storagev1alpha1.CSIStorageCapacity, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cSIStorageCapacitiesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &storagev1alpha1.CSIStorageCapacity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), err
}
