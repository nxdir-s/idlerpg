//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validatePooler_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validatePooler_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePooler_ManifestParameters(props *PoolerProps) error {
	return nil
}

func validatePooler_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewPoolerParameters(scope constructs.Construct, id *string, props *PoolerProps) error {
	return nil
}

