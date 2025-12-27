# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-12-27

### Added
- Initial stable release of the `dg-observability` plugin.
- **OpenTelemetry Integration**: Complete OTel SDK configuration and setup.
- **Prometheus Exporter**: Built-in Prometheus metrics exporter.
- **Auto-Discovery**: Automatic detection and instrumentation of active plugins.
- **Metrics Collection**: Standardized metrics for all framework components.
- **Container Integration**: Auto-registration with Injectable pattern.

### Features
- Zero-configuration observability for framework plugins
- Automatic metric registration for HTTP, database, cache, queue, and filesystem
- Prometheus-compatible metrics endpoint
- Configurable service name and environment
- Production-ready with minimal overhead

### Documentation
- Complete README with setup instructions
- Prometheus integration guide
- Metrics reference documentation

---

## Development History

The following versions represent the development journey leading to v1.0.0:

### 2025-11-24
- Initial implementation with OTel SDK
- Prometheus exporter integration
- Auto-discovery mechanism for plugins
