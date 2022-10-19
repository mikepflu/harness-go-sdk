/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type V1RepositoryQuery struct {
	// Account Identifier for the Entity.
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Agent identifier for entity.
	AgentIdentifier string `json:"agentIdentifier,omitempty"`
	Identifier      string `json:"identifier,omitempty"`
	SearchTerm      string `json:"searchTerm,omitempty"`
	PageSize        int32  `json:"pageSize,omitempty"`
	PageIndex       int32  `json:"pageIndex,omitempty"`
	// Filters for Repositories. Eg. \"identifier\": { \"$in\": [\"id1\", \"id2\"]
	Filter      *interface{} `json:"filter,omitempty"`
	RepoCredsId string       `json:"repoCredsId,omitempty"`
}
