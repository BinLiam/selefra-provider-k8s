package k8s_client

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/constants"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"testing"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) kubernetes.Interface, _ TestOptions) {
	ctrl := gomock.NewController(t)
	testProvider := newTestProvider(t, ctrl, table, builder)
	config := constants.Testtest
	test_helper.RunProviderPullTables(testProvider, config, constants.Constants_8, constants.Constants_9)
}

func newTestProvider(t *testing.T, ctrl *gomock.Controller, table *schema.Table, builder func(*testing.T, *gomock.Controller) kubernetes.Interface) *provider.Provider {
	return &provider.Provider{
		Name:      constants.Ks,
		Version:   constants.V,
		TableList: []*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				services := builder(t, ctrl)
				client := &Client{
					Context: constants.TestContext,
				}
				client.SetServices(map[string]kubernetes.Interface{constants.TestContext: services})
				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return ``
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{},
			DataSourcePullResultAutoExpand:       true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
