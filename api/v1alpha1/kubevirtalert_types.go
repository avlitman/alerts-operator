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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubevirtAlertSpec defines the desired state of KubevirtAlert
type KubevirtAlertSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Metric     MetricSpec     `json:"metric"`
	RecordRule RecordRuleSpec `json:"recordRule"`
}

// MetricSpec defines the attributes for a Prometheus metric
type MetricSpec struct {
	Name string `json:"name"`
	Type string `json:"type"` // e.g., Gauge, Counter, etc.
	Help string `json:"help"`
}

// RecordRuleSpec defines the attributes for a Prometheus recording rule
type RecordRuleSpec struct {
	Record string            `json:"record"`
	Expr   string            `json:"expr"` // PromQL expression
	Labels map[string]string `json:"labels,omitempty"`
}

// KubevirtAlertStatus defines the observed state of KubevirtAlert
type KubevirtAlertStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KubevirtAlert is the Schema for the kubevirtalerts API
type KubevirtAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubevirtAlertSpec   `json:"spec,omitempty"`
	Status KubevirtAlertStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KubevirtAlertList contains a list of KubevirtAlert
type KubevirtAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubevirtAlert `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubevirtAlert{}, &KubevirtAlertList{})
}
