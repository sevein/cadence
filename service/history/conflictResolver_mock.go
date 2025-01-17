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
// Source: conflictResolver.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	persistence "github.com/uber/cadence/common/persistence"
	reflect "reflect"
)

// MockconflictResolver is a mock of conflictResolver interface
type MockconflictResolver struct {
	ctrl     *gomock.Controller
	recorder *MockconflictResolverMockRecorder
}

// MockconflictResolverMockRecorder is the mock recorder for MockconflictResolver
type MockconflictResolverMockRecorder struct {
	mock *MockconflictResolver
}

// NewMockconflictResolver creates a new mock instance
func NewMockconflictResolver(ctrl *gomock.Controller) *MockconflictResolver {
	mock := &MockconflictResolver{ctrl: ctrl}
	mock.recorder = &MockconflictResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockconflictResolver) EXPECT() *MockconflictResolverMockRecorder {
	return m.recorder
}

// reset mocks base method
func (m *MockconflictResolver) reset(prevRunID string, prevLastWriteVersion int64, prevState int, requestID string, replayEventID int64, info *persistence.WorkflowExecutionInfo, updateCondition int64) (mutableState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "reset", prevRunID, prevLastWriteVersion, prevState, requestID, replayEventID, info, updateCondition)
	ret0, _ := ret[0].(mutableState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// reset indicates an expected call of reset
func (mr *MockconflictResolverMockRecorder) reset(prevRunID, prevLastWriteVersion, prevState, requestID, replayEventID, info, updateCondition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "reset", reflect.TypeOf((*MockconflictResolver)(nil).reset), prevRunID, prevLastWriteVersion, prevState, requestID, replayEventID, info, updateCondition)
}
