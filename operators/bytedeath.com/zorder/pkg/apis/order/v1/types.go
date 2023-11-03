package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OrderSpec defines the desired state of Order
type OrderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS -- desired state of cluster
	Name string `json:"name"`
	Fee  int32  `json:"fee"`
}

// OrderStatus defines the observed state of Order.
// It should always be reconstructable from the state of the cluster and/or outside world.
type OrderStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of cluster
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Order is the Schema for the orders API
// ...
type Order struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrderSpec   `json:"spec,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrderList contains a list of Order
type OrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Order `json:"items"`
}
