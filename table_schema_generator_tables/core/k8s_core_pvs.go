package core

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sCorePvsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sCorePvsGenerator{}

func (x *TableK8sCorePvsGenerator) GetTableName() string {
	return "k8s_core_pvs"
}

func (x *TableK8sCorePvsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sCorePvsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sCorePvsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableK8sCorePvsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().CoreV1().PersistentVolumes()

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

func (x *TableK8sCorePvsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sCorePvsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("spec_volume_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.VolumeMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_capacity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Capacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_persistent_volume_reclaim_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.PersistentVolumeReclaimPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.Reason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_persistent_volume_source").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.PersistentVolumeSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_claim_ref").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.ClaimRef")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_storage_class_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.StorageClassName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_mount_options").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.MountOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_phase").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.Phase")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_access_modes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.AccessModes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_node_affinity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.NodeAffinity")).Build(),
	}
}

func (x *TableK8sCorePvsGenerator) GetSubTables() []*schema.Table {
	return nil
}
