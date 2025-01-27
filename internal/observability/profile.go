package observability

import (
	"runtime"

	"github.com/grafana/pyroscope-go"
)

type ProfileConfig struct {
	ApplicationName string
	ServerAddress   string
	AuthUser        string
	AuthPassword    string
}

func NewProfiler(config *ProfileConfig) (*pyroscope.Profiler, error) {
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	profileCfg := pyroscope.Config{
		ApplicationName:   config.ApplicationName,
		ServerAddress:     config.ServerAddress,
		BasicAuthUser:     config.AuthUser,
		BasicAuthPassword: config.AuthPassword,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	}

	return pyroscope.Start(profileCfg)
}
