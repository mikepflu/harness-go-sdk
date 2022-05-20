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

// This is the Kerberos configuration details, defined in Harness.
type KerberosConfig struct {
	Type_ string `json:"type"`
	// This is the authorization role, the user/service has in the realm.
	Principal string `json:"principal"`
	// Name of the Realm.
	Realm               string                `json:"realm"`
	TgtGenerationMethod string                `json:"tgtGenerationMethod,omitempty"`
	Spec                *TgtGenerationSpecDto `json:"spec,omitempty"`
}
