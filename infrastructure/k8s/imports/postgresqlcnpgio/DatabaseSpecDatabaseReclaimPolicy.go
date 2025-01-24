package postgresqlcnpgio


// The policy for end-of-life maintenance of this database.
type DatabaseSpecDatabaseReclaimPolicy string

const (
	// delete.
	DatabaseSpecDatabaseReclaimPolicy_DELETE DatabaseSpecDatabaseReclaimPolicy = "DELETE"
	// retain.
	DatabaseSpecDatabaseReclaimPolicy_RETAIN DatabaseSpecDatabaseReclaimPolicy = "RETAIN"
)

