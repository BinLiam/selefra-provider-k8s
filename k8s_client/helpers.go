package k8s_client

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type OwnerReferences struct {
	ResourceUID types.UID
	v1.OwnerReference
}

func ContextExpand() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, cli any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		client := cli.(*Client)
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)
		for _, ctxName := range client.contexts {
			clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
				Client: client.CopyWithContext(ctxName),
				Task:   task,
			})
		}
		return clientTaskContextSlice
	}
}

func ContextExtractor() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, cli any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		client := cli.(*Client)
		return client.Context, nil
	})
}
