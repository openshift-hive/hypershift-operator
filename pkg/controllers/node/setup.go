package node

import (
	"k8s.io/client-go/informers"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/openshift-hive/hypershift-operator/pkg/cmd/operator"
	"github.com/openshift-hive/hypershift-operator/pkg/controllers"
)

func Setup(cfg *operator.ControlPlaneOperatorConfig) error {
	informerFactory := informers.NewSharedInformerFactory(cfg.TargetKubeClient(), controllers.DefaultResync)
	cfg.Manager().Add(manager.RunnableFunc(func(stopCh <-chan struct{}) error {
		informerFactory.Start(stopCh)
		return nil
	}))
	nodes := informerFactory.Core().V1().Nodes()
	reconciler := &NodeReconciler{
		Lister:     nodes.Lister(),
		KubeClient: cfg.TargetKubeClient(),
		Log:        cfg.Logger().WithName("Node"),
	}
	c, err := controller.New("node", cfg.Manager(), controller.Options{Reconciler: reconciler})
	if err != nil {
		return err
	}
	if err := c.Watch(&source.Informer{Informer: nodes.Informer()}, &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}
	return nil
}
