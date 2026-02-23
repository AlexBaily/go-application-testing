package telemetry

import (
	"context"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"google.golang.org/grpc/credentials/insecure"
)

var Tracer = otel.Tracer("app-testing")

// Initialize tracer provider (do this once at startup)
func InitTracer(traceUrl string, ctx context.Context) (*trace.TracerProvider, error) {

	slog.Debug("Initalising tracer...")

	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(traceUrl+":4317"),
		otlptracegrpc.WithTLSCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName("application-testing"),
			semconv.ServiceVersion("1.0.0"),
		),
	)
	if err != nil {
		return nil, err
	}

	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(provider)

	return provider, nil
}

func TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := Tracer.Start(r.Context(), r.URL.Path)
		defer span.End()

		next.ServeHTTP(w, r.WithContext(ctx))

		span.SetAttributes(
			attribute.String("http.method", r.Method),
			attribute.String("http.route", r.URL.Path),
		)

	})
}
