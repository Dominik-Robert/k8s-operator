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
type MatchingPeople struct {
	PersonName                    string          `json:"personName,omitempty"`
	Skills                        []SkillResolved `json:"skills,omitempty"`
	PercentageStrict              string          `json:"percentageStrict,omitempty"`
	PercentageWithoutExperience   string          `json:"percentageWithoutExperience,omitempty"`
	PercentageWithoutAlternatives string          `json:"percentageWithoutAlternatives,omitempty"`
}

type SkillResolved struct {
	SkillSpecs                       SkillSpecDetails `json:"skillSpecs,omitempty"`
	SkillIsAlternative               bool             `json:"skillIsAlternative,omitempty"`
	SkillAvailable                   bool             `json:"skillAvailable,omitempty"`
	SkillAvailableByAlternative      bool             `json:"skillAvailableByAlternative,omitempty"`
	ExperienceFulfilled              bool             `json:"experienceFulfilled,omitempty"`
	ExperienceFulfilledByAlternative bool             `json:"fulfilledByAlternative,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	// company specific
	Company      string `json:"company,omitempty"`
	Department   string `json:"department,omitempty"`
	Organization string `json:"organization,omitempty"`
	Job          string `json:"job,omitempty"`
	Description  string `json:"description,omitempty"`

	// Buzzwords
	OpenPositions  int                `json:"openPositions,omitempty"`
	NeededSkills   []SkillSpecDetails `json:"neededSkills,omitempty"`
	OptionalSkills []SkillSpecDetails `json:"optionalSkills,omitempty"`
}

// ProjectStatus defines the observed state of Project
type ProjectStatus struct {
	PossiblePeople []MatchingPeople `json:"possiblePeople"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Project is the Schema for the projects API
// +kubebuilder:printcolumn:name="LastName",type=string,JSONPath=".spec.company"
// +kubebuilder:printcolumn:name="Job",type=string,JSONPath=".spec.job"
// +kubebuilder:printcolumn:name="Job",type=string,JSONPath=".spec.openPositions"
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProjectList contains a list of Project
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
