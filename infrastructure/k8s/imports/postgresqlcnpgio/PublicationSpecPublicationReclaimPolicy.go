package postgresqlcnpgio


// The policy for end-of-life maintenance of this publication.
type PublicationSpecPublicationReclaimPolicy string

const (
	// delete.
	PublicationSpecPublicationReclaimPolicy_DELETE PublicationSpecPublicationReclaimPolicy = "DELETE"
	// retain.
	PublicationSpecPublicationReclaimPolicy_RETAIN PublicationSpecPublicationReclaimPolicy = "RETAIN"
)

