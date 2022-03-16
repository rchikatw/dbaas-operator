/*
Copyright 2021.

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

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/RHEcosystemAppEng/dbaas-operator/api/v1alpha1"
)

// DBaaSConnectionReconciler reconciles a DBaaSConnection object
type DBaaSMetricReconciler struct {
	*DBaaSReconciler
}

//+kubebuilder:rbac:groups=dbaas.redhat.com,resources=*,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dbaas.redhat.com,resources=*/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dbaas.redhat.com,resources=*/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *DBaaSMetricReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrl.LoggerFrom(ctx)

	logger.Info("First mission is successful")

	return ctrl.Result{}, nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *DBaaSMetricReconciler) SetupWithManager(mgr ctrl.Manager) (controller.Controller, error) {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.DBaaSConnection{}).
		WithOptions(
			controller.Options{MaxConcurrentReconciles: 2},
		).
		Build(r)
}
