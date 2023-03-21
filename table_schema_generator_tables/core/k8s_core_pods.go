package core

import (
	"context"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sCorePodsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sCorePodsGenerator{}

func (x *TableK8sCorePodsGenerator) GetTableName() string {
	return "k8s_core_pods"
}

func (x *TableK8sCorePodsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sCorePodsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sCorePodsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"uid",
		},
	}
}

func (x *TableK8sCorePodsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*k8s_client.Client).Client().CoreV1().Pods("")

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

func (x *TableK8sCorePodsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sCorePodsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_host_ip").ColumnType(schema.ColumnTypeIp).
			Extractor(column_value_extractor.StructSelector("Status.HostIP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_pod_ips").ColumnType(schema.ColumnTypeIpArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ObjectMeta")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_pod_ip").ColumnType(schema.ColumnTypeIp).
			Extractor(column_value_extractor.StructSelector("Status.PodIP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_meta").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TypeMeta")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableK8sCorePodsGenerator) GetSubTables() []*schema.Table {
	return nil
}
