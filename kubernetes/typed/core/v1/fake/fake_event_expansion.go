/*
Copyright 2014 The Kubernetes Authors.
Modifications Copyright 2022 The KCP Authors.

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

package fake

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	core "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var eventsResource = schema.GroupVersionResource{Group: "events.k8s.io", Version: "v1beta1", Resource: "events"}
var eventsKind = schema.GroupVersionKind{Group: "events.k8s.io", Version: "v1beta1", Kind: "Event"}

func (c *eventScopedClient) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := core.NewRootCreateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace() != "" {
		action = core.NewCreateAction(eventsResource, c.ClusterPath, c.Namespace(), event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
func (c *eventScopedClient) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := core.NewRootUpdateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace() != "" {
		action = core.NewUpdateAction(eventsResource, c.ClusterPath, c.Namespace(), event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// PatchWithEventNamespace patches an existing event. Returns the copy of the event the server returns, or an error.
// TODO: Should take a PatchType as an argument probably.
func (c *eventScopedClient) PatchWithEventNamespace(event *v1.Event, data []byte) (*v1.Event, error) {
	// TODO: Should be configurable to support additional patch strategies.
	pt := types.StrategicMergePatchType
	action := core.NewRootPatchAction(eventsResource, c.ClusterPath, event.Name, pt, data)
	if c.Namespace() != "" {
		action = core.NewPatchAction(eventsResource, c.ClusterPath, c.Namespace(), event.Name, pt, data)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Search returns a list of events matching the specified object.
func (c *eventScopedClient) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	action := core.NewRootListAction(eventsResource, eventsKind, c.ClusterPath, metav1.ListOptions{})
	if c.Namespace() != "" {
		action = core.NewListAction(eventsResource, eventsKind, c.ClusterPath, c.Namespace(), metav1.ListOptions{})
	}
	obj, err := c.Fake.Invokes(action, &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.EventList), err
}

func (c *eventScopedClient) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	action := core.GenericActionImpl{}
	action.Verb = "get-field-selector"
	action.Resource = eventsResource
	action.ClusterPath = c.ClusterPath

	_, _ = c.Fake.Invokes(action, nil)
	return fields.Everything()
}
