package handlers

import (
	"log/slog"
	"net/http"
)

func HandleTest(w http.ResponseWriter, r *http.Request) {
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
