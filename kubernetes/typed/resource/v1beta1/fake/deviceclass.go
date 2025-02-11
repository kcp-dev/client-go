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

	resourcev1beta1 "k8s.io/api/resource/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsresourcev1beta1 "k8s.io/client-go/applyconfigurations/resource/v1beta1"
	resourcev1beta1client "k8s.io/client-go/kubernetes/typed/resource/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var deviceClassesResource = schema.GroupVersionResource{Group: "resource.k8s.io", Version: "v1beta1", Resource: "deviceclasses"}
var deviceClassesKind = schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1beta1", Kind: "DeviceClass"}

type deviceClassesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *deviceClassesClusterClient) Cluster(clusterPath logicalcluster.Path) resourcev1beta1client.DeviceClassInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &deviceClassesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of DeviceClasses that match those selectors across all clusters.
func (c *deviceClassesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1beta1.DeviceClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(deviceClassesResource, deviceClassesKind, logicalcluster.Wildcard, opts), &resourcev1beta1.DeviceClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1beta1.DeviceClassList{ListMeta: obj.(*resourcev1beta1.DeviceClassList).ListMeta}
	for _, item := range obj.(*resourcev1beta1.DeviceClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested DeviceClasses across all clusters.
func (c *deviceClassesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(deviceClassesResource, logicalcluster.Wildcard, opts))
}

type deviceClassesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *deviceClassesClient) Create(ctx context.Context, deviceClass *resourcev1beta1.DeviceClass, opts metav1.CreateOptions) (*resourcev1beta1.DeviceClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(deviceClassesResource, c.ClusterPath, deviceClass), &resourcev1beta1.DeviceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1beta1.DeviceClass), err
}

func (c *deviceClassesClient) Update(ctx context.Context, deviceClass *resourcev1beta1.DeviceClass, opts metav1.UpdateOptions) (*resourcev1beta1.DeviceClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(deviceClassesResource, c.ClusterPath, deviceClass), &resourcev1beta1.DeviceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1beta1.DeviceClass), err
}

func (c *deviceClassesClient) UpdateStatus(ctx context.Context, deviceClass *resourcev1beta1.DeviceClass, opts metav1.UpdateOptions) (*resourcev1beta1.DeviceClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(deviceClassesResource, c.ClusterPath, "status", deviceClass), &resourcev1beta1.DeviceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1beta1.DeviceClass), err
}

func (c *deviceClassesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(deviceClassesResource, c.ClusterPath, name, opts), &resourcev1beta1.DeviceClass{})
	return err
}

func (c *deviceClassesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(deviceClassesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &resourcev1beta1.DeviceClassList{})
	return err
}

func (c *deviceClassesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*resourcev1beta1.DeviceClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(deviceClassesResource, c.ClusterPath, name), &resourcev1beta1.DeviceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1beta1.DeviceClass), err
}

// List takes label and field selectors, and returns the list of DeviceClasses that match those selectors.
func (c *deviceClassesClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1beta1.DeviceClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(deviceClassesResource, deviceClassesKind, c.ClusterPath, opts), &resourcev1beta1.DeviceClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1beta1.DeviceClassList{ListMeta: obj.(*resourcev1beta1.DeviceClassList).ListMeta}
	for _, item := range obj.(*resourcev1beta1.DeviceClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *deviceClassesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(deviceClassesResource, c.ClusterPath, opts))
}

func (c *deviceClassesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*resourcev1beta1.DeviceClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(deviceClassesResource, c.ClusterPath, name, pt, data, subresources...), &resourcev1beta1.DeviceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1beta1.DeviceClass), err
}

func (c *deviceClassesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1beta1.DeviceClassApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1beta1.DeviceClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(deviceClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &resourcev1beta1.DeviceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1beta1.DeviceClass), err
}

func (c *deviceClassesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1beta1.DeviceClassApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1beta1.DeviceClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(deviceClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &resourcev1beta1.DeviceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1beta1.DeviceClass), err
}
