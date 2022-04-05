// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	configv1alpha1 "github.com/openshift/cert-manager-operator/apis/config/v1alpha1"
	versioned "github.com/openshift/cert-manager-operator/pkg/config/clientset/versioned"
	internalinterfaces "github.com/openshift/cert-manager-operator/pkg/config/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/cert-manager-operator/pkg/config/listers/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CertManagerInformer provides access to a shared informer and lister for
// CertManagers.
type CertManagerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CertManagerLister
}

type certManagerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCertManagerInformer constructs a new informer for CertManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertManagerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertManagerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCertManagerInformer constructs a new informer for CertManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertManagerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().CertManagers().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().CertManagers().Watch(context.TODO(), options)
			},
		},
		&configv1alpha1.CertManager{},
		resyncPeriod,
		indexers,
	)
}

func (f *certManagerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertManagerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *certManagerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha1.CertManager{}, f.defaultInformer)
}

func (f *certManagerInformer) Lister() v1alpha1.CertManagerLister {
	return v1alpha1.NewCertManagerLister(f.Informer().GetIndexer())
}
