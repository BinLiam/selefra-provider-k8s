package storage

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sStorageCsiDriversGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sStorageCsiDriversGenerator{}

func (x *TableK8sStorageCsiDriversGenerator) GetTableName() string {
	return "k8s_storage_csi_drivers"
}

func (x *TableK8sStorageCsiDriversGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sStorageCsiDriversGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sStorageCsiDriversGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableK8sStorageCsiDriversGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().StorageV1().CSIDrivers()

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

func (x *TableK8sStorageCsiDriversGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sStorageCsiDriversGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("spec_requires_republish").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.RequiresRepublish")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_se_linux_mount").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.SELinuxMount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_token_requests").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.TokenRequests")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_pod_info_on_mount").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.PodInfoOnMount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_storage_capacity").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.StorageCapacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_attach_required").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.AttachRequired")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_volume_lifecycle_modes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.VolumeLifecycleModes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_fs_group_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.FSGroupPolicy")).Build(),
	}
}

func (x *TableK8sStorageCsiDriversGenerator) GetSubTables() []*schema.Table {
	return nil
}
