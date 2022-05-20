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

// A variation of a flag that can be returned to a target
type Variation struct {
	// A description of the variation
	Description string `json:"description,omitempty"`
	// The unique identifier for the variation
	Identifier string `json:"identifier"`
	// The user friendly name of the variation
	Name string `json:"name,omitempty"`
	// The variation value to serve such as true or false for a boolean flag
	Value string `json:"value"`
}
