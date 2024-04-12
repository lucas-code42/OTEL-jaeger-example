package serviceb

import (
	"log"
	"net/http"

	"github.com/lucas-code42/OTEL-impl-example/pkg/otel"
	"go.opentelemetry.io/otel/propagation"
)

const thisServerName = "server-b"

func RunServiceB() {
	log.Printf("start service [B]")

	http.HandleFunc("/serviceb/ping", func(w http.ResponseWriter, r *http.Request) {
		propgator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
		ctx := propgator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))

		tracer := otel.InitializeTracer(r.Context(), thisServerName)
		_, span := tracer.Start(ctx, "handler B")
		span.SetName("handler B")
		defer span.End()

		w.Write([]byte("service_a -> ping -> service_b"))
	})
	http.ListenAndServe(":8181", nil)
}
