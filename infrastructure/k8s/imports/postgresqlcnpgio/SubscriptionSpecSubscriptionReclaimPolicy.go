package postgresqlcnpgio


// The policy for end-of-life maintenance of this subscription.
type SubscriptionSpecSubscriptionReclaimPolicy string

const (
	// delete.
	SubscriptionSpecSubscriptionReclaimPolicy_DELETE SubscriptionSpecSubscriptionReclaimPolicy = "DELETE"
	// retain.
	SubscriptionSpecSubscriptionReclaimPolicy_RETAIN SubscriptionSpecSubscriptionReclaimPolicy = "RETAIN"
)

