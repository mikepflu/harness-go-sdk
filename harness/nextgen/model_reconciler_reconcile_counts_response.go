/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ReconcilerReconcileCountsResponse struct {
	ApplicationCount int32 `json:"applicationCount,omitempty"`
	ClusterCount int32 `json:"clusterCount,omitempty"`
	RepositoryCount int32 `json:"repositoryCount,omitempty"`
	RepositoryCertificateCount int32 `json:"repositoryCertificateCount,omitempty"`
	GnuPGPublicKeyCount int32 `json:"gnuPGPublicKeyCount,omitempty"`
	RepoCredsCount int32 `json:"repoCredsCount,omitempty"`
	ApplicationPerProjectCount map[string]int32 `json:"applicationPerProjectCount,omitempty"`
	ClusterPerProjectCount map[string]int32 `json:"clusterPerProjectCount,omitempty"`
	RepositoryPerProjectCount map[string]int32 `json:"repositoryPerProjectCount,omitempty"`
}
