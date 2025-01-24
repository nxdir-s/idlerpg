package postgresqlcnpgio


// ServiceSelectorType describes a valid value for generating the service selectors.
//
// It indicates which type of service the selector applies to, such as read-write, read, or read-only.
type ClusterSpecManagedServicesDisabledDefaultServices string

const (
	// rw.
	ClusterSpecManagedServicesDisabledDefaultServices_RW ClusterSpecManagedServicesDisabledDefaultServices = "RW"
	// r.
	ClusterSpecManagedServicesDisabledDefaultServices_R ClusterSpecManagedServicesDisabledDefaultServices = "R"
	// ro.
	ClusterSpecManagedServicesDisabledDefaultServices_RO ClusterSpecManagedServicesDisabledDefaultServices = "RO"
)

