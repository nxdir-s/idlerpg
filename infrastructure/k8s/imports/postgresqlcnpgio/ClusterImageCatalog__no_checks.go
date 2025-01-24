//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterImageCatalog_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateClusterImageCatalog_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterImageCatalog_ManifestParameters(props *ClusterImageCatalogProps) error {
	return nil
}

func validateClusterImageCatalog_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewClusterImageCatalogParameters(scope constructs.Construct, id *string, props *ClusterImageCatalogProps) error {
	return nil
}

