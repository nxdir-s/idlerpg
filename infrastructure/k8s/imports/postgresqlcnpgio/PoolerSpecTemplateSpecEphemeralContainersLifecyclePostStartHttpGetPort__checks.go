//go:build !no_runtime_type_checking

package postgresqlcnpgio

import (
	"fmt"
)

func validatePoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

