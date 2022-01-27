/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This contains details of the Generic Git authentication information used via HTTP connections
type GitHttpAuthenticationDto struct {
	Type_       string `json:"-"`
	Username    string `json:"username,omitempty"`
	UsernameRef string `json:"usernameRef,omitempty"`
	PasswordRef string `json:"passwordRef"`
}
