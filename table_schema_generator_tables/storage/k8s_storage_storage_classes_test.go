package storage

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-k8s/faker"

	"github.com/selefra/selefra-provider-k8s/k8s_client"

	"github.com/selefra/selefra-provider-k8s/mocks"
	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/storage/v1"

	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/storage/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func createStorageClasses(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.StorageClass{}

	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)

	}

	resourceClient := resourcemock.NewMockStorageClassInterface(ctrl)

	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.StorageClassList{Items: []resource.StorageClass{r}}, nil,
	)

	serviceClient := resourcemock.NewMockStorageV1Interface(ctrl)

	serviceClient.EXPECT().StorageClasses().AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().StorageV1().AnyTimes().Return(serviceClient)

	return cl
}

func TestStorageClasses(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sStorageStorageClassesGenerator{}), createStorageClasses, k8s_client.TestOptions{})

}
