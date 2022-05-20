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

type AsgMinimal struct {
	Id                string               `json:"id,omitempty"`
	Name              string               `json:"name,omitempty"`
	Desired           int32                `json:"desired,omitempty"`
	Min               int32                `json:"min,omitempty"`
	Max               int32                `json:"max,omitempty"`
	OnDemand          int32                `json:"on_demand,omitempty"`
	Spot              int32                `json:"spot,omitempty"`
	MixedInstance     bool                 `json:"mixed_instance,omitempty"`
	CloudAccountId    string               `json:"cloud_account_id,omitempty"`
	ProviderName      string               `json:"provider_name,omitempty"`
	TargetGroups      []TargetGroupMinimal `json:"target_groups,omitempty"`
	Region            string               `json:"region,omitempty"`
	AvailabilityZones []string             `json:"availability_zones,omitempty"`
	Status            string               `json:"status,omitempty"`
	Meta              *interface{}         `json:"meta,omitempty"`
}
