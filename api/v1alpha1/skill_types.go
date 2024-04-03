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

// SkillSpec defines the desired state of Skill

type SkillSpecDetails struct {
	Name                   string `json:"name"`
	Experience             string `json:"experience,omitempty"`
	Description            string `json:"description,omitempty"`
	Fulfilled              bool   `json:"fulfilled,omitempty"`
	FulfilledByAlternative bool   `json:"fulfilledByAlternative,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Alternatives []SkillSpecDetails `json:"alternatives,omitempty"`
}

type SkillSpec struct {
	Name         string             `json:"name"`
	Experience   string             `json:"experience,omitempty"`
	Description  string             `json:"description,omitempty"`
	Alternatives []SkillSpecDetails `json:"alternatives,omitempty"`
}

// SkillStatus defines the observed state of Skill
type SkillStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Skill is the Schema for the skills API
type Skill struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SkillSpec   `json:"spec,omitempty"`
	Status SkillStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SkillList contains a list of Skill
type SkillList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Skill `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Skill{}, &SkillList{})
}
