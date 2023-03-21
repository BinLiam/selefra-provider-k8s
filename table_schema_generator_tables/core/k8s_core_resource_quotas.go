package core

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sCoreResourceQuotasGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sCoreResourceQuotasGenerator{}

func (x *TableK8sCoreResourceQuotasGenerator) GetTableName() string {
	return "k8s_core_resource_quotas"
}

func (x *TableK8sCoreResourceQuotasGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sCoreResourceQuotasGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sCoreResourceQuotasGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableK8sCoreResourceQuotasGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().CoreV1().ResourceQuotas("")

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

func (x *TableK8sCoreResourceQuotasGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sCoreResourceQuotasGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_scope_selector").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.ScopeSelector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_hard").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.Hard")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_scopes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Spec.Scopes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_used").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.Used")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_hard").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Hard")).Build(),
	}
}

func (x *TableK8sCoreResourceQuotasGenerator) GetSubTables() []*schema.Table {
	return nil
}
