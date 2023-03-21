

package autoscaling







import (
	"github.com/selefra/selefra-provider-k8s/constants"
	"testing"









	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8s_client"




	"github.com/selefra/selefra-provider-k8s/mocks"


	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/autoscaling/v1"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"


	"k8s.io/client-go/kubernetes"


)





func createHpas(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {


	r := resource.HorizontalPodAutoscaler{}


	if err := faker.FakeObject(&r); err != nil {




		t.Fatal(err)




	}





	resourceClient := resourcemock.NewMockHorizontalPodAutoscalerInterface(ctrl)




	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.HorizontalPodAutoscalerList{Items: []resource.HorizontalPodAutoscaler{r}}, nil,




	)





	serviceClient := resourcemock.NewMockAutoscalingV1Interface(ctrl)



	serviceClient.EXPECT().HorizontalPodAutoscalers(constants.Constants_20).AnyTimes().Return(resourceClient)





	cl := mocks.NewMockInterface(ctrl)




	cl.EXPECT().AutoscalingV1().AnyTimes().Return(serviceClient)







	return cl
}



func TestHpas(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sAutoscalingHpasGenerator{}), createHpas, k8s_client.TestOptions{})
}




