package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockCertificatesV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockCertificatesV1beta1InterfaceMockRecorder
}

type MockCertificatesV1beta1InterfaceMockRecorder struct {
	mock *MockCertificatesV1beta1Interface
}

func NewMockCertificatesV1beta1Interface(ctrl *gomock.Controller) *MockCertificatesV1beta1Interface {
	mock := &MockCertificatesV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockCertificatesV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockCertificatesV1beta1Interface) EXPECT() *MockCertificatesV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockCertificatesV1beta1Interface) CertificateSigningRequests() v1beta1.CertificateSigningRequestInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CertificateSigningRequests)
	ret0, _ := ret[0].(v1beta1.CertificateSigningRequestInterface)
	return ret0
}

func (mr *MockCertificatesV1beta1InterfaceMockRecorder) CertificateSigningRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CertificateSigningRequests, reflect.TypeOf((*MockCertificatesV1beta1Interface)(nil).CertificateSigningRequests))
}

func (m *MockCertificatesV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockCertificatesV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockCertificatesV1beta1Interface)(nil).RESTClient))
}
