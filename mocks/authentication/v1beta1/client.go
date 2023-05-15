package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockAuthenticationV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationV1beta1InterfaceMockRecorder
}

type MockAuthenticationV1beta1InterfaceMockRecorder struct {
	mock *MockAuthenticationV1beta1Interface
}

func NewMockAuthenticationV1beta1Interface(ctrl *gomock.Controller) *MockAuthenticationV1beta1Interface {
	mock := &MockAuthenticationV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAuthenticationV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAuthenticationV1beta1Interface) EXPECT() *MockAuthenticationV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAuthenticationV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAuthenticationV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAuthenticationV1beta1Interface)(nil).RESTClient))
}

func (m *MockAuthenticationV1beta1Interface) TokenReviews() v1beta1.TokenReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.TokenReviews)
	ret0, _ := ret[0].(v1beta1.TokenReviewInterface)
	return ret0
}

func (mr *MockAuthenticationV1beta1InterfaceMockRecorder) TokenReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.TokenReviews, reflect.TypeOf((*MockAuthenticationV1beta1Interface)(nil).TokenReviews))
}
