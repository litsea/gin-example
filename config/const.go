package config

const (
	// server

	KeyHost             = "server.host"
	KeyPort             = "server.port"
	KeyReadTimeout      = "server.read-timeout"
	KeyWriteTimeout     = "server.write-timeout"
	KeyRequestTimeout   = "server.request-timeout"
	KeyStopTimeout      = "server.stop-timeout"
	KeyCORSAllowOrigins = "server.cors.allow-origins"
	KeyPprofToken       = "server.pprof-token" //nolint:gosec

	// log

	KeyLog = "log"
)
