
// Code generated by mockery v2.23.1. DO NOT EDIT.

package service_test

import (
	entity "api-beer-challenge/internal/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "api-beer-challenge/internal/model"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx
func (_m *MockRepository) Count(ctx context.Context) (uint32, error) {
	ret := _m.Called(ctx)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint32, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint32); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockRepository_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepository_Expecter) Count(ctx interface{}) *MockRepository_Count_Call {
	return &MockRepository_Count_Call{Call: _e.mock.On("Count", ctx)}
}

func (_c *MockRepository_Count_Call) Run(run func(ctx context.Context)) *MockRepository_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepository_Count_Call) Return(_a0 uint32, _a1 error) *MockRepository_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Count_Call) RunAndReturn(run func(context.Context) (uint32, error)) *MockRepository_Count_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBeerByID provides a mock function with given fields: ctx, id
func (_m *MockRepository) DeleteBeerByID(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_DeleteBeerByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBeerByID'
type MockRepository_DeleteBeerByID_Call struct {
	*mock.Call
}

// DeleteBeerByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockRepository_Expecter) DeleteBeerByID(ctx interface{}, id interface{}) *MockRepository_DeleteBeerByID_Call {
	return &MockRepository_DeleteBeerByID_Call{Call: _e.mock.On("DeleteBeerByID", ctx, id)}
}

func (_c *MockRepository_DeleteBeerByID_Call) Run(run func(ctx context.Context, id uint64)) *MockRepository_DeleteBeerByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockRepository_DeleteBeerByID_Call) Return(_a0 error) *MockRepository_DeleteBeerByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_DeleteBeerByID_Call) RunAndReturn(run func(context.Context, uint64) error) *MockRepository_DeleteBeerByID_Call {
	_c.Call.Return(run)
	return _c
}

// FindBeerByID provides a mock function with given fields: ctx, id
func (_m *MockRepository) FindBeerByID(ctx context.Context, id uint64) (*entity.Beer, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*entity.Beer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *entity.Beer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Beer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindBeerByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBeerByID'
type MockRepository_FindBeerByID_Call struct {
	*mock.Call
}

// FindBeerByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockRepository_Expecter) FindBeerByID(ctx interface{}, id interface{}) *MockRepository_FindBeerByID_Call {
	return &MockRepository_FindBeerByID_Call{Call: _e.mock.On("FindBeerByID", ctx, id)}
}

func (_c *MockRepository_FindBeerByID_Call) Run(run func(ctx context.Context, id uint64)) *MockRepository_FindBeerByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockRepository_FindBeerByID_Call) Return(_a0 *entity.Beer, _a1 error) *MockRepository_FindBeerByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindBeerByID_Call) RunAndReturn(run func(context.Context, uint64) (*entity.Beer, error)) *MockRepository_FindBeerByID_Call {
	_c.Call.Return(run)
	return _c
}

// FindBeers provides a mock function with given fields: ctx, skip, limit
func (_m *MockRepository) FindBeers(ctx context.Context, skip uint32, limit uint32) ([]entity.Beer, error) {
	ret := _m.Called(ctx, skip, limit)

	var r0 []entity.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) ([]entity.Beer, error)); ok {
		return rf(ctx, skip, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) []entity.Beer); ok {
		r0 = rf(ctx, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Beer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint32) error); ok {
		r1 = rf(ctx, skip, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindBeers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBeers'
type MockRepository_FindBeers_Call struct {
	*mock.Call
}

// FindBeers is a helper method to define mock.On call
//   - ctx context.Context
//   - skip uint32
//   - limit uint32
func (_e *MockRepository_Expecter) FindBeers(ctx interface{}, skip interface{}, limit interface{}) *MockRepository_FindBeers_Call {
	return &MockRepository_FindBeers_Call{Call: _e.mock.On("FindBeers", ctx, skip, limit)}
}

func (_c *MockRepository_FindBeers_Call) Run(run func(ctx context.Context, skip uint32, limit uint32)) *MockRepository_FindBeers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *MockRepository_FindBeers_Call) Return(_a0 []entity.Beer, _a1 error) *MockRepository_FindBeers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindBeers_Call) RunAndReturn(run func(context.Context, uint32, uint32) ([]entity.Beer, error)) *MockRepository_FindBeers_Call {
	_c.Call.Return(run)
	return _c
}

