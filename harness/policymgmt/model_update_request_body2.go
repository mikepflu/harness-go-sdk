/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt

type UpdateRequestBody2 struct {
	// Action that triggers the policy set
	Action string `json:"action,omitempty"`
	// Description of the policy set
	Description string `json:"description,omitempty"`
	// Only enabled policy sets are evaluated when evaluating by type/action
	Enabled bool `json:"enabled,omitempty"`
	// A string enum value which determines which entities the policy set applies to during evaluation. This feature is not available for all accounts, Contact support if you wish to have it enabled.
	EntitySelector string `json:"entity_selector,omitempty"`
	// Name of the policy set
	Name string `json:"name,omitempty"`
	// Policies linked to this policy set
	Policies []Linkedpolicyidentifier `json:"policies,omitempty"`
	// Resource groups that contain the resources that this policy set should be evaluated for. Resource groups are not supported for flag or custom policy sets. This feature is not available for all accounts, Contact support if you wish to have it enabled.
	ResourceGroups []ResourceGroupIdentifier `json:"resource_groups,omitempty"`
	// Type of input suitable for the policy set
	Type_ string `json:"type,omitempty"`
}
