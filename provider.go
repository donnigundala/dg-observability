package dgobservability

import (
	"context"
	"fmt"

	"github.com/donnigundala/dg-core/contracts/foundation"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// ObservabilityServiceProvider implements the PluginProvider interface.
type ObservabilityServiceProvider struct {
	Config        Config `config:"observability"`
	meterProvider *metric.MeterProvider
}

// NewObservabilityServiceProvider creates a new observability service provider.
func NewObservabilityServiceProvider() *ObservabilityServiceProvider {
	return &ObservabilityServiceProvider{}
}

// Name returns the provider name.
func (p *ObservabilityServiceProvider) Name() string {
	return Binding
}

// Version returns the provider version.
func (p *ObservabilityServiceProvider) Version() string {
	return Version
}

// Dependencies returns the provider dependencies.
func (p *ObservabilityServiceProvider) Dependencies() []string {
	return []string{}
}

// Register registers the observability services.
func (p *ObservabilityServiceProvider) Register(app foundation.Application) error {
	app.Singleton(Binding, func() (interface{}, error) {
		return p, nil
	})
	return nil
}

// Boot boots the observability services by initializing OpenTelemetry.
func (p *ObservabilityServiceProvider) Boot(app foundation.Application) error {
	if !p.Config.Enabled {
		return nil
	}

	// 1. Create Prometheus exporter (default for now)
	exporter, err := prometheus.New()
	if err != nil {
		return fmt.Errorf("failed to create prometheus exporter: %w", err)
	}

	// 2. Create resource
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(p.Config.ServiceName),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	// 3. Create and configure MeterProvider
	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(res),
	)

	// 4. Set global MeterProvider
	otel.SetMeterProvider(provider)
	p.meterProvider = provider

	return nil
}

// Shutdown gracefully shuts down the observability services.
func (p *ObservabilityServiceProvider) Shutdown(app foundation.Application) error {
	if p.meterProvider != nil {
		return p.meterProvider.Shutdown(context.Background())
	}
	return nil
}
