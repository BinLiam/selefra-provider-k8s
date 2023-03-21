package core

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sCoreReplicationControllersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sCoreReplicationControllersGenerator{}

func (x *TableK8sCoreReplicationControllersGenerator) GetTableName() string {
	return "k8s_core_replication_controllers"
}

func (x *TableK8sCoreReplicationControllersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sCoreReplicationControllersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sCoreReplicationControllersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableK8sCoreReplicationControllersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().CoreV1().ReplicationControllers("")

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

func (x *TableK8sCoreReplicationControllersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sCoreReplicationControllersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("status_ready_replicas").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.ReadyReplicas")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_observed_generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.ObservedGeneration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_replicas").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.Replicas")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_available_replicas").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.AvailableReplicas")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_template").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Template")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_min_ready_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.MinReadySeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_selector").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Selector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_replicas").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.Replicas")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_fully_labeled_replicas").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.FullyLabeledReplicas")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_conditions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.Conditions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableK8sCoreReplicationControllersGenerator) GetSubTables() []*schema.Table {
	return nil
}
