package postgresqlcnpgio


// CatalogImage defines the image and major version.
type ImageCatalogSpecImages struct {
	// The image reference.
	Image *string `field:"required" json:"image" yaml:"image"`
	// The PostgreSQL major version of the image.
	//
	// Must be unique within the catalog.
	Major *float64 `field:"required" json:"major" yaml:"major"`
}

