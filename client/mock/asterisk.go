// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: Asterisk)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Asterisk interface
type MockAsterisk struct {
	ctrl     *gomock.Controller
	recorder *_MockAsteriskRecorder
}

// Recorder for MockAsterisk (not exported)
type _MockAsteriskRecorder struct {
	mock *MockAsterisk
}

func NewMockAsterisk(ctrl *gomock.Controller) *MockAsterisk {
	mock := &MockAsterisk{ctrl: ctrl}
	mock.recorder = &_MockAsteriskRecorder{mock}
	return mock
}

func (_m *MockAsterisk) EXPECT() *_MockAsteriskRecorder {
	return _m.recorder
}

func (_m *MockAsterisk) Info(_param0 string) (*ari.AsteriskInfo, error) {
	ret := _m.ctrl.Call(_m, "Info", _param0)
	ret0, _ := ret[0].(*ari.AsteriskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAsteriskRecorder) Info(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Info", arg0)
}

func (_m *MockAsterisk) Logging() ari.Logging {
	ret := _m.ctrl.Call(_m, "Logging")
	ret0, _ := ret[0].(ari.Logging)
	return ret0
}

func (_mr *_MockAsteriskRecorder) Logging() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Logging")
}

func (_m *MockAsterisk) Modules() ari.Modules {
	ret := _m.ctrl.Call(_m, "Modules")
	ret0, _ := ret[0].(ari.Modules)
	return ret0
}

func (_mr *_MockAsteriskRecorder) Modules() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Modules")
}

func (_m *MockAsterisk) ReloadModule(_param0 string) error {
	ret := _m.ctrl.Call(_m, "ReloadModule", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAsteriskRecorder) ReloadModule(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReloadModule", arg0)
}

func (_m *MockAsterisk) Variables() ari.Variables {
	ret := _m.ctrl.Call(_m, "Variables")
	ret0, _ := ret[0].(ari.Variables)
	return ret0
}

func (_mr *_MockAsteriskRecorder) Variables() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Variables")
}
