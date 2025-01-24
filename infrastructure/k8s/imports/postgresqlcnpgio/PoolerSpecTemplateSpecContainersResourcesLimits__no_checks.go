//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validatePoolerSpecTemplateSpecContainersResourcesLimits_FromNumberParameters(value *float64) error {
	return nil
}

func validatePoolerSpecTemplateSpecContainersResourcesLimits_FromStringParameters(value *string) error {
	return nil
}

