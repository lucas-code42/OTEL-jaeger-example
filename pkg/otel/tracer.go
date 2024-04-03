package otel

import (
	"context"
	"errors"
	"log"

	"github.com/lucas-code42/OTEL-impl-example/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/trace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var errFoo = errors.New("error while creating tracer exporter")

func InitializeTracer(ctx context.Context) trace.Tracer {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpointURL(config.GetEnv("OTEL_EXPORTER_ENDPOINT")),
	)

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatalf("%s: %s", errFoo, err.Error())
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)

	return tracerProvider.Tracer(config.GetEnv("OTEL_SERVICE_NAME"))
}
