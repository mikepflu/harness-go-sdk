/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type OpenapiMergePullReq struct {
	BypassRules bool             `json:"bypass_rules,omitempty"`
	DryRun      bool             `json:"dry_run,omitempty"`
	Method      *EnumMergeMethod `json:"method,omitempty"`
	SourceSha   string           `json:"source_sha,omitempty"`
}
