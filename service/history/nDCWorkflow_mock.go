// The MIT License (MIT)
//
// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: nDCWorkflow.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MocknDCWorkflow is a mock of nDCWorkflow interface
type MocknDCWorkflow struct {
	ctrl     *gomock.Controller
	recorder *MocknDCWorkflowMockRecorder
}

// MocknDCWorkflowMockRecorder is the mock recorder for MocknDCWorkflow
type MocknDCWorkflowMockRecorder struct {
	mock *MocknDCWorkflow
}

// NewMocknDCWorkflow creates a new mock instance
func NewMocknDCWorkflow(ctrl *gomock.Controller) *MocknDCWorkflow {
	mock := &MocknDCWorkflow{ctrl: ctrl}
	mock.recorder = &MocknDCWorkflowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocknDCWorkflow) EXPECT() *MocknDCWorkflowMockRecorder {
	return m.recorder
}

// getContext mocks base method
func (m *MocknDCWorkflow) getContext() workflowExecutionContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getContext")
	ret0, _ := ret[0].(workflowExecutionContext)
	return ret0
}

// getContext indicates an expected call of getContext
func (mr *MocknDCWorkflowMockRecorder) getContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getContext", reflect.TypeOf((*MocknDCWorkflow)(nil).getContext))
}

// getMutableState mocks base method
func (m *MocknDCWorkflow) getMutableState() mutableState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getMutableState")
	ret0, _ := ret[0].(mutableState)
	return ret0
}

// getMutableState indicates an expected call of getMutableState
func (mr *MocknDCWorkflowMockRecorder) getMutableState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getMutableState", reflect.TypeOf((*MocknDCWorkflow)(nil).getMutableState))
}

// getReleaseFn mocks base method
func (m *MocknDCWorkflow) getReleaseFn() releaseWorkflowExecutionFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getReleaseFn")
	ret0, _ := ret[0].(releaseWorkflowExecutionFunc)
	return ret0
}

// getReleaseFn indicates an expected call of getReleaseFn
func (mr *MocknDCWorkflowMockRecorder) getReleaseFn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getReleaseFn", reflect.TypeOf((*MocknDCWorkflow)(nil).getReleaseFn))
}

// getVectorClock mocks base method
func (m *MocknDCWorkflow) getVectorClock() (int64, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getVectorClock")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// getVectorClock indicates an expected call of getVectorClock
func (mr *MocknDCWorkflowMockRecorder) getVectorClock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getVectorClock", reflect.TypeOf((*MocknDCWorkflow)(nil).getVectorClock))
}

// happensAfter mocks base method
func (m *MocknDCWorkflow) happensAfter(that nDCWorkflow) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "happensAfter", that)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// happensAfter indicates an expected call of happensAfter
func (mr *MocknDCWorkflowMockRecorder) happensAfter(that interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "happensAfter", reflect.TypeOf((*MocknDCWorkflow)(nil).happensAfter), that)
}

// revive mocks base method
func (m *MocknDCWorkflow) revive() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "revive")
	ret0, _ := ret[0].(error)
	return ret0
}

// revive indicates an expected call of revive
func (mr *MocknDCWorkflowMockRecorder) revive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "revive", reflect.TypeOf((*MocknDCWorkflow)(nil).revive))
}

// suppressBy mocks base method
func (m *MocknDCWorkflow) suppressBy(incomingWorkflow nDCWorkflow) (transactionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "suppressBy", incomingWorkflow)
	ret0, _ := ret[0].(transactionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// suppressBy indicates an expected call of suppressBy
func (mr *MocknDCWorkflowMockRecorder) suppressBy(incomingWorkflow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "suppressBy", reflect.TypeOf((*MocknDCWorkflow)(nil).suppressBy), incomingWorkflow)
}

// flushBufferedEvents mocks base method
func (m *MocknDCWorkflow) flushBufferedEvents() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "flushBufferedEvents")
	ret0, _ := ret[0].(error)
	return ret0
}

// flushBufferedEvents indicates an expected call of flushBufferedEvents
func (mr *MocknDCWorkflowMockRecorder) flushBufferedEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "flushBufferedEvents", reflect.TypeOf((*MocknDCWorkflow)(nil).flushBufferedEvents))
}
