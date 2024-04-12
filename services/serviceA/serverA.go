package servicea

import (
	"log"
	"net/http"

	serviceHttp "github.com/lucas-code42/OTEL-impl-example/pkg/http"
	"github.com/lucas-code42/OTEL-impl-example/pkg/otel"
	"github.com/lucas-code42/OTEL-impl-example/services/serviceA/fakeRepository"
)

const thisServerName = "server-a"

func RunServiceA() {
	log.Printf("start service [A]")

	http.HandleFunc("/servicea/ping", func(w http.ResponseWriter, r *http.Request) {
		tracer := otel.InitializeTracer(r.Context(), thisServerName)

		ctx, span := tracer.Start(r.Context(), "http.request.A")
		// span.SetName("handler A")
		defer span.End()

		ctx, response := serviceHttp.RequestService(ctx, tracer, "8181", "serviceb")
		fakeRepository.SimulateSQLQuery(ctx, tracer)

		w.Write([]byte(response))
	})
	http.ListenAndServe(":8080", nil)
}
