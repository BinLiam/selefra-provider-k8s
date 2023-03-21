package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	discovery "k8s.io/client-go/discovery"
	v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	v1alpha10 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	v10 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v1beta10 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	v11 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	v1alpha11 "k8s.io/client-go/kubernetes/typed/authentication/v1alpha1"
	v1beta11 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	v12 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	v1beta12 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	v13 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	v2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	v2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	v2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	v14 "k8s.io/client-go/kubernetes/typed/batch/v1"
	v1beta13 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	v15 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	v1beta14 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	v16 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	v1beta15 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	v17 "k8s.io/client-go/kubernetes/typed/core/v1"
	v18 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	v1beta16 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	v19 "k8s.io/client-go/kubernetes/typed/events/v1"
	v1beta17 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	v1beta18 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	v1alpha12 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	v1beta19 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	v1beta20 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	v1beta3 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta3"
	v110 "k8s.io/client-go/kubernetes/typed/networking/v1"
	v1alpha13 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	v1beta110 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	v111 "k8s.io/client-go/kubernetes/typed/node/v1"
	v1alpha14 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	v1beta111 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	v112 "k8s.io/client-go/kubernetes/typed/policy/v1"
	v1beta112 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	v113 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	v1alpha15 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	v1beta113 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	v1alpha16 "k8s.io/client-go/kubernetes/typed/resource/v1alpha1"
	v114 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	v1alpha17 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	v1beta114 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	v115 "k8s.io/client-go/kubernetes/typed/storage/v1"
	v1alpha18 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	v1beta115 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

type MockInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockInterfaceMockRecorder
}

type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

func (m *MockInterface) AdmissionregistrationV1() v1.AdmissionregistrationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AdmissionregistrationV)
	ret0, _ := ret[0].(v1.AdmissionregistrationV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AdmissionregistrationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AdmissionregistrationV, reflect.TypeOf((*MockInterface)(nil).AdmissionregistrationV1))
}

func (m *MockInterface) AdmissionregistrationV1alpha1() v1alpha1.AdmissionregistrationV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AdmissionregistrationValpha)
	ret0, _ := ret[0].(v1alpha1.AdmissionregistrationV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AdmissionregistrationV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AdmissionregistrationValpha, reflect.TypeOf((*MockInterface)(nil).AdmissionregistrationV1alpha1))
}

func (m *MockInterface) AdmissionregistrationV1beta1() v1beta1.AdmissionregistrationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AdmissionregistrationVbeta)
	ret0, _ := ret[0].(v1beta1.AdmissionregistrationV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AdmissionregistrationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AdmissionregistrationVbeta, reflect.TypeOf((*MockInterface)(nil).AdmissionregistrationV1beta1))
}

func (m *MockInterface) AppsV1() v10.AppsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AppsV)
	ret0, _ := ret[0].(v10.AppsV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AppsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AppsV, reflect.TypeOf((*MockInterface)(nil).AppsV1))
}

func (m *MockInterface) AppsV1beta1() v1beta10.AppsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AppsVbeta)
	ret0, _ := ret[0].(v1beta10.AppsV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AppsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AppsVbeta, reflect.TypeOf((*MockInterface)(nil).AppsV1beta1))
}

func (m *MockInterface) AppsV1beta2() v1beta2.AppsV1beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1beta2")
	ret0, _ := ret[0].(v1beta2.AppsV1beta2Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AppsV1beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1beta2", reflect.TypeOf((*MockInterface)(nil).AppsV1beta2))
}

func (m *MockInterface) AuthenticationV1() v11.AuthenticationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AuthenticationV)
	ret0, _ := ret[0].(v11.AuthenticationV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AuthenticationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AuthenticationV, reflect.TypeOf((*MockInterface)(nil).AuthenticationV1))
}

func (m *MockInterface) AuthenticationV1alpha1() v1alpha11.AuthenticationV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AuthenticationValpha)
	ret0, _ := ret[0].(v1alpha11.AuthenticationV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AuthenticationV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AuthenticationValpha, reflect.TypeOf((*MockInterface)(nil).AuthenticationV1alpha1))
}

