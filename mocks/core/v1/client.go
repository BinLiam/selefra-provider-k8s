package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
)

type MockCoreV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockCoreV1InterfaceMockRecorder
}

type MockCoreV1InterfaceMockRecorder struct {
	mock *MockCoreV1Interface
}

func NewMockCoreV1Interface(ctrl *gomock.Controller) *MockCoreV1Interface {
	mock := &MockCoreV1Interface{ctrl: ctrl}
	mock.recorder = &MockCoreV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockCoreV1Interface) EXPECT() *MockCoreV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockCoreV1Interface) ComponentStatuses() v1.ComponentStatusInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ComponentStatuses)
	ret0, _ := ret[0].(v1.ComponentStatusInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) ComponentStatuses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ComponentStatuses, reflect.TypeOf((*MockCoreV1Interface)(nil).ComponentStatuses))
}

func (m *MockCoreV1Interface) ConfigMaps(arg0 string) v1.ConfigMapInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ConfigMaps, arg0)
	ret0, _ := ret[0].(v1.ConfigMapInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) ConfigMaps(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ConfigMaps, reflect.TypeOf((*MockCoreV1Interface)(nil).ConfigMaps), arg0)
}

func (m *MockCoreV1Interface) Endpoints(arg0 string) v1.EndpointsInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Endpoints, arg0)
	ret0, _ := ret[0].(v1.EndpointsInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) Endpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Endpoints, reflect.TypeOf((*MockCoreV1Interface)(nil).Endpoints), arg0)
}

func (m *MockCoreV1Interface) Events(arg0 string) v1.EventInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Events, arg0)
	ret0, _ := ret[0].(v1.EventInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) Events(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Events, reflect.TypeOf((*MockCoreV1Interface)(nil).Events), arg0)
}

func (m *MockCoreV1Interface) LimitRanges(arg0 string) v1.LimitRangeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.LimitRanges, arg0)
	ret0, _ := ret[0].(v1.LimitRangeInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) LimitRanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.LimitRanges, reflect.TypeOf((*MockCoreV1Interface)(nil).LimitRanges), arg0)
}

func (m *MockCoreV1Interface) Namespaces() v1.NamespaceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Namespaces)
	ret0, _ := ret[0].(v1.NamespaceInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) Namespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Namespaces, reflect.TypeOf((*MockCoreV1Interface)(nil).Namespaces))
}

func (m *MockCoreV1Interface) Nodes() v1.NodeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Nodes)
	ret0, _ := ret[0].(v1.NodeInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Nodes, reflect.TypeOf((*MockCoreV1Interface)(nil).Nodes))
}

func (m *MockCoreV1Interface) PersistentVolumeClaims(arg0 string) v1.PersistentVolumeClaimInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PersistentVolumeClaims, arg0)
	ret0, _ := ret[0].(v1.PersistentVolumeClaimInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) PersistentVolumeClaims(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PersistentVolumeClaims, reflect.TypeOf((*MockCoreV1Interface)(nil).PersistentVolumeClaims), arg0)
}

func (m *MockCoreV1Interface) PersistentVolumes() v1.PersistentVolumeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PersistentVolumes)
	ret0, _ := ret[0].(v1.PersistentVolumeInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) PersistentVolumes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PersistentVolumes, reflect.TypeOf((*MockCoreV1Interface)(nil).PersistentVolumes))
}

func (m *MockCoreV1Interface) PodTemplates(arg0 string) v1.PodTemplateInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PodTemplates, arg0)
	ret0, _ := ret[0].(v1.PodTemplateInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) PodTemplates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PodTemplates, reflect.TypeOf((*MockCoreV1Interface)(nil).PodTemplates), arg0)
}

func (m *MockCoreV1Interface) Pods(arg0 string) v1.PodInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Pods, arg0)
	ret0, _ := ret[0].(v1.PodInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) Pods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Pods, reflect.TypeOf((*MockCoreV1Interface)(nil).Pods), arg0)
}

func (m *MockCoreV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockCoreV1Interface)(nil).RESTClient))
}

func (m *MockCoreV1Interface) ReplicationControllers(arg0 string) v1.ReplicationControllerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ReplicationControllers, arg0)
	ret0, _ := ret[0].(v1.ReplicationControllerInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) ReplicationControllers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ReplicationControllers, reflect.TypeOf((*MockCoreV1Interface)(nil).ReplicationControllers), arg0)
}

func (m *MockCoreV1Interface) ResourceQuotas(arg0 string) v1.ResourceQuotaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ResourceQuotas, arg0)
	ret0, _ := ret[0].(v1.ResourceQuotaInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) ResourceQuotas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ResourceQuotas, reflect.TypeOf((*MockCoreV1Interface)(nil).ResourceQuotas), arg0)
}

func (m *MockCoreV1Interface) Secrets(arg0 string) v1.SecretInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Secrets, arg0)
	ret0, _ := ret[0].(v1.SecretInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) Secrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Secrets, reflect.TypeOf((*MockCoreV1Interface)(nil).Secrets), arg0)
}

func (m *MockCoreV1Interface) ServiceAccounts(arg0 string) v1.ServiceAccountInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServiceAccounts, arg0)
	ret0, _ := ret[0].(v1.ServiceAccountInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) ServiceAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServiceAccounts, reflect.TypeOf((*MockCoreV1Interface)(nil).ServiceAccounts), arg0)
}

func (m *MockCoreV1Interface) Services(arg0 string) v1.ServiceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Services, arg0)
	ret0, _ := ret[0].(v1.ServiceInterface)
	return ret0
}

func (mr *MockCoreV1InterfaceMockRecorder) Services(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Services, reflect.TypeOf((*MockCoreV1Interface)(nil).Services), arg0)
}
