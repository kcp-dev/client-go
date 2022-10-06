/*
Copyright 2022 The KCP Authors.

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
	"fmt"
	"sync"

	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/apimachinery/pkg/runtime"
)

// Constructor is a wrapper around a constructor method for the fake client of type R.
type Constructor[R any] struct {
	NewSimpleClientset func(objects ...runtime.Object) R
}

// Cache is a fake client factory that caches previous results.
type Cache[R any] interface {
	Cluster(name logicalcluster.Name, objects ...runtime.Object) R
}

// NewCache creates a new fake client factory cache using the given constructor.
func NewCache[R any](constructor Constructor[R]) Cache[R] {
	return &cache[R]{constructor: constructor}
}

type cache[R any] struct {
	constructor Constructor[R]

	*sync.RWMutex
	fakesByCluster map[logicalcluster.Name]R
}

// Cluster initializes a new fake client, or returns a previously initialized fake client scoped to the given logical cluster.
// Note: attempting to initialize the client more than once will cause a panic.
func (c *cache[R]) Cluster(name logicalcluster.Name, objects ...runtime.Object) R {
	c.Lock()
	defer c.Unlock()
	if cachedFake, exists := c.fakesByCluster[name]; exists {
		if len(objects) > 0 {
			panic(fmt.Sprintf("programmer error: attempting to initialize fake client for cluster %v more than once!", name))
		}
		return cachedFake
	}

	instance := c.constructor.NewSimpleClientset(objects...)
	c.fakesByCluster[name] = instance
	return instance
}
