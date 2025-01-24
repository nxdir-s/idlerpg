//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validateBackup_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateBackup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBackup_ManifestParameters(props *BackupProps) error {
	return nil
}

func validateBackup_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewBackupParameters(scope constructs.Construct, id *string, props *BackupProps) error {
	return nil
}

