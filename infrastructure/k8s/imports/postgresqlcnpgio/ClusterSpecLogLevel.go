package postgresqlcnpgio


// The instances' log level, one of the following values: error, warning, info (default), debug, trace.
type ClusterSpecLogLevel string

const (
	// error.
	ClusterSpecLogLevel_ERROR ClusterSpecLogLevel = "ERROR"
	// warning.
	ClusterSpecLogLevel_WARNING ClusterSpecLogLevel = "WARNING"
	// info.
	ClusterSpecLogLevel_INFO ClusterSpecLogLevel = "INFO"
	// debug.
	ClusterSpecLogLevel_DEBUG ClusterSpecLogLevel = "DEBUG"
	// trace.
	ClusterSpecLogLevel_TRACE ClusterSpecLogLevel = "TRACE"
)

