package postgresqlcnpgio


// Specification of the desired behavior of the ImageCatalog.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
type ImageCatalogSpec struct {
	// List of CatalogImages available in the catalog.
	Images *[]*ImageCatalogSpecImages `field:"required" json:"images" yaml:"images"`
}

