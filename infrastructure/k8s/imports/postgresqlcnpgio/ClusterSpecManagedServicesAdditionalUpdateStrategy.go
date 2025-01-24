package postgresqlcnpgio


// UpdateStrategy describes how the service differences should be reconciled.
type ClusterSpecManagedServicesAdditionalUpdateStrategy string

const (
	// patch.
	ClusterSpecManagedServicesAdditionalUpdateStrategy_PATCH ClusterSpecManagedServicesAdditionalUpdateStrategy = "PATCH"
	// replace.
	ClusterSpecManagedServicesAdditionalUpdateStrategy_REPLACE ClusterSpecManagedServicesAdditionalUpdateStrategy = "REPLACE"
)

