package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type MockInternalV1alpha1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockInternalV1alpha1InterfaceMockRecorder
}

type MockInternalV1alpha1InterfaceMockRecorder struct {
	mock *MockInternalV1alpha1Interface
}

func NewMockInternalV1alpha1Interface(ctrl *gomock.Controller) *MockInternalV1alpha1Interface {
	mock := &MockInternalV1alpha1Interface{ctrl: ctrl}
	mock.recorder = &MockInternalV1alpha1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockInternalV1alpha1Interface) EXPECT() *MockInternalV1alpha1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockInternalV1alpha1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockInternalV1alpha1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockInternalV1alpha1Interface)(nil).RESTClient))
}

func (m *MockInternalV1alpha1Interface) StorageVersions() v1alpha1.StorageVersionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StorageVersions)
	ret0, _ := ret[0].(v1alpha1.StorageVersionInterface)
	return ret0
}

func (mr *MockInternalV1alpha1InterfaceMockRecorder) StorageVersions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StorageVersions, reflect.TypeOf((*MockInternalV1alpha1Interface)(nil).StorageVersions))
}
