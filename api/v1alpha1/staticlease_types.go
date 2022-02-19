package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StaticLeaseSpec defines the desired state of StaticLease
type StaticLeaseSpec struct {
	ServerClassName string `json:"serverClassName"`
	LeaseData       `json:",inline"`
}

//+kubebuilder:object:root=true

// StaticLease is the Schema for the staticleases API
type StaticLease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec StaticLeaseSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// StaticLeaseList contains a list of StaticLease
type StaticLeaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticLease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StaticLease{}, &StaticLeaseList{})
}
