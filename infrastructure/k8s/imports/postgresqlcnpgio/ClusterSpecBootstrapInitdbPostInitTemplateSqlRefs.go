package postgresqlcnpgio


// List of references to ConfigMaps or Secrets containing SQL files to be executed as a superuser in the `template1` database right after the cluster has been created.
//
// The references are processed in a specific order:
// first, all Secrets are processed, followed by all ConfigMaps.
// Within each group, the processing order follows the sequence specified
// in their respective arrays.
// (by default empty).
type ClusterSpecBootstrapInitdbPostInitTemplateSqlRefs struct {
	// ConfigMapRefs holds a list of references to ConfigMaps.
	ConfigMapRefs *[]*ClusterSpecBootstrapInitdbPostInitTemplateSqlRefsConfigMapRefs `field:"optional" json:"configMapRefs" yaml:"configMapRefs"`
	// SecretRefs holds a list of references to Secrets.
	SecretRefs *[]*ClusterSpecBootstrapInitdbPostInitTemplateSqlRefsSecretRefs `field:"optional" json:"secretRefs" yaml:"secretRefs"`
}

