package provider

import (
	"context"
	"github.com/selefra/selefra-provider-k8s/constants"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"strings"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Ks,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				k8sConfigPath := config.GetString(constants.Configpath)
				if k8sConfigPath != constants.Constants_10 {
					if !strings.HasPrefix(k8sConfigPath, constants.Constants_11) && !strings.HasPrefix(k8sConfigPath, constants.Constants_12) {
						return nil, diagnostics.AddErrorMsg(constants.Configpathmuststartwithorsisnotok, k8sConfigPath)
					}
				}
				client, err := k8s_client.NewClient(ctx, k8sConfigPath, nil)
				if err != nil {
					return nil, schema.NewDiagnostics().AddError(err)
				}
				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# config-path: <YOUR_K8S_CONFIG_FILE>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_13,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
