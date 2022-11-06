// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	sessions "github.com/gorilla/sessions"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// SessionStore is an autogenerated mock type for the SessionStore type
type SessionStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: r, name
func (_m *SessionStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	ret := _m.Called(r, name)

	var r0 *sessions.Session
	if rf, ok := ret.Get(0).(func(*http.Request, string) *sessions.Session); ok {
		r0 = rf(r, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sessions.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(r, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSessionStore creates a new instance of SessionStore. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewSessionStore(t testing.TB) *SessionStore {
	mock := &SessionStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}