//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validateImageCatalog_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateImageCatalog_IsConstructParameters(x interface{}) error {
	return nil
}

func validateImageCatalog_ManifestParameters(props *ImageCatalogProps) error {
	return nil
}

func validateImageCatalog_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewImageCatalogParameters(scope constructs.Construct, id *string, props *ImageCatalogProps) error {
	return nil
}

