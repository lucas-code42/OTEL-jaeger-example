package otel

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/trace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitializeTracer(ctx context.Context, serviceName string) trace.Tracer {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpointURL("http://jaeger:4318/v1/traces"),
	)

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatalf("error while creating tracer exporter: %s", err.Error())
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)

	return tracerProvider.Tracer(serviceName)
}