func (m *MockInterface) AuthenticationV1beta1() v1beta11.AuthenticationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AuthenticationVbeta)
	ret0, _ := ret[0].(v1beta11.AuthenticationV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AuthenticationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AuthenticationVbeta, reflect.TypeOf((*MockInterface)(nil).AuthenticationV1beta1))
}

func (m *MockInterface) AuthorizationV1() v12.AuthorizationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AuthorizationV)
	ret0, _ := ret[0].(v12.AuthorizationV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AuthorizationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AuthorizationV, reflect.TypeOf((*MockInterface)(nil).AuthorizationV1))
}

func (m *MockInterface) AuthorizationV1beta1() v1beta12.AuthorizationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AuthorizationVbeta)
	ret0, _ := ret[0].(v1beta12.AuthorizationV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AuthorizationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AuthorizationVbeta, reflect.TypeOf((*MockInterface)(nil).AuthorizationV1beta1))
}

func (m *MockInterface) AutoscalingV1() v13.AutoscalingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AutoscalingV)
	ret0, _ := ret[0].(v13.AutoscalingV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AutoscalingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AutoscalingV, reflect.TypeOf((*MockInterface)(nil).AutoscalingV1))
}

func (m *MockInterface) AutoscalingV2() v2.AutoscalingV2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2")
	ret0, _ := ret[0].(v2.AutoscalingV2Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AutoscalingV2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2", reflect.TypeOf((*MockInterface)(nil).AutoscalingV2))
}

func (m *MockInterface) AutoscalingV2beta1() v2beta1.AutoscalingV2beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.AutoscalingVbeta)
	ret0, _ := ret[0].(v2beta1.AutoscalingV2beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AutoscalingV2beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.AutoscalingVbeta, reflect.TypeOf((*MockInterface)(nil).AutoscalingV2beta1))
}

func (m *MockInterface) AutoscalingV2beta2() v2beta2.AutoscalingV2beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2beta2")
	ret0, _ := ret[0].(v2beta2.AutoscalingV2beta2Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) AutoscalingV2beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2beta2", reflect.TypeOf((*MockInterface)(nil).AutoscalingV2beta2))
}

func (m *MockInterface) BatchV1() v14.BatchV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.BatchV)
	ret0, _ := ret[0].(v14.BatchV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) BatchV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchV, reflect.TypeOf((*MockInterface)(nil).BatchV1))
}

func (m *MockInterface) BatchV1beta1() v1beta13.BatchV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.BatchVbeta)
	ret0, _ := ret[0].(v1beta13.BatchV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) BatchV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.BatchVbeta, reflect.TypeOf((*MockInterface)(nil).BatchV1beta1))
}

func (m *MockInterface) CertificatesV1() v15.CertificatesV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CertificatesV)
	ret0, _ := ret[0].(v15.CertificatesV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) CertificatesV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CertificatesV, reflect.TypeOf((*MockInterface)(nil).CertificatesV1))
}

func (m *MockInterface) CertificatesV1beta1() v1beta14.CertificatesV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CertificatesVbeta)
	ret0, _ := ret[0].(v1beta14.CertificatesV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) CertificatesV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CertificatesVbeta, reflect.TypeOf((*MockInterface)(nil).CertificatesV1beta1))
}

func (m *MockInterface) CoordinationV1() v16.CoordinationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CoordinationV)
	ret0, _ := ret[0].(v16.CoordinationV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) CoordinationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CoordinationV, reflect.TypeOf((*MockInterface)(nil).CoordinationV1))
}

func (m *MockInterface) CoordinationV1beta1() v1beta15.CoordinationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CoordinationVbeta)
	ret0, _ := ret[0].(v1beta15.CoordinationV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) CoordinationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CoordinationVbeta, reflect.TypeOf((*MockInterface)(nil).CoordinationV1beta1))
}

func (m *MockInterface) CoreV1() v17.CoreV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CoreV)
	ret0, _ := ret[0].(v17.CoreV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) CoreV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CoreV, reflect.TypeOf((*MockInterface)(nil).CoreV1))
}

func (m *MockInterface) Discovery() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Discovery)
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) Discovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Discovery, reflect.TypeOf((*MockInterface)(nil).Discovery))
}

