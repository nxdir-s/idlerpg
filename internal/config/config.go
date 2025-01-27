package config

import "os"

type ErrMissingEnv struct {
	envVar string
}

func (e *ErrMissingEnv) Error() string {
	return "missing env var: " + e.envVar
}

type ConfigOpt func(c *Config) error

// WithListenerAddr checks the env for LISTENER_ADDRESS
func WithListenerAddr() ConfigOpt {
	return func(c *Config) error {
		listenerAddr := os.Getenv("LISTENER_ADDRESS")
		if listenerAddr == "" {
			return &ErrMissingEnv{"LISTENER_ADDRESS"}
		}

		c.ListenerAddr = listenerAddr

		return nil
	}
}

// WithBrokers checks the env for BROKERS
func WithBrokers() ConfigOpt {
	return func(c *Config) error {
		brokers := os.Getenv("BROKERS")
		if brokers == "" {
			return &ErrMissingEnv{"BROKERS"}
		}

		c.Brokers = brokers

		return nil
	}
}

// WithRedPandaUsr checks the env for REDPANDA_SASL_USERNAME
func WithRedPandaUsr() ConfigOpt {
	return func(c *Config) error {
		rpUser := os.Getenv("REDPANDA_SASL_USERNAME")
		if rpUser == "" {
			return &ErrMissingEnv{"REDPANDA_SASL_USERNAME"}
		}

		c.RedPandaUsr = rpUser

		return nil
	}
}

// WithRedPandaPass checks the env for REDPANDA_SASL_PASSWORD
func WithRedPandaPass() ConfigOpt {
	return func(c *Config) error {
		rpPass := os.Getenv("REDPANDA_SASL_PASSWORD")
		if rpPass == "" {
			return &ErrMissingEnv{"REDPANDA_SASL_PASSWORD"}
		}

		c.RedPandaPass = rpPass

		return nil
	}
}

// WithOtelServiceName checks the env for OTEL_SERVICE_NAME
func WithOtelServiceName() ConfigOpt {
	return func(c *Config) error {
		serviceName := os.Getenv("OTEL_SERVICE_NAME")
		if serviceName == "" {
			return &ErrMissingEnv{"OTEL_SERVICE_NAME"}
		}

		c.OtelService = serviceName

		return nil
	}
}

// WithOtelEndpoint checks the env for OTEL_EXPORTER_OTLP_ENDPOINT
func WithOtelEndpoint() ConfigOpt {
	return func(c *Config) error {
		otelEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
		if otelEndpoint == "" {
			return &ErrMissingEnv{"OTEL_EXPORTER_OTLP_ENDPOINT"}
		}

		c.OtelEndpoint = otelEndpoint

		return nil
	}
}

// WithProfileURL checks the env for PROFILE_URL
func WithProfileURL() ConfigOpt {
	return func(c *Config) error {
		profileUrl := os.Getenv("PROFILE_URL")
		if profileUrl == "" {
			return &ErrMissingEnv{"PROFILE_URL"}
		}

		c.ProfileURL = profileUrl

		return nil
	}
}

// WithGrafanaUsr checks the env for GCLOUD_USER
func WithGrafanaUsr() ConfigOpt {
	return func(c *Config) error {
		gcUser := os.Getenv("GCLOUD_USER")
		if gcUser == "" {
			return &ErrMissingEnv{"GCLOUD_USER"}
		}

		c.GrafanaUsr = gcUser

		return nil
	}
}

// WithGrafanaPass checks the env for GCLOUD_PASSWORD
func WithGrafanaPass() ConfigOpt {
	return func(c *Config) error {
		gcPass := os.Getenv("GCLOUD_PASSWORD")
		if gcPass == "" {
			return &ErrMissingEnv{"GCLOUD_PASSWORD"}
		}

		c.GrafanaPass = gcPass

		return nil
	}
}

// WithConsumerName checks the env for CONSUMER_GROUP_NAME
func WithConsumerName() ConfigOpt {
	return func(c *Config) error {
		name := os.Getenv("CONSUMER_GROUP_NAME")
		if name == "" {
			return &ErrMissingEnv{"CONSUMER_GROUP_NAME"}
		}

		c.ConsumerName = name

		return nil
	}
}

// WithGoogleClientID checks the env for GOOGLE_CLIENT_ID
func WithGoogleClientID() ConfigOpt {
	return func(c *Config) error {
		id := os.Getenv("GOOGLE_CLIENT_ID")
		if id == "" {
			return &ErrMissingEnv{"GOOGLE_CLIENT_ID"}
		}

		c.GoogleClientID = id

		return nil
	}
}

// WithGoogleClientSecret checks the env for GOOGLE_CLIENT_SECRET
func WithGoogleClientSecret() ConfigOpt {
	return func(c *Config) error {
		secret := os.Getenv("GOOGLE_CLIENT_SECRET")
		if secret == "" {
			return &ErrMissingEnv{"GOOGLE_CLIENT_SECRET"}
		}

		c.GoogleClientSecret = secret

		return nil
	}
}

type Config struct {
	ListenerAddr       string
	Brokers            string
	RedPandaUsr        string
	RedPandaPass       string
	OtelService        string
	OtelEndpoint       string
	ProfileURL         string
	GrafanaUsr         string
	GrafanaPass        string
	ConsumerName       string
	GoogleClientID     string
	GoogleClientSecret string
}

func New(opts ...ConfigOpt) (*Config, error) {
	config := &Config{}

	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}
