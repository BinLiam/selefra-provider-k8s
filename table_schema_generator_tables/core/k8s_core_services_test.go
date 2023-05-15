package core

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/mocks"

	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/core/v1"

	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func createServices(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.Service{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)

	}

	r.Spec.ClusterIP = constants.Constants_39

	r.Spec.ClusterIPs = []string{constants.Constants_40}

	r.Spec.ExternalIPs = []string{constants.Constants_41}
	r.Spec.LoadBalancerIP = constants.Constants_42

	r.Spec.Ports = []resource.ServicePort{}

	resourceClient := resourcemock.NewMockServiceInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.ServiceList{Items: []resource.Service{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)

	serviceClient.EXPECT().Services(constants.Constants_43).AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().AnyTimes().Return(serviceClient)

	return cl

}

func TestServices(t *testing.T) {

	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCoreServicesGenerator{}), createServices, k8s_client.TestOptions{})

}
