/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta2

import (
	flowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// PriorityLevelConfigurationLister helps list PriorityLevelConfigurations.
// All objects returned here must be treated as read-only.
type PriorityLevelConfigurationLister interface {
	// List lists all PriorityLevelConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*flowcontrolv1beta2.PriorityLevelConfiguration, err error)
	// Get retrieves the PriorityLevelConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*flowcontrolv1beta2.PriorityLevelConfiguration, error)
	PriorityLevelConfigurationListerExpansion
}

// priorityLevelConfigurationLister implements the PriorityLevelConfigurationLister interface.
type priorityLevelConfigurationLister struct {
	listers.ResourceIndexer[*flowcontrolv1beta2.PriorityLevelConfiguration]
}

// NewPriorityLevelConfigurationLister returns a new PriorityLevelConfigurationLister.
func NewPriorityLevelConfigurationLister(indexer cache.Indexer) PriorityLevelConfigurationLister {
	return &priorityLevelConfigurationLister{listers.New[*flowcontrolv1beta2.PriorityLevelConfiguration](indexer, flowcontrolv1beta2.Resource("prioritylevelconfiguration"))}
}
