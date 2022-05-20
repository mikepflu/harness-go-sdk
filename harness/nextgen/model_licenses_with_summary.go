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

// This contains details of the License With Summary defined in Harness
type LicensesWithSummary struct {
	Edition       string `json:"edition,omitempty"`
	LicenseType   string `json:"licenseType,omitempty"`
	ModuleType    string `json:"moduleType,omitempty"`
	MaxExpiryTime int64  `json:"maxExpiryTime,omitempty"`
}