// FindBoxPriceBeer provides a mock function with given fields: ctx, id, quantity, currency
func (_m *MockRepository) FindBoxPriceBeer(ctx context.Context, id uint64, quantity uint64, currency string) (float64, error) {
	ret := _m.Called(ctx, id, quantity, currency)

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, string) (float64, error)); ok {
		return rf(ctx, id, quantity, currency)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, string) float64); ok {
		r0 = rf(ctx, id, quantity, currency)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, string) error); ok {
		r1 = rf(ctx, id, quantity, currency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindBoxPriceBeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBoxPriceBeer'
type MockRepository_FindBoxPriceBeer_Call struct {
	*mock.Call
}

// FindBoxPriceBeer is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
//   - quantity uint64
//   - currency string
func (_e *MockRepository_Expecter) FindBoxPriceBeer(ctx interface{}, id interface{}, quantity interface{}, currency interface{}) *MockRepository_FindBoxPriceBeer_Call {
	return &MockRepository_FindBoxPriceBeer_Call{Call: _e.mock.On("FindBoxPriceBeer", ctx, id, quantity, currency)}
}

func (_c *MockRepository_FindBoxPriceBeer_Call) Run(run func(ctx context.Context, id uint64, quantity uint64, currency string)) *MockRepository_FindBoxPriceBeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64), args[3].(string))
	})
	return _c
}

func (_c *MockRepository_FindBoxPriceBeer_Call) Return(_a0 float64, _a1 error) *MockRepository_FindBoxPriceBeer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindBoxPriceBeer_Call) RunAndReturn(run func(context.Context, uint64, uint64, string) (float64, error)) *MockRepository_FindBoxPriceBeer_Call {
	_c.Call.Return(run)
	return _c
}

// InsertBeer provides a mock function with given fields: ctx, input
func (_m *MockRepository) InsertBeer(ctx context.Context, input *model.InputBeer) (*entity.Beer, error) {
	ret := _m.Called(ctx, input)

	var r0 *entity.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.InputBeer) (*entity.Beer, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.InputBeer) *entity.Beer); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Beer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.InputBeer) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_InsertBeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertBeer'
type MockRepository_InsertBeer_Call struct {
	*mock.Call
}

// InsertBeer is a helper method to define mock.On call
//   - ctx context.Context
//   - input *model.InputBeer
func (_e *MockRepository_Expecter) InsertBeer(ctx interface{}, input interface{}) *MockRepository_InsertBeer_Call {
	return &MockRepository_InsertBeer_Call{Call: _e.mock.On("InsertBeer", ctx, input)}
}

func (_c *MockRepository_InsertBeer_Call) Run(run func(ctx context.Context, input *model.InputBeer)) *MockRepository_InsertBeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.InputBeer))
	})
	return _c
}

func (_c *MockRepository_InsertBeer_Call) Return(_a0 *entity.Beer, _a1 error) *MockRepository_InsertBeer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_InsertBeer_Call) RunAndReturn(run func(context.Context, *model.InputBeer) (*entity.Beer, error)) *MockRepository_InsertBeer_Call {
	_c.Call.Return(run)
	return _c
}

// RestartTable provides a mock function with given fields: ctx, src
func (_m *MockRepository) RestartTable(ctx context.Context, src string) error {
	ret := _m.Called(ctx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_RestartTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestartTable'
type MockRepository_RestartTable_Call struct {
	*mock.Call
}

// RestartTable is a helper method to define mock.On call
//   - ctx context.Context
//   - src string
func (_e *MockRepository_Expecter) RestartTable(ctx interface{}, src interface{}) *MockRepository_RestartTable_Call {
	return &MockRepository_RestartTable_Call{Call: _e.mock.On("RestartTable", ctx, src)}
}

func (_c *MockRepository_RestartTable_Call) Run(run func(ctx context.Context, src string)) *MockRepository_RestartTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_RestartTable_Call) Return(_a0 error) *MockRepository_RestartTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_RestartTable_Call) RunAndReturn(run func(context.Context, string) error) *MockRepository_RestartTable_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBeerByID provides a mock function with given fields: ctx, id, input
func (_m *MockRepository) UpdateBeerByID(ctx context.Context, id uint64, input *model.InputUBeer) (*entity.Beer, error) {
	ret := _m.Called(ctx, id, input)

	var r0 *entity.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *model.InputUBeer) (*entity.Beer, error)); ok {
		return rf(ctx, id, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *model.InputUBeer) *entity.Beer); ok {
		r0 = rf(ctx, id, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Beer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, *model.InputUBeer) error); ok {
		r1 = rf(ctx, id, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_UpdateBeerByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBeerByID'
type MockRepository_UpdateBeerByID_Call struct {
	*mock.Call
}

// UpdateBeerByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
//   - input *model.InputUBeer
func (_e *MockRepository_Expecter) UpdateBeerByID(ctx interface{}, id interface{}, input interface{}) *MockRepository_UpdateBeerByID_Call {
	return &MockRepository_UpdateBeerByID_Call{Call: _e.mock.On("UpdateBeerByID", ctx, id, input)}
}

func (_c *MockRepository_UpdateBeerByID_Call) Run(run func(ctx context.Context, id uint64, input *model.InputUBeer)) *MockRepository_UpdateBeerByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*model.InputUBeer))
	})
	return _c
}

func (_c *MockRepository_UpdateBeerByID_Call) Return(_a0 *entity.Beer, _a1 error) *MockRepository_UpdateBeerByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_UpdateBeerByID_Call) RunAndReturn(run func(context.Context, uint64, *model.InputUBeer) (*entity.Beer, error)) *MockRepository_UpdateBeerByID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
