//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validateScheduledBackup_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateScheduledBackup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScheduledBackup_ManifestParameters(props *ScheduledBackupProps) error {
	return nil
}

func validateScheduledBackup_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewScheduledBackupParameters(scope constructs.Construct, id *string, props *ScheduledBackupProps) error {
	return nil
}

