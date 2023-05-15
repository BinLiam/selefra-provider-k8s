package networking

import (
	"github.com/selefra/selefra-provider-k8s/constants"

	"testing"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/mocks"

	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/networking/v1"

	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func createIngresses(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.Ingress{}

	if err := faker.FakeObject(&r); err != nil {

		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockIngressInterface(ctrl)

	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(

		&resource.IngressList{Items: []resource.Ingress{r}}, nil,
	)

	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)

	serviceClient.EXPECT().Ingresses(constants.Constants_45).AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)

	cl.EXPECT().NetworkingV1().AnyTimes().Return(serviceClient)

	return cl

}

func TestIngresses(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sNetworkingIngressesGenerator{}), createIngresses, k8s_client.TestOptions{})

}
