package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var (
	errMountRequest = errors.New("error while creating request for service")
	errRequest      = errors.New("error while requesting for service")
	errReadBody     = errors.New("failed to read response body")
)

func RequestService(
	ctx context.Context,
	tracer trace.Tracer,
	port string,
	path string,
) (context.Context, string) {
	ctx, span := tracer.Start(ctx, "htt.ping")
	defer span.End()

	url := fmt.Sprintf("http://localhost:%s/%s/ping", port, path)
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		span.RecordError(errMountRequest)
		return ctx, ""
	}

	span.AddEvent("requestMount", trace.WithAttributes(
		attribute.KeyValue{
			Key:   attribute.Key("status"),
			Value: attribute.StringValue("OK"),
		},
	))

	propagator := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{},
	)
	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		span.RecordError(errRequest)
		return ctx, ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(errReadBody)
		return ctx, ""
	}

	return ctx, string(body)
}
