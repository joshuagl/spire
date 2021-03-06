// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/spire/server/datastore (interfaces: DataStore,DataStoreServer)

// Package mock_datastore is a generated GoMock package.
package mock_datastore

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	datastore "github.com/spiffe/spire/proto/spire/server/datastore"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AppendBundle mocks base method
func (m *MockDataStore) AppendBundle(arg0 context.Context, arg1 *datastore.AppendBundleRequest) (*datastore.AppendBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.AppendBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendBundle indicates an expected call of AppendBundle
func (mr *MockDataStoreMockRecorder) AppendBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendBundle", reflect.TypeOf((*MockDataStore)(nil).AppendBundle), arg0, arg1)
}

// CreateAttestedNode mocks base method
func (m *MockDataStore) CreateAttestedNode(arg0 context.Context, arg1 *datastore.CreateAttestedNodeRequest) (*datastore.CreateAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttestedNode indicates an expected call of CreateAttestedNode
func (mr *MockDataStoreMockRecorder) CreateAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttestedNode", reflect.TypeOf((*MockDataStore)(nil).CreateAttestedNode), arg0, arg1)
}

// CreateBundle mocks base method
func (m *MockDataStore) CreateBundle(arg0 context.Context, arg1 *datastore.CreateBundleRequest) (*datastore.CreateBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBundle indicates an expected call of CreateBundle
func (mr *MockDataStoreMockRecorder) CreateBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBundle", reflect.TypeOf((*MockDataStore)(nil).CreateBundle), arg0, arg1)
}

// CreateJoinToken mocks base method
func (m *MockDataStore) CreateJoinToken(arg0 context.Context, arg1 *datastore.CreateJoinTokenRequest) (*datastore.CreateJoinTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJoinToken", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateJoinTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJoinToken indicates an expected call of CreateJoinToken
func (mr *MockDataStoreMockRecorder) CreateJoinToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJoinToken", reflect.TypeOf((*MockDataStore)(nil).CreateJoinToken), arg0, arg1)
}

// CreateRegistrationEntry mocks base method
func (m *MockDataStore) CreateRegistrationEntry(arg0 context.Context, arg1 *datastore.CreateRegistrationEntryRequest) (*datastore.CreateRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRegistrationEntry indicates an expected call of CreateRegistrationEntry
func (mr *MockDataStoreMockRecorder) CreateRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRegistrationEntry", reflect.TypeOf((*MockDataStore)(nil).CreateRegistrationEntry), arg0, arg1)
}

// DeleteAttestedNode mocks base method
func (m *MockDataStore) DeleteAttestedNode(arg0 context.Context, arg1 *datastore.DeleteAttestedNodeRequest) (*datastore.DeleteAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAttestedNode indicates an expected call of DeleteAttestedNode
func (mr *MockDataStoreMockRecorder) DeleteAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttestedNode", reflect.TypeOf((*MockDataStore)(nil).DeleteAttestedNode), arg0, arg1)
}

// DeleteBundle mocks base method
func (m *MockDataStore) DeleteBundle(arg0 context.Context, arg1 *datastore.DeleteBundleRequest) (*datastore.DeleteBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBundle indicates an expected call of DeleteBundle
func (mr *MockDataStoreMockRecorder) DeleteBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBundle", reflect.TypeOf((*MockDataStore)(nil).DeleteBundle), arg0, arg1)
}

// DeleteJoinToken mocks base method
func (m *MockDataStore) DeleteJoinToken(arg0 context.Context, arg1 *datastore.DeleteJoinTokenRequest) (*datastore.DeleteJoinTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJoinToken", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteJoinTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteJoinToken indicates an expected call of DeleteJoinToken
func (mr *MockDataStoreMockRecorder) DeleteJoinToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJoinToken", reflect.TypeOf((*MockDataStore)(nil).DeleteJoinToken), arg0, arg1)
}

// DeleteRegistrationEntry mocks base method
func (m *MockDataStore) DeleteRegistrationEntry(arg0 context.Context, arg1 *datastore.DeleteRegistrationEntryRequest) (*datastore.DeleteRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRegistrationEntry indicates an expected call of DeleteRegistrationEntry
func (mr *MockDataStoreMockRecorder) DeleteRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegistrationEntry", reflect.TypeOf((*MockDataStore)(nil).DeleteRegistrationEntry), arg0, arg1)
}

