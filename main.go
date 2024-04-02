package main

import (
	"context"

	"github.com/lucas-code42/OTEL-impl-example/pkg/otel"
	servicea "github.com/lucas-code42/OTEL-impl-example/services/serviceA"
	serviceb "github.com/lucas-code42/OTEL-impl-example/services/serviceB"
)

func main() {
	tracer_a := otel.InitializeTracer(context.Background(), "serviceA")
	tracer_b := otel.InitializeTracer(context.Background(), "serviceB")

	go serviceb.RunServiceB(tracer_b)
	servicea.RunServiceA(tracer_a)
}
