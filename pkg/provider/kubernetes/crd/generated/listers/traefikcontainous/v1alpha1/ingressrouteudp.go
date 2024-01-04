/*
The MIT License (MIT)

Copyright (c) 2016-2020 Containous SAS; 2020-2024 Traefik Labs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikcontainous/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IngressRouteUDPLister helps list IngressRouteUDPs.
// All objects returned here must be treated as read-only.
type IngressRouteUDPLister interface {
	// List lists all IngressRouteUDPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IngressRouteUDP, err error)
	// IngressRouteUDPs returns an object that can list and get IngressRouteUDPs.
	IngressRouteUDPs(namespace string) IngressRouteUDPNamespaceLister
	IngressRouteUDPListerExpansion
}

// ingressRouteUDPLister implements the IngressRouteUDPLister interface.
type ingressRouteUDPLister struct {
	indexer cache.Indexer
}

// NewIngressRouteUDPLister returns a new IngressRouteUDPLister.
func NewIngressRouteUDPLister(indexer cache.Indexer) IngressRouteUDPLister {
	return &ingressRouteUDPLister{indexer: indexer}
}

// List lists all IngressRouteUDPs in the indexer.
func (s *ingressRouteUDPLister) List(selector labels.Selector) (ret []*v1alpha1.IngressRouteUDP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IngressRouteUDP))
	})
	return ret, err
}

// IngressRouteUDPs returns an object that can list and get IngressRouteUDPs.
func (s *ingressRouteUDPLister) IngressRouteUDPs(namespace string) IngressRouteUDPNamespaceLister {
	return ingressRouteUDPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IngressRouteUDPNamespaceLister helps list and get IngressRouteUDPs.
// All objects returned here must be treated as read-only.
type IngressRouteUDPNamespaceLister interface {
	// List lists all IngressRouteUDPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IngressRouteUDP, err error)
	// Get retrieves the IngressRouteUDP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IngressRouteUDP, error)
	IngressRouteUDPNamespaceListerExpansion
}

// ingressRouteUDPNamespaceLister implements the IngressRouteUDPNamespaceLister
// interface.
type ingressRouteUDPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IngressRouteUDPs in the indexer for a given namespace.
func (s ingressRouteUDPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IngressRouteUDP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IngressRouteUDP))
	})
	return ret, err
}

// Get retrieves the IngressRouteUDP from the indexer for a given namespace and name.
func (s ingressRouteUDPNamespaceLister) Get(name string) (*v1alpha1.IngressRouteUDP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ingressrouteudp"), name)
	}
	return obj.(*v1alpha1.IngressRouteUDP), nil
}
