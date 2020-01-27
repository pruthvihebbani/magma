// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	policydb "magma/feg/gateway/policydb"

	mock "github.com/stretchr/testify/mock"

	protos "magma/lte/cloud/go/protos"
)

// PolicyDBClient is an autogenerated mock type for the PolicyDBClient type
type PolicyDBClient struct {
	mock.Mock
}

// GetChargingKeysForRules provides a mock function with given fields: ruleIDs, ruleDefs
func (_m *PolicyDBClient) GetChargingKeysForRules(ruleIDs []string, ruleDefs []*protos.PolicyRule) ([]policydb.ChargingKey, error) {
	ret := _m.Called(ruleIDs, ruleDefs)

	var r0 []policydb.ChargingKey
	if rf, ok := ret.Get(0).(func([]string, []*protos.PolicyRule) []policydb.ChargingKey); ok {
		r0 = rf(ruleIDs, ruleDefs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]policydb.ChargingKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, []*protos.PolicyRule) error); ok {
		r1 = rf(ruleIDs, ruleDefs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOmnipresentRules provides a mock function with given fields:
func (_m *PolicyDBClient) GetOmnipresentRules() ([]string, []string) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func() []string); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	return r0, r1
}

// GetPolicyRuleByID provides a mock function with given fields: id
func (_m *PolicyDBClient) GetPolicyRuleByID(id string) (*protos.PolicyRule, error) {
	ret := _m.Called(id)

	var r0 *protos.PolicyRule
	if rf, ok := ret.Get(0).(func(string) *protos.PolicyRule); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.PolicyRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRuleIDsForBaseNames provides a mock function with given fields: baseNames
func (_m *PolicyDBClient) GetRuleIDsForBaseNames(baseNames []string) []string {
	ret := _m.Called(baseNames)

	var r0 []string
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(baseNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}
