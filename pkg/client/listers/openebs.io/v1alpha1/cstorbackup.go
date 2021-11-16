/*
Copyright 2021 The OpenEBS Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/openebs/api/v3/pkg/apis/openebs.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CStorBackupLister helps list CStorBackups.
// All objects returned here must be treated as read-only.
type CStorBackupLister interface {
	// List lists all CStorBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CStorBackup, err error)
	// CStorBackups returns an object that can list and get CStorBackups.
	CStorBackups(namespace string) CStorBackupNamespaceLister
	CStorBackupListerExpansion
}

// cStorBackupLister implements the CStorBackupLister interface.
type cStorBackupLister struct {
	indexer cache.Indexer
}

// NewCStorBackupLister returns a new CStorBackupLister.
func NewCStorBackupLister(indexer cache.Indexer) CStorBackupLister {
	return &cStorBackupLister{indexer: indexer}
}

// List lists all CStorBackups in the indexer.
func (s *cStorBackupLister) List(selector labels.Selector) (ret []*v1alpha1.CStorBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CStorBackup))
	})
	return ret, err
}

// CStorBackups returns an object that can list and get CStorBackups.
func (s *cStorBackupLister) CStorBackups(namespace string) CStorBackupNamespaceLister {
	return cStorBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CStorBackupNamespaceLister helps list and get CStorBackups.
// All objects returned here must be treated as read-only.
type CStorBackupNamespaceLister interface {
	// List lists all CStorBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CStorBackup, err error)
	// Get retrieves the CStorBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CStorBackup, error)
	CStorBackupNamespaceListerExpansion
}

// cStorBackupNamespaceLister implements the CStorBackupNamespaceLister
// interface.
type cStorBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CStorBackups in the indexer for a given namespace.
func (s cStorBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CStorBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CStorBackup))
	})
	return ret, err
}

// Get retrieves the CStorBackup from the indexer for a given namespace and name.
func (s cStorBackupNamespaceLister) Get(name string) (*v1alpha1.CStorBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cstorbackup"), name)
	}
	return obj.(*v1alpha1.CStorBackup), nil
}
