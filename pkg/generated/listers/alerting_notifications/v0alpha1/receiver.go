// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by lister-gen. DO NOT EDIT.

package v0alpha1

import (
	v0alpha1 "github.com/grafana/grafana/pkg/apis/alerting_notifications/v0alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ReceiverLister helps list Receivers.
// All objects returned here must be treated as read-only.
type ReceiverLister interface {
	// List lists all Receivers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v0alpha1.Receiver, err error)
	// Receivers returns an object that can list and get Receivers.
	Receivers(namespace string) ReceiverNamespaceLister
	ReceiverListerExpansion
}

// receiverLister implements the ReceiverLister interface.
type receiverLister struct {
	listers.ResourceIndexer[*v0alpha1.Receiver]
}

// NewReceiverLister returns a new ReceiverLister.
func NewReceiverLister(indexer cache.Indexer) ReceiverLister {
	return &receiverLister{listers.New[*v0alpha1.Receiver](indexer, v0alpha1.Resource("receiver"))}
}

// Receivers returns an object that can list and get Receivers.
func (s *receiverLister) Receivers(namespace string) ReceiverNamespaceLister {
	return receiverNamespaceLister{listers.NewNamespaced[*v0alpha1.Receiver](s.ResourceIndexer, namespace)}
}

// ReceiverNamespaceLister helps list and get Receivers.
// All objects returned here must be treated as read-only.
type ReceiverNamespaceLister interface {
	// List lists all Receivers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v0alpha1.Receiver, err error)
	// Get retrieves the Receiver from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v0alpha1.Receiver, error)
	ReceiverNamespaceListerExpansion
}

// receiverNamespaceLister implements the ReceiverNamespaceLister
// interface.
type receiverNamespaceLister struct {
	listers.ResourceIndexer[*v0alpha1.Receiver]
}
