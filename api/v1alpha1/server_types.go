package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AddressRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
	CIDR  int    `json:"cidr"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	ServerClassName string `json:"serverClassName"`

	Interfaces []string `json:"interfaces"`

	Range AddressRange `json:"range"`

	LeaseTime metav1.Duration `json:"leaseTime"`

	Router string   `json:"router"`
	DNS    []string `json:"dns"`
}

// ServerStatus defines the observed state of Server
//type ServerStatus struct {
//	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
//	// Important: Run "make" to regenerate code after modifying this file
//}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Server is the Schema for the servers API
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServerSpec `json:"spec,omitempty"`
	//Status ServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServerList contains a list of Server
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
