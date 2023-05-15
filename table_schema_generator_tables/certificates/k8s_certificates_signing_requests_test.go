package certificates

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8s_client"

	"github.com/selefra/selefra-provider-k8s/mocks"
	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/certificates/v1"

	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func createSigningRequests(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.CertificateSigningRequest{}

	if err := faker.FakeObject(&r); err != nil {

		t.Fatal(err)

	}

	resourceClient := resourcemock.NewMockCertificateSigningRequestInterface(ctrl)

	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(

		&resource.CertificateSigningRequestList{Items: []resource.CertificateSigningRequest{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCertificatesV1Interface(ctrl)

	serviceClient.EXPECT().CertificateSigningRequests().AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CertificatesV1().AnyTimes().Return(serviceClient)

	return cl

}

func TestSigningRequests(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCertificatesSigningRequestsGenerator{}), createSigningRequests, k8s_client.TestOptions{})

}
