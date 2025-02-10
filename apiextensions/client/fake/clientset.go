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

	client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	clientscheme "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"

	kcpclient "github.com/kcp-dev/client-go/apiextensions/client"
	kcpapiextensionsv1 "github.com/kcp-dev/client-go/apiextensions/client/typed/apiextensions/v1"
	fakeapiextensionsv1 "github.com/kcp-dev/client-go/apiextensions/client/typed/apiextensions/v1/fake"
	kcpapiextensionsv1beta1 "github.com/kcp-dev/client-go/apiextensions/client/typed/apiextensions/v1beta1"
	fakeapiextensionsv1beta1 "github.com/kcp-dev/client-go/apiextensions/client/typed/apiextensions/v1beta1/fake"
	kcpfakediscovery "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/discovery/fake"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *ClusterClientset {
	o := kcptesting.NewObjectTracker(clientscheme.Scheme, clientscheme.Codecs.UniversalDecoder())
	o.AddAll(objects...)

	cs := &ClusterClientset{Fake: &kcptesting.Fake{}, tracker: o}
	cs.discovery = &kcpfakediscovery.FakeDiscovery{Fake: cs.Fake, ClusterPath: logicalcluster.Wildcard}
	cs.AddReactor("*", "*", kcptesting.ObjectReaction(o))
	cs.AddWatchReactor("*", kcptesting.WatchReaction(o))

	return cs
}

var _ kcpclient.ClusterInterface = (*ClusterClientset)(nil)

// ClusterClientset contains the clients for groups.
type ClusterClientset struct {
	*kcptesting.Fake
	discovery *kcpfakediscovery.FakeDiscovery
	tracker   kcptesting.ObjectTracker
}

// Discovery retrieves the DiscoveryClient
func (c *ClusterClientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *ClusterClientset) Tracker() kcptesting.ObjectTracker {
	return c.tracker
}

// ApiextensionsV1 retrieves the ApiextensionsV1ClusterClient.
func (c *ClusterClientset) ApiextensionsV1() kcpapiextensionsv1.ApiextensionsV1ClusterInterface {
	return &fakeapiextensionsv1.ApiextensionsV1ClusterClient{Fake: c.Fake}
}

// ApiextensionsV1beta1 retrieves the ApiextensionsV1beta1ClusterClient.
func (c *ClusterClientset) ApiextensionsV1beta1() kcpapiextensionsv1beta1.ApiextensionsV1beta1ClusterInterface {
	return &fakeapiextensionsv1beta1.ApiextensionsV1beta1ClusterClient{Fake: c.Fake}
}

// Cluster scopes this clientset to one cluster.
func (c *ClusterClientset) Cluster(clusterPath logicalcluster.Path) client.Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &Clientset{
		Fake:        c.Fake,
		discovery:   &kcpfakediscovery.FakeDiscovery{Fake: c.Fake, ClusterPath: clusterPath},
		tracker:     c.tracker.Cluster(clusterPath),
		clusterPath: clusterPath,
	}
}

var _ client.Interface = (*Clientset)(nil)

// Clientset contains the clients for groups.
type Clientset struct {
	*kcptesting.Fake
	discovery   *kcpfakediscovery.FakeDiscovery
	tracker     kcptesting.ScopedObjectTracker
	clusterPath logicalcluster.Path
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() kcptesting.ScopedObjectTracker {
	return c.tracker
}

// ApiextensionsV1 retrieves the ApiextensionsV1Client.
func (c *Clientset) ApiextensionsV1() apiextensionsv1.ApiextensionsV1Interface {
	return &fakeapiextensionsv1.ApiextensionsV1Client{Fake: c.Fake, ClusterPath: c.clusterPath}
}

// ApiextensionsV1beta1 retrieves the ApiextensionsV1beta1Client.
func (c *Clientset) ApiextensionsV1beta1() apiextensionsv1beta1.ApiextensionsV1beta1Interface {
	return &fakeapiextensionsv1beta1.ApiextensionsV1beta1Client{Fake: c.Fake, ClusterPath: c.clusterPath}
}
