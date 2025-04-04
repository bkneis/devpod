// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apisstoragev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	versioned "github.com/loft-sh/api/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/api/v4/pkg/informers/externalversions/internalinterfaces"
	storagev1 "github.com/loft-sh/api/v4/pkg/listers/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SharedSecretInformer provides access to a shared informer and lister for
// SharedSecrets.
type SharedSecretInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() storagev1.SharedSecretLister
}

type sharedSecretInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSharedSecretInformer constructs a new informer for SharedSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSharedSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSharedSecretInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSharedSecretInformer constructs a new informer for SharedSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSharedSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().SharedSecrets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().SharedSecrets(namespace).Watch(context.TODO(), options)
			},
		},
		&apisstoragev1.SharedSecret{},
		resyncPeriod,
		indexers,
	)
}

func (f *sharedSecretInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSharedSecretInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sharedSecretInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisstoragev1.SharedSecret{}, f.defaultInformer)
}

func (f *sharedSecretInformer) Lister() storagev1.SharedSecretLister {
	return storagev1.NewSharedSecretLister(f.Informer().GetIndexer())
}
