package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	rest "k8s.io/client-go/rest"
)

type MockAuthenticationV1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockAuthenticationV1InterfaceMockRecorder
}

type MockAuthenticationV1InterfaceMockRecorder struct {
	mock *MockAuthenticationV1Interface
}

func NewMockAuthenticationV1Interface(ctrl *gomock.Controller) *MockAuthenticationV1Interface {
	mock := &MockAuthenticationV1Interface{ctrl: ctrl}
	mock.recorder = &MockAuthenticationV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAuthenticationV1Interface) EXPECT() *MockAuthenticationV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAuthenticationV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAuthenticationV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAuthenticationV1Interface)(nil).RESTClient))
}

func (m *MockAuthenticationV1Interface) TokenReviews() v1.TokenReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.TokenReviews)
	ret0, _ := ret[0].(v1.TokenReviewInterface)
	return ret0
}

func (mr *MockAuthenticationV1InterfaceMockRecorder) TokenReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.TokenReviews, reflect.TypeOf((*MockAuthenticationV1Interface)(nil).TokenReviews))
}
