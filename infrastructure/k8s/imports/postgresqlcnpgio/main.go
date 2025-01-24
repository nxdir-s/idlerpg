// postgresqlcnpgio
package postgresqlcnpgio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"postgresqlcnpgio.Backup",
		reflect.TypeOf((*Backup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Backup{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.BackupProps",
		reflect.TypeOf((*BackupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.BackupSpec",
		reflect.TypeOf((*BackupSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.BackupSpecCluster",
		reflect.TypeOf((*BackupSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.BackupSpecMethod",
		reflect.TypeOf((*BackupSpecMethod)(nil)).Elem(),
		map[string]interface{}{
			"BARMAN_OBJECT_STORE": BackupSpecMethod_BARMAN_OBJECT_STORE,
			"VOLUME_SNAPSHOT": BackupSpecMethod_VOLUME_SNAPSHOT,
			"PLUGIN": BackupSpecMethod_PLUGIN,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.BackupSpecOnlineConfiguration",
		reflect.TypeOf((*BackupSpecOnlineConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.BackupSpecPluginConfiguration",
		reflect.TypeOf((*BackupSpecPluginConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.BackupSpecTarget",
		reflect.TypeOf((*BackupSpecTarget)(nil)).Elem(),
		map[string]interface{}{
			"PRIMARY": BackupSpecTarget_PRIMARY,
			"PREFER_HYPHEN_STANDBY": BackupSpecTarget_PREFER_HYPHEN_STANDBY,
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Cluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterImageCatalog",
		reflect.TypeOf((*ClusterImageCatalog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterImageCatalog{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterImageCatalogProps",
		reflect.TypeOf((*ClusterImageCatalogProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterImageCatalogSpec",
		reflect.TypeOf((*ClusterImageCatalogSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterImageCatalogSpecImages",
		reflect.TypeOf((*ClusterImageCatalogSpecImages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpec",
		reflect.TypeOf((*ClusterSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinity",
		reflect.TypeOf((*ClusterSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinity",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinity",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityAdditionalPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinity",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*ClusterSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecAffinityTolerations",
		reflect.TypeOf((*ClusterSpecAffinityTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackup",
		reflect.TypeOf((*ClusterSpecBackup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStore",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreAzureCredentials",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreAzureCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreAzureCredentialsConnectionString",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreAzureCredentialsConnectionString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreAzureCredentialsStorageAccount",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreAzureCredentialsStorageAccount)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreAzureCredentialsStorageKey",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreAzureCredentialsStorageKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreAzureCredentialsStorageSasToken",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreAzureCredentialsStorageSasToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreData",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreData)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreDataCompression",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreDataCompression)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": ClusterSpecBackupBarmanObjectStoreDataCompression_GZIP,
			"BZIP2": ClusterSpecBackupBarmanObjectStoreDataCompression_BZIP2,
			"SNAPPY": ClusterSpecBackupBarmanObjectStoreDataCompression_SNAPPY,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreDataEncryption",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreDataEncryption)(nil)).Elem(),
		map[string]interface{}{
			"AES256": ClusterSpecBackupBarmanObjectStoreDataEncryption_AES256,
			"AWS_KMS": ClusterSpecBackupBarmanObjectStoreDataEncryption_AWS_KMS,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreEndpointCa",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreEndpointCa)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreGoogleCredentials",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreGoogleCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreGoogleCredentialsApplicationCredentials",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreGoogleCredentialsApplicationCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreS3Credentials",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreS3Credentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreS3CredentialsAccessKeyId",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreS3CredentialsAccessKeyId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreS3CredentialsRegion",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreS3CredentialsRegion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreS3CredentialsSecretAccessKey",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreS3CredentialsSecretAccessKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreS3CredentialsSessionToken",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreS3CredentialsSessionToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreWal",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreWal)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreWalCompression",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreWalCompression)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": ClusterSpecBackupBarmanObjectStoreWalCompression_GZIP,
			"BZIP2": ClusterSpecBackupBarmanObjectStoreWalCompression_BZIP2,
			"SNAPPY": ClusterSpecBackupBarmanObjectStoreWalCompression_SNAPPY,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecBackupBarmanObjectStoreWalEncryption",
		reflect.TypeOf((*ClusterSpecBackupBarmanObjectStoreWalEncryption)(nil)).Elem(),
		map[string]interface{}{
			"AES256": ClusterSpecBackupBarmanObjectStoreWalEncryption_AES256,
			"AWS_KMS": ClusterSpecBackupBarmanObjectStoreWalEncryption_AWS_KMS,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecBackupTarget",
		reflect.TypeOf((*ClusterSpecBackupTarget)(nil)).Elem(),
		map[string]interface{}{
			"PRIMARY": ClusterSpecBackupTarget_PRIMARY,
			"PREFER_HYPHEN_STANDBY": ClusterSpecBackupTarget_PREFER_HYPHEN_STANDBY,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupVolumeSnapshot",
		reflect.TypeOf((*ClusterSpecBackupVolumeSnapshot)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBackupVolumeSnapshotOnlineConfiguration",
		reflect.TypeOf((*ClusterSpecBackupVolumeSnapshotOnlineConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference",
		reflect.TypeOf((*ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference_NONE,
			"CLUSTER": ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference_CLUSTER,
			"BACKUP": ClusterSpecBackupVolumeSnapshotSnapshotOwnerReference_BACKUP,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrap",
		reflect.TypeOf((*ClusterSpecBootstrap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdb",
		reflect.TypeOf((*ClusterSpecBootstrapInitdb)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbImport",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbImport)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbImportSource",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbImportSource)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbImportType",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbImportType)(nil)).Elem(),
		map[string]interface{}{
			"MICROSERVICE": ClusterSpecBootstrapInitdbImportType_MICROSERVICE,
			"MONOLITH": ClusterSpecBootstrapInitdbImportType_MONOLITH,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitApplicationSqlRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitApplicationSqlRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitApplicationSqlRefsConfigMapRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitApplicationSqlRefsConfigMapRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitApplicationSqlRefsSecretRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitApplicationSqlRefsSecretRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitSqlRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitSqlRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitSqlRefsConfigMapRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitSqlRefsConfigMapRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitSqlRefsSecretRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitSqlRefsSecretRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitTemplateSqlRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitTemplateSqlRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitTemplateSqlRefsConfigMapRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitTemplateSqlRefsConfigMapRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbPostInitTemplateSqlRefsSecretRefs",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbPostInitTemplateSqlRefsSecretRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapInitdbSecret",
		reflect.TypeOf((*ClusterSpecBootstrapInitdbSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapPgBasebackup",
		reflect.TypeOf((*ClusterSpecBootstrapPgBasebackup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapPgBasebackupSecret",
		reflect.TypeOf((*ClusterSpecBootstrapPgBasebackupSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecovery",
		reflect.TypeOf((*ClusterSpecBootstrapRecovery)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoveryBackup",
		reflect.TypeOf((*ClusterSpecBootstrapRecoveryBackup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoveryBackupEndpointCa",
		reflect.TypeOf((*ClusterSpecBootstrapRecoveryBackupEndpointCa)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoveryRecoveryTarget",
		reflect.TypeOf((*ClusterSpecBootstrapRecoveryRecoveryTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoverySecret",
		reflect.TypeOf((*ClusterSpecBootstrapRecoverySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoveryVolumeSnapshots",
		reflect.TypeOf((*ClusterSpecBootstrapRecoveryVolumeSnapshots)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoveryVolumeSnapshotsStorage",
		reflect.TypeOf((*ClusterSpecBootstrapRecoveryVolumeSnapshotsStorage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoveryVolumeSnapshotsTablespaceStorage",
		reflect.TypeOf((*ClusterSpecBootstrapRecoveryVolumeSnapshotsTablespaceStorage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecBootstrapRecoveryVolumeSnapshotsWalStorage",
		reflect.TypeOf((*ClusterSpecBootstrapRecoveryVolumeSnapshotsWalStorage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecCertificates",
		reflect.TypeOf((*ClusterSpecCertificates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnv",
		reflect.TypeOf((*ClusterSpecEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvFrom",
		reflect.TypeOf((*ClusterSpecEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvFromConfigMapRef",
		reflect.TypeOf((*ClusterSpecEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvFromSecretRef",
		reflect.TypeOf((*ClusterSpecEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvValueFrom",
		reflect.TypeOf((*ClusterSpecEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*ClusterSpecEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvValueFromFieldRef",
		reflect.TypeOf((*ClusterSpecEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvValueFromResourceFieldRef",
		reflect.TypeOf((*ClusterSpecEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*ClusterSpecEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEnvValueFromSecretKeyRef",
		reflect.TypeOf((*ClusterSpecEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSource",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplate",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpec",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecEphemeralVolumeSourceVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumesSizeLimit",
		reflect.TypeOf((*ClusterSpecEphemeralVolumesSizeLimit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumesSizeLimitShm",
		reflect.TypeOf((*ClusterSpecEphemeralVolumesSizeLimitShm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecEphemeralVolumesSizeLimitShm{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecEphemeralVolumesSizeLimitTemporaryData",
		reflect.TypeOf((*ClusterSpecEphemeralVolumesSizeLimitTemporaryData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecEphemeralVolumesSizeLimitTemporaryData{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClusters",
		reflect.TypeOf((*ClusterSpecExternalClusters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStore",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreAzureCredentials",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreAzureCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsConnectionString",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsConnectionString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageAccount",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageAccount)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageKey",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageSasToken",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreAzureCredentialsStorageSasToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreData",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreData)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreDataCompression",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreDataCompression)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": ClusterSpecExternalClustersBarmanObjectStoreDataCompression_GZIP,
			"BZIP2": ClusterSpecExternalClustersBarmanObjectStoreDataCompression_BZIP2,
			"SNAPPY": ClusterSpecExternalClustersBarmanObjectStoreDataCompression_SNAPPY,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreDataEncryption",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreDataEncryption)(nil)).Elem(),
		map[string]interface{}{
			"AES256": ClusterSpecExternalClustersBarmanObjectStoreDataEncryption_AES256,
			"AWS_KMS": ClusterSpecExternalClustersBarmanObjectStoreDataEncryption_AWS_KMS,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreEndpointCa",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreEndpointCa)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreGoogleCredentials",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreGoogleCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreGoogleCredentialsApplicationCredentials",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreGoogleCredentialsApplicationCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreS3Credentials",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreS3Credentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsAccessKeyId",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsAccessKeyId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsRegion",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsRegion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsSecretAccessKey",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsSecretAccessKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsSessionToken",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreS3CredentialsSessionToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreWal",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreWal)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreWalCompression",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreWalCompression)(nil)).Elem(),
		map[string]interface{}{
			"GZIP": ClusterSpecExternalClustersBarmanObjectStoreWalCompression_GZIP,
			"BZIP2": ClusterSpecExternalClustersBarmanObjectStoreWalCompression_BZIP2,
			"SNAPPY": ClusterSpecExternalClustersBarmanObjectStoreWalCompression_SNAPPY,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecExternalClustersBarmanObjectStoreWalEncryption",
		reflect.TypeOf((*ClusterSpecExternalClustersBarmanObjectStoreWalEncryption)(nil)).Elem(),
		map[string]interface{}{
			"AES256": ClusterSpecExternalClustersBarmanObjectStoreWalEncryption_AES256,
			"AWS_KMS": ClusterSpecExternalClustersBarmanObjectStoreWalEncryption_AWS_KMS,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersPassword",
		reflect.TypeOf((*ClusterSpecExternalClustersPassword)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersPlugin",
		reflect.TypeOf((*ClusterSpecExternalClustersPlugin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersSslCert",
		reflect.TypeOf((*ClusterSpecExternalClustersSslCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersSslKey",
		reflect.TypeOf((*ClusterSpecExternalClustersSslKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecExternalClustersSslRootCert",
		reflect.TypeOf((*ClusterSpecExternalClustersSslRootCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecImageCatalogRef",
		reflect.TypeOf((*ClusterSpecImageCatalogRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecImagePullSecrets",
		reflect.TypeOf((*ClusterSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecInheritedMetadata",
		reflect.TypeOf((*ClusterSpecInheritedMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecLogLevel",
		reflect.TypeOf((*ClusterSpecLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"ERROR": ClusterSpecLogLevel_ERROR,
			"WARNING": ClusterSpecLogLevel_WARNING,
			"INFO": ClusterSpecLogLevel_INFO,
			"DEBUG": ClusterSpecLogLevel_DEBUG,
			"TRACE": ClusterSpecLogLevel_TRACE,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManaged",
		reflect.TypeOf((*ClusterSpecManaged)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedRoles",
		reflect.TypeOf((*ClusterSpecManagedRoles)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecManagedRolesEnsure",
		reflect.TypeOf((*ClusterSpecManagedRolesEnsure)(nil)).Elem(),
		map[string]interface{}{
			"PRESENT": ClusterSpecManagedRolesEnsure_PRESENT,
			"ABSENT": ClusterSpecManagedRolesEnsure_ABSENT,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedRolesPasswordSecret",
		reflect.TypeOf((*ClusterSpecManagedRolesPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServices",
		reflect.TypeOf((*ClusterSpecManagedServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditional",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditional)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalSelectorType",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalSelectorType)(nil)).Elem(),
		map[string]interface{}{
			"RW": ClusterSpecManagedServicesAdditionalSelectorType_RW,
			"R": ClusterSpecManagedServicesAdditionalSelectorType_R,
			"RO": ClusterSpecManagedServicesAdditionalSelectorType_RO,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplate",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalServiceTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateMetadata",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalServiceTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateSpec",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalServiceTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateSpecPorts",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalServiceTemplateSpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecManagedServicesAdditionalServiceTemplateSpecPortsTargetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateSpecSessionAffinityConfig",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalServiceTemplateSpecSessionAffinityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalServiceTemplateSpecSessionAffinityConfigClientIp",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalServiceTemplateSpecSessionAffinityConfigClientIp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecManagedServicesAdditionalUpdateStrategy",
		reflect.TypeOf((*ClusterSpecManagedServicesAdditionalUpdateStrategy)(nil)).Elem(),
		map[string]interface{}{
			"PATCH": ClusterSpecManagedServicesAdditionalUpdateStrategy_PATCH,
			"REPLACE": ClusterSpecManagedServicesAdditionalUpdateStrategy_REPLACE,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecManagedServicesDisabledDefaultServices",
		reflect.TypeOf((*ClusterSpecManagedServicesDisabledDefaultServices)(nil)).Elem(),
		map[string]interface{}{
			"RW": ClusterSpecManagedServicesDisabledDefaultServices_RW,
			"R": ClusterSpecManagedServicesDisabledDefaultServices_R,
			"RO": ClusterSpecManagedServicesDisabledDefaultServices_RO,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecMonitoring",
		reflect.TypeOf((*ClusterSpecMonitoring)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecMonitoringCustomQueriesConfigMap",
		reflect.TypeOf((*ClusterSpecMonitoringCustomQueriesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecMonitoringCustomQueriesSecret",
		reflect.TypeOf((*ClusterSpecMonitoringCustomQueriesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecMonitoringPodMonitorMetricRelabelings",
		reflect.TypeOf((*ClusterSpecMonitoringPodMonitorMetricRelabelings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecMonitoringPodMonitorMetricRelabelingsAction",
		reflect.TypeOf((*ClusterSpecMonitoringPodMonitorMetricRelabelingsAction)(nil)).Elem(),
		map[string]interface{}{
			"REPLACE": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_REPLACE,
			"KEEP": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_KEEP,
			"DROP": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_DROP,
			"HASHMOD": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_HASHMOD,
			"LABELMAP": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LABELMAP,
			"LABELDROP": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LABELDROP,
			"LABELKEEP": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LABELKEEP,
			"LOWERCASE": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_LOWERCASE,
			"UPPERCASE": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_UPPERCASE,
			"KEEPEQUAL": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_KEEPEQUAL,
			"DROPEQUAL": ClusterSpecMonitoringPodMonitorMetricRelabelingsAction_DROPEQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecMonitoringPodMonitorRelabelings",
		reflect.TypeOf((*ClusterSpecMonitoringPodMonitorRelabelings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecMonitoringPodMonitorRelabelingsAction",
		reflect.TypeOf((*ClusterSpecMonitoringPodMonitorRelabelingsAction)(nil)).Elem(),
		map[string]interface{}{
			"REPLACE": ClusterSpecMonitoringPodMonitorRelabelingsAction_REPLACE,
			"KEEP": ClusterSpecMonitoringPodMonitorRelabelingsAction_KEEP,
			"DROP": ClusterSpecMonitoringPodMonitorRelabelingsAction_DROP,
			"HASHMOD": ClusterSpecMonitoringPodMonitorRelabelingsAction_HASHMOD,
			"LABELMAP": ClusterSpecMonitoringPodMonitorRelabelingsAction_LABELMAP,
			"LABELDROP": ClusterSpecMonitoringPodMonitorRelabelingsAction_LABELDROP,
			"LABELKEEP": ClusterSpecMonitoringPodMonitorRelabelingsAction_LABELKEEP,
			"LOWERCASE": ClusterSpecMonitoringPodMonitorRelabelingsAction_LOWERCASE,
			"UPPERCASE": ClusterSpecMonitoringPodMonitorRelabelingsAction_UPPERCASE,
			"KEEPEQUAL": ClusterSpecMonitoringPodMonitorRelabelingsAction_KEEPEQUAL,
			"DROPEQUAL": ClusterSpecMonitoringPodMonitorRelabelingsAction_DROPEQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecMonitoringTls",
		reflect.TypeOf((*ClusterSpecMonitoringTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecNodeMaintenanceWindow",
		reflect.TypeOf((*ClusterSpecNodeMaintenanceWindow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPlugins",
		reflect.TypeOf((*ClusterSpecPlugins)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPostgresql",
		reflect.TypeOf((*ClusterSpecPostgresql)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPostgresqlLdap",
		reflect.TypeOf((*ClusterSpecPostgresqlLdap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPostgresqlLdapBindAsAuth",
		reflect.TypeOf((*ClusterSpecPostgresqlLdapBindAsAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPostgresqlLdapBindSearchAuth",
		reflect.TypeOf((*ClusterSpecPostgresqlLdapBindSearchAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPostgresqlLdapBindSearchAuthBindPassword",
		reflect.TypeOf((*ClusterSpecPostgresqlLdapBindSearchAuthBindPassword)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecPostgresqlLdapScheme",
		reflect.TypeOf((*ClusterSpecPostgresqlLdapScheme)(nil)).Elem(),
		map[string]interface{}{
			"LDAP": ClusterSpecPostgresqlLdapScheme_LDAP,
			"LDAPS": ClusterSpecPostgresqlLdapScheme_LDAPS,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPostgresqlSyncReplicaElectionConstraint",
		reflect.TypeOf((*ClusterSpecPostgresqlSyncReplicaElectionConstraint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecPostgresqlSynchronous",
		reflect.TypeOf((*ClusterSpecPostgresqlSynchronous)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecPostgresqlSynchronousDataDurability",
		reflect.TypeOf((*ClusterSpecPostgresqlSynchronousDataDurability)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED": ClusterSpecPostgresqlSynchronousDataDurability_REQUIRED,
			"PREFERRED": ClusterSpecPostgresqlSynchronousDataDurability_PREFERRED,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecPostgresqlSynchronousMethod",
		reflect.TypeOf((*ClusterSpecPostgresqlSynchronousMethod)(nil)).Elem(),
		map[string]interface{}{
			"ANY": ClusterSpecPostgresqlSynchronousMethod_ANY,
			"FIRST": ClusterSpecPostgresqlSynchronousMethod_FIRST,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecPrimaryUpdateMethod",
		reflect.TypeOf((*ClusterSpecPrimaryUpdateMethod)(nil)).Elem(),
		map[string]interface{}{
			"SWITCHOVER": ClusterSpecPrimaryUpdateMethod_SWITCHOVER,
			"RESTART": ClusterSpecPrimaryUpdateMethod_RESTART,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ClusterSpecPrimaryUpdateStrategy",
		reflect.TypeOf((*ClusterSpecPrimaryUpdateStrategy)(nil)).Elem(),
		map[string]interface{}{
			"UNSUPERVISED": ClusterSpecPrimaryUpdateStrategy_UNSUPERVISED,
			"SUPERVISED": ClusterSpecPrimaryUpdateStrategy_SUPERVISED,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProbes",
		reflect.TypeOf((*ClusterSpecProbes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProbesLiveness",
		reflect.TypeOf((*ClusterSpecProbesLiveness)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProbesReadiness",
		reflect.TypeOf((*ClusterSpecProbesReadiness)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProbesStartup",
		reflect.TypeOf((*ClusterSpecProbesStartup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplate",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSources",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesClusterTrustBundle",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesClusterTrustBundle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesClusterTrustBundleLabelSelector",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesClusterTrustBundleLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesClusterTrustBundleLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesClusterTrustBundleLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesConfigMap",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesConfigMapItems",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesDownwardApi",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItems",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecProjectedVolumeTemplateSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesSecret",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesSecretItems",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecProjectedVolumeTemplateSourcesServiceAccountToken",
		reflect.TypeOf((*ClusterSpecProjectedVolumeTemplateSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecReplica",
		reflect.TypeOf((*ClusterSpecReplica)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecReplicationSlots",
		reflect.TypeOf((*ClusterSpecReplicationSlots)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecReplicationSlotsHighAvailability",
		reflect.TypeOf((*ClusterSpecReplicationSlotsHighAvailability)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecReplicationSlotsSynchronizeReplicas",
		reflect.TypeOf((*ClusterSpecReplicationSlotsSynchronizeReplicas)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecResources",
		reflect.TypeOf((*ClusterSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecResourcesClaims",
		reflect.TypeOf((*ClusterSpecResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecResourcesLimits",
		reflect.TypeOf((*ClusterSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecResourcesRequests",
		reflect.TypeOf((*ClusterSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecSeccompProfile",
		reflect.TypeOf((*ClusterSpecSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecServiceAccountTemplate",
		reflect.TypeOf((*ClusterSpecServiceAccountTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecServiceAccountTemplateMetadata",
		reflect.TypeOf((*ClusterSpecServiceAccountTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecStorage",
		reflect.TypeOf((*ClusterSpecStorage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplate",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateDataSource",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplateDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateDataSourceRef",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplateDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateResources",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplateResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateResourcesLimits",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplateResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecStoragePvcTemplateResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateResourcesRequests",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplateResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecStoragePvcTemplateResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateSelector",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplateSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecStoragePvcTemplateSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecStoragePvcTemplateSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecSuperuserSecret",
		reflect.TypeOf((*ClusterSpecSuperuserSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespaces",
		reflect.TypeOf((*ClusterSpecTablespaces)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesOwner",
		reflect.TypeOf((*ClusterSpecTablespacesOwner)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesStorage",
		reflect.TypeOf((*ClusterSpecTablespacesStorage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplate",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateDataSource",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplateDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateDataSourceRef",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplateDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateResources",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplateResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateResourcesLimits",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplateResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecTablespacesStoragePvcTemplateResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateResourcesRequests",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplateResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecTablespacesStoragePvcTemplateResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateSelector",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplateSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTablespacesStoragePvcTemplateSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecTablespacesStoragePvcTemplateSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTopologySpreadConstraints",
		reflect.TypeOf((*ClusterSpecTopologySpreadConstraints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTopologySpreadConstraintsLabelSelector",
		reflect.TypeOf((*ClusterSpecTopologySpreadConstraintsLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecTopologySpreadConstraintsLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecTopologySpreadConstraintsLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecWalStorage",
		reflect.TypeOf((*ClusterSpecWalStorage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplate",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateDataSource",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplateDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateDataSourceRef",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplateDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateResources",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplateResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateResourcesLimits",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplateResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecWalStoragePvcTemplateResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateResourcesRequests",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplateResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClusterSpecWalStoragePvcTemplateResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateSelector",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplateSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ClusterSpecWalStoragePvcTemplateSelectorMatchExpressions",
		reflect.TypeOf((*ClusterSpecWalStoragePvcTemplateSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.Database",
		reflect.TypeOf((*Database)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Database{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.DatabaseProps",
		reflect.TypeOf((*DatabaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.DatabaseSpec",
		reflect.TypeOf((*DatabaseSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.DatabaseSpecCluster",
		reflect.TypeOf((*DatabaseSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.DatabaseSpecDatabaseReclaimPolicy",
		reflect.TypeOf((*DatabaseSpecDatabaseReclaimPolicy)(nil)).Elem(),
		map[string]interface{}{
			"DELETE": DatabaseSpecDatabaseReclaimPolicy_DELETE,
			"RETAIN": DatabaseSpecDatabaseReclaimPolicy_RETAIN,
		},
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.DatabaseSpecEnsure",
		reflect.TypeOf((*DatabaseSpecEnsure)(nil)).Elem(),
		map[string]interface{}{
			"PRESENT": DatabaseSpecEnsure_PRESENT,
			"ABSENT": DatabaseSpecEnsure_ABSENT,
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ImageCatalog",
		reflect.TypeOf((*ImageCatalog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ImageCatalog{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ImageCatalogProps",
		reflect.TypeOf((*ImageCatalogProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ImageCatalogSpec",
		reflect.TypeOf((*ImageCatalogSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ImageCatalogSpecImages",
		reflect.TypeOf((*ImageCatalogSpecImages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.Pooler",
		reflect.TypeOf((*Pooler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Pooler{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerProps",
		reflect.TypeOf((*PoolerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpec",
		reflect.TypeOf((*PoolerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecCluster",
		reflect.TypeOf((*PoolerSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategy",
		reflect.TypeOf((*PoolerSpecDeploymentStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategyRollingUpdate",
		reflect.TypeOf((*PoolerSpecDeploymentStrategyRollingUpdate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategyRollingUpdateMaxSurge",
		reflect.TypeOf((*PoolerSpecDeploymentStrategyRollingUpdateMaxSurge)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecDeploymentStrategyRollingUpdateMaxSurge{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable",
		reflect.TypeOf((*PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecDeploymentStrategyRollingUpdateMaxUnavailable{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecMonitoring",
		reflect.TypeOf((*PoolerSpecMonitoring)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecMonitoringPodMonitorMetricRelabelings",
		reflect.TypeOf((*PoolerSpecMonitoringPodMonitorMetricRelabelings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.PoolerSpecMonitoringPodMonitorMetricRelabelingsAction",
		reflect.TypeOf((*PoolerSpecMonitoringPodMonitorMetricRelabelingsAction)(nil)).Elem(),
		map[string]interface{}{
			"REPLACE": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_REPLACE,
			"KEEP": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_KEEP,
			"DROP": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_DROP,
			"HASHMOD": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_HASHMOD,
			"LABELMAP": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LABELMAP,
			"LABELDROP": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LABELDROP,
			"LABELKEEP": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LABELKEEP,
			"LOWERCASE": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_LOWERCASE,
			"UPPERCASE": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_UPPERCASE,
			"KEEPEQUAL": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_KEEPEQUAL,
			"DROPEQUAL": PoolerSpecMonitoringPodMonitorMetricRelabelingsAction_DROPEQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecMonitoringPodMonitorRelabelings",
		reflect.TypeOf((*PoolerSpecMonitoringPodMonitorRelabelings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.PoolerSpecMonitoringPodMonitorRelabelingsAction",
		reflect.TypeOf((*PoolerSpecMonitoringPodMonitorRelabelingsAction)(nil)).Elem(),
		map[string]interface{}{
			"REPLACE": PoolerSpecMonitoringPodMonitorRelabelingsAction_REPLACE,
			"KEEP": PoolerSpecMonitoringPodMonitorRelabelingsAction_KEEP,
			"DROP": PoolerSpecMonitoringPodMonitorRelabelingsAction_DROP,
			"HASHMOD": PoolerSpecMonitoringPodMonitorRelabelingsAction_HASHMOD,
			"LABELMAP": PoolerSpecMonitoringPodMonitorRelabelingsAction_LABELMAP,
			"LABELDROP": PoolerSpecMonitoringPodMonitorRelabelingsAction_LABELDROP,
			"LABELKEEP": PoolerSpecMonitoringPodMonitorRelabelingsAction_LABELKEEP,
			"LOWERCASE": PoolerSpecMonitoringPodMonitorRelabelingsAction_LOWERCASE,
			"UPPERCASE": PoolerSpecMonitoringPodMonitorRelabelingsAction_UPPERCASE,
			"KEEPEQUAL": PoolerSpecMonitoringPodMonitorRelabelingsAction_KEEPEQUAL,
			"DROPEQUAL": PoolerSpecMonitoringPodMonitorRelabelingsAction_DROPEQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecPgbouncer",
		reflect.TypeOf((*PoolerSpecPgbouncer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecPgbouncerAuthQuerySecret",
		reflect.TypeOf((*PoolerSpecPgbouncerAuthQuerySecret)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.PoolerSpecPgbouncerPoolMode",
		reflect.TypeOf((*PoolerSpecPgbouncerPoolMode)(nil)).Elem(),
		map[string]interface{}{
			"SESSION": PoolerSpecPgbouncerPoolMode_SESSION,
			"TRANSACTION": PoolerSpecPgbouncerPoolMode_TRANSACTION,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecServiceTemplate",
		reflect.TypeOf((*PoolerSpecServiceTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecServiceTemplateMetadata",
		reflect.TypeOf((*PoolerSpecServiceTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecServiceTemplateSpec",
		reflect.TypeOf((*PoolerSpecServiceTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecServiceTemplateSpecPorts",
		reflect.TypeOf((*PoolerSpecServiceTemplateSpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecServiceTemplateSpecPortsTargetPort",
		reflect.TypeOf((*PoolerSpecServiceTemplateSpecPortsTargetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecServiceTemplateSpecPortsTargetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecServiceTemplateSpecSessionAffinityConfig",
		reflect.TypeOf((*PoolerSpecServiceTemplateSpecSessionAffinityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecServiceTemplateSpecSessionAffinityConfigClientIp",
		reflect.TypeOf((*PoolerSpecServiceTemplateSpecSessionAffinityConfigClientIp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplate",
		reflect.TypeOf((*PoolerSpecTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateMetadata",
		reflect.TypeOf((*PoolerSpecTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpec",
		reflect.TypeOf((*PoolerSpecTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinity",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinity",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinity",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainers",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnv",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvFrom",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvFromConfigMapRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvFromSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFrom",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFromFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecycle",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStart",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartSleep",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStartSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStop",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStopExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStopSleep",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStopSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersPorts",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResizePolicy",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersResizePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResources",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResourcesClaims",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResourcesLimits",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersResourcesRequests",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersSecurityContext",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersSecurityContextAppArmorProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersSecurityContextCapabilities",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersVolumeDevices",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecContainersVolumeMounts",
		reflect.TypeOf((*PoolerSpecTemplateSpecContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecDnsConfig",
		reflect.TypeOf((*PoolerSpecTemplateSpecDnsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecDnsConfigOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecDnsConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainers",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnv",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvFrom",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvFromConfigMapRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvFromSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFrom",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFromFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecycle",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStart",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartSleep",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStop",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopSleep",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersPorts",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResizePolicy",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersResizePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResources",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResourcesClaims",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResourcesLimits",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersResourcesRequests",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersSecurityContext",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersSecurityContextAppArmorProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersSecurityContextCapabilities",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersVolumeDevices",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecEphemeralContainersVolumeMounts",
		reflect.TypeOf((*PoolerSpecTemplateSpecEphemeralContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecHostAliases",
		reflect.TypeOf((*PoolerSpecTemplateSpecHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecImagePullSecrets",
		reflect.TypeOf((*PoolerSpecTemplateSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainers",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnv",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvFrom",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvFromSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvValueFrom",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecycle",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStart",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStartSleep",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStartSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStop",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStopSleep",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStopSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersPorts",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersResizePolicy",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersResizePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersResources",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersResourcesClaims",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersResourcesLimits",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersResourcesRequests",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersSecurityContext",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersSecurityContextAppArmorProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbe",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeExec",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeGrpc",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersVolumeDevices",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecInitContainersVolumeMounts",
		reflect.TypeOf((*PoolerSpecTemplateSpecInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecOs",
		reflect.TypeOf((*PoolerSpecTemplateSpecOs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecOverhead",
		reflect.TypeOf((*PoolerSpecTemplateSpecOverhead)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecOverhead{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecReadinessGates",
		reflect.TypeOf((*PoolerSpecTemplateSpecReadinessGates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourceClaims",
		reflect.TypeOf((*PoolerSpecTemplateSpecResourceClaims)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResources",
		reflect.TypeOf((*PoolerSpecTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourcesClaims",
		reflect.TypeOf((*PoolerSpecTemplateSpecResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourcesLimits",
		reflect.TypeOf((*PoolerSpecTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecResourcesRequests",
		reflect.TypeOf((*PoolerSpecTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecSchedulingGates",
		reflect.TypeOf((*PoolerSpecTemplateSpecSchedulingGates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecSecurityContext",
		reflect.TypeOf((*PoolerSpecTemplateSpecSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecSecurityContextAppArmorProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecSecurityContextSeLinuxOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecSecurityContextSeccompProfile",
		reflect.TypeOf((*PoolerSpecTemplateSpecSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecSecurityContextSysctls",
		reflect.TypeOf((*PoolerSpecTemplateSpecSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecSecurityContextWindowsOptions",
		reflect.TypeOf((*PoolerSpecTemplateSpecSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecTolerations",
		reflect.TypeOf((*PoolerSpecTemplateSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecTopologySpreadConstraints",
		reflect.TypeOf((*PoolerSpecTemplateSpecTopologySpreadConstraints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecTopologySpreadConstraintsLabelSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecTopologySpreadConstraintsLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecTopologySpreadConstraintsLabelSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecTopologySpreadConstraintsLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumes",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesAwsElasticBlockStore",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesAzureDisk",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesAzureFile",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesCephfs",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesCephfsSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesCinder",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesCinderSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesConfigMap",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesConfigMapItems",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesCsi",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesDownwardApi",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesDownwardApiItems",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEmptyDir",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeral",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesFc",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesFlexVolume",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesFlocker",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesGcePersistentDisk",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesGitRepo",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesGlusterfs",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesHostPath",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesImage",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesIscsi",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesIscsiSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesNfs",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesPersistentVolumeClaim",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesPhotonPersistentDisk",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesPortworxVolume",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjected",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSources",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundle",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelector",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelectorMatchExpressions",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PoolerSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesSecret",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesQuobyte",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesRbd",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesRbdSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesScaleIo",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesScaleIoSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesSecret",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesSecretItems",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesStorageos",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesStorageosSecretRef",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PoolerSpecTemplateSpecVolumesVsphereVolume",
		reflect.TypeOf((*PoolerSpecTemplateSpecVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.PoolerSpecType",
		reflect.TypeOf((*PoolerSpecType)(nil)).Elem(),
		map[string]interface{}{
			"RW": PoolerSpecType_RW,
			"RO": PoolerSpecType_RO,
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.Publication",
		reflect.TypeOf((*Publication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Publication{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PublicationProps",
		reflect.TypeOf((*PublicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PublicationSpec",
		reflect.TypeOf((*PublicationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PublicationSpecCluster",
		reflect.TypeOf((*PublicationSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.PublicationSpecPublicationReclaimPolicy",
		reflect.TypeOf((*PublicationSpecPublicationReclaimPolicy)(nil)).Elem(),
		map[string]interface{}{
			"DELETE": PublicationSpecPublicationReclaimPolicy_DELETE,
			"RETAIN": PublicationSpecPublicationReclaimPolicy_RETAIN,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PublicationSpecTarget",
		reflect.TypeOf((*PublicationSpecTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PublicationSpecTargetObjects",
		reflect.TypeOf((*PublicationSpecTargetObjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.PublicationSpecTargetObjectsTable",
		reflect.TypeOf((*PublicationSpecTargetObjectsTable)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.ScheduledBackup",
		reflect.TypeOf((*ScheduledBackup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ScheduledBackup{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ScheduledBackupProps",
		reflect.TypeOf((*ScheduledBackupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ScheduledBackupSpec",
		reflect.TypeOf((*ScheduledBackupSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ScheduledBackupSpecBackupOwnerReference",
		reflect.TypeOf((*ScheduledBackupSpecBackupOwnerReference)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ScheduledBackupSpecBackupOwnerReference_NONE,
			"SELF": ScheduledBackupSpecBackupOwnerReference_SELF,
			"CLUSTER": ScheduledBackupSpecBackupOwnerReference_CLUSTER,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ScheduledBackupSpecCluster",
		reflect.TypeOf((*ScheduledBackupSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ScheduledBackupSpecMethod",
		reflect.TypeOf((*ScheduledBackupSpecMethod)(nil)).Elem(),
		map[string]interface{}{
			"BARMAN_OBJECT_STORE": ScheduledBackupSpecMethod_BARMAN_OBJECT_STORE,
			"VOLUME_SNAPSHOT": ScheduledBackupSpecMethod_VOLUME_SNAPSHOT,
			"PLUGIN": ScheduledBackupSpecMethod_PLUGIN,
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ScheduledBackupSpecOnlineConfiguration",
		reflect.TypeOf((*ScheduledBackupSpecOnlineConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.ScheduledBackupSpecPluginConfiguration",
		reflect.TypeOf((*ScheduledBackupSpecPluginConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.ScheduledBackupSpecTarget",
		reflect.TypeOf((*ScheduledBackupSpecTarget)(nil)).Elem(),
		map[string]interface{}{
			"PRIMARY": ScheduledBackupSpecTarget_PRIMARY,
			"PREFER_HYPHEN_STANDBY": ScheduledBackupSpecTarget_PREFER_HYPHEN_STANDBY,
		},
	)
	_jsii_.RegisterClass(
		"postgresqlcnpgio.Subscription",
		reflect.TypeOf((*Subscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Subscription{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.SubscriptionProps",
		reflect.TypeOf((*SubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.SubscriptionSpec",
		reflect.TypeOf((*SubscriptionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"postgresqlcnpgio.SubscriptionSpecCluster",
		reflect.TypeOf((*SubscriptionSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"postgresqlcnpgio.SubscriptionSpecSubscriptionReclaimPolicy",
		reflect.TypeOf((*SubscriptionSpecSubscriptionReclaimPolicy)(nil)).Elem(),
		map[string]interface{}{
			"DELETE": SubscriptionSpecSubscriptionReclaimPolicy_DELETE,
			"RETAIN": SubscriptionSpecSubscriptionReclaimPolicy_RETAIN,
		},
	)
}