// FetchAttestedNode mocks base method
func (m *MockDataStore) FetchAttestedNode(arg0 context.Context, arg1 *datastore.FetchAttestedNodeRequest) (*datastore.FetchAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAttestedNode indicates an expected call of FetchAttestedNode
func (mr *MockDataStoreMockRecorder) FetchAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttestedNode", reflect.TypeOf((*MockDataStore)(nil).FetchAttestedNode), arg0, arg1)
}

// FetchBundle mocks base method
func (m *MockDataStore) FetchBundle(arg0 context.Context, arg1 *datastore.FetchBundleRequest) (*datastore.FetchBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBundle indicates an expected call of FetchBundle
func (mr *MockDataStoreMockRecorder) FetchBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBundle", reflect.TypeOf((*MockDataStore)(nil).FetchBundle), arg0, arg1)
}

// FetchJoinToken mocks base method
func (m *MockDataStore) FetchJoinToken(arg0 context.Context, arg1 *datastore.FetchJoinTokenRequest) (*datastore.FetchJoinTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJoinToken", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchJoinTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJoinToken indicates an expected call of FetchJoinToken
func (mr *MockDataStoreMockRecorder) FetchJoinToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJoinToken", reflect.TypeOf((*MockDataStore)(nil).FetchJoinToken), arg0, arg1)
}

// FetchRegistrationEntry mocks base method
func (m *MockDataStore) FetchRegistrationEntry(arg0 context.Context, arg1 *datastore.FetchRegistrationEntryRequest) (*datastore.FetchRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRegistrationEntry indicates an expected call of FetchRegistrationEntry
func (mr *MockDataStoreMockRecorder) FetchRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRegistrationEntry", reflect.TypeOf((*MockDataStore)(nil).FetchRegistrationEntry), arg0, arg1)
}

