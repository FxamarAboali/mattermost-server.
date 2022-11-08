// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// PostPriorityStore is an autogenerated mock type for the PostPriorityStore type
type PostPriorityStore struct {
	mock.Mock
}

// GetForPost provides a mock function with given fields: postId
func (_m *PostPriorityStore) GetForPost(postId string) (*model.PostPriority, error) {
	ret := _m.Called(postId)

	var r0 *model.PostPriority
	if rf, ok := ret.Get(0).(func(string) *model.PostPriority); ok {
		r0 = rf(postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostPriority)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPersistentNotificationsPosts provides a mock function with given fields: minCreateAt
func (_m *PostPriorityStore) GetPersistentNotificationsPosts(minCreateAt int64) ([]*model.PostPersistentNotifications, error) {
	ret := _m.Called(minCreateAt)

	var r0 []*model.PostPersistentNotifications
	if rf, ok := ret.Get(0).(func(int64) []*model.PostPersistentNotifications); ok {
		r0 = rf(minCreateAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostPersistentNotifications)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(minCreateAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
