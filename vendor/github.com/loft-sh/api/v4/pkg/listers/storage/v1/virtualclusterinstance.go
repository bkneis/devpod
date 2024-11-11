// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualClusterInstanceLister helps list VirtualClusterInstances.
// All objects returned here must be treated as read-only.
type VirtualClusterInstanceLister interface {
	// List lists all VirtualClusterInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualClusterInstance, err error)
	// VirtualClusterInstances returns an object that can list and get VirtualClusterInstances.
	VirtualClusterInstances(namespace string) VirtualClusterInstanceNamespaceLister
	VirtualClusterInstanceListerExpansion
}

// virtualClusterInstanceLister implements the VirtualClusterInstanceLister interface.
type virtualClusterInstanceLister struct {
	indexer cache.Indexer
}

// NewVirtualClusterInstanceLister returns a new VirtualClusterInstanceLister.
func NewVirtualClusterInstanceLister(indexer cache.Indexer) VirtualClusterInstanceLister {
	return &virtualClusterInstanceLister{indexer: indexer}
}

// List lists all VirtualClusterInstances in the indexer.
func (s *virtualClusterInstanceLister) List(selector labels.Selector) (ret []*v1.VirtualClusterInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualClusterInstance))
	})
	return ret, err
}

// VirtualClusterInstances returns an object that can list and get VirtualClusterInstances.
func (s *virtualClusterInstanceLister) VirtualClusterInstances(namespace string) VirtualClusterInstanceNamespaceLister {
	return virtualClusterInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualClusterInstanceNamespaceLister helps list and get VirtualClusterInstances.
// All objects returned here must be treated as read-only.
type VirtualClusterInstanceNamespaceLister interface {
	// List lists all VirtualClusterInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualClusterInstance, err error)
	// Get retrieves the VirtualClusterInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.VirtualClusterInstance, error)
	VirtualClusterInstanceNamespaceListerExpansion
}

// virtualClusterInstanceNamespaceLister implements the VirtualClusterInstanceNamespaceLister
// interface.
type virtualClusterInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualClusterInstances in the indexer for a given namespace.
func (s virtualClusterInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1.VirtualClusterInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualClusterInstance))
	})
	return ret, err
}

// Get retrieves the VirtualClusterInstance from the indexer for a given namespace and name.
func (s virtualClusterInstanceNamespaceLister) Get(name string) (*v1.VirtualClusterInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("virtualclusterinstance"), name)
	}
	return obj.(*v1.VirtualClusterInstance), nil
}