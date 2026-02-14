package main

import (
	"context"
	"go-application-testing/internal/logging"
	"go-application-testing/internal/telemetry"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("app-testing")

func main() {
	slog.Debug("Starting server...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tracerProvider, err := telemetry.InitTracer("tempo", ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create tracer", "error", err)
	}
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			slog.Error("Failed to shutdown tracer", "error", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", handleTest)
	mux.HandleFunc("GET /health", handleHealth)

	handler := otelhttp.NewHandler(
		logging.HttpLogger(mux),
		"go-app",
	)
	slog.Info("Server listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		slog.Error("Server failed", "error", err)
	}

}

func handleTest(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "handleTest")
	defer span.End()

	span.SetAttributes(
		attribute.String("http.method", r.Method),
		attribute.String("http.route", "/test"),
	)
	span.SetAttributes(attribute.Int("http.status_code", 200))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
		slog.Error("Failed to write response", "error", err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		slog.Error("Failed to write response", "error", err)
	}
}
