package dgobservability

// Config defines the configuration for the observability provider.
type Config struct {
	// Enabled indicates if observability is enabled.
	Enabled bool `config:"enabled"`
	// ServiceName is the name of the service for metrics.
	ServiceName string `config:"service_name"`
	// Exporter defines the metrics exporter to use (e.g., "prometheus").
	Exporter string `config:"exporter"`
}

// DefaultConfig returns the default configuration for observability.
func DefaultConfig() Config {
	return Config{
		Enabled:     true,
		ServiceName: "dg-framework",
		Exporter:    "prometheus",
	}
}
