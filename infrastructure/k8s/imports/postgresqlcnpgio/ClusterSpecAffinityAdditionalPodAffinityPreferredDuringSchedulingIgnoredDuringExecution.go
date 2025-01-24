package postgresqlcnpgio


// The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s).
type ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// Required.
	//
	// A pod affinity term, associated with the corresponding weight.
	PodAffinityTerm *ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

