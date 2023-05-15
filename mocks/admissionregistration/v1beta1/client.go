package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockAdmissionregistrationV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAdmissionregistrationV1beta1InterfaceMockRecorder
}

type MockAdmissionregistrationV1beta1InterfaceMockRecorder struct {
	mock *MockAdmissionregistrationV1beta1Interface
}

func NewMockAdmissionregistrationV1beta1Interface(ctrl *gomock.Controller) *MockAdmissionregistrationV1beta1Interface {
	mock := &MockAdmissionregistrationV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAdmissionregistrationV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAdmissionregistrationV1beta1Interface) EXPECT() *MockAdmissionregistrationV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAdmissionregistrationV1beta1Interface) MutatingWebhookConfigurations() v1beta1.MutatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.MutatingWebhookConfigurations)
	ret0, _ := ret[0].(v1beta1.MutatingWebhookConfigurationInterface)
	return ret0
}

func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) MutatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.MutatingWebhookConfigurations, reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).MutatingWebhookConfigurations))
}

func (m *MockAdmissionregistrationV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).RESTClient))
}

func (m *MockAdmissionregistrationV1beta1Interface) ValidatingWebhookConfigurations() v1beta1.ValidatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ValidatingWebhookConfigurations)
	ret0, _ := ret[0].(v1beta1.ValidatingWebhookConfigurationInterface)
	return ret0
}

func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) ValidatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ValidatingWebhookConfigurations, reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).ValidatingWebhookConfigurations))
}
