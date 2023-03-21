package v1alpha1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type MockNodeV1alpha1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockNodeV1alpha1InterfaceMockRecorder
}

type MockNodeV1alpha1InterfaceMockRecorder struct {
	mock *MockNodeV1alpha1Interface
}

func NewMockNodeV1alpha1Interface(ctrl *gomock.Controller) *MockNodeV1alpha1Interface {
	mock := &MockNodeV1alpha1Interface{ctrl: ctrl}
	mock.recorder = &MockNodeV1alpha1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockNodeV1alpha1Interface) EXPECT() *MockNodeV1alpha1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockNodeV1alpha1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockNodeV1alpha1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockNodeV1alpha1Interface)(nil).RESTClient))
}

func (m *MockNodeV1alpha1Interface) RuntimeClasses() v1alpha1.RuntimeClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RuntimeClasses)
	ret0, _ := ret[0].(v1alpha1.RuntimeClassInterface)
	return ret0
}

func (mr *MockNodeV1alpha1InterfaceMockRecorder) RuntimeClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RuntimeClasses, reflect.TypeOf((*MockNodeV1alpha1Interface)(nil).RuntimeClasses))
}
