/*
Copyright 2018 Samsung SDS Cloud Native Computing Team.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type JenkinsServer struct {
	metav1.TypeMeta   			`json:",inline"`
	metav1.ObjectMeta 			`json:"metadata,omitempty"`

	Spec   JenkinsServerSpec	`json:"spec"`
	Status JenkinsServerStatus	`json:"status"`
}

// +k8s:openapi-gen=true
type JenkinsServerSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// +k8s:openapi-gen=true
type JenkinsServerStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type JenkinsServerList struct {
	metav1.TypeMeta 		`json:",inline"`
	metav1.ListMeta 		`json:"metadata"`

	Items []JenkinsServer	`json:"items"`
}