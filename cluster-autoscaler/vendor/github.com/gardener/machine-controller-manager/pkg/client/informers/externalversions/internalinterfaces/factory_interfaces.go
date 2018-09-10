// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package internalinterfaces

import (
	time "time"

	versioned "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned"
	v1 "k8s.io2/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io2/apimachinery/pkg/runtime"
	cache "k8s.io2/client-go/tools/cache"
)

type NewInformerFunc func(versioned.Interface, time.Duration) cache.SharedIndexInformer

// SharedInformerFactory a small interface to allow for adding an informer without an import cycle
type SharedInformerFactory interface {
	Start(stopCh <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer
}

type TweakListOptionsFunc func(*v1.ListOptions)
