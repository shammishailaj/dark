// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/K-Phoen/dark/internal/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GrafanaDashboards returns a GrafanaDashboardInformer.
	GrafanaDashboards() GrafanaDashboardInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// GrafanaDashboards returns a GrafanaDashboardInformer.
func (v *version) GrafanaDashboards() GrafanaDashboardInformer {
	return &grafanaDashboardInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}