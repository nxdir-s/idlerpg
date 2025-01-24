package postgresqlcnpgio


// Ensure the PostgreSQL database is `present` or `absent` - defaults to "present".
type DatabaseSpecEnsure string

const (
	// present.
	DatabaseSpecEnsure_PRESENT DatabaseSpecEnsure = "PRESENT"
	// absent.
	DatabaseSpecEnsure_ABSENT DatabaseSpecEnsure = "ABSENT"
)

