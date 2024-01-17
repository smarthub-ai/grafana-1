// Code generated by mockery v2.32.0. DO NOT EDIT.

package publicdashboards

import (
	context "context"

	dashboards "github.com/grafana/grafana/pkg/services/dashboards"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/services/publicdashboards/models"
)

// FakePublicDashboardStore is an autogenerated mock type for the Store type
type FakePublicDashboardStore struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, cmd
func (_m *FakePublicDashboardStore) Create(ctx context.Context, cmd models.SavePublicDashboardCommand) (int64, error) {
	ret := _m.Called(ctx, cmd)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SavePublicDashboardCommand) (int64, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.SavePublicDashboardCommand) int64); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.SavePublicDashboardCommand) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, uid
func (_m *FakePublicDashboardStore) Delete(ctx context.Context, uid string) (int64, error) {
	ret := _m.Called(ctx, uid)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistsEnabledByAccessToken provides a mock function with given fields: ctx, accessToken
func (_m *FakePublicDashboardStore) ExistsEnabledByAccessToken(ctx context.Context, accessToken string) (bool, error) {
	ret := _m.Called(ctx, accessToken)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, accessToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, accessToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistsEnabledByDashboardUid provides a mock function with given fields: ctx, dashboardUid
func (_m *FakePublicDashboardStore) ExistsEnabledByDashboardUid(ctx context.Context, dashboardUid string) (bool, error) {
	ret := _m.Called(ctx, dashboardUid)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, dashboardUid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, dashboardUid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dashboardUid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, uid
func (_m *FakePublicDashboardStore) Find(ctx context.Context, uid string) (*models.PublicDashboard, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.PublicDashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.PublicDashboard, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.PublicDashboard); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicDashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllWithPagination provides a mock function with given fields: ctx, query
func (_m *FakePublicDashboardStore) FindAllWithPagination(ctx context.Context, query *models.PublicDashboardListQuery) (*models.PublicDashboardListResponseWithPagination, error) {
	ret := _m.Called(ctx, query)

	var r0 *models.PublicDashboardListResponseWithPagination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.PublicDashboardListQuery) (*models.PublicDashboardListResponseWithPagination, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.PublicDashboardListQuery) *models.PublicDashboardListResponseWithPagination); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicDashboardListResponseWithPagination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.PublicDashboardListQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByAccessToken provides a mock function with given fields: ctx, accessToken
func (_m *FakePublicDashboardStore) FindByAccessToken(ctx context.Context, accessToken string) (*models.PublicDashboard, error) {
	ret := _m.Called(ctx, accessToken)

	var r0 *models.PublicDashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.PublicDashboard, error)); ok {
		return rf(ctx, accessToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.PublicDashboard); ok {
		r0 = rf(ctx, accessToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicDashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByDashboardFolder provides a mock function with given fields: ctx, dashboard
func (_m *FakePublicDashboardStore) FindByDashboardFolder(ctx context.Context, dashboard *dashboards.Dashboard) ([]*models.PublicDashboard, error) {
	ret := _m.Called(ctx, dashboard)

	var r0 []*models.PublicDashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dashboards.Dashboard) ([]*models.PublicDashboard, error)); ok {
		return rf(ctx, dashboard)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dashboards.Dashboard) []*models.PublicDashboard); ok {
		r0 = rf(ctx, dashboard)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.PublicDashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dashboards.Dashboard) error); ok {
		r1 = rf(ctx, dashboard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByDashboardUid provides a mock function with given fields: ctx, orgId, dashboardUid
func (_m *FakePublicDashboardStore) FindByDashboardUid(ctx context.Context, orgId int64, dashboardUid string) (*models.PublicDashboard, error) {
	ret := _m.Called(ctx, orgId, dashboardUid)

	var r0 *models.PublicDashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (*models.PublicDashboard, error)); ok {
		return rf(ctx, orgId, dashboardUid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *models.PublicDashboard); ok {
		r0 = rf(ctx, orgId, dashboardUid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicDashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgId, dashboardUid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetrics provides a mock function with given fields: ctx
func (_m *FakePublicDashboardStore) GetMetrics(ctx context.Context) (*models.Metrics, error) {
	ret := _m.Called(ctx)

	var r0 *models.Metrics
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*models.Metrics, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *models.Metrics); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Metrics)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrgIdByAccessToken provides a mock function with given fields: ctx, accessToken
func (_m *FakePublicDashboardStore) GetOrgIdByAccessToken(ctx context.Context, accessToken string) (int64, error) {
	ret := _m.Called(ctx, accessToken)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, accessToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, accessToken)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, cmd
func (_m *FakePublicDashboardStore) Update(ctx context.Context, cmd models.SavePublicDashboardCommand) (int64, error) {
	ret := _m.Called(ctx, cmd)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SavePublicDashboardCommand) (int64, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.SavePublicDashboardCommand) int64); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.SavePublicDashboardCommand) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFakePublicDashboardStore creates a new instance of FakePublicDashboardStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFakePublicDashboardStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *FakePublicDashboardStore {
	mock := &FakePublicDashboardStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
