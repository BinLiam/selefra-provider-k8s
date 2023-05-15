package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockBatchV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockBatchV1beta1InterfaceMockRecorder
}

type MockBatchV1beta1InterfaceMockRecorder struct {
	mock *MockBatchV1beta1Interface
}

func NewMockBatchV1beta1Interface(ctrl *gomock.Controller) *MockBatchV1beta1Interface {
	mock := &MockBatchV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockBatchV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockBatchV1beta1Interface) EXPECT() *MockBatchV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockBatchV1beta1Interface) CronJobs(arg0 string) v1beta1.CronJobInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CronJobs, arg0)
	ret0, _ := ret[0].(v1beta1.CronJobInterface)
	return ret0
}

func (mr *MockBatchV1beta1InterfaceMockRecorder) CronJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CronJobs, reflect.TypeOf((*MockBatchV1beta1Interface)(nil).CronJobs), arg0)
}

func (m *MockBatchV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockBatchV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockBatchV1beta1Interface)(nil).RESTClient))
}
