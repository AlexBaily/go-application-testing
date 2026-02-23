package telemetry

import (
	"log/slog"

	"github.com/grafana/pyroscope-go"
)

// InitProfiler initializes Pyroscope profiling
func InitProfiler(serviceName, pyroscopeURL string) error {
	slog.Debug("Initializing profiler...")

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: serviceName,
		ServerAddress:   pyroscopeURL,

		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})
	if err != nil {
		return err
	}

	slog.Info("Profiler initialized", "service", serviceName, "url", pyroscopeURL)
	return nil
}
