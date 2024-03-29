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

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	// Name is the name of the database to create
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Owner is the name of the user that will own the database
	// +optional
	Owner string `json:"owner,omitempty"`
	// Charset is the character set for the database
	// +optional
	Charset string `json:"charset,omitempty"`
	// Collation is the collation for the database
	// +optional
	Collation string `json:"collation,omitempty"`

	// HostRef is a reference to a DatabaseHost object in the same namespace
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	DatabaseHostRef string `json:"databaseHostRef"`
}

// DatabaseStatus defines the observed state of Database
type DatabaseStatus struct {
	CreationTime   metav1.Time `json:"creationTime,omitempty"`
	CreationStatus string      `json:"creationStatus,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Database is the Schema for the databases API
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseSpec   `json:"spec,omitempty"`
	Status DatabaseStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DatabaseList contains a list of Database
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