// GetNodeSelectors mocks base method
func (m *MockDataStore) GetNodeSelectors(arg0 context.Context, arg1 *datastore.GetNodeSelectorsRequest) (*datastore.GetNodeSelectorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeSelectors", arg0, arg1)
	ret0, _ := ret[0].(*datastore.GetNodeSelectorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeSelectors indicates an expected call of GetNodeSelectors
func (mr *MockDataStoreMockRecorder) GetNodeSelectors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeSelectors", reflect.TypeOf((*MockDataStore)(nil).GetNodeSelectors), arg0, arg1)
}

// ListAttestedNodes mocks base method
func (m *MockDataStore) ListAttestedNodes(arg0 context.Context, arg1 *datastore.ListAttestedNodesRequest) (*datastore.ListAttestedNodesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttestedNodes", arg0, arg1)
	ret0, _ := ret[0].(*datastore.ListAttestedNodesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttestedNodes indicates an expected call of ListAttestedNodes
func (mr *MockDataStoreMockRecorder) ListAttestedNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttestedNodes", reflect.TypeOf((*MockDataStore)(nil).ListAttestedNodes), arg0, arg1)
}

// ListBundles mocks base method
func (m *MockDataStore) ListBundles(arg0 context.Context, arg1 *datastore.ListBundlesRequest) (*datastore.ListBundlesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBundles", arg0, arg1)
	ret0, _ := ret[0].(*datastore.ListBundlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBundles indicates an expected call of ListBundles
func (mr *MockDataStoreMockRecorder) ListBundles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBundles", reflect.TypeOf((*MockDataStore)(nil).ListBundles), arg0, arg1)
}

// ListRegistrationEntries mocks base method
func (m *MockDataStore) ListRegistrationEntries(arg0 context.Context, arg1 *datastore.ListRegistrationEntriesRequest) (*datastore.ListRegistrationEntriesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRegistrationEntries", arg0, arg1)
	ret0, _ := ret[0].(*datastore.ListRegistrationEntriesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegistrationEntries indicates an expected call of ListRegistrationEntries
func (mr *MockDataStoreMockRecorder) ListRegistrationEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegistrationEntries", reflect.TypeOf((*MockDataStore)(nil).ListRegistrationEntries), arg0, arg1)
}

// PruneJoinTokens mocks base method
func (m *MockDataStore) PruneJoinTokens(arg0 context.Context, arg1 *datastore.PruneJoinTokensRequest) (*datastore.PruneJoinTokensResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneJoinTokens", arg0, arg1)
	ret0, _ := ret[0].(*datastore.PruneJoinTokensResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PruneJoinTokens indicates an expected call of PruneJoinTokens
func (mr *MockDataStoreMockRecorder) PruneJoinTokens(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneJoinTokens", reflect.TypeOf((*MockDataStore)(nil).PruneJoinTokens), arg0, arg1)
}

// PruneRegistrationEntries mocks base method
func (m *MockDataStore) PruneRegistrationEntries(arg0 context.Context, arg1 *datastore.PruneRegistrationEntriesRequest) (*datastore.PruneRegistrationEntriesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneRegistrationEntries", arg0, arg1)
	ret0, _ := ret[0].(*datastore.PruneRegistrationEntriesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PruneRegistrationEntries indicates an expected call of PruneRegistrationEntries
func (mr *MockDataStoreMockRecorder) PruneRegistrationEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneRegistrationEntries", reflect.TypeOf((*MockDataStore)(nil).PruneRegistrationEntries), arg0, arg1)
}

// SetNodeSelectors mocks base method
func (m *MockDataStore) SetNodeSelectors(arg0 context.Context, arg1 *datastore.SetNodeSelectorsRequest) (*datastore.SetNodeSelectorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNodeSelectors", arg0, arg1)
	ret0, _ := ret[0].(*datastore.SetNodeSelectorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNodeSelectors indicates an expected call of SetNodeSelectors
func (mr *MockDataStoreMockRecorder) SetNodeSelectors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeSelectors", reflect.TypeOf((*MockDataStore)(nil).SetNodeSelectors), arg0, arg1)
}

// UpdateAttestedNode mocks base method
func (m *MockDataStore) UpdateAttestedNode(arg0 context.Context, arg1 *datastore.UpdateAttestedNodeRequest) (*datastore.UpdateAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.UpdateAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAttestedNode indicates an expected call of UpdateAttestedNode
func (mr *MockDataStoreMockRecorder) UpdateAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttestedNode", reflect.TypeOf((*MockDataStore)(nil).UpdateAttestedNode), arg0, arg1)
}

// UpdateBundle mocks base method
func (m *MockDataStore) UpdateBundle(arg0 context.Context, arg1 *datastore.UpdateBundleRequest) (*datastore.UpdateBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.UpdateBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBundle indicates an expected call of UpdateBundle
func (mr *MockDataStoreMockRecorder) UpdateBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBundle", reflect.TypeOf((*MockDataStore)(nil).UpdateBundle), arg0, arg1)
}

// UpdateRegistrationEntry mocks base method
func (m *MockDataStore) UpdateRegistrationEntry(arg0 context.Context, arg1 *datastore.UpdateRegistrationEntryRequest) (*datastore.UpdateRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.UpdateRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRegistrationEntry indicates an expected call of UpdateRegistrationEntry
func (mr *MockDataStoreMockRecorder) UpdateRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRegistrationEntry", reflect.TypeOf((*MockDataStore)(nil).UpdateRegistrationEntry), arg0, arg1)
}

// MockDataStoreServer is a mock of DataStoreServer interface
type MockDataStoreServer struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreServerMockRecorder
}

// MockDataStoreServerMockRecorder is the mock recorder for MockDataStoreServer
type MockDataStoreServerMockRecorder struct {
	mock *MockDataStoreServer
}

// NewMockDataStoreServer creates a new mock instance
func NewMockDataStoreServer(ctrl *gomock.Controller) *MockDataStoreServer {
	mock := &MockDataStoreServer{ctrl: ctrl}
	mock.recorder = &MockDataStoreServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStoreServer) EXPECT() *MockDataStoreServerMockRecorder {
	return m.recorder
}

// AppendBundle mocks base method
func (m *MockDataStoreServer) AppendBundle(arg0 context.Context, arg1 *datastore.AppendBundleRequest) (*datastore.AppendBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.AppendBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendBundle indicates an expected call of AppendBundle
func (mr *MockDataStoreServerMockRecorder) AppendBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendBundle", reflect.TypeOf((*MockDataStoreServer)(nil).AppendBundle), arg0, arg1)
}

// Configure mocks base method
func (m *MockDataStoreServer) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockDataStoreServerMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockDataStoreServer)(nil).Configure), arg0, arg1)
}

// CreateAttestedNode mocks base method
func (m *MockDataStoreServer) CreateAttestedNode(arg0 context.Context, arg1 *datastore.CreateAttestedNodeRequest) (*datastore.CreateAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttestedNode indicates an expected call of CreateAttestedNode
func (mr *MockDataStoreServerMockRecorder) CreateAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttestedNode", reflect.TypeOf((*MockDataStoreServer)(nil).CreateAttestedNode), arg0, arg1)
}

// CreateBundle mocks base method
func (m *MockDataStoreServer) CreateBundle(arg0 context.Context, arg1 *datastore.CreateBundleRequest) (*datastore.CreateBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBundle indicates an expected call of CreateBundle
func (mr *MockDataStoreServerMockRecorder) CreateBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBundle", reflect.TypeOf((*MockDataStoreServer)(nil).CreateBundle), arg0, arg1)
}

// CreateJoinToken mocks base method
func (m *MockDataStoreServer) CreateJoinToken(arg0 context.Context, arg1 *datastore.CreateJoinTokenRequest) (*datastore.CreateJoinTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJoinToken", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateJoinTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJoinToken indicates an expected call of CreateJoinToken
func (mr *MockDataStoreServerMockRecorder) CreateJoinToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJoinToken", reflect.TypeOf((*MockDataStoreServer)(nil).CreateJoinToken), arg0, arg1)
}

// CreateRegistrationEntry mocks base method
func (m *MockDataStoreServer) CreateRegistrationEntry(arg0 context.Context, arg1 *datastore.CreateRegistrationEntryRequest) (*datastore.CreateRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.CreateRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRegistrationEntry indicates an expected call of CreateRegistrationEntry
func (mr *MockDataStoreServerMockRecorder) CreateRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRegistrationEntry", reflect.TypeOf((*MockDataStoreServer)(nil).CreateRegistrationEntry), arg0, arg1)
}

// DeleteAttestedNode mocks base method
func (m *MockDataStoreServer) DeleteAttestedNode(arg0 context.Context, arg1 *datastore.DeleteAttestedNodeRequest) (*datastore.DeleteAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAttestedNode indicates an expected call of DeleteAttestedNode
func (mr *MockDataStoreServerMockRecorder) DeleteAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttestedNode", reflect.TypeOf((*MockDataStoreServer)(nil).DeleteAttestedNode), arg0, arg1)
}

// DeleteBundle mocks base method
func (m *MockDataStoreServer) DeleteBundle(arg0 context.Context, arg1 *datastore.DeleteBundleRequest) (*datastore.DeleteBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBundle indicates an expected call of DeleteBundle
func (mr *MockDataStoreServerMockRecorder) DeleteBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBundle", reflect.TypeOf((*MockDataStoreServer)(nil).DeleteBundle), arg0, arg1)
}

// DeleteJoinToken mocks base method
func (m *MockDataStoreServer) DeleteJoinToken(arg0 context.Context, arg1 *datastore.DeleteJoinTokenRequest) (*datastore.DeleteJoinTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJoinToken", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteJoinTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteJoinToken indicates an expected call of DeleteJoinToken
func (mr *MockDataStoreServerMockRecorder) DeleteJoinToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJoinToken", reflect.TypeOf((*MockDataStoreServer)(nil).DeleteJoinToken), arg0, arg1)
}

// DeleteRegistrationEntry mocks base method
func (m *MockDataStoreServer) DeleteRegistrationEntry(arg0 context.Context, arg1 *datastore.DeleteRegistrationEntryRequest) (*datastore.DeleteRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.DeleteRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRegistrationEntry indicates an expected call of DeleteRegistrationEntry
func (mr *MockDataStoreServerMockRecorder) DeleteRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegistrationEntry", reflect.TypeOf((*MockDataStoreServer)(nil).DeleteRegistrationEntry), arg0, arg1)
}

// FetchAttestedNode mocks base method
func (m *MockDataStoreServer) FetchAttestedNode(arg0 context.Context, arg1 *datastore.FetchAttestedNodeRequest) (*datastore.FetchAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAttestedNode indicates an expected call of FetchAttestedNode
func (mr *MockDataStoreServerMockRecorder) FetchAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttestedNode", reflect.TypeOf((*MockDataStoreServer)(nil).FetchAttestedNode), arg0, arg1)
}

// FetchBundle mocks base method
func (m *MockDataStoreServer) FetchBundle(arg0 context.Context, arg1 *datastore.FetchBundleRequest) (*datastore.FetchBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBundle indicates an expected call of FetchBundle
func (mr *MockDataStoreServerMockRecorder) FetchBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBundle", reflect.TypeOf((*MockDataStoreServer)(nil).FetchBundle), arg0, arg1)
}

// FetchJoinToken mocks base method
func (m *MockDataStoreServer) FetchJoinToken(arg0 context.Context, arg1 *datastore.FetchJoinTokenRequest) (*datastore.FetchJoinTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJoinToken", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchJoinTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJoinToken indicates an expected call of FetchJoinToken
func (mr *MockDataStoreServerMockRecorder) FetchJoinToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJoinToken", reflect.TypeOf((*MockDataStoreServer)(nil).FetchJoinToken), arg0, arg1)
}

// FetchRegistrationEntry mocks base method
func (m *MockDataStoreServer) FetchRegistrationEntry(arg0 context.Context, arg1 *datastore.FetchRegistrationEntryRequest) (*datastore.FetchRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.FetchRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRegistrationEntry indicates an expected call of FetchRegistrationEntry
func (mr *MockDataStoreServerMockRecorder) FetchRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRegistrationEntry", reflect.TypeOf((*MockDataStoreServer)(nil).FetchRegistrationEntry), arg0, arg1)
}

// GetNodeSelectors mocks base method
func (m *MockDataStoreServer) GetNodeSelectors(arg0 context.Context, arg1 *datastore.GetNodeSelectorsRequest) (*datastore.GetNodeSelectorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeSelectors", arg0, arg1)
	ret0, _ := ret[0].(*datastore.GetNodeSelectorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeSelectors indicates an expected call of GetNodeSelectors
func (mr *MockDataStoreServerMockRecorder) GetNodeSelectors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeSelectors", reflect.TypeOf((*MockDataStoreServer)(nil).GetNodeSelectors), arg0, arg1)
}

// GetPluginInfo mocks base method
func (m *MockDataStoreServer) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockDataStoreServerMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockDataStoreServer)(nil).GetPluginInfo), arg0, arg1)
}

