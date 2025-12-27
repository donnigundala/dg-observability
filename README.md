# dg-observability

> Centralized observability plugin for the DG Framework, providing a unified OpenTelemetry-based metrics and tracing system.

The `dg-observability` plugin acts as the central telemetry hub for the framework. Other plugins (Database, Cache, HTTP, etc.) automatically detect its presence and export their metrics through it.

## Features

- 📊 **OpenTelemetry Native** - Built on the industry-standard OTel SDK.
- 🔗 **Automatic Discovery** - Plugins automatically hook into the observer if enabled.
- 🔌 **Multiple Exporters** - Support for Prometheus (default), OTLP (gRPC/HTTP), and Stdout.
- 🏗️ **Structured Metrics** - Standardized naming conventions across the ecosystem.

## Installation

```bash
go get github.com/donnigundala/dg-observability
```

## Usage

### 1. Register the Provider

```go
package main

import (
    "github.com/donnigundala/dg-core/foundation"
    "github.com/donnigundala/dg-observability"
)

func main() {
    app := foundation.New(".")
    
    // Register the observability provider
    app.Register(dgobservability.NewObservabilityServiceProvider())
    
    app.Start()
}
```

### 2. Manual Instrumentation

While most plugins handle this for you, you can also emit your own metrics:

```go
import "go.opentelemetry.io/otel"

meter := otel.GetMeterProvider().Meter("my-service")
counter, _ := meter.Int64Counter("my_custom_event_total")
counter.Add(ctx, 1)
```

## Configuration

The plugin uses the `observability` key in your configuration file.

### Configuration Mapping (YAML vs ENV)

| YAML Key | Environment Variable | Default | Description |
| :--- | :--- | :--- | :--- |
| `observability.enabled` | `OBSERVABILITY_ENABLED` | `true` | Enable telemetry |
| `observability.service_name` | `OBSERVABILITY_SERVICE_NAME` | - | Service name for OTel |
| `observability.exporter` | `OBSERVABILITY_EXPORTER` | `prometheus` | `prometheus`, `otlp`, `stdout` |
| `observability.endpoint` | `OBSERVABILITY_ENDPOINT` | - | OTLP Collector endpoint |

### Example YAML

```yaml
observability:
  enabled: true
  service_name: "my-awesome-app"
  exporter: "prometheus"
```

## Supported Metrics

When enabled, the following plugins will automatically export metrics:

- **dg-http**: Request count, duration, status codes.
- **dg-database**: Connection pool stats, latency.
- **dg-cache**: Hits, misses, evictions, items.
- **dg-queue**: Job processing counts, depth, worker status.
- **dg-filesystem**: Disk operations, throughput.

## License

MIT License - see [LICENSE](LICENSE) file for details.
