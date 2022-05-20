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

// This contains details of connectivity mode and whether Git Sync is enabled
type GitEnabled struct {
	// This checks if Git Sync is enabled for a given scope
	IsGitSyncEnabled bool `json:"isGitSyncEnabled,omitempty"`
	// This is the Git Sync connectivity mode
	ConnectivityMode string `json:"connectivityMode,omitempty"`
	GitSyncEnabled   bool   `json:"gitSyncEnabled,omitempty"`
}
