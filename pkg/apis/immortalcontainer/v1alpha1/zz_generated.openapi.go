// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/flugel-it/k8s-go-operator-sdk-operator/pkg/apis/immortalcontainer/v1alpha1.ImmortalContainer":       schema_pkg_apis_immortalcontainer_v1alpha1_ImmortalContainer(ref),
		"github.com/flugel-it/k8s-go-operator-sdk-operator/pkg/apis/immortalcontainer/v1alpha1.ImmortalContainerSpec":   schema_pkg_apis_immortalcontainer_v1alpha1_ImmortalContainerSpec(ref),
		"github.com/flugel-it/k8s-go-operator-sdk-operator/pkg/apis/immortalcontainer/v1alpha1.ImmortalContainerStatus": schema_pkg_apis_immortalcontainer_v1alpha1_ImmortalContainerStatus(ref),
	}
}

func schema_pkg_apis_immortalcontainer_v1alpha1_ImmortalContainer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ImmortalContainer is the Schema for the immortalcontainers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/flugel-it/k8s-go-operator-sdk-operator/pkg/apis/immortalcontainer/v1alpha1.ImmortalContainerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/flugel-it/k8s-go-operator-sdk-operator/pkg/apis/immortalcontainer/v1alpha1.ImmortalContainerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/flugel-it/k8s-go-operator-sdk-operator/pkg/apis/immortalcontainer/v1alpha1.ImmortalContainerSpec", "github.com/flugel-it/k8s-go-operator-sdk-operator/pkg/apis/immortalcontainer/v1alpha1.ImmortalContainerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_immortalcontainer_v1alpha1_ImmortalContainerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ImmortalContainerSpec defines the desired state of ImmortalContainer",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"image"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_immortalcontainer_v1alpha1_ImmortalContainerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ImmortalContainerStatus defines the observed state of ImmortalContainer",
				Properties: map[string]spec.Schema{
					"currentPod": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startTimes": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
