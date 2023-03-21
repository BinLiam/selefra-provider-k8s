package networking

import (


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









func createIngressClasses(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.IngressClass{}
	if err := faker.FakeObject(&r); err != nil {




		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockIngressClassInterface(ctrl)




	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.IngressClassList{Items: []resource.IngressClass{r}}, nil,


	)



	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)





	serviceClient.EXPECT().IngressClasses().AnyTimes().Return(resourceClient)





	cl := mocks.NewMockInterface(ctrl)




	cl.EXPECT().NetworkingV1().AnyTimes().Return(serviceClient)





	return cl


}



func TestIngressClasses(t *testing.T) {


	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sNetworkingIngressClassesGenerator{}), createIngressClasses, k8s_client.TestOptions{})




}




