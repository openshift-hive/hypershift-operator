package kubeadminpwd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/openshift-hive/hypershift-operator/pkg/cmd/operator"
	"github.com/openshift-hive/hypershift-operator/pkg/controllers"
)

const (
	KubeAdminSecret = "kubeadmin"
)

func Setup(cfg *operator.ControlPlaneOperatorConfig) error {
	informerFactory := informers.NewSharedInformerFactoryWithOptions(cfg.TargetKubeClient(), controllers.DefaultResync, informers.WithNamespace(metav1.NamespaceSystem))
	cfg.Manager().Add(manager.RunnableFunc(func(stopCh <-chan struct{}) error {
		informerFactory.Start(stopCh)
		return nil
	}))
	secrets := informerFactory.Core().V1().Secrets()
	reconciler := &OAuthRestarter{
		Client:       cfg.KubeClient(),
		Namespace:    cfg.Namespace(),
		SecretLister: secrets.Lister(),
		Log:          cfg.Logger().WithName("OAuthRestarter"),
	}
	c, err := controller.New("oauth-restarter", cfg.Manager(), controller.Options{Reconciler: reconciler})
	if err != nil {
		return err
	}
	if err := c.Watch(&source.Informer{Informer: secrets.Informer()}, &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}
	return nil
}
