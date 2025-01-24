package postgresqlcnpgio


// PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc.
//
// The handler is not called if the
// container crashes or exits. The Pod's termination grace period countdown begins before the
// PreStop hook is executed. Regardless of the outcome of the handler, the
// container will eventually terminate within the Pod's termination grace
// period (unless delayed by finalizers). Other management of the container blocks until the hook completes
// or until the termination grace period is reached.
// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
type PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStop struct {
	// Exec specifies a command to execute in the container.
	Exec *PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	// HTTPGet specifies an HTTP GET request to perform.
	HttpGet *PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Sleep represents a duration that the container should sleep.
	Sleep *PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopSleep `field:"optional" json:"sleep" yaml:"sleep"`
	// Deprecated.
	//
	// TCPSocket is NOT supported as a LifecycleHandler and kept
	// for backward compatibility. There is no validation of this field and
	// lifecycle hooks will fail at runtime when it is specified.
	TcpSocket *PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

