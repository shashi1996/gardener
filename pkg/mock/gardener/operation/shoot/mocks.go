// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/operation/shoot (interfaces: ExtensionContainerRuntime,ExtensionWorker,ExtensionInfrastructure)

// Package shoot is a generated GoMock package.
package shoot

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	v1alpha10 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	shoot "github.com/gardener/gardener/pkg/operation/shoot"
	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// MockExtensionContainerRuntime is a mock of ExtensionContainerRuntime interface.
type MockExtensionContainerRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionContainerRuntimeMockRecorder
}

// MockExtensionContainerRuntimeMockRecorder is the mock recorder for MockExtensionContainerRuntime.
type MockExtensionContainerRuntimeMockRecorder struct {
	mock *MockExtensionContainerRuntime
}

// NewMockExtensionContainerRuntime creates a new mock instance.
func NewMockExtensionContainerRuntime(ctrl *gomock.Controller) *MockExtensionContainerRuntime {
	mock := &MockExtensionContainerRuntime{ctrl: ctrl}
	mock.recorder = &MockExtensionContainerRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtensionContainerRuntime) EXPECT() *MockExtensionContainerRuntimeMockRecorder {
	return m.recorder
}

// DeleteStaleResources mocks base method.
func (m *MockExtensionContainerRuntime) DeleteStaleResources(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStaleResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStaleResources indicates an expected call of DeleteStaleResources.
func (mr *MockExtensionContainerRuntimeMockRecorder) DeleteStaleResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStaleResources", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).DeleteStaleResources), arg0)
}

// Deploy mocks base method.
func (m *MockExtensionContainerRuntime) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockExtensionContainerRuntimeMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockExtensionContainerRuntime) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockExtensionContainerRuntimeMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).Destroy), arg0)
}

// Migrate mocks base method.
func (m *MockExtensionContainerRuntime) Migrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockExtensionContainerRuntimeMockRecorder) Migrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).Migrate), arg0)
}

// Restore mocks base method.
func (m *MockExtensionContainerRuntime) Restore(arg0 context.Context, arg1 *v1alpha1.ShootState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockExtensionContainerRuntimeMockRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).Restore), arg0, arg1)
}

// Wait mocks base method.
func (m *MockExtensionContainerRuntime) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockExtensionContainerRuntimeMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockExtensionContainerRuntime) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockExtensionContainerRuntimeMockRecorder) WaitCleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).WaitCleanup), arg0)
}

// WaitMigrate mocks base method.
func (m *MockExtensionContainerRuntime) WaitMigrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitMigrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitMigrate indicates an expected call of WaitMigrate.
func (mr *MockExtensionContainerRuntimeMockRecorder) WaitMigrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitMigrate", reflect.TypeOf((*MockExtensionContainerRuntime)(nil).WaitMigrate), arg0)
}

// MockExtensionWorker is a mock of ExtensionWorker interface.
type MockExtensionWorker struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionWorkerMockRecorder
}

// MockExtensionWorkerMockRecorder is the mock recorder for MockExtensionWorker.
type MockExtensionWorkerMockRecorder struct {
	mock *MockExtensionWorker
}

// NewMockExtensionWorker creates a new mock instance.
func NewMockExtensionWorker(ctrl *gomock.Controller) *MockExtensionWorker {
	mock := &MockExtensionWorker{ctrl: ctrl}
	mock.recorder = &MockExtensionWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtensionWorker) EXPECT() *MockExtensionWorkerMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockExtensionWorker) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockExtensionWorkerMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockExtensionWorker)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockExtensionWorker) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockExtensionWorkerMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockExtensionWorker)(nil).Destroy), arg0)
}

// MachineDeployments mocks base method.
func (m *MockExtensionWorker) MachineDeployments() []v1alpha10.MachineDeployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineDeployments")
	ret0, _ := ret[0].([]v1alpha10.MachineDeployment)
	return ret0
}

// MachineDeployments indicates an expected call of MachineDeployments.
func (mr *MockExtensionWorkerMockRecorder) MachineDeployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineDeployments", reflect.TypeOf((*MockExtensionWorker)(nil).MachineDeployments))
}

// Migrate mocks base method.
func (m *MockExtensionWorker) Migrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockExtensionWorkerMockRecorder) Migrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockExtensionWorker)(nil).Migrate), arg0)
}

// Restore mocks base method.
func (m *MockExtensionWorker) Restore(arg0 context.Context, arg1 *v1alpha1.ShootState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockExtensionWorkerMockRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockExtensionWorker)(nil).Restore), arg0, arg1)
}

// SetInfrastructureProviderStatus mocks base method.
func (m *MockExtensionWorker) SetInfrastructureProviderStatus(arg0 *runtime.RawExtension) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInfrastructureProviderStatus", arg0)
}

// SetInfrastructureProviderStatus indicates an expected call of SetInfrastructureProviderStatus.
func (mr *MockExtensionWorkerMockRecorder) SetInfrastructureProviderStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInfrastructureProviderStatus", reflect.TypeOf((*MockExtensionWorker)(nil).SetInfrastructureProviderStatus), arg0)
}

