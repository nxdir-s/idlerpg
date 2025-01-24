package postgresqlcnpgio


// Requirements to be met by sync replicas.
//
// This will affect how the "synchronous_standby_names" parameter will be
// set up.
type ClusterSpecPostgresqlSyncReplicaElectionConstraint struct {
	// This flag enables the constraints for sync replicas.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// A list of node labels values to extract and compare to evaluate if the pods reside in the same topology or not.
	NodeLabelsAntiAffinity *[]*string `field:"optional" json:"nodeLabelsAntiAffinity" yaml:"nodeLabelsAntiAffinity"`
}

