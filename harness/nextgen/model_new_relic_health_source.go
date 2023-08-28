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

// This is the NewRelic Metric Health Source spec entity defined in Harness
type NewRelicHealthSource struct {
	ConnectorRef              string                     `json:"connectorRef"`
	MetricPacks               []TimeSeriesMetricPackDto  `json:"metricPacks,omitempty"`
	ApplicationName           string                     `json:"applicationName,omitempty"`
	ApplicationId             string                     `json:"applicationId,omitempty"`
	Feature                   string                     `json:"feature,omitempty"`
	NewRelicMetricDefinitions []NewRelicMetricDefinition `json:"newRelicMetricDefinitions,omitempty"`
}
