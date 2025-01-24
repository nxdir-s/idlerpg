package postgresqlcnpgio


// Select all ClusterTrustBundles that match this label selector.
//
// Only has
// effect if signerName is set.  Mutually-exclusive with name.  If unset,
// interpreted as "match nothing".  If set but empty, interpreted as "match
// everything".
type PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelector struct {
	// matchExpressions is a list of label selector requirements.
	//
	// The requirements are ANDed.
	MatchExpressions *[]*PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// matchLabels is a map of {key,value} pairs.
	//
	// A single {key,value} in the matchLabels
	// map is equivalent to an element of matchExpressions, whose key field is "key", the
	// operator is "In", and the values array contains only "value". The requirements are ANDed.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

