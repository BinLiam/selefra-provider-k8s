package k8s_client

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	"context"
	"fmt"
	"strings"

	"k8s.io/client-go/kubernetes"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	services	map[string]kubernetes.Interface
	contexts	[]string
	paths		map[string]struct{}

	Context	string

	KubeConfig	clientcmd.ClientConfig
}

func (c *Client) Client() kubernetes.Interface {
	return c.services[c.Context]
}

func (c Client) CopyWithContext(k8sContext string) *Client {
	return &Client{
		services:	c.services,
		contexts:	c.contexts,
		paths:		c.paths,
		KubeConfig:	c.KubeConfig,
		Context:	k8sContext,
	}
}

func (c *Client) SetServices(s map[string]kubernetes.Interface) {
	c.services = s
	contexts := make([]string, 0, len(s))
	for k := range s {
		contexts = append(contexts, k)
	}
	c.contexts = contexts
}

func NewClient(ctx context.Context, k8sConfigFile string, k8sContextSlice []string) (*Client, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	loadingRules.ExplicitPath = k8sConfigFile
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides,
	)
	rawKubeConfig, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	var contexts []string
	switch len(k8sContextSlice) {
	case 0:
		contexts = []string{rawKubeConfig.CurrentContext}
	case 1:
		if k8sContextSlice[0] == constants.Constants_7 {
			for cName := range rawKubeConfig.Contexts {
				contexts = append(contexts, cName)
			}
		} else {
			if _, ok := rawKubeConfig.Contexts[k8sContextSlice[0]]; !ok {
				return nil, fmt.Errorf(constants.Contextqdoesntexistinkubeconfiguration, k8sContextSlice[0])
			}
			contexts = []string{k8sContextSlice[0]}
		}
	default:
		for _, cName := range k8sContextSlice {
			if _, ok := rawKubeConfig.Contexts[cName]; !ok {
				return nil, fmt.Errorf(constants.Contextqdoesntexistinkubeconfiguration, cName)
			}
			contexts = append(contexts, cName)
		}
	}

	if len(contexts) == 0 {
		return nil, fmt.Errorf(constants.CouldnotfindanycontextTrytoaddcontexthttpskubernetesiodocsreferencekubectlcheatsheetkubectlcontextandconfiguration)
	}

	c := Client{
		services:	make(map[string]kubernetes.Interface),
		contexts:	contexts,
		Context:	contexts[0],
		paths:		make(map[string]struct{}),
		KubeConfig:	kubeConfig,
	}

	for _, ctxName := range contexts {
		kClient, err := buildKubeClient(rawKubeConfig, ctxName)
		if err != nil {
			return nil, fmt.Errorf(constants.Failedtobuildksclientforcontextqw, ctxName, err)
		}
		c.paths, err = getAPIsMap(kClient)
		if err != nil {
		}
		c.services[ctxName] = kClient
	}

	return &c, nil
}

func buildKubeClient(kubeConfig api.Config, ctx string) (*kubernetes.Clientset, error) {
	override := &clientcmd.ConfigOverrides{CurrentContext: ctx}
	clientConfig := clientcmd.NewNonInteractiveClientConfig(
		kubeConfig,
		override.CurrentContext,
		override,
		&clientcmd.ClientConfigLoadingRules{},
	)
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(restConfig)
}

func getAPIsMap(client *kubernetes.Clientset) (map[string]struct{}, error) {
	doc, err := client.OpenAPISchema()
	if err != nil {
		return nil, err
	}
	paths := make(map[string]struct{})
	for _, p := range doc.Paths.Path {
		path := p.Name
		if strings.HasPrefix(path, constants.Apis) {
			paths[path] = struct{}{}
		}
	}
	return paths, nil
}
