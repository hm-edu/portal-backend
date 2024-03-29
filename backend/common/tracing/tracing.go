package tracing

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
)

// LoggingHandler is a simple error handler that logs the error.
type LoggingHandler struct {
}

// Handle captures the logs to the zap logger.
func (l *LoggingHandler) Handle(_ error) {

}

// InitTracer performs the initialization of the traceprovider.
// By default this tries to init a jeager tracer.
func InitTracer(logger *zap.Logger, name string) *sdktrace.TracerProvider {
	//exporter, err := stdout.New(stdout.WithPrettyPrint())
	logger.Debug("Setting up tracing provider")
	var exporter sdktrace.SpanExporter
	var err error
	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") == "" {
		exporter, err = stdout.New(stdout.WithWriter(io.Discard))
	} else {
		opts := []otlptracegrpc.Option{otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))}
		client := otlptracegrpc.NewClient(opts...)
		exporter, err = otlptrace.New(context.Background(), client)
	}

	if err != nil {
		logger.Fatal("Error creating collector.", zap.Error(err))
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(name),
			)),
	)

	otelProm, err := otelprom.New()
	if err != nil {
		logger.Panic("failed to initialize prometheus exporter", zap.Error(err))
	}
	provider := metric.NewMeterProvider(metric.WithReader(otelProm))

	otel.SetMeterProvider(provider)
	otel.SetTracerProvider(tp)
	b3 := b3.New()
	otel.SetTextMapPropagator(b3)
	otel.SetErrorHandler(&LoggingHandler{})

	http.Handle("/", promhttp.Handler())
	go func() {
		_ = http.ListenAndServe(":2222", nil) // nolint:gosec // we expect don't expose this interface to the internet
	}()

	logger.Info("Prometheus server running on :2222")
	return tp
}
