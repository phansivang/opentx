package config

type Config struct {
	OpenTxTarget string // OTLP endpoint (e.g., "localhost:4317")
	ServiceName  string // Service name for tracing
}
