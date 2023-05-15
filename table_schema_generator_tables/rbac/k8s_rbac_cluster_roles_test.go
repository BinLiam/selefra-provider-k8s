package rbac

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8s_client"

	"github.com/selefra/selefra-provider-k8s/mocks"

	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/rbac/v1"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/rbac/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func createClusterRoles(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.ClusterRole{}

	if err := faker.FakeObject(&r); err != nil {

		t.Fatal(err)

	}

	resourceClient := resourcemock.NewMockClusterRoleInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.ClusterRoleList{Items: []resource.ClusterRole{r}}, nil,
	)

	serviceClient := resourcemock.NewMockRbacV1Interface(ctrl)

	serviceClient.EXPECT().ClusterRoles().AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)

	cl.EXPECT().RbacV1().AnyTimes().Return(serviceClient)

	return cl

}

func TestClusterRoles(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sRbacClusterRolesGenerator{}), createClusterRoles, k8s_client.TestOptions{})
}
