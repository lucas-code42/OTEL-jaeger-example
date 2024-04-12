package otel

import (
	"context"
	"errors"
	"log"

	"github.com/lucas-code42/OTEL-impl-example/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var errFoo = errors.New("error while creating tracer exporter")

func newResource(service string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(service),
		semconv.ServiceVersion("0.0.1"),
	)
}

func InitializeTracer(ctx context.Context, serviceName string) trace.Tracer {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpointURL(config.GetEnv("OTEL_EXPORTER_ENDPOINT")),
	)

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatalf("%s: %s", errFoo, err.Error())
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource(serviceName)),
	)

	// otel.SetTracerProvider(tracerProvider)
	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
	// 	propagation.TraceContext{}, propagation.Baggage{}),
	// )

	return tracerProvider.Tracer(serviceName, trace.WithInstrumentationVersion("1.0"))
}
