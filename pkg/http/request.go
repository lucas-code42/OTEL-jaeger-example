package http

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func RequestService(
	ctx context.Context,
	servicePort string,
	serviceName string,
	tracer trace.Tracer,
) string {
	_, span := tracer.Start(ctx, "ping-call")
	defer span.End()

	url := fmt.Sprintf("http://localhost:%s/%s/ping", servicePort, serviceName)
	fmt.Println(url)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		log.Fatalf("error while creating request for service %s", serviceName)
	}

	// propagator := propagation.TraceContext{}
	// propagator.Inject(ctx, propagation.MapCarrier{"msg": "some body"})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error while requesting for service %s", serviceName)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %s", err.Error())
	}

	return string(body)
}
