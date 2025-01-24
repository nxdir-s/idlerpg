package postgresqlcnpgio


// SelectorType specifies the type of selectors that the service will have.
//
// Valid values are "rw", "r", and "ro", representing read-write, read, and read-only services.
type ClusterSpecManagedServicesAdditionalSelectorType string

const (
	// rw.
	ClusterSpecManagedServicesAdditionalSelectorType_RW ClusterSpecManagedServicesAdditionalSelectorType = "RW"
	// r.
	ClusterSpecManagedServicesAdditionalSelectorType_R ClusterSpecManagedServicesAdditionalSelectorType = "R"
	// ro.
	ClusterSpecManagedServicesAdditionalSelectorType_RO ClusterSpecManagedServicesAdditionalSelectorType = "RO"
)

