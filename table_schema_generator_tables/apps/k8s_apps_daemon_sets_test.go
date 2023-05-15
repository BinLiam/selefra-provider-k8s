package apps

import (
	"github.com/selefra/selefra-provider-k8s/constants"

	"testing"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-k8s/faker"

	"github.com/selefra/selefra-provider-k8s/k8s_client"

	mocks "github.com/selefra/selefra-provider-k8s/mocks"

	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/apps/v1"

	"github.com/selefra/selefra-provider-k8s/table_schema_generator"

	resource "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func createDaemonSets(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.DaemonSet{}

	if err := faker.FakeObject(&r); err != nil {

		t.Fatal(err)

	}

	r.Spec.Template = corev1.PodTemplateSpec{}
	r.Spec.UpdateStrategy = resource.DaemonSetUpdateStrategy{}

	resourceClient := resourcemock.NewMockDaemonSetInterface(ctrl)

	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(

		&resource.DaemonSetList{Items: []resource.DaemonSet{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAppsV1Interface(ctrl)

	serviceClient.EXPECT().DaemonSets(constants.Constants_16).AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AppsV1().AnyTimes().Return(serviceClient)

	return cl
}

func TestDaemonSets(t *testing.T) {

	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sAppsDaemonSetsGenerator{}), createDaemonSets, k8s_client.TestOptions{})
}
