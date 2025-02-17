// Code generated by MockGen. DO NOT EDIT.
// Source: adapter.go
//
// Generated by this command:
//
//	mockgen -source adapter.go -package mocks -destination ../mocks/adapter.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockRandomAdapter is a mock of RandomAdapter interface.
type MockRandomAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockRandomAdapterMockRecorder
	isgomock struct{}
}

// MockRandomAdapterMockRecorder is the mock recorder for MockRandomAdapter.
type MockRandomAdapterMockRecorder struct {
	mock *MockRandomAdapter
}

// NewMockRandomAdapter creates a new mock instance.
func NewMockRandomAdapter(ctrl *gomock.Controller) *MockRandomAdapter {
	mock := &MockRandomAdapter{ctrl: ctrl}
	mock.recorder = &MockRandomAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRandomAdapter) EXPECT() *MockRandomAdapterMockRecorder {
	return m.recorder
}

// RandomIntn mocks base method.
func (m *MockRandomAdapter) RandomIntn(n int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomIntn", n)
	ret0, _ := ret[0].(int)
	return ret0
}

// RandomIntn indicates an expected call of RandomIntn.
func (mr *MockRandomAdapterMockRecorder) RandomIntn(n any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomIntn", reflect.TypeOf((*MockRandomAdapter)(nil).RandomIntn), n)
}

// RandomString mocks base method.
func (m *MockRandomAdapter) RandomString(length int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomString", length)
	ret0, _ := ret[0].(string)
	return ret0
}

// RandomString indicates an expected call of RandomString.
func (mr *MockRandomAdapterMockRecorder) RandomString(length any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomString", reflect.TypeOf((*MockRandomAdapter)(nil).RandomString), length)
}

// RandomStringWithTimeNanoSeed mocks base method.
func (m *MockRandomAdapter) RandomStringWithTimeNanoSeed(length int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomStringWithTimeNanoSeed", length)
	ret0, _ := ret[0].(string)
	return ret0
}

// RandomStringWithTimeNanoSeed indicates an expected call of RandomStringWithTimeNanoSeed.
func (mr *MockRandomAdapterMockRecorder) RandomStringWithTimeNanoSeed(length any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomStringWithTimeNanoSeed", reflect.TypeOf((*MockRandomAdapter)(nil).RandomStringWithTimeNanoSeed), length)
}

// RandomToken mocks base method.
func (m *MockRandomAdapter) RandomToken(length int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomToken", length)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RandomToken indicates an expected call of RandomToken.
func (mr *MockRandomAdapterMockRecorder) RandomToken(length any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomToken", reflect.TypeOf((*MockRandomAdapter)(nil).RandomToken), length)
}

// MockTimeAdapter is a mock of TimeAdapter interface.
type MockTimeAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockTimeAdapterMockRecorder
	isgomock struct{}
}

// MockTimeAdapterMockRecorder is the mock recorder for MockTimeAdapter.
type MockTimeAdapterMockRecorder struct {
	mock *MockTimeAdapter
}

// NewMockTimeAdapter creates a new mock instance.
func NewMockTimeAdapter(ctrl *gomock.Controller) *MockTimeAdapter {
	mock := &MockTimeAdapter{ctrl: ctrl}
	mock.recorder = &MockTimeAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimeAdapter) EXPECT() *MockTimeAdapterMockRecorder {
	return m.recorder
}

// Locale mocks base method.
func (m *MockTimeAdapter) Locale() *time.Location {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Locale")
	ret0, _ := ret[0].(*time.Location)
	return ret0
}

// Locale indicates an expected call of Locale.
func (mr *MockTimeAdapterMockRecorder) Locale() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Locale", reflect.TypeOf((*MockTimeAdapter)(nil).Locale))
}

// MillisecondsToTime mocks base method.
func (m *MockTimeAdapter) MillisecondsToTime(ms int64) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MillisecondsToTime", ms)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// MillisecondsToTime indicates an expected call of MillisecondsToTime.
func (mr *MockTimeAdapterMockRecorder) MillisecondsToTime(ms any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MillisecondsToTime", reflect.TypeOf((*MockTimeAdapter)(nil).MillisecondsToTime), ms)
}

// Now mocks base method.
func (m *MockTimeAdapter) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockTimeAdapterMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockTimeAdapter)(nil).Now))
}

// TimeMidnight mocks base method.
func (m *MockTimeAdapter) TimeMidnight(t time.Time) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeMidnight", t)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// TimeMidnight indicates an expected call of TimeMidnight.
func (mr *MockTimeAdapterMockRecorder) TimeMidnight(t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeMidnight", reflect.TypeOf((*MockTimeAdapter)(nil).TimeMidnight), t)
}

// TodayMidnight mocks base method.
func (m *MockTimeAdapter) TodayMidnight() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodayMidnight")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// TodayMidnight indicates an expected call of TodayMidnight.
func (mr *MockTimeAdapterMockRecorder) TodayMidnight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodayMidnight", reflect.TypeOf((*MockTimeAdapter)(nil).TodayMidnight))
}
