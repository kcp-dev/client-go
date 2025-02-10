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

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta1 "k8s.io/client-go/applyconfigurations/apps/v1beta1"
	appsv1beta1client "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"k8s.io/client-go/testing"

	kcpappsv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/apps/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var controllerRevisionsResource = schema.GroupVersionResource{Group: "apps", Version: "v1beta1", Resource: "controllerrevisions"}
var controllerRevisionsKind = schema.GroupVersionKind{Group: "apps", Version: "v1beta1", Kind: "ControllerRevision"}

type controllerRevisionsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *controllerRevisionsClusterClient) Cluster(clusterPath logicalcluster.Path) kcpappsv1beta1.ControllerRevisionsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &controllerRevisionsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ControllerRevisions that match those selectors across all clusters.
func (c *controllerRevisionsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta1.ControllerRevisionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(controllerRevisionsResource, controllerRevisionsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &appsv1beta1.ControllerRevisionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1beta1.ControllerRevisionList{ListMeta: obj.(*appsv1beta1.ControllerRevisionList).ListMeta}
	for _, item := range obj.(*appsv1beta1.ControllerRevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ControllerRevisions across all clusters.
func (c *controllerRevisionsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(controllerRevisionsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type controllerRevisionsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *controllerRevisionsNamespacer) Namespace(namespace string) appsv1beta1client.ControllerRevisionInterface {
	return &controllerRevisionsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type controllerRevisionsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *controllerRevisionsClient) Create(ctx context.Context, controllerRevision *appsv1beta1.ControllerRevision, opts metav1.CreateOptions) (*appsv1beta1.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, controllerRevision), &appsv1beta1.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.ControllerRevision), err
}

func (c *controllerRevisionsClient) Update(ctx context.Context, controllerRevision *appsv1beta1.ControllerRevision, opts metav1.UpdateOptions) (*appsv1beta1.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, controllerRevision), &appsv1beta1.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.ControllerRevision), err
}

func (c *controllerRevisionsClient) UpdateStatus(ctx context.Context, controllerRevision *appsv1beta1.ControllerRevision, opts metav1.UpdateOptions) (*appsv1beta1.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(controllerRevisionsResource, c.ClusterPath, "status", c.Namespace, controllerRevision), &appsv1beta1.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.ControllerRevision), err
}

func (c *controllerRevisionsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(controllerRevisionsResource, c.ClusterPath, c.Namespace, name, opts), &appsv1beta1.ControllerRevision{})
	return err
}

func (c *controllerRevisionsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &appsv1beta1.ControllerRevisionList{})
	return err
}

func (c *controllerRevisionsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*appsv1beta1.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, name), &appsv1beta1.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.ControllerRevision), err
}

// List takes label and field selectors, and returns the list of ControllerRevisions that match those selectors.
func (c *controllerRevisionsClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta1.ControllerRevisionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(controllerRevisionsResource, controllerRevisionsKind, c.ClusterPath, c.Namespace, opts), &appsv1beta1.ControllerRevisionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1beta1.ControllerRevisionList{ListMeta: obj.(*appsv1beta1.ControllerRevisionList).ListMeta}
	for _, item := range obj.(*appsv1beta1.ControllerRevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *controllerRevisionsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *controllerRevisionsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appsv1beta1.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &appsv1beta1.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.ControllerRevision), err
}

func (c *controllerRevisionsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsappsv1beta1.ControllerRevisionApplyConfiguration, opts metav1.ApplyOptions) (*appsv1beta1.ControllerRevision, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &appsv1beta1.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.ControllerRevision), err
}

func (c *controllerRevisionsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsappsv1beta1.ControllerRevisionApplyConfiguration, opts metav1.ApplyOptions) (*appsv1beta1.ControllerRevision, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(controllerRevisionsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &appsv1beta1.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.ControllerRevision), err
}
