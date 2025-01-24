package postgresqlcnpgio


// Exec specifies a command to execute in the container.
type PoolerSpecTemplateSpecContainersLifecyclePreStopExec struct {
	// Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.
	//
	// The command is simply exec'd, it is
	// not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
	// a shell, you need to explicitly call out to that shell.
	// Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
}

