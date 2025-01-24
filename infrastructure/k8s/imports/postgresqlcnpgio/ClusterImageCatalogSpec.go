package postgresqlcnpgio


// Specification of the desired behavior of the ClusterImageCatalog.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
type ClusterImageCatalogSpec struct {
	// List of CatalogImages available in the catalog.
	Images *[]*ClusterImageCatalogSpecImages `field:"required" json:"images" yaml:"images"`
}