// SetOperatingSystemConfigMaps mocks base method.
func (m *MockExtensionWorker) SetOperatingSystemConfigMaps(arg0 map[string]shoot.OperatingSystemConfigs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOperatingSystemConfigMaps", arg0)
}

// SetOperatingSystemConfigMaps indicates an expected call of SetOperatingSystemConfigMaps.
func (mr *MockExtensionWorkerMockRecorder) SetOperatingSystemConfigMaps(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOperatingSystemConfigMaps", reflect.TypeOf((*MockExtensionWorker)(nil).SetOperatingSystemConfigMaps), arg0)
}

// SetSSHPublicKey mocks base method.
func (m *MockExtensionWorker) SetSSHPublicKey(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSSHPublicKey", arg0)
}

// SetSSHPublicKey indicates an expected call of SetSSHPublicKey.
func (mr *MockExtensionWorkerMockRecorder) SetSSHPublicKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSSHPublicKey", reflect.TypeOf((*MockExtensionWorker)(nil).SetSSHPublicKey), arg0)
}

// Wait mocks base method.
func (m *MockExtensionWorker) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockExtensionWorkerMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockExtensionWorker)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockExtensionWorker) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockExtensionWorkerMockRecorder) WaitCleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockExtensionWorker)(nil).WaitCleanup), arg0)
}

// WaitMigrate mocks base method.
func (m *MockExtensionWorker) WaitMigrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitMigrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitMigrate indicates an expected call of WaitMigrate.
func (mr *MockExtensionWorkerMockRecorder) WaitMigrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitMigrate", reflect.TypeOf((*MockExtensionWorker)(nil).WaitMigrate), arg0)
}

// MockExtensionInfrastructure is a mock of ExtensionInfrastructure interface.
type MockExtensionInfrastructure struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionInfrastructureMockRecorder
}

// MockExtensionInfrastructureMockRecorder is the mock recorder for MockExtensionInfrastructure.
type MockExtensionInfrastructureMockRecorder struct {
	mock *MockExtensionInfrastructure
}

// NewMockExtensionInfrastructure creates a new mock instance.
func NewMockExtensionInfrastructure(ctrl *gomock.Controller) *MockExtensionInfrastructure {
	mock := &MockExtensionInfrastructure{ctrl: ctrl}
	mock.recorder = &MockExtensionInfrastructureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtensionInfrastructure) EXPECT() *MockExtensionInfrastructureMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockExtensionInfrastructure) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockExtensionInfrastructureMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockExtensionInfrastructure)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockExtensionInfrastructure) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockExtensionInfrastructureMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockExtensionInfrastructure)(nil).Destroy), arg0)
}

// Migrate mocks base method.
func (m *MockExtensionInfrastructure) Migrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockExtensionInfrastructureMockRecorder) Migrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockExtensionInfrastructure)(nil).Migrate), arg0)
}

// NodesCIDR mocks base method.
func (m *MockExtensionInfrastructure) NodesCIDR() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodesCIDR")
	ret0, _ := ret[0].(*string)
	return ret0
}

// NodesCIDR indicates an expected call of NodesCIDR.
func (mr *MockExtensionInfrastructureMockRecorder) NodesCIDR() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodesCIDR", reflect.TypeOf((*MockExtensionInfrastructure)(nil).NodesCIDR))
}

// ProviderStatus mocks base method.
func (m *MockExtensionInfrastructure) ProviderStatus() *runtime.RawExtension {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderStatus")
	ret0, _ := ret[0].(*runtime.RawExtension)
	return ret0
}

// ProviderStatus indicates an expected call of ProviderStatus.
func (mr *MockExtensionInfrastructureMockRecorder) ProviderStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderStatus", reflect.TypeOf((*MockExtensionInfrastructure)(nil).ProviderStatus))
}

// Restore mocks base method.
func (m *MockExtensionInfrastructure) Restore(arg0 context.Context, arg1 *v1alpha1.ShootState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockExtensionInfrastructureMockRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockExtensionInfrastructure)(nil).Restore), arg0, arg1)
}

// SetSSHPublicKey mocks base method.
func (m *MockExtensionInfrastructure) SetSSHPublicKey(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSSHPublicKey", arg0)
}

// SetSSHPublicKey indicates an expected call of SetSSHPublicKey.
func (mr *MockExtensionInfrastructureMockRecorder) SetSSHPublicKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSSHPublicKey", reflect.TypeOf((*MockExtensionInfrastructure)(nil).SetSSHPublicKey), arg0)
}

// Wait mocks base method.
func (m *MockExtensionInfrastructure) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockExtensionInfrastructureMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockExtensionInfrastructure)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockExtensionInfrastructure) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockExtensionInfrastructureMockRecorder) WaitCleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockExtensionInfrastructure)(nil).WaitCleanup), arg0)
}

// WaitMigrate mocks base method.
func (m *MockExtensionInfrastructure) WaitMigrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitMigrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitMigrate indicates an expected call of WaitMigrate.
func (mr *MockExtensionInfrastructureMockRecorder) WaitMigrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitMigrate", reflect.TypeOf((*MockExtensionInfrastructure)(nil).WaitMigrate), arg0)
}