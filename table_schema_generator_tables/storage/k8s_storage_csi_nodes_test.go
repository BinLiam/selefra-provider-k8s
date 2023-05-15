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

func createCsiNodes(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {

	r := resource.CSINode{}

	if err := faker.FakeObject(&r); err != nil {

		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockCSINodeInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(

		&resource.CSINodeList{Items: []resource.CSINode{r}}, nil,
	)

	serviceClient := resourcemock.NewMockStorageV1Interface(ctrl)

	serviceClient.EXPECT().CSINodes().AnyTimes().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().StorageV1().AnyTimes().Return(serviceClient)

	return cl

}

func TestCsiNodes(t *testing.T) {

	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sStorageCsiNodesGenerator{}), createCsiNodes, k8s_client.TestOptions{})

}
