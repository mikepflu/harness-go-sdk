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

type ApplicationsApplication struct {
	Metadata  *V1ObjectMeta                  `json:"metadata,omitempty"`
	Spec      *ApplicationsApplicationSpec   `json:"spec,omitempty"`
	Status    *ApplicationsApplicationStatus `json:"status,omitempty"`
	Operation *ApplicationsOperation         `json:"operation,omitempty"`
}
