package config

const (
	KeyEnv = "env"

	// server

	KeyHost             = "server.host"
	KeyPort             = "server.port"
	KeyReadTimeout      = "server.read-timeout"
	KeyWriteTimeout     = "server.write-timeout"
	KeyRequestTimeout   = "server.request-timeout"
	KeyStopTimeout      = "server.stop-timeout"
	KeyCORSAllowOrigins = "server.cors.allow-origins"
	KeyPprofToken       = "server.pprof-token" //nolint:gosec

	// profiler

	KeyProfilerServerAddress = "profiler.server-address"
	KeyProfilerAuthUsername  = "profiler.auth-username"
	KeyProfilerAuthPassword  = "profiler.auth-password" //nolint:gosec
	KeyProfilerDebug         = "profiler.debug"

	// log

	KeyLog = "log"
)
