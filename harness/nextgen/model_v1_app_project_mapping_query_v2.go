/*
 * Harness NextGen GitOps API Reference
 *
 * This the Open Api Spec 3 for the GitoOps service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type V1AppProjectMappingQueryV2 struct {
	// app project mapping identifier.
	Identifier string `json:"identifier,omitempty"`
	// Agent identifier for entity.
	AgentIdentifier string `json:"agentIdentifier,omitempty"`
	// Account Identifier for the Entity.
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier     string `json:"projectIdentifier,omitempty"`
	ArgoProjectIdentifier string `json:"argoProjectIdentifier,omitempty"`
}
