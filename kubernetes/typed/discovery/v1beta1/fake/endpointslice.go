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

	discoveryv1beta1 "k8s.io/api/discovery/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsdiscoveryv1beta1 "k8s.io/client-go/applyconfigurations/discovery/v1beta1"
	discoveryv1beta1client "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	"k8s.io/client-go/testing"

	kcpdiscoveryv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/discovery/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var endpointSlicesResource = schema.GroupVersionResource{Group: "discovery.k8s.io", Version: "v1beta1", Resource: "endpointslices"}
var endpointSlicesKind = schema.GroupVersionKind{Group: "discovery.k8s.io", Version: "v1beta1", Kind: "EndpointSlice"}

type endpointSlicesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *endpointSlicesClusterClient) Cluster(clusterPath logicalcluster.Path) kcpdiscoveryv1beta1.EndpointSlicesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &endpointSlicesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of EndpointSlices that match those selectors across all clusters.
func (c *endpointSlicesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*discoveryv1beta1.EndpointSliceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(endpointSlicesResource, endpointSlicesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &discoveryv1beta1.EndpointSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &discoveryv1beta1.EndpointSliceList{ListMeta: obj.(*discoveryv1beta1.EndpointSliceList).ListMeta}
	for _, item := range obj.(*discoveryv1beta1.EndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested EndpointSlices across all clusters.
func (c *endpointSlicesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(endpointSlicesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type endpointSlicesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *endpointSlicesNamespacer) Namespace(namespace string) discoveryv1beta1client.EndpointSliceInterface {
	return &endpointSlicesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type endpointSlicesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *endpointSlicesClient) Create(ctx context.Context, endpointSlice *discoveryv1beta1.EndpointSlice, opts metav1.CreateOptions) (*discoveryv1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(endpointSlicesResource, c.ClusterPath, c.Namespace, endpointSlice), &discoveryv1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*discoveryv1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) Update(ctx context.Context, endpointSlice *discoveryv1beta1.EndpointSlice, opts metav1.UpdateOptions) (*discoveryv1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(endpointSlicesResource, c.ClusterPath, c.Namespace, endpointSlice), &discoveryv1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*discoveryv1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) UpdateStatus(ctx context.Context, endpointSlice *discoveryv1beta1.EndpointSlice, opts metav1.UpdateOptions) (*discoveryv1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(endpointSlicesResource, c.ClusterPath, "status", c.Namespace, endpointSlice), &discoveryv1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*discoveryv1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(endpointSlicesResource, c.ClusterPath, c.Namespace, name, opts), &discoveryv1beta1.EndpointSlice{})
	return err
}

func (c *endpointSlicesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(endpointSlicesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &discoveryv1beta1.EndpointSliceList{})
	return err
}

func (c *endpointSlicesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*discoveryv1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(endpointSlicesResource, c.ClusterPath, c.Namespace, name), &discoveryv1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*discoveryv1beta1.EndpointSlice), err
}

// List takes label and field selectors, and returns the list of EndpointSlices that match those selectors.
func (c *endpointSlicesClient) List(ctx context.Context, opts metav1.ListOptions) (*discoveryv1beta1.EndpointSliceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(endpointSlicesResource, endpointSlicesKind, c.ClusterPath, c.Namespace, opts), &discoveryv1beta1.EndpointSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &discoveryv1beta1.EndpointSliceList{ListMeta: obj.(*discoveryv1beta1.EndpointSliceList).ListMeta}
	for _, item := range obj.(*discoveryv1beta1.EndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *endpointSlicesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(endpointSlicesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *endpointSlicesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*discoveryv1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(endpointSlicesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &discoveryv1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*discoveryv1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsdiscoveryv1beta1.EndpointSliceApplyConfiguration, opts metav1.ApplyOptions) (*discoveryv1beta1.EndpointSlice, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(endpointSlicesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &discoveryv1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*discoveryv1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsdiscoveryv1beta1.EndpointSliceApplyConfiguration, opts metav1.ApplyOptions) (*discoveryv1beta1.EndpointSlice, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(endpointSlicesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &discoveryv1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*discoveryv1beta1.EndpointSlice), err
}
