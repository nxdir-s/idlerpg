package postgresqlcnpgio


// Deprecated.
//
// TCPSocket is NOT supported as a LifecycleHandler and kept
// for backward compatibility. There is no validation of this field and
// lifecycle hooks will fail at runtime when it is specified.
type PoolerSpecTemplateSpecContainersLifecyclePreStopTcpSocket struct {
	// Number or name of the port to access on the container.
	//
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port PoolerSpecTemplateSpecContainersLifecyclePreStopTcpSocketPort `field:"required" json:"port" yaml:"port"`
	// Optional: Host name to connect to, defaults to the pod IP.
	Host *string `field:"optional" json:"host" yaml:"host"`
}
