package serviceb

import (
	"log"
	"net/http"

	"github.com/lucas-code42/OTEL-impl-example/pkg/otel"
)

func RunServiceB() {
	log.Printf("start service [B]")

	http.HandleFunc("/serviceb/ping", func(w http.ResponseWriter, r *http.Request) {
		tracer := otel.InitializeTracer(r.Context())
		_, span := tracer.Start(r.Context(), "handler B")
		span.SetName("handler B")
		defer span.End()

		w.Write([]byte("service_a -> ping -> service_b"))
	})
	http.ListenAndServe(":8181", nil)
}
