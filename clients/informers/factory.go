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

package informers

import (
	"reflect"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	admissionregistrationinformers "github.com/kcp-dev/client-go/clients/informers/admissionregistration"
	apiserverinternalinformers "github.com/kcp-dev/client-go/clients/informers/apiserverinternal"
	appsinformers "github.com/kcp-dev/client-go/clients/informers/apps"
	autoscalinginformers "github.com/kcp-dev/client-go/clients/informers/autoscaling"
	batchinformers "github.com/kcp-dev/client-go/clients/informers/batch"
	certificatesinformers "github.com/kcp-dev/client-go/clients/informers/certificates"
	coordinationinformers "github.com/kcp-dev/client-go/clients/informers/coordination"
	coreinformers "github.com/kcp-dev/client-go/clients/informers/core"
	discoveryinformers "github.com/kcp-dev/client-go/clients/informers/discovery"
	eventsinformers "github.com/kcp-dev/client-go/clients/informers/events"
	extensionsinformers "github.com/kcp-dev/client-go/clients/informers/extensions"
	flowcontrolinformers "github.com/kcp-dev/client-go/clients/informers/flowcontrol"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	networkinginformers "github.com/kcp-dev/client-go/clients/informers/networking"
	nodeinformers "github.com/kcp-dev/client-go/clients/informers/node"
	policyinformers "github.com/kcp-dev/client-go/clients/informers/policy"
	rbacinformers "github.com/kcp-dev/client-go/clients/informers/rbac"
	schedulinginformers "github.com/kcp-dev/client-go/clients/informers/scheduling"
	storageinformers "github.com/kcp-dev/client-go/clients/informers/storage"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client           clientset.ClusterInterface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[metav1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of SharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client clientset.ClusterInterface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client clientset.ClusterInterface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Admissionregistration() admissionregistrationinformers.ClusterInterface
	Internal() apiserverinternalinformers.ClusterInterface
	Apps() appsinformers.ClusterInterface
	Autoscaling() autoscalinginformers.ClusterInterface
	Batch() batchinformers.ClusterInterface
	Certificates() certificatesinformers.ClusterInterface
	Coordination() coordinationinformers.ClusterInterface
	Core() coreinformers.ClusterInterface
	Discovery() discoveryinformers.ClusterInterface
	Events() eventsinformers.ClusterInterface
	Extensions() extensionsinformers.ClusterInterface
	Flowcontrol() flowcontrolinformers.ClusterInterface
	Networking() networkinginformers.ClusterInterface
	Node() nodeinformers.ClusterInterface
	Policy() policyinformers.ClusterInterface
	Rbac() rbacinformers.ClusterInterface
	Scheduling() schedulinginformers.ClusterInterface
	Storage() storageinformers.ClusterInterface
}

func (f *sharedInformerFactory) Admissionregistration() admissionregistrationinformers.ClusterInterface {
	return admissionregistrationinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Internal() apiserverinternalinformers.ClusterInterface {
	return apiserverinternalinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Apps() appsinformers.ClusterInterface {
	return appsinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Autoscaling() autoscalinginformers.ClusterInterface {
	return autoscalinginformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Batch() batchinformers.ClusterInterface {
	return batchinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Certificates() certificatesinformers.ClusterInterface {
	return certificatesinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Coordination() coordinationinformers.ClusterInterface {
	return coordinationinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Core() coreinformers.ClusterInterface {
	return coreinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Discovery() discoveryinformers.ClusterInterface {
	return discoveryinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Events() eventsinformers.ClusterInterface {
	return eventsinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Extensions() extensionsinformers.ClusterInterface {
	return extensionsinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Flowcontrol() flowcontrolinformers.ClusterInterface {
	return flowcontrolinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Networking() networkinginformers.ClusterInterface {
	return networkinginformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Node() nodeinformers.ClusterInterface {
	return nodeinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Policy() policyinformers.ClusterInterface {
	return policyinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Rbac() rbacinformers.ClusterInterface {
	return rbacinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Scheduling() schedulinginformers.ClusterInterface {
	return schedulinginformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Storage() storageinformers.ClusterInterface {
	return storageinformers.New(f, f.tweakListOptions)
}
