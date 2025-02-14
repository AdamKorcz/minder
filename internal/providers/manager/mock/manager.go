// Code generated by MockGen. DO NOT EDIT.
// Source: ./manager.go
//
// Generated by this command:
//
//	mockgen -package mock_manager -destination=./mock/manager.go -source=./manager.go
//

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	json "encoding/json"
	reflect "reflect"

	uuid "github.com/google/uuid"
	db "github.com/stacklok/minder/internal/db"
	v1 "github.com/stacklok/minder/pkg/providers/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockProviderManager is a mock of ProviderManager interface.
type MockProviderManager struct {
	ctrl     *gomock.Controller
	recorder *MockProviderManagerMockRecorder
}

// MockProviderManagerMockRecorder is the mock recorder for MockProviderManager.
type MockProviderManagerMockRecorder struct {
	mock *MockProviderManager
}

// NewMockProviderManager creates a new mock instance.
func NewMockProviderManager(ctrl *gomock.Controller) *MockProviderManager {
	mock := &MockProviderManager{ctrl: ctrl}
	mock.recorder = &MockProviderManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderManager) EXPECT() *MockProviderManagerMockRecorder {
	return m.recorder
}

// BulkInstantiateByTrait mocks base method.
func (m *MockProviderManager) BulkInstantiateByTrait(ctx context.Context, projectID uuid.UUID, trait db.ProviderType, name string) (map[string]v1.Provider, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkInstantiateByTrait", ctx, projectID, trait, name)
	ret0, _ := ret[0].(map[string]v1.Provider)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BulkInstantiateByTrait indicates an expected call of BulkInstantiateByTrait.
func (mr *MockProviderManagerMockRecorder) BulkInstantiateByTrait(ctx, projectID, trait, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkInstantiateByTrait", reflect.TypeOf((*MockProviderManager)(nil).BulkInstantiateByTrait), ctx, projectID, trait, name)
}

// CreateFromConfig mocks base method.
func (m *MockProviderManager) CreateFromConfig(ctx context.Context, providerClass db.ProviderClass, projectID uuid.UUID, name string, config json.RawMessage) (*db.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromConfig", ctx, providerClass, projectID, name, config)
	ret0, _ := ret[0].(*db.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFromConfig indicates an expected call of CreateFromConfig.
func (mr *MockProviderManagerMockRecorder) CreateFromConfig(ctx, providerClass, projectID, name, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromConfig", reflect.TypeOf((*MockProviderManager)(nil).CreateFromConfig), ctx, providerClass, projectID, name, config)
}

// DeleteByID mocks base method.
func (m *MockProviderManager) DeleteByID(ctx context.Context, providerID, projectID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, providerID, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockProviderManagerMockRecorder) DeleteByID(ctx, providerID, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockProviderManager)(nil).DeleteByID), ctx, providerID, projectID)
}

// DeleteByName mocks base method.
func (m *MockProviderManager) DeleteByName(ctx context.Context, name string, projectID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByName", ctx, name, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByName indicates an expected call of DeleteByName.
func (mr *MockProviderManagerMockRecorder) DeleteByName(ctx, name, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByName", reflect.TypeOf((*MockProviderManager)(nil).DeleteByName), ctx, name, projectID)
}

// InstantiateFromID mocks base method.
func (m *MockProviderManager) InstantiateFromID(ctx context.Context, providerID uuid.UUID) (v1.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstantiateFromID", ctx, providerID)
	ret0, _ := ret[0].(v1.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstantiateFromID indicates an expected call of InstantiateFromID.
func (mr *MockProviderManagerMockRecorder) InstantiateFromID(ctx, providerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstantiateFromID", reflect.TypeOf((*MockProviderManager)(nil).InstantiateFromID), ctx, providerID)
}

// InstantiateFromNameProject mocks base method.
func (m *MockProviderManager) InstantiateFromNameProject(ctx context.Context, name string, projectID uuid.UUID) (v1.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstantiateFromNameProject", ctx, name, projectID)
	ret0, _ := ret[0].(v1.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstantiateFromNameProject indicates an expected call of InstantiateFromNameProject.
func (mr *MockProviderManagerMockRecorder) InstantiateFromNameProject(ctx, name, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstantiateFromNameProject", reflect.TypeOf((*MockProviderManager)(nil).InstantiateFromNameProject), ctx, name, projectID)
}

// MockProviderClassManager is a mock of ProviderClassManager interface.
type MockProviderClassManager struct {
	ctrl     *gomock.Controller
	recorder *MockProviderClassManagerMockRecorder
}

// MockProviderClassManagerMockRecorder is the mock recorder for MockProviderClassManager.
type MockProviderClassManagerMockRecorder struct {
	mock *MockProviderClassManager
}

// NewMockProviderClassManager creates a new mock instance.
func NewMockProviderClassManager(ctrl *gomock.Controller) *MockProviderClassManager {
	mock := &MockProviderClassManager{ctrl: ctrl}
	mock.recorder = &MockProviderClassManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderClassManager) EXPECT() *MockProviderClassManagerMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockProviderClassManager) Build(ctx context.Context, config *db.Provider) (v1.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, config)
	ret0, _ := ret[0].(v1.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockProviderClassManagerMockRecorder) Build(ctx, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockProviderClassManager)(nil).Build), ctx, config)
}

// Delete mocks base method.
func (m *MockProviderClassManager) Delete(ctx context.Context, config *db.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProviderClassManagerMockRecorder) Delete(ctx, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProviderClassManager)(nil).Delete), ctx, config)
}

// GetConfig mocks base method.
func (m *MockProviderClassManager) GetConfig(ctx context.Context, class db.ProviderClass, userConfig json.RawMessage) (json.RawMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, class, userConfig)
	ret0, _ := ret[0].(json.RawMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockProviderClassManagerMockRecorder) GetConfig(ctx, class, userConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockProviderClassManager)(nil).GetConfig), ctx, class, userConfig)
}

// GetSupportedClasses mocks base method.
func (m *MockProviderClassManager) GetSupportedClasses() []db.ProviderClass {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedClasses")
	ret0, _ := ret[0].([]db.ProviderClass)
	return ret0
}

// GetSupportedClasses indicates an expected call of GetSupportedClasses.
func (mr *MockProviderClassManagerMockRecorder) GetSupportedClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedClasses", reflect.TypeOf((*MockProviderClassManager)(nil).GetSupportedClasses))
}
