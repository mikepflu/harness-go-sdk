/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type GitPrCreateRequest struct {
	// Source Branch for pull request
	SourceBranch string `json:"sourceBranch"`
	// Target Branch for pull request
	TargetBranch string `json:"targetBranch"`
	// PR title
	Title string `json:"title"`
	// Git Sync Config Id
	YamlGitConfigRef string `json:"yamlGitConfigRef"`
	// Account Identifier for the Entity.
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Specifies which token to use. If True, the SCM token will be used, else the Git Connector token will be used
	UseUserFromToken bool `json:"useUserFromToken,omitempty"`
}
