//go:build !no_runtime_type_checking

package postgresqlcnpgio

import (
	"fmt"
)

func validateClusterSpecEphemeralVolumesSizeLimitTemporaryData_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateClusterSpecEphemeralVolumesSizeLimitTemporaryData_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

