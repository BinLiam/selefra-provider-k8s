package provider

import (
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/admissionregistration"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/apps"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/autoscaling"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/batch"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/certificates"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/coordination"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/core"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/discovery"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/networking"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/nodes"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/rbac"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/storage"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&autoscaling.TableK8sAutoscalingHpasGenerator{}),
		table_schema_generator.GenTableSchema(&batch.TableK8sBatchCronJobsGenerator{}),
		table_schema_generator.GenTableSchema(&batch.TableK8sBatchJobsGenerator{}),
		table_schema_generator.GenTableSchema(&certificates.TableK8sCertificatesSigningRequestsGenerator{}),
		table_schema_generator.GenTableSchema(&coordination.TableK8sCoordinationLeasesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCorePodsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCorePvcsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreServiceAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreEndpointsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreServicesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreComponentStatusesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreNodesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreConfigMapsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCorePvsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreNamespacesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreReplicationControllersGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreLimitRangesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreResourceQuotasGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreEventsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreSecretsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCorePodsSecurityPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&admissionregistration.TableK8sAdmissionregistrationMutatingWebhookConfigurationsGenerator{}),
		table_schema_generator.GenTableSchema(&admissionregistration.TableK8sAdmissionregistrationValidatingWebhookConfigurationsGenerator{}),
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsStatefulSetsGenerator{}),
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsDaemonSetsGenerator{}),
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsDeploymentsGenerator{}),
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsReplicaSetsGenerator{}),
		table_schema_generator.GenTableSchema(&discovery.TableK8sDiscoveryEndpointSlicesGenerator{}),
		table_schema_generator.GenTableSchema(&networking.TableK8sNetworkingIngressClassesGenerator{}),
		table_schema_generator.GenTableSchema(&networking.TableK8sNetworkingIngressesGenerator{}),
		table_schema_generator.GenTableSchema(&networking.TableK8sNetworkingNetworkPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&nodes.TableK8sNodesRuntimeClassesGenerator{}),
		table_schema_generator.GenTableSchema(&rbac.TableK8sRbacClusterRoleBindingsGenerator{}),
		table_schema_generator.GenTableSchema(&rbac.TableK8sRbacClusterRolesGenerator{}),
		table_schema_generator.GenTableSchema(&rbac.TableK8sRbacRoleBindingsGenerator{}),
		table_schema_generator.GenTableSchema(&rbac.TableK8sRbacRolesGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableK8sStorageStorageClassesGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableK8sStorageVolumeAttachmentsGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableK8sStorageCsiDriversGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableK8sStorageCsiNodesGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableK8sStorageCsiStorageCapacitiesGenerator{}),
	}
}
