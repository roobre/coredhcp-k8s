package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LeaseData struct {
	HardwareAddress string `json:"hardwareAddress"`
	Address         string `json:"address"`
	Hostname        string `json:"hostname,omitempty"`
}

//+kubebuilder:object:root=true

// Lease is the Schema for the leases API
type Lease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ServerName string `json:"serverName"`

	LeaseData `json:",inline"`
	Request   LeaseData       `json:"request,omitempty"`
	Duration  metav1.Duration `json:"duration"`
	Renewed   metav1.Time     `json:"renewed"`
}

//+kubebuilder:object:root=true

// LeaseList contains a list of Lease
type LeaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Lease{}, &LeaseList{})
}
