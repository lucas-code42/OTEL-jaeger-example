package servicea

import (
	"log"
	"net/http"

	serviceHttp "github.com/lucas-code42/OTEL-impl-example/pkg/http"
	"github.com/lucas-code42/OTEL-impl-example/services/serviceA/fakeRepository"
	"go.opentelemetry.io/otel/trace"
)

const PATH = "servicea/ping"

func RunServiceA(tracer trace.Tracer) {
	log.Printf("start service A")
	http.HandleFunc("/"+PATH, func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), PATH)
		defer span.End()

		response := serviceHttp.RequestService(ctx, "8181", "serviceb", tracer)
		fakeRepository.SimulateSQLQuery(ctx, tracer)
		w.Write([]byte(response))
	})
	http.ListenAndServe(":8080", nil)
}
