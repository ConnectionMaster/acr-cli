// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import acr "github.com/Azure/acr-cli/acr"

import autorest "github.com/Azure/go-autorest/autorest"
import context "context"
import mock "github.com/stretchr/testify/mock"

// AcrCLIClientInterface is an autogenerated mock type for the AcrCLIClientInterface type
type AcrCLIClientInterface struct {
	mock.Mock
}

// DeleteAcrTag provides a mock function with given fields: ctx, repoName, reference
func (_m *AcrCLIClientInterface) DeleteAcrTag(ctx context.Context, repoName string, reference string) (*autorest.Response, error) {
	ret := _m.Called(ctx, repoName, reference)

	var r0 *autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *autorest.Response); ok {
		r0 = rf(ctx, repoName, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autorest.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, repoName, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteManifest provides a mock function with given fields: ctx, repoName, reference
func (_m *AcrCLIClientInterface) DeleteManifest(ctx context.Context, repoName string, reference string) (*autorest.Response, error) {
	ret := _m.Called(ctx, repoName, reference)

	var r0 *autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *autorest.Response); ok {
		r0 = rf(ctx, repoName, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autorest.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, repoName, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrManifests provides a mock function with given fields: ctx, repoName, orderBy, last
func (_m *AcrCLIClientInterface) GetAcrManifests(ctx context.Context, repoName string, orderBy string, last string) (*acr.Manifests, error) {
	ret := _m.Called(ctx, repoName, orderBy, last)

	var r0 *acr.Manifests
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *acr.Manifests); ok {
		r0 = rf(ctx, repoName, orderBy, last)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acr.Manifests)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, repoName, orderBy, last)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrTags provides a mock function with given fields: ctx, repoName, orderBy, last
func (_m *AcrCLIClientInterface) GetAcrTags(ctx context.Context, repoName string, orderBy string, last string) (*acr.RepositoryTagsType, error) {
	ret := _m.Called(ctx, repoName, orderBy, last)

	var r0 *acr.RepositoryTagsType
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *acr.RepositoryTagsType); ok {
		r0 = rf(ctx, repoName, orderBy, last)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acr.RepositoryTagsType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, repoName, orderBy, last)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManifest provides a mock function with given fields: ctx, repoName, reference
func (_m *AcrCLIClientInterface) GetManifest(ctx context.Context, repoName string, reference string) ([]byte, error) {
	ret := _m.Called(ctx, repoName, reference)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []byte); ok {
		r0 = rf(ctx, repoName, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, repoName, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrManifestAttributes provides a mock function with given fields: ctx, repoName, reference
func (_m *AcrCLIClientInterface) GetAcrManifestAttributes(ctx context.Context, repoName string, reference string) (*acr.ManifestAttributes, error) {
	ret := _m.Called(ctx, repoName, reference)

	var r0 *acr.ManifestAttributes
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *acr.ManifestAttributes); ok {
		r0 = rf(ctx, repoName, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acr.ManifestAttributes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, repoName, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrTagAttributes provides a mock function with given fields: ctx, repoName, reference, value
func (_m *AcrCLIClientInterface) UpdateAcrTagAttributes(ctx context.Context, repoName string, reference string, value *acr.ChangeableAttributes) (*autorest.Response, error) {
	ret := _m.Called(ctx, repoName, reference, value)

	var r0 *autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *acr.ChangeableAttributes) *autorest.Response); ok {
		r0 = rf(ctx, repoName, reference, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autorest.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *acr.ChangeableAttributes) error); ok {
		r1 = rf(ctx, repoName, reference, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrManifestAttributes provides a mock function with given fields: ctx, repoName, reference, value
func (_m *AcrCLIClientInterface) UpdateAcrManifestAttributes(ctx context.Context, repoName string, reference string, value *acr.ChangeableAttributes) (*autorest.Response, error) {
	ret := _m.Called(ctx, repoName, reference, value)

	var r0 *autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *acr.ChangeableAttributes) *autorest.Response); ok {
		r0 = rf(ctx, repoName, reference, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autorest.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *acr.ChangeableAttributes) error); ok {
		r1 = rf(ctx, repoName, reference, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
