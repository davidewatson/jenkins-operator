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
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: "jenkins.jenkinsoperator.maratoid.github.com", Version: "v1alpha1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&JenkinsInstance{},
		&JenkinsInstanceList{},
		&JenkinsPlugin{},
		&JenkinsPluginList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type JenkinsInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JenkinsInstance `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type JenkinsPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JenkinsPlugin `json:"items"`
}

// CRD Generation
func getFloat(f float64) *float64 {
	return &f
}

func getInt(i int64) *int64 {
	return &i
}

var (
	// Define CRDs for resources
	JenkinsInstanceCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "jenkinsinstances.jenkins.jenkinsoperator.maratoid.github.com",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "jenkins.jenkinsoperator.maratoid.github.com",
			Version: "v1alpha1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "JenkinsInstance",
				Plural: "jenkinsinstances",
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"adminemail": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"adminsecret": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"agentport": v1beta1.JSONSchemaProps{
									Type:   "integer",
									Format: "int32",
								},
								"config": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "string",
										},
									},
								},
								"env": v1beta1.JSONSchemaProps{
									Type: "object",
								},
								"executors": v1beta1.JSONSchemaProps{
									Type:   "integer",
									Format: "int32",
								},
								"image": v1beta1.JSONSchemaProps{
									Pattern: ".+:.+",
									Type:    "string",
								},
								"location": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"masterport": v1beta1.JSONSchemaProps{
									Type:   "integer",
									Format: "int32",
								},
								"name": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"pullpolicy": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"replicas": v1beta1.JSONSchemaProps{
									Type:   "integer",
									Format: "int32",
								},
								"servicetype": v1beta1.JSONSchemaProps{
									Type: "string",
								},
							},
						},
						"status": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"adminsecret": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"api": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"phase": v1beta1.JSONSchemaProps{
									Type: "string",
								},
							},
							Required: []string{
								"phase",
							}},
					},
				},
			},
		},
	}
	// Define CRDs for resources
	JenkinsPluginCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "jenkinsplugins.jenkins.jenkinsoperator.maratoid.github.com",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "jenkins.jenkinsoperator.maratoid.github.com",
			Version: "v1alpha1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "JenkinsPlugin",
				Plural: "jenkinsplugins",
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"config": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "string",
										},
									},
								},
								"jenkinsinstance": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"name": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"pluginid": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"pluginversion": v1beta1.JSONSchemaProps{
									Type: "string",
								},
							},
						},
						"status": v1beta1.JSONSchemaProps{
							Type:       "object",
							Properties: map[string]v1beta1.JSONSchemaProps{},
						},
					},
				},
			},
		},
	}
)
