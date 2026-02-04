package main

import (
	"context"
	"go-application-testing/internal/telemetry"
	"log/slog"

	"go.opentelemetry.io/otel"
)

func main() {
	slog.Debug("Starting server...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tracerProvider, err := telemetry.InitTracer()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create tracer", "error", err)
	}
	defer tracerProvider.Shutdown(ctx)

	traceTest(ctx)

}

func traceTest(ctx context.Context) {
	slog.InfoContext(ctx, "running traceTest...")

	ctx, span := otel.Tracer("app-testing").Start(ctx, "traceTest")
	defer span.End()

}
