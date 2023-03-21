package apps

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sAppsDaemonSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sAppsDaemonSetsGenerator{}

func (x *TableK8sAppsDaemonSetsGenerator) GetTableName() string {
	return "k8s_apps_daemon_sets"
}

func (x *TableK8sAppsDaemonSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sAppsDaemonSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sAppsDaemonSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableK8sAppsDaemonSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().AppsV1().DaemonSets("")

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

func (x *TableK8sAppsDaemonSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sAppsDaemonSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_collision_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.CollisionCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_number_available").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.NumberAvailable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_min_ready_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.MinReadySeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_revision_history_limit").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.RevisionHistoryLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_desired_number_scheduled").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.DesiredNumberScheduled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_selector").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Selector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_updated_number_scheduled").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.UpdatedNumberScheduled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_observed_generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.ObservedGeneration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_number_unavailable").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.NumberUnavailable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_conditions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.Conditions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_current_number_scheduled").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.CurrentNumberScheduled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_number_ready").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.NumberReady")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_template").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Template")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_update_strategy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.UpdateStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_number_misscheduled").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.NumberMisscheduled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableK8sAppsDaemonSetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
