/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MetricsSourceSpec defines the desired state of MetricsSource
type MetricsSourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	MetricsName string                    `json:"metricsName"`
	Labels      map[string]string         `json:"labels,omitempty"`
	Metrics     []MetricsSourceSpecMetric `json:"metrics,omitempty"`
}

type MetricsSourceSpecMetric struct {
	Start    string          `json:"start"`
	Duration metav1.Duration `json:"duration"`
	Value    int             `json:"value"`
}

// MetricsSourceStatus defines the observed state of MetricsSource
type MetricsSourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	CurrentValue int `json:"currentValue,omitempty"`

	// +optional
	LastSchedule metav1.Time `json:"lastSchedule,omitempty"`

	// +optional
	NextSchedule metav1.Time `json:"nextSchedule,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MetricsSource is the Schema for the metricssources API
type MetricsSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetricsSourceSpec   `json:"spec,omitempty"`
	Status MetricsSourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MetricsSourceList contains a list of MetricsSource
type MetricsSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricsSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetricsSource{}, &MetricsSourceList{})
}
