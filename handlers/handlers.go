package handlers

import (
	"log/slog"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("app-testing")

func HandleTest(w http.ResponseWriter, r *http.Request) {
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

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		slog.Error("Failed to write response", "error", err)
	}
}
