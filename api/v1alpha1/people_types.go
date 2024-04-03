/*
Copyright 2024.

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

// PeopleSpec defines the desired state of People
type PeopleSpec struct {
	// company specific
	Company      string `json:"company,omitempty"`
	Department   string `json:"department,omitempty"`
	Organization string `json:"organization,omitempty"`
	Job          string `json:"job,omitempty"`

	// person specific
	FirstName     string             `json:"firstName,omitempty"`
	LastName      string             `json:"lastName,omitempty"`
	Street        string             `json:"street,omitempty"`
	HouseNumber   string             `json:"houseNumber,omitempty"`
	City          string             `json:"city,omitempty"`
	PostalAddress string             `json:"postalAddress,omitempty"`
	Skills        []SkillSpecDetails `json:"skills,omitempty"`
}

type PeopleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// People is the Schema for the people API
// +kubebuilder:printcolumn:name="FirstName",type=string,JSONPath=".spec.firstName"
// +kubebuilder:printcolumn:name="LastName",type=string,JSONPath=".spec.lastName"
// +kubebuilder:printcolumn:name="Job",type=string,JSONPath=".spec.job"
type People struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PeopleSpec   `json:"spec,omitempty"`
	Status PeopleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PeopleList contains a list of People
type PeopleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []People `json:"items"`
}

func init() {
	SchemeBuilder.Register(&People{}, &PeopleList{})
}
