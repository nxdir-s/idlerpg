package storage

import (
	"example.com/charts/imports/k8s"
	cnpg "example.com/charts/imports/postgresqlcnpgio"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

const (
	PostgresReplicas   int    = 1
	PostgresApiVersion string = "postgresql.cnpg.io/v1"
)

type PostgresProps struct {
	Namespace k8s.KubeNamespace
	Replicas  *float64
}

func NewPostgresCluster(scope constructs.Construct, id *string, props *PostgresProps) cnpg.Cluster {
	database := constructs.NewConstruct(scope, id)

	replicas := props.Replicas
	if replicas == nil {
		replicas = jsii.Number(PostgresReplicas)
	}

	labels := map[string]*string{"name": id}

	cluster := cnpg.NewCluster(database, id, &cnpg.ClusterProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      id,
			Namespace: props.Namespace.Name(),
			Labels:    &labels,
		},
		Spec: &cnpg.ClusterSpec{
			Instances: replicas,
			Storage: &cnpg.ClusterSpecStorage{
				Size: jsii.String("256Mi"),
			},
			Monitoring: &cnpg.ClusterSpecMonitoring{
				EnablePodMonitor: jsii.Bool(true),
			},
		},
	})

	cluster.AddDependency(props.Namespace)

	return cluster
}
