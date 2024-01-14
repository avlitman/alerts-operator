/*
Copyright 2023.

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
	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	alertsv1alpha1 "github.com/kubevirt/alerts-operator/api/v1alpha1"
)

// KubevirtAlertReconciler reconciles a KubevirtAlert object
type KubevirtAlertReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=alerts.kubevirt.io,resources=kubevirtalerts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=alerts.kubevirt.io,resources=kubevirtalerts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=alerts.kubevirt.io,resources=kubevirtalerts/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KubevirtAlert object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *KubevirtAlertReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	var alert alertsv1alpha1.KubevirtAlert
	if err := r.Get(ctx, req.NamespacedName, &alert); err != nil {
		logger.Error(err, "unable to fetch KubevirtAlert")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	alertExists, err := r.checkIfAlertExists(ctx, &alert)
	if err != nil {
		logger.Error(err, "error checking if alert exists in Prometheus")
		return ctrl.Result{}, err // not sure if client.IgnoreNotFound needed here
	}

	if !alertExists {
		// Create the alert
		err := r.createPrometheusAlert(ctx, &alert)
		if err != nil {
			logger.Error(err, "error creating Prometheus alert")
			return ctrl.Result{}, err // not sure if client.IgnoreNotFound needed here
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KubevirtAlertReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alertsv1alpha1.KubevirtAlert{}).
		Complete(r)
}

func (r *KubevirtAlertReconciler) checkIfAlertExists(ctx context.Context, alert *alertsv1alpha1.KubevirtAlert) (bool, error) {
	ruleName := alert.Spec.Metric.Name + "_alert"
	var promRule promv1.PrometheusRule

	// try to get the alert from kubernetes API
	err := r.Client.Get(ctx, client.ObjectKey{
		Name:      ruleName,
		Namespace: alert.Namespace,
	}, &promRule)

	if err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *KubevirtAlertReconciler) createPrometheusAlert(ctx context.Context, alert *alertsv1alpha1.KubevirtAlert) error {
	// Constructing the rule name from the KubevirtAlert resource
	ruleName := alert.Spec.Metric.Name + "_alert"

	// Creating a PrometheusRule object
	promRule := &promv1.PrometheusRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ruleName,
			Namespace: alert.Namespace,
		},
		Spec: promv1.PrometheusRuleSpec{
			Groups: []promv1.RuleGroup{
				{
					Name: ruleName + "_group",
					Rules: []promv1.Rule{
						{
							Alert: ruleName,
							Expr:  intstr.FromString(alert.Spec.RecordRule.Expr),
							For:   ptr.To(promv1.Duration("5m")),
							Labels: map[string]string{
								"severity": "warning",
							},
							Annotations: map[string]string{
								"summary":     "High number of pods outdated",
								"description": "This alert fires when pods outdated rate is more then 90%\"",
							},
						},
					},
				},
			},
		},
	}

	// Create the alert rule in kubernetes cluster, which prometheus operator will then process
	if err := r.Create(ctx, promRule); err != nil { //not sure what to put here
		return err // not sure if client.IgnoreNotFound needed here
	}

	return nil
}
