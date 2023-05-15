package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	rest "k8s.io/client-go/rest"
)

type MockBatchV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockBatchV1InterfaceMockRecorder
}

type MockBatchV1InterfaceMockRecorder struct {
	mock *MockBatchV1Interface
}

func NewMockBatchV1Interface(ctrl *gomock.Controller) *MockBatchV1Interface {
	mock := &MockBatchV1Interface{ctrl: ctrl}
	mock.recorder = &MockBatchV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockBatchV1Interface) EXPECT() *MockBatchV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockBatchV1Interface) CronJobs(arg0 string) v1.CronJobInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CronJobs, arg0)
	ret0, _ := ret[0].(v1.CronJobInterface)
	return ret0
}

func (mr *MockBatchV1InterfaceMockRecorder) CronJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CronJobs, reflect.TypeOf((*MockBatchV1Interface)(nil).CronJobs), arg0)
}

func (m *MockBatchV1Interface) Jobs(arg0 string) v1.JobInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Jobs, arg0)
	ret0, _ := ret[0].(v1.JobInterface)
	return ret0
}

func (mr *MockBatchV1InterfaceMockRecorder) Jobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Jobs, reflect.TypeOf((*MockBatchV1Interface)(nil).Jobs), arg0)
}

func (m *MockBatchV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockBatchV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockBatchV1Interface)(nil).RESTClient))
}
