package serviceb

import (
	"log"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

const PATH = "serviceb/ping"

func RunServiceB(tracer trace.Tracer) {
	log.Printf("start service B")
	http.HandleFunc("/"+PATH, func(w http.ResponseWriter, r *http.Request) {
		_, span := tracer.Start(r.Context(), PATH)
		defer span.End()

		w.Write([]byte("service_a -> ping -> service_b"))
	})
	http.ListenAndServe(":8181", nil)
}
