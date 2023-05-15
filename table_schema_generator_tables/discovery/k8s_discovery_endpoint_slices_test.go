package discovery

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/faker"

	"github.com/selefra/selefra-provider-k8s/k8s_client"

	"github.com/selefra/selefra-provider-k8s/mocks"
	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/discovery/v1"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/discovery/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func createEndpointSlices(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.EndpointSlice{}

	if err := faker.FakeObject(&r); err != nil {

		t.Fatal(err)

	}

	resourceClient := resourcemock.NewMockEndpointSliceInterface(ctrl)

	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(

		&resource.EndpointSliceList{Items: []resource.EndpointSlice{r}}, nil,
	)

	serviceClient := resourcemock.NewMockDiscoveryV1Interface(ctrl)

	serviceClient.EXPECT().EndpointSlices(constants.Constants_44).AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)

	cl.EXPECT().DiscoveryV1().AnyTimes().Return(serviceClient)

	return cl

}

func TestEndpointSlices(t *testing.T) {

	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sDiscoveryEndpointSlicesGenerator{}), createEndpointSlices, k8s_client.TestOptions{})

}
