//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validatePublication_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validatePublication_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePublication_ManifestParameters(props *PublicationProps) error {
	return nil
}

func validatePublication_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewPublicationParameters(scope constructs.Construct, id *string, props *PublicationProps) error {
	return nil
}