// ListAttestedNodes mocks base method
func (m *MockDataStoreServer) ListAttestedNodes(arg0 context.Context, arg1 *datastore.ListAttestedNodesRequest) (*datastore.ListAttestedNodesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttestedNodes", arg0, arg1)
	ret0, _ := ret[0].(*datastore.ListAttestedNodesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttestedNodes indicates an expected call of ListAttestedNodes
func (mr *MockDataStoreServerMockRecorder) ListAttestedNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttestedNodes", reflect.TypeOf((*MockDataStoreServer)(nil).ListAttestedNodes), arg0, arg1)
}

// ListBundles mocks base method
func (m *MockDataStoreServer) ListBundles(arg0 context.Context, arg1 *datastore.ListBundlesRequest) (*datastore.ListBundlesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBundles", arg0, arg1)
	ret0, _ := ret[0].(*datastore.ListBundlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBundles indicates an expected call of ListBundles
func (mr *MockDataStoreServerMockRecorder) ListBundles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBundles", reflect.TypeOf((*MockDataStoreServer)(nil).ListBundles), arg0, arg1)
}

// ListRegistrationEntries mocks base method
func (m *MockDataStoreServer) ListRegistrationEntries(arg0 context.Context, arg1 *datastore.ListRegistrationEntriesRequest) (*datastore.ListRegistrationEntriesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRegistrationEntries", arg0, arg1)
	ret0, _ := ret[0].(*datastore.ListRegistrationEntriesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegistrationEntries indicates an expected call of ListRegistrationEntries
func (mr *MockDataStoreServerMockRecorder) ListRegistrationEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegistrationEntries", reflect.TypeOf((*MockDataStoreServer)(nil).ListRegistrationEntries), arg0, arg1)
}

// PruneJoinTokens mocks base method
func (m *MockDataStoreServer) PruneJoinTokens(arg0 context.Context, arg1 *datastore.PruneJoinTokensRequest) (*datastore.PruneJoinTokensResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneJoinTokens", arg0, arg1)
	ret0, _ := ret[0].(*datastore.PruneJoinTokensResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PruneJoinTokens indicates an expected call of PruneJoinTokens
func (mr *MockDataStoreServerMockRecorder) PruneJoinTokens(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneJoinTokens", reflect.TypeOf((*MockDataStoreServer)(nil).PruneJoinTokens), arg0, arg1)
}

// PruneRegistrationEntries mocks base method
func (m *MockDataStoreServer) PruneRegistrationEntries(arg0 context.Context, arg1 *datastore.PruneRegistrationEntriesRequest) (*datastore.PruneRegistrationEntriesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneRegistrationEntries", arg0, arg1)
	ret0, _ := ret[0].(*datastore.PruneRegistrationEntriesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PruneRegistrationEntries indicates an expected call of PruneRegistrationEntries
func (mr *MockDataStoreServerMockRecorder) PruneRegistrationEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneRegistrationEntries", reflect.TypeOf((*MockDataStoreServer)(nil).PruneRegistrationEntries), arg0, arg1)
}

// SetNodeSelectors mocks base method
func (m *MockDataStoreServer) SetNodeSelectors(arg0 context.Context, arg1 *datastore.SetNodeSelectorsRequest) (*datastore.SetNodeSelectorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNodeSelectors", arg0, arg1)
	ret0, _ := ret[0].(*datastore.SetNodeSelectorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNodeSelectors indicates an expected call of SetNodeSelectors
func (mr *MockDataStoreServerMockRecorder) SetNodeSelectors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeSelectors", reflect.TypeOf((*MockDataStoreServer)(nil).SetNodeSelectors), arg0, arg1)
}

// UpdateAttestedNode mocks base method
func (m *MockDataStoreServer) UpdateAttestedNode(arg0 context.Context, arg1 *datastore.UpdateAttestedNodeRequest) (*datastore.UpdateAttestedNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttestedNode", arg0, arg1)
	ret0, _ := ret[0].(*datastore.UpdateAttestedNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAttestedNode indicates an expected call of UpdateAttestedNode
func (mr *MockDataStoreServerMockRecorder) UpdateAttestedNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttestedNode", reflect.TypeOf((*MockDataStoreServer)(nil).UpdateAttestedNode), arg0, arg1)
}

// UpdateBundle mocks base method
func (m *MockDataStoreServer) UpdateBundle(arg0 context.Context, arg1 *datastore.UpdateBundleRequest) (*datastore.UpdateBundleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBundle", arg0, arg1)
	ret0, _ := ret[0].(*datastore.UpdateBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBundle indicates an expected call of UpdateBundle
func (mr *MockDataStoreServerMockRecorder) UpdateBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBundle", reflect.TypeOf((*MockDataStoreServer)(nil).UpdateBundle), arg0, arg1)
}

// UpdateRegistrationEntry mocks base method
func (m *MockDataStoreServer) UpdateRegistrationEntry(arg0 context.Context, arg1 *datastore.UpdateRegistrationEntryRequest) (*datastore.UpdateRegistrationEntryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRegistrationEntry", arg0, arg1)
	ret0, _ := ret[0].(*datastore.UpdateRegistrationEntryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRegistrationEntry indicates an expected call of UpdateRegistrationEntry
func (mr *MockDataStoreServerMockRecorder) UpdateRegistrationEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRegistrationEntry", reflect.TypeOf((*MockDataStoreServer)(nil).UpdateRegistrationEntry), arg0, arg1)
}
