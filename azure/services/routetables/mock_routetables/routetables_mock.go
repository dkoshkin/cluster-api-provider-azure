/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../routetables.go
//
// Generated by this command:
//
//	mockgen -destination routetables_mock.go -package mock_routetables -source ../routetables.go RouteTableScope
//
// Package mock_routetables is a generated GoMock package.
package mock_routetables

import (
	reflect "reflect"
	time "time"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	gomock "go.uber.org/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockRouteTableScope is a mock of RouteTableScope interface.
type MockRouteTableScope struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTableScopeMockRecorder
}

// MockRouteTableScopeMockRecorder is the mock recorder for MockRouteTableScope.
type MockRouteTableScopeMockRecorder struct {
	mock *MockRouteTableScope
}

// NewMockRouteTableScope creates a new mock instance.
func NewMockRouteTableScope(ctrl *gomock.Controller) *MockRouteTableScope {
	mock := &MockRouteTableScope{ctrl: ctrl}
	mock.recorder = &MockRouteTableScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteTableScope) EXPECT() *MockRouteTableScopeMockRecorder {
	return m.recorder
}

// BaseURI mocks base method.
func (m *MockRouteTableScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockRouteTableScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockRouteTableScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockRouteTableScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockRouteTableScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockRouteTableScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockRouteTableScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockRouteTableScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockRouteTableScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockRouteTableScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockRouteTableScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockRouteTableScope)(nil).CloudEnvironment))
}

// DefaultedAzureCallTimeout mocks base method.
func (m *MockRouteTableScope) DefaultedAzureCallTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureCallTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureCallTimeout indicates an expected call of DefaultedAzureCallTimeout.
func (mr *MockRouteTableScopeMockRecorder) DefaultedAzureCallTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureCallTimeout", reflect.TypeOf((*MockRouteTableScope)(nil).DefaultedAzureCallTimeout))
}

// DefaultedAzureServiceReconcileTimeout mocks base method.
func (m *MockRouteTableScope) DefaultedAzureServiceReconcileTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureServiceReconcileTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureServiceReconcileTimeout indicates an expected call of DefaultedAzureServiceReconcileTimeout.
func (mr *MockRouteTableScopeMockRecorder) DefaultedAzureServiceReconcileTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureServiceReconcileTimeout", reflect.TypeOf((*MockRouteTableScope)(nil).DefaultedAzureServiceReconcileTimeout))
}

// DefaultedReconcilerRequeue mocks base method.
func (m *MockRouteTableScope) DefaultedReconcilerRequeue() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedReconcilerRequeue")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedReconcilerRequeue indicates an expected call of DefaultedReconcilerRequeue.
func (mr *MockRouteTableScopeMockRecorder) DefaultedReconcilerRequeue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedReconcilerRequeue", reflect.TypeOf((*MockRouteTableScope)(nil).DefaultedReconcilerRequeue))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockRouteTableScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockRouteTableScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockRouteTableScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// GetLongRunningOperationState mocks base method.
func (m *MockRouteTableScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockRouteTableScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockRouteTableScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// HashKey mocks base method.
func (m *MockRouteTableScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockRouteTableScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockRouteTableScope)(nil).HashKey))
}

// IsVnetManaged mocks base method.
func (m *MockRouteTableScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockRouteTableScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockRouteTableScope)(nil).IsVnetManaged))
}

// RouteTableSpecs mocks base method.
func (m *MockRouteTableScope) RouteTableSpecs() []azure.ResourceSpecGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteTableSpecs")
	ret0, _ := ret[0].([]azure.ResourceSpecGetter)
	return ret0
}

// RouteTableSpecs indicates an expected call of RouteTableSpecs.
func (mr *MockRouteTableScopeMockRecorder) RouteTableSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteTableSpecs", reflect.TypeOf((*MockRouteTableScope)(nil).RouteTableSpecs))
}

// SetLongRunningOperationState mocks base method.
func (m *MockRouteTableScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockRouteTableScopeMockRecorder) SetLongRunningOperationState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockRouteTableScope)(nil).SetLongRunningOperationState), arg0)
}

// SubscriptionID mocks base method.
func (m *MockRouteTableScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockRouteTableScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockRouteTableScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockRouteTableScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockRouteTableScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockRouteTableScope)(nil).TenantID))
}

// Token mocks base method.
func (m *MockRouteTableScope) Token() azcore.TokenCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(azcore.TokenCredential)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockRouteTableScopeMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockRouteTableScope)(nil).Token))
}

// UpdateDeleteStatus mocks base method.
func (m *MockRouteTableScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockRouteTableScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockRouteTableScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockRouteTableScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockRouteTableScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockRouteTableScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockRouteTableScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockRouteTableScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockRouteTableScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}
