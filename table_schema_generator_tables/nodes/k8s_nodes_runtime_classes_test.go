

package nodes



import (


	"testing"





	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-k8s/faker"




	"github.com/selefra/selefra-provider-k8s/k8s_client"




	"github.com/selefra/selefra-provider-k8s/mocks"


	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/node/v1"


	"github.com/selefra/selefra-provider-k8s/table_schema_generator"


	resource "k8s.io/api/node/v1"




	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"




	"k8s.io/client-go/kubernetes"
)







func createRuntimeClasses(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {




	r := resource.RuntimeClass{}


	if err := faker.FakeObject(&r); err != nil {




		t.Fatal(err)




	}









	resourceClient := resourcemock.NewMockRuntimeClassInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(




		&resource.RuntimeClassList{Items: []resource.RuntimeClass{r}}, nil,


	)





	serviceClient := resourcemock.NewMockNodeV1Interface(ctrl)







	serviceClient.EXPECT().RuntimeClasses().AnyTimes().Return(resourceClient)



	cl := mocks.NewMockInterface(ctrl)




	cl.EXPECT().NodeV1().AnyTimes().Return(serviceClient)



	return cl




}

func TestRuntimeClasses(t *testing.T) {


	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sNodesRuntimeClassesGenerator{}), createRuntimeClasses, k8s_client.TestOptions{})
}




