package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
)

var (
	tracerProvider *sdktrace.TracerProvider
)

func Setup(openTxTarget, serviceName string) {
	ctx := context.Background()
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(openTxTarget),
	)
	if err != nil {
		log.Println("OpenTX initialize failed!")
		return
	}

	// Create a trace provider with the exporter
	tracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	log.Println("OpenTX initialized!")
}

func Shutdown(ctx context.Context) error {
	if tracerProvider != nil {
		return tracerProvider.Shutdown(ctx)
	}
	return nil
}

func StartSpan(ctx context.Context, spanName string) (context.Context, func()) {
	tracer := otel.Tracer("opentx")
	ctx, span := tracer.Start(ctx, spanName)

	log.Println("Created new span with Trace ID:", span.SpanContext().TraceID())

	return ctx, func() { span.End() }
}
