package v1alpha1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type MockNetworkingV1alpha1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockNetworkingV1alpha1InterfaceMockRecorder
}

type MockNetworkingV1alpha1InterfaceMockRecorder struct {
	mock *MockNetworkingV1alpha1Interface
}

func NewMockNetworkingV1alpha1Interface(ctrl *gomock.Controller) *MockNetworkingV1alpha1Interface {
	mock := &MockNetworkingV1alpha1Interface{ctrl: ctrl}
	mock.recorder = &MockNetworkingV1alpha1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockNetworkingV1alpha1Interface) EXPECT() *MockNetworkingV1alpha1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockNetworkingV1alpha1Interface) ClusterCIDRs() v1alpha1.ClusterCIDRInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ClusterCIDRs)
	ret0, _ := ret[0].(v1alpha1.ClusterCIDRInterface)
	return ret0
}

func (mr *MockNetworkingV1alpha1InterfaceMockRecorder) ClusterCIDRs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ClusterCIDRs, reflect.TypeOf((*MockNetworkingV1alpha1Interface)(nil).ClusterCIDRs))
}

func (m *MockNetworkingV1alpha1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockNetworkingV1alpha1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockNetworkingV1alpha1Interface)(nil).RESTClient))
}
