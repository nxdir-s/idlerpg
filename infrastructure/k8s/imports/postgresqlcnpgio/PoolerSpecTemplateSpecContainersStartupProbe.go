package postgresqlcnpgio


// StartupProbe indicates that the Pod has successfully initialized.
//
// If specified, no other probes are executed until this completes successfully.
// If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.
// This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,
// when it might take a long time to load data or warm a cache, than during steady-state operation.
// This cannot be updated.
// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
type PoolerSpecTemplateSpecContainersStartupProbe struct {
	// Exec specifies a command to execute in the container.
	Exec *PoolerSpecTemplateSpecContainersStartupProbeExec `field:"optional" json:"exec" yaml:"exec"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	// Default: 3. Minimum value is 1.
	//
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// GRPC specifies a GRPC HealthCheckRequest.
	Grpc *PoolerSpecTemplateSpecContainersStartupProbeGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	// HTTPGet specifies an HTTP GET request to perform.
	HttpGet *PoolerSpecTemplateSpecContainersStartupProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Number of seconds after the container has started before liveness probes are initiated.
	//
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Default: 10 seconds. Minimum value is 1.
	//
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	//
	// Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
	// Default: 1. Must be 1 for liveness and startup. Minimum value is 1.
	//
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// TCPSocket specifies a connection to a TCP port.
	TcpSocket *PoolerSpecTemplateSpecContainersStartupProbeTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	// Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
	//
	// The grace period is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
	// value overrides the value provided by the pod spec.
	// Value must be non-negative integer. The value zero indicates stop immediately via
	// the kill signal (no opportunity to shut down).
	// This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
	// Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.
	TerminationGracePeriodSeconds *float64 `field:"optional" json:"terminationGracePeriodSeconds" yaml:"terminationGracePeriodSeconds"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// Default: 1 second. Minimum value is 1.
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