func (m *MockInterface) DiscoveryV1() v18.DiscoveryV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DiscoveryV)
	ret0, _ := ret[0].(v18.DiscoveryV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) DiscoveryV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DiscoveryV, reflect.TypeOf((*MockInterface)(nil).DiscoveryV1))
}

func (m *MockInterface) DiscoveryV1beta1() v1beta16.DiscoveryV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DiscoveryVbeta)
	ret0, _ := ret[0].(v1beta16.DiscoveryV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) DiscoveryV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DiscoveryVbeta, reflect.TypeOf((*MockInterface)(nil).DiscoveryV1beta1))
}

func (m *MockInterface) EventsV1() v19.EventsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.EventsV)
	ret0, _ := ret[0].(v19.EventsV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) EventsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.EventsV, reflect.TypeOf((*MockInterface)(nil).EventsV1))
}

func (m *MockInterface) EventsV1beta1() v1beta17.EventsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.EventsVbeta)
	ret0, _ := ret[0].(v1beta17.EventsV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) EventsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.EventsVbeta, reflect.TypeOf((*MockInterface)(nil).EventsV1beta1))
}

func (m *MockInterface) ExtensionsV1beta1() v1beta18.ExtensionsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ExtensionsVbeta)
	ret0, _ := ret[0].(v1beta18.ExtensionsV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) ExtensionsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ExtensionsVbeta, reflect.TypeOf((*MockInterface)(nil).ExtensionsV1beta1))
}

func (m *MockInterface) FlowcontrolV1alpha1() v1alpha12.FlowcontrolV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.FlowcontrolValpha)
	ret0, _ := ret[0].(v1alpha12.FlowcontrolV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) FlowcontrolV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.FlowcontrolValpha, reflect.TypeOf((*MockInterface)(nil).FlowcontrolV1alpha1))
}

func (m *MockInterface) FlowcontrolV1beta1() v1beta19.FlowcontrolV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.FlowcontrolVbeta)
	ret0, _ := ret[0].(v1beta19.FlowcontrolV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) FlowcontrolV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.FlowcontrolVbeta, reflect.TypeOf((*MockInterface)(nil).FlowcontrolV1beta1))
}

func (m *MockInterface) FlowcontrolV1beta2() v1beta20.FlowcontrolV1beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowcontrolV1beta2")
	ret0, _ := ret[0].(v1beta20.FlowcontrolV1beta2Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) FlowcontrolV1beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowcontrolV1beta2", reflect.TypeOf((*MockInterface)(nil).FlowcontrolV1beta2))
}

func (m *MockInterface) FlowcontrolV1beta3() v1beta3.FlowcontrolV1beta3Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowcontrolV1beta3")
	ret0, _ := ret[0].(v1beta3.FlowcontrolV1beta3Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) FlowcontrolV1beta3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowcontrolV1beta3", reflect.TypeOf((*MockInterface)(nil).FlowcontrolV1beta3))
}

func (m *MockInterface) InternalV1alpha1() v1alpha10.InternalV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.InternalValpha)
	ret0, _ := ret[0].(v1alpha10.InternalV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) InternalV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.InternalValpha, reflect.TypeOf((*MockInterface)(nil).InternalV1alpha1))
}

func (m *MockInterface) NetworkingV1() v110.NetworkingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NetworkingV)
	ret0, _ := ret[0].(v110.NetworkingV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) NetworkingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NetworkingV, reflect.TypeOf((*MockInterface)(nil).NetworkingV1))
}

func (m *MockInterface) NetworkingV1alpha1() v1alpha13.NetworkingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NetworkingValpha)
	ret0, _ := ret[0].(v1alpha13.NetworkingV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) NetworkingV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NetworkingValpha, reflect.TypeOf((*MockInterface)(nil).NetworkingV1alpha1))
}

func (m *MockInterface) NetworkingV1beta1() v1beta110.NetworkingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NetworkingVbeta)
	ret0, _ := ret[0].(v1beta110.NetworkingV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) NetworkingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NetworkingVbeta, reflect.TypeOf((*MockInterface)(nil).NetworkingV1beta1))
}

