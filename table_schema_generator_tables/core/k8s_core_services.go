package core

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sCoreServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sCoreServicesGenerator{}

func (x *TableK8sCoreServicesGenerator) GetTableName() string {
	return "k8s_core_services"
}

func (x *TableK8sCoreServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sCoreServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sCoreServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableK8sCoreServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().CoreV1().Services("")

			opts := metav1.ListOptions{}
			for {
				result, err := cl.List(ctx, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.Items
				if result.GetContinue() == "" {
					return nil
				}
				opts.Continue = result.GetContinue()
			}
		},
	}
}

func (x *TableK8sCoreServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sCoreServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_selector").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Selector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_load_balancer_source_ranges").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.LoadBalancerSourceRanges")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_external_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.ExternalName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_load_balancer").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.LoadBalancer")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_load_balancer_ip").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.LoadBalancerIP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_session_affinity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.SessionAffinity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_ip_families").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.IPFamilies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_load_balancer_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.LoadBalancerClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_internal_traffic_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.InternalTrafficPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_ports").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Ports")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_session_affinity_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.SessionAffinityConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_allocate_load_balancer_node_ports").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.AllocateLoadBalancerNodePorts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_cluster_ip").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.ClusterIP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_ip_family_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.IPFamilyPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_conditions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.Conditions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_external_ips").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.ExternalIPs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_health_check_node_port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.HealthCheckNodePort")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_external_traffic_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.ExternalTrafficPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_publish_not_ready_addresses").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.PublishNotReadyAddresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_cluster_ips").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.ClusterIPs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.Type")).Build(),
	}
}

func (x *TableK8sCoreServicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
