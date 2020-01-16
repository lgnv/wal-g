// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	oplog "github.com/wal-g/wal-g/internal/databases/mongo/oplog"

	sync "sync"
)

// FromFetcher is an autogenerated mock type for the FromFetcher type
type FromFetcher struct {
	mock.Mock
}

// OplogFrom provides a mock function with given fields: _a0, _a1, _a2
func (_m *FromFetcher) OplogFrom(_a0 context.Context, _a1 oplog.Timestamp, _a2 *sync.WaitGroup) (chan oplog.Record, chan error, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 chan oplog.Record
	if rf, ok := ret.Get(0).(func(context.Context, oplog.Timestamp, *sync.WaitGroup) chan oplog.Record); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan oplog.Record)
		}
	}

	var r1 chan error
	if rf, ok := ret.Get(1).(func(context.Context, oplog.Timestamp, *sync.WaitGroup) chan error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, oplog.Timestamp, *sync.WaitGroup) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}