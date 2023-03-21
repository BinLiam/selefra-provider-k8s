package admissionregistration





import (




	"testing"



	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-k8s/faker"




	"github.com/selefra/selefra-provider-k8s/k8s_client"
	mocks "github.com/selefra/selefra-provider-k8s/mocks"


	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/admissionregistration/v1"


	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/admissionregistration/v1"




	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"


	"k8s.io/client-go/kubernetes"




)







func createValidatingWebhookConfigurations(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.ValidatingWebhookConfiguration{}




	if err := faker.FakeObject(&r); err != nil {


		t.Fatal(err)


	}



	resourceClient := resourcemock.NewMockValidatingWebhookConfigurationInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(


		&resource.ValidatingWebhookConfigurationList{Items: []resource.ValidatingWebhookConfiguration{r}}, nil,




	)





	serviceClient := resourcemock.NewMockAdmissionregistrationV1Interface(ctrl)







	serviceClient.EXPECT().ValidatingWebhookConfigurations().AnyTimes().Return(resourceClient)



	cl := mocks.NewMockInterface(ctrl)




	cl.EXPECT().AdmissionregistrationV1().AnyTimes().Return(serviceClient)





	return cl




}

func TestValidatingWebhookConfigurations(t *testing.T) {




	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sAdmissionregistrationValidatingWebhookConfigurationsGenerator{}), createValidatingWebhookConfigurations, k8s_client.TestOptions{})
}
