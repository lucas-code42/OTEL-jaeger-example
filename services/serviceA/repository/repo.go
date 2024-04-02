package repository

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func SimulateSQLQuery(ctx context.Context, tracer trace.Tracer) error {
	_, span := tracer.Start(ctx, "SQL query")
	defer span.End()

	span.AddEvent("SQL query", trace.WithAttributes(
		attribute.KeyValue{
			Key:   "eventKey",
			Value: attribute.StringValue("eventValue"),
		},
	))

	span.SetAttributes(attribute.KeyValue{
		Key:   "attributeKey",
		Value: attribute.StringValue("attributeValue"),
	})

	time.Sleep(2 * time.Second)
	return nil
}
