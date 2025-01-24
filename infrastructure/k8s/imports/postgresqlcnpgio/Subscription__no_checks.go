//go:build no_runtime_type_checking

package postgresqlcnpgio

// Building without runtime type checking enabled, so all the below just return nil

func validateSubscription_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateSubscription_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubscription_ManifestParameters(props *SubscriptionProps) error {
	return nil
}

func validateSubscription_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewSubscriptionParameters(scope constructs.Construct, id *string, props *SubscriptionProps) error {
	return nil
}

