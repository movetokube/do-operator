// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/do/v1alpha1.DNS":       schema_pkg_apis_do_v1alpha1_DNS(ref),
		"./pkg/apis/do/v1alpha1.DNSSpec":   schema_pkg_apis_do_v1alpha1_DNSSpec(ref),
		"./pkg/apis/do/v1alpha1.DNSStatus": schema_pkg_apis_do_v1alpha1_DNSStatus(ref),
	}
}

func schema_pkg_apis_do_v1alpha1_DNS(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DNS is the Schema for the dns API",
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
							Ref: ref("./pkg/apis/do/v1alpha1.DNSSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/do/v1alpha1.DNSStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/do/v1alpha1.DNSSpec", "./pkg/apis/do/v1alpha1.DNSStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_do_v1alpha1_DNSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DNSSpec defines the desired state of DNS",
				Properties: map[string]spec.Schema{
					"domainName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"recordType": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"hostname": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/do/v1alpha1.DNSRecordValue"),
						},
					},
					"ttl": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"priority": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"flag": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"weight": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"tag": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"domainName", "recordType", "hostname", "value"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/do/v1alpha1.DNSRecordValue"},
	}
}

func schema_pkg_apis_do_v1alpha1_DNSStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DNSStatus defines the observed state of DNS",
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"id": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"state", "id"},
			},
		},
		Dependencies: []string{},
	}
}
