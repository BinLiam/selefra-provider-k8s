



package batch





import (
	"github.com/selefra/selefra-provider-k8s/constants"




	"testing"









	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-k8s/faker"


	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/mocks"
	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/batch/v1"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"




	resource "k8s.io/api/batch/v1"


	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"


	"k8s.io/client-go/kubernetes"


)





func createJobs(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {




	r := resource.Job{}




	if err := faker.FakeObject(&r); err != nil {


		t.Fatal(err)


	}



	r.Spec.Template = corev1.PodTemplateSpec{}

	resourceClient := resourcemock.NewMockJobInterface(ctrl)


	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.JobList{Items: []resource.Job{r}}, nil,


	)









	serviceClient := resourcemock.NewMockBatchV1Interface(ctrl)





	serviceClient.EXPECT().Jobs(constants.Constants_22).AnyTimes().Return(resourceClient)





	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().BatchV1().AnyTimes().Return(serviceClient)

	return cl


}







func TestJobs(t *testing.T) {




	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sBatchJobsGenerator{}), createJobs, k8s_client.TestOptions{})


}
