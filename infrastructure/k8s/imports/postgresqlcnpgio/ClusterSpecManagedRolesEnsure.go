package postgresqlcnpgio


// Ensure the role is `present` or `absent` - defaults to "present".
type ClusterSpecManagedRolesEnsure string

const (
	// present.
	ClusterSpecManagedRolesEnsure_PRESENT ClusterSpecManagedRolesEnsure = "PRESENT"
	// absent.
	ClusterSpecManagedRolesEnsure_ABSENT ClusterSpecManagedRolesEnsure = "ABSENT"
)

