package postgresqlcnpgio


// The import type.
//
// Can be `microservice` or `monolith`.
type ClusterSpecBootstrapInitdbImportType string

const (
	// microservice.
	ClusterSpecBootstrapInitdbImportType_MICROSERVICE ClusterSpecBootstrapInitdbImportType = "MICROSERVICE"
	// monolith.
	ClusterSpecBootstrapInitdbImportType_MONOLITH ClusterSpecBootstrapInitdbImportType = "MONOLITH"
)

