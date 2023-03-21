package main

import (
	"github.com/selefra/selefra-provider-k8s/provider"
	"github.com/selefra/selefra-provider-sdk/grpc/serve"
)

func main() {

	myProvider := provider.GetProvider()
	serve.Serve(myProvider.Name, myProvider)

}
