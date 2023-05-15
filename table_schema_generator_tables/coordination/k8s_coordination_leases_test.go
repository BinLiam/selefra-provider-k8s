package coordination

import (
	"github.com/selefra/selefra-provider-k8s/constants"

	"testing"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8s_client"

	"github.com/selefra/selefra-provider-k8s/mocks"

	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/coordination/v1"

	"github.com/selefra/selefra-provider-k8s/table_schema_generator"

	resource "k8s.io/api/coordination/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createLeases(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.Lease{}

	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)

	}

	resourceClient := resourcemock.NewMockLeaseInterface(ctrl)

	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(

		&resource.LeaseList{Items: []resource.Lease{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoordinationV1Interface(ctrl)

	serviceClient.EXPECT().Leases(constants.Constants_23).AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoordinationV1().AnyTimes().Return(serviceClient)

	return cl
}

func TestLeases(t *testing.T) {

	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCoordinationLeasesGenerator{}), createLeases, k8s_client.TestOptions{})

}