func (m *MockInterface) NodeV1() v111.NodeV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NodeV)
	ret0, _ := ret[0].(v111.NodeV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) NodeV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NodeV, reflect.TypeOf((*MockInterface)(nil).NodeV1))
}

func (m *MockInterface) NodeV1alpha1() v1alpha14.NodeV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NodeValpha)
	ret0, _ := ret[0].(v1alpha14.NodeV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) NodeV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NodeValpha, reflect.TypeOf((*MockInterface)(nil).NodeV1alpha1))
}

func (m *MockInterface) NodeV1beta1() v1beta111.NodeV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NodeVbeta)
	ret0, _ := ret[0].(v1beta111.NodeV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) NodeV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NodeVbeta, reflect.TypeOf((*MockInterface)(nil).NodeV1beta1))
}

func (m *MockInterface) PolicyV1() v112.PolicyV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PolicyV)
	ret0, _ := ret[0].(v112.PolicyV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) PolicyV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PolicyV, reflect.TypeOf((*MockInterface)(nil).PolicyV1))
}

func (m *MockInterface) PolicyV1beta1() v1beta112.PolicyV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PolicyVbeta)
	ret0, _ := ret[0].(v1beta112.PolicyV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) PolicyV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PolicyVbeta, reflect.TypeOf((*MockInterface)(nil).PolicyV1beta1))
}

func (m *MockInterface) RbacV1() v113.RbacV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RbacV)
	ret0, _ := ret[0].(v113.RbacV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) RbacV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RbacV, reflect.TypeOf((*MockInterface)(nil).RbacV1))
}

func (m *MockInterface) RbacV1alpha1() v1alpha15.RbacV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RbacValpha)
	ret0, _ := ret[0].(v1alpha15.RbacV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) RbacV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RbacValpha, reflect.TypeOf((*MockInterface)(nil).RbacV1alpha1))
}

func (m *MockInterface) RbacV1beta1() v1beta113.RbacV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RbacVbeta)
	ret0, _ := ret[0].(v1beta113.RbacV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) RbacV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RbacVbeta, reflect.TypeOf((*MockInterface)(nil).RbacV1beta1))
}

func (m *MockInterface) ResourceV1alpha1() v1alpha16.ResourceV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ResourceValpha)
	ret0, _ := ret[0].(v1alpha16.ResourceV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) ResourceV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ResourceValpha, reflect.TypeOf((*MockInterface)(nil).ResourceV1alpha1))
}

func (m *MockInterface) SchedulingV1() v114.SchedulingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SchedulingV)
	ret0, _ := ret[0].(v114.SchedulingV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) SchedulingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SchedulingV, reflect.TypeOf((*MockInterface)(nil).SchedulingV1))
}

func (m *MockInterface) SchedulingV1alpha1() v1alpha17.SchedulingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SchedulingValpha)
	ret0, _ := ret[0].(v1alpha17.SchedulingV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) SchedulingV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SchedulingValpha, reflect.TypeOf((*MockInterface)(nil).SchedulingV1alpha1))
}

func (m *MockInterface) SchedulingV1beta1() v1beta114.SchedulingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SchedulingVbeta)
	ret0, _ := ret[0].(v1beta114.SchedulingV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) SchedulingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SchedulingVbeta, reflect.TypeOf((*MockInterface)(nil).SchedulingV1beta1))
}

func (m *MockInterface) StorageV1() v115.StorageV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StorageV)
	ret0, _ := ret[0].(v115.StorageV1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) StorageV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StorageV, reflect.TypeOf((*MockInterface)(nil).StorageV1))
}

func (m *MockInterface) StorageV1alpha1() v1alpha18.StorageV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StorageValpha)
	ret0, _ := ret[0].(v1alpha18.StorageV1alpha1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) StorageV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StorageValpha, reflect.TypeOf((*MockInterface)(nil).StorageV1alpha1))
}

func (m *MockInterface) StorageV1beta1() v1beta115.StorageV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StorageVbeta)
	ret0, _ := ret[0].(v1beta115.StorageV1beta1Interface)
	return ret0
}

func (mr *MockInterfaceMockRecorder) StorageV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StorageVbeta, reflect.TypeOf((*MockInterface)(nil).StorageV1beta1))
}
