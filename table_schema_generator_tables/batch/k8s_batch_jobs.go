package batch

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sBatchJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sBatchJobsGenerator{}

func (x *TableK8sBatchJobsGenerator) GetTableName() string {
	return "k8s_batch_jobs"
}

func (x *TableK8sBatchJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sBatchJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sBatchJobsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableK8sBatchJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().BatchV1().Jobs("")

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

func (x *TableK8sBatchJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sBatchJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_manual_selector").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.ManualSelector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_template").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Template")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_completion_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Status.CompletionTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_succeeded").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.Succeeded")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_failed").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.Failed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_pod_failure_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.PodFailurePolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_suspend").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.Suspend")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Status.StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_parallelism").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.Parallelism")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_ttl_seconds_after_finished").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.TTLSecondsAfterFinished")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_completion_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.CompletionMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_ready").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.Ready")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_active").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status.Active")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_conditions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.Conditions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_completed_indexes").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.CompletedIndexes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_uncounted_terminated_pods").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.UncountedTerminatedPods")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_active_deadline_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.ActiveDeadlineSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_backoff_limit").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.BackoffLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_selector").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Selector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_completions").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Spec.Completions")).Build(),
	}
}

func (x *TableK8sBatchJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
