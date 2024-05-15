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

	storagemigrationv1alpha1 "k8s.io/api/storagemigration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragemigrationv1alpha1 "k8s.io/client-go/applyconfigurations/storagemigration/v1alpha1"
	storagemigrationv1alpha1client "k8s.io/client-go/kubernetes/typed/storagemigration/v1alpha1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var storageVersionMigrationsResource = schema.GroupVersionResource{Group: "storagemigration.k8s.io", Version: "v1alpha1", Resource: "storageversionmigrations"}
var storageVersionMigrationsKind = schema.GroupVersionKind{Group: "storagemigration.k8s.io", Version: "v1alpha1", Kind: "StorageVersionMigration"}

type storageVersionMigrationsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *storageVersionMigrationsClusterClient) Cluster(clusterPath logicalcluster.Path) storagemigrationv1alpha1client.StorageVersionMigrationInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &storageVersionMigrationsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of StorageVersionMigrations that match those selectors across all clusters.
func (c *storageVersionMigrationsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*storagemigrationv1alpha1.StorageVersionMigrationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(storageVersionMigrationsResource, storageVersionMigrationsKind, logicalcluster.Wildcard, opts), &storagemigrationv1alpha1.StorageVersionMigrationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagemigrationv1alpha1.StorageVersionMigrationList{ListMeta: obj.(*storagemigrationv1alpha1.StorageVersionMigrationList).ListMeta}
	for _, item := range obj.(*storagemigrationv1alpha1.StorageVersionMigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested StorageVersionMigrations across all clusters.
func (c *storageVersionMigrationsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(storageVersionMigrationsResource, logicalcluster.Wildcard, opts))
}

type storageVersionMigrationsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *storageVersionMigrationsClient) Create(ctx context.Context, storageVersionMigration *storagemigrationv1alpha1.StorageVersionMigration, opts metav1.CreateOptions) (*storagemigrationv1alpha1.StorageVersionMigration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(storageVersionMigrationsResource, c.ClusterPath, storageVersionMigration), &storagemigrationv1alpha1.StorageVersionMigration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagemigrationv1alpha1.StorageVersionMigration), err
}

func (c *storageVersionMigrationsClient) Update(ctx context.Context, storageVersionMigration *storagemigrationv1alpha1.StorageVersionMigration, opts metav1.UpdateOptions) (*storagemigrationv1alpha1.StorageVersionMigration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(storageVersionMigrationsResource, c.ClusterPath, storageVersionMigration), &storagemigrationv1alpha1.StorageVersionMigration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagemigrationv1alpha1.StorageVersionMigration), err
}

func (c *storageVersionMigrationsClient) UpdateStatus(ctx context.Context, storageVersionMigration *storagemigrationv1alpha1.StorageVersionMigration, opts metav1.UpdateOptions) (*storagemigrationv1alpha1.StorageVersionMigration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(storageVersionMigrationsResource, c.ClusterPath, "status", storageVersionMigration), &storagemigrationv1alpha1.StorageVersionMigration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagemigrationv1alpha1.StorageVersionMigration), err
}

func (c *storageVersionMigrationsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(storageVersionMigrationsResource, c.ClusterPath, name, opts), &storagemigrationv1alpha1.StorageVersionMigration{})
	return err
}

func (c *storageVersionMigrationsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(storageVersionMigrationsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &storagemigrationv1alpha1.StorageVersionMigrationList{})
	return err
}

func (c *storageVersionMigrationsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*storagemigrationv1alpha1.StorageVersionMigration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(storageVersionMigrationsResource, c.ClusterPath, name), &storagemigrationv1alpha1.StorageVersionMigration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagemigrationv1alpha1.StorageVersionMigration), err
}

// List takes label and field selectors, and returns the list of StorageVersionMigrations that match those selectors.
func (c *storageVersionMigrationsClient) List(ctx context.Context, opts metav1.ListOptions) (*storagemigrationv1alpha1.StorageVersionMigrationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(storageVersionMigrationsResource, storageVersionMigrationsKind, c.ClusterPath, opts), &storagemigrationv1alpha1.StorageVersionMigrationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagemigrationv1alpha1.StorageVersionMigrationList{ListMeta: obj.(*storagemigrationv1alpha1.StorageVersionMigrationList).ListMeta}
	for _, item := range obj.(*storagemigrationv1alpha1.StorageVersionMigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *storageVersionMigrationsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(storageVersionMigrationsResource, c.ClusterPath, opts))
}

func (c *storageVersionMigrationsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagemigrationv1alpha1.StorageVersionMigration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(storageVersionMigrationsResource, c.ClusterPath, name, pt, data, subresources...), &storagemigrationv1alpha1.StorageVersionMigration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagemigrationv1alpha1.StorageVersionMigration), err
}

func (c *storageVersionMigrationsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsstoragemigrationv1alpha1.StorageVersionMigrationApplyConfiguration, opts metav1.ApplyOptions) (*storagemigrationv1alpha1.StorageVersionMigration, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(storageVersionMigrationsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &storagemigrationv1alpha1.StorageVersionMigration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagemigrationv1alpha1.StorageVersionMigration), err
}

func (c *storageVersionMigrationsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsstoragemigrationv1alpha1.StorageVersionMigrationApplyConfiguration, opts metav1.ApplyOptions) (*storagemigrationv1alpha1.StorageVersionMigration, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(storageVersionMigrationsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &storagemigrationv1alpha1.StorageVersionMigration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagemigrationv1alpha1.StorageVersionMigration), err
}
