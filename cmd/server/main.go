package main

import (
	"context"
	"go-application-testing/handlers"
	"go-application-testing/internal/logging"
	"go-application-testing/internal/telemetry"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

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
	mux.HandleFunc("GET /test", handlers.HandleTest)
	mux.HandleFunc("GET /health", handlers.HandleHealth)

	handler := otelhttp.NewHandler(
		logging.HttpLogger(mux),
		"go-app",
	)
	slog.Info("Server listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		slog.Error("Server failed", "error", err)
	}

}
