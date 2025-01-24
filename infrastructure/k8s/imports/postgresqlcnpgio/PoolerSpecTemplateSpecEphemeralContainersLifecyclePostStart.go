package postgresqlcnpgio


// PostStart is called immediately after a container is created.
//
// If the handler fails,
// the container is terminated and restarted according to its restart policy.
// Other management of the container blocks until the hook completes.
// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
type PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStart struct {
	// Exec specifies a command to execute in the container.
	Exec *PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	// HTTPGet specifies an HTTP GET request to perform.
	HttpGet *PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Sleep represents a duration that the container should sleep.
	Sleep *PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartSleep `field:"optional" json:"sleep" yaml:"sleep"`
	// Deprecated.
	//
	// TCPSocket is NOT supported as a LifecycleHandler and kept
	// for backward compatibility. There is no validation of this field and
	// lifecycle hooks will fail at runtime when it is specified.
	TcpSocket *PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

