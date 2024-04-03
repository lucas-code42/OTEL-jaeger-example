package serviceb

import (
	"log"
	"net/http"

	"github.com/lucas-code42/OTEL-impl-example/pkg/otel"
)

const PATH = "serviceb/ping"

func RunServiceB() {
	log.Printf("start service B")
	http.HandleFunc("/"+PATH, func(w http.ResponseWriter, r *http.Request) {
		tracer := otel.InitializeTracer(r.Context())
		_, span := tracer.Start(r.Context(), PATH)
		span.SetName("handler B")
		defer span.End()

		w.Write([]byte("service_a -> ping -> service_b"))
	})
	http.ListenAndServe(":8181", nil)
}
