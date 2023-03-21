package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	rest "k8s.io/client-go/rest"
)

type MockCertificatesV1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockCertificatesV1InterfaceMockRecorder
}

type MockCertificatesV1InterfaceMockRecorder struct {
	mock *MockCertificatesV1Interface
}

func NewMockCertificatesV1Interface(ctrl *gomock.Controller) *MockCertificatesV1Interface {
	mock := &MockCertificatesV1Interface{ctrl: ctrl}
	mock.recorder = &MockCertificatesV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockCertificatesV1Interface) EXPECT() *MockCertificatesV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockCertificatesV1Interface) CertificateSigningRequests() v1.CertificateSigningRequestInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CertificateSigningRequests)
	ret0, _ := ret[0].(v1.CertificateSigningRequestInterface)
	return ret0
}

func (mr *MockCertificatesV1InterfaceMockRecorder) CertificateSigningRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CertificateSigningRequests, reflect.TypeOf((*MockCertificatesV1Interface)(nil).CertificateSigningRequests))
}

func (m *MockCertificatesV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockCertificatesV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockCertificatesV1Interface)(nil).RESTClient))
}
