

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

func createNetworkPolicies(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {




	r := resource.NetworkPolicy{}




	if err := faker.FakeObject(&r); err != nil {


		t.Fatal(err)
	}





	r.Spec.Ingress = []resource.NetworkPolicyIngressRule{}


	r.Spec.Egress = []resource.NetworkPolicyEgressRule{}



	resourceClient := resourcemock.NewMockNetworkPolicyInterface(ctrl)


	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.NetworkPolicyList{Items: []resource.NetworkPolicy{r}}, nil,


	)

	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)





	serviceClient.EXPECT().NetworkPolicies(constants.Constants_46).AnyTimes().Return(resourceClient)







	cl := mocks.NewMockInterface(ctrl)


	cl.EXPECT().NetworkingV1().AnyTimes().Return(serviceClient)







	return cl
}







func TestNetworkPolicies(t *testing.T) {




	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sNetworkingNetworkPoliciesGenerator{}), createNetworkPolicies, k8s_client.TestOptions{})


}




