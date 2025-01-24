package postgresqlcnpgio


// Specification of the desired behavior of the cluster.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
type ClusterSpec struct {
	// Number of instances required in the cluster.
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// Affinity/Anti-affinity rules for Pods.
	Affinity *ClusterSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	// The configuration to be used for backups.
	Backup *ClusterSpecBackup `field:"optional" json:"backup" yaml:"backup"`
	// Instructions to bootstrap this cluster.
	Bootstrap *ClusterSpecBootstrap `field:"optional" json:"bootstrap" yaml:"bootstrap"`
	// The configuration for the CA and related certificates.
	Certificates *ClusterSpecCertificates `field:"optional" json:"certificates" yaml:"certificates"`
	// Description of this PostgreSQL cluster.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Manage the `PodDisruptionBudget` resources within the cluster.
	//
	// When
	// configured as `true` (default setting), the pod disruption budgets
	// will safeguard the primary node from being terminated. Conversely,
	// setting it to `false` will result in the absence of any
	// `PodDisruptionBudget` resource, permitting the shutdown of all nodes
	// hosting the PostgreSQL cluster. This latter configuration is
	// advisable for any PostgreSQL cluster employed for
	// development/staging purposes.
	EnablePdb *bool `field:"optional" json:"enablePdb" yaml:"enablePdb"`
	// When this option is enabled, the operator will use the `SuperuserSecret` to update the `postgres` user password (if the secret is not present, the operator will automatically create one).
	//
	// When this
	// option is disabled, the operator will ignore the `SuperuserSecret` content, delete
	// it when automatically created, and then blank the password of the `postgres`
	// user by setting it to `NULL`. Disabled by default.
	EnableSuperuserAccess *bool `field:"optional" json:"enableSuperuserAccess" yaml:"enableSuperuserAccess"`
	// Env follows the Env format to pass environment variables to the pods created in the cluster.
	Env *[]*ClusterSpecEnv `field:"optional" json:"env" yaml:"env"`
	// EnvFrom follows the EnvFrom format to pass environment variables sources to the pods to be used by Env.
	EnvFrom *[]*ClusterSpecEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	// EphemeralVolumeSource allows the user to configure the source of ephemeral volumes.
	EphemeralVolumeSource *ClusterSpecEphemeralVolumeSource `field:"optional" json:"ephemeralVolumeSource" yaml:"ephemeralVolumeSource"`
	// EphemeralVolumesSizeLimit allows the user to set the limits for the ephemeral volumes.
	EphemeralVolumesSizeLimit *ClusterSpecEphemeralVolumesSizeLimit `field:"optional" json:"ephemeralVolumesSizeLimit" yaml:"ephemeralVolumesSizeLimit"`
	// The list of external clusters which are used in the configuration.
	ExternalClusters *[]*ClusterSpecExternalClusters `field:"optional" json:"externalClusters" yaml:"externalClusters"`
	// The amount of time (in seconds) to wait before triggering a failover after the primary PostgreSQL instance in the cluster was detected to be unhealthy.
	FailoverDelay *float64 `field:"optional" json:"failoverDelay" yaml:"failoverDelay"`
	// Defines the major PostgreSQL version we want to use within an ImageCatalog.
	ImageCatalogRef *ClusterSpecImageCatalogRef `field:"optional" json:"imageCatalogRef" yaml:"imageCatalogRef"`
	// Name of the container image, supporting both tags (`<image>:<tag>`) and digests for deterministic and repeatable deployments (`<image>:<tag>@sha256:<digestValue>`).
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Image pull policy.
	//
	// One of `Always`, `Never` or `IfNotPresent`.
	// If not defined, it defaults to `IfNotPresent`.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// The list of pull secrets to be used to pull the images.
	ImagePullSecrets *[]*ClusterSpecImagePullSecrets `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	// Metadata that will be inherited by all objects related to the Cluster.
	InheritedMetadata *ClusterSpecInheritedMetadata `field:"optional" json:"inheritedMetadata" yaml:"inheritedMetadata"`
	// LivenessProbeTimeout is the time (in seconds) that is allowed for a PostgreSQL instance to successfully respond to the liveness probe (default 30).
	//
	// The Liveness probe failure threshold is derived from this value using the formula:
	// ceiling(livenessProbe / 10).
	LivenessProbeTimeout *float64 `field:"optional" json:"livenessProbeTimeout" yaml:"livenessProbeTimeout"`
	// The instances' log level, one of the following values: error, warning, info (default), debug, trace.
	LogLevel ClusterSpecLogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The configuration that is used by the portions of PostgreSQL that are managed by the instance manager.
	Managed *ClusterSpecManaged `field:"optional" json:"managed" yaml:"managed"`
	// The target value for the synchronous replication quorum, that can be decreased if the number of ready standbys is lower than this.
	//
	// Undefined or 0 disable synchronous replication.
	MaxSyncReplicas *float64 `field:"optional" json:"maxSyncReplicas" yaml:"maxSyncReplicas"`
	// Minimum number of instances required in synchronous replication with the primary.
	//
	// Undefined or 0 allow writes to complete when no standby is
	// available.
	MinSyncReplicas *float64 `field:"optional" json:"minSyncReplicas" yaml:"minSyncReplicas"`
	// The configuration of the monitoring infrastructure of this cluster.
	Monitoring *ClusterSpecMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Define a maintenance window for the Kubernetes nodes.
	NodeMaintenanceWindow *ClusterSpecNodeMaintenanceWindow `field:"optional" json:"nodeMaintenanceWindow" yaml:"nodeMaintenanceWindow"`
	// The plugins configuration, containing any plugin to be loaded with the corresponding configuration.
	Plugins *[]*ClusterSpecPlugins `field:"optional" json:"plugins" yaml:"plugins"`
	// The GID of the `postgres` user inside the image, defaults to `26`.
	PostgresGid *float64 `field:"optional" json:"postgresGid" yaml:"postgresGid"`
	// Configuration of the PostgreSQL server.
	Postgresql *ClusterSpecPostgresql `field:"optional" json:"postgresql" yaml:"postgresql"`
	// The UID of the `postgres` user inside the image, defaults to `26`.
	PostgresUid *float64 `field:"optional" json:"postgresUid" yaml:"postgresUid"`
	// Method to follow to upgrade the primary server during a rolling update procedure, after all replicas have been successfully updated: it can be with a switchover (`switchover`) or in-place (`restart` - default).
	PrimaryUpdateMethod ClusterSpecPrimaryUpdateMethod `field:"optional" json:"primaryUpdateMethod" yaml:"primaryUpdateMethod"`
	// Deployment strategy to follow to upgrade the primary server during a rolling update procedure, after all replicas have been successfully updated: it can be automated (`unsupervised` - default) or manual (`supervised`).
	PrimaryUpdateStrategy ClusterSpecPrimaryUpdateStrategy `field:"optional" json:"primaryUpdateStrategy" yaml:"primaryUpdateStrategy"`
	// Name of the priority class which will be used in every generated Pod, if the PriorityClass specified does not exist, the pod will not be able to schedule.
	//
	// Please refer to
	// https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#priorityclass
	// for more information.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// The configuration of the probes to be injected in the PostgreSQL Pods.
	Probes *ClusterSpecProbes `field:"optional" json:"probes" yaml:"probes"`
	// Template to be used to define projected volumes, projected volumes will be mounted under `/projected` base folder.
	ProjectedVolumeTemplate *ClusterSpecProjectedVolumeTemplate `field:"optional" json:"projectedVolumeTemplate" yaml:"projectedVolumeTemplate"`
	// Replica cluster configuration.
	Replica *ClusterSpecReplica `field:"optional" json:"replica" yaml:"replica"`
	// Replication slots management configuration.
	ReplicationSlots *ClusterSpecReplicationSlots `field:"optional" json:"replicationSlots" yaml:"replicationSlots"`
	// Resources requirements of every generated Pod.
	//
	// Please refer to
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// for more information.
	Resources *ClusterSpecResources `field:"optional" json:"resources" yaml:"resources"`
	// If specified, the pod will be dispatched by specified Kubernetes scheduler.
	//
	// If not specified, the pod will be dispatched by the default
	// scheduler. More info:
	// https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	// The SeccompProfile applied to every Pod and Container.
	//
	// Defaults to: `RuntimeDefault`.
	// Default: RuntimeDefault`.
	//
	SeccompProfile *ClusterSpecSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	// Configure the generation of the service account.
	ServiceAccountTemplate *ClusterSpecServiceAccountTemplate `field:"optional" json:"serviceAccountTemplate" yaml:"serviceAccountTemplate"`
	// The time in seconds that controls the window of time reserved for the smart shutdown of Postgres to complete.
	//
	// Make sure you reserve enough time for the operator to request a fast shutdown of Postgres
	// (that is: `stopDelay` - `smartShutdownTimeout`).
	SmartShutdownTimeout *float64 `field:"optional" json:"smartShutdownTimeout" yaml:"smartShutdownTimeout"`
	// The time in seconds that is allowed for a PostgreSQL instance to successfully start up (default 3600).
	//
	// The startup probe failure threshold is derived from this value using the formula:
	// ceiling(startDelay / 10).
	StartDelay *float64 `field:"optional" json:"startDelay" yaml:"startDelay"`
	// The time in seconds that is allowed for a PostgreSQL instance to gracefully shutdown (default 1800).
	StopDelay *float64 `field:"optional" json:"stopDelay" yaml:"stopDelay"`
	// Configuration of the storage of the instances.
	Storage *ClusterSpecStorage `field:"optional" json:"storage" yaml:"storage"`
	// The secret containing the superuser password.
	//
	// If not defined a new
	// secret will be created with a randomly generated password.
	SuperuserSecret *ClusterSpecSuperuserSecret `field:"optional" json:"superuserSecret" yaml:"superuserSecret"`
	// The time in seconds that is allowed for a primary PostgreSQL instance to gracefully shutdown during a switchover.
	//
	// Default value is 3600 seconds (1 hour).
	SwitchoverDelay *float64 `field:"optional" json:"switchoverDelay" yaml:"switchoverDelay"`
	// The tablespaces configuration.
	Tablespaces *[]*ClusterSpecTablespaces `field:"optional" json:"tablespaces" yaml:"tablespaces"`
	// TopologySpreadConstraints specifies how to spread matching pods among the given topology.
	//
	// More info:
	// https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/
	TopologySpreadConstraints *[]*ClusterSpecTopologySpreadConstraints `field:"optional" json:"topologySpreadConstraints" yaml:"topologySpreadConstraints"`
	// Configuration of the storage for PostgreSQL WAL (Write-Ahead Log).
	WalStorage *ClusterSpecWalStorage `field:"optional" json:"walStorage" yaml:"walStorage"`
}

