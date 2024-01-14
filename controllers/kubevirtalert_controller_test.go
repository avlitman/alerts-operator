package controllers

import (
	"context"
	"github.com/kubevirt/alerts-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("Create alert rule", func() {
	It("Should create alert rule", func() {
		alert := &v1alpha1.KubevirtAlert{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "alert1",
				Namespace: "alert1ns",
			},
			Spec: v1alpha1.KubevirtAlertSpec{
				Metric: v1alpha1.MetricSpec{
					Name: "test_metrics",
					Type: "Counter",
					Help: "test metric",
				},
				RecordRule: v1alpha1.RecordRuleSpec{
					Record: "record_test",
					Expr:   "rule_test > 0",
				},
			},
		}
		sch := runtime.NewScheme()
		Expect(v1alpha1.AddToScheme(sch)).To(Succeed())
		Expect(promv1.AddToScheme(sch)).To(Succeed())
		cli := fake.NewClientBuilder().WithObjects(alert).WithScheme(sch).Build()
		r := &KubevirtAlertReconciler{
			Client: cli,
			Scheme: sch,
		}
		result, err := r.Reconcile(context.TODO(), ctrl.Request{
			NamespacedName: types.NamespacedName{
				Name:      "alert1",
				Namespace: "alert1ns",
			},
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(result).ToNot(BeNil())

		palert := &promv1.PrometheusRule{}
		err = cli.Get(context.TODO(), client.ObjectKey{
			Name:      "test_metrics_alert",
			Namespace: "alert1ns",
		}, palert)
		Expect(err).ToNot(HaveOccurred())
		Expect(palert.Spec.Groups).To(HaveLen(1))
		Expect(palert.Spec.Groups[0].Rules).To(HaveLen(1))
		Expect(palert.Spec.Groups[0].Rules[0].Labels).To(HaveKeyWithValue("severity", "warning"))
	})
})
