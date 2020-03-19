// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSpaceQuotaActor struct {
	GetSpaceQuotaByNameStub        func(string, string) (v7action.SpaceQuota, v7action.Warnings, error)
	getSpaceQuotaByNameMutex       sync.RWMutex
	getSpaceQuotaByNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceQuotaByNameReturns struct {
		result1 v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}
	getSpaceQuotaByNameReturnsOnCall map[int]struct {
		result1 v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceQuotaActor) GetSpaceQuotaByName(arg1 string, arg2 string) (v7action.SpaceQuota, v7action.Warnings, error) {
	fake.getSpaceQuotaByNameMutex.Lock()
	ret, specificReturn := fake.getSpaceQuotaByNameReturnsOnCall[len(fake.getSpaceQuotaByNameArgsForCall)]
	fake.getSpaceQuotaByNameArgsForCall = append(fake.getSpaceQuotaByNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceQuotaByName", []interface{}{arg1, arg2})
	fake.getSpaceQuotaByNameMutex.Unlock()
	if fake.GetSpaceQuotaByNameStub != nil {
		return fake.GetSpaceQuotaByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceQuotaByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceQuotaActor) GetSpaceQuotaByNameCallCount() int {
	fake.getSpaceQuotaByNameMutex.RLock()
	defer fake.getSpaceQuotaByNameMutex.RUnlock()
	return len(fake.getSpaceQuotaByNameArgsForCall)
}

func (fake *FakeSpaceQuotaActor) GetSpaceQuotaByNameCalls(stub func(string, string) (v7action.SpaceQuota, v7action.Warnings, error)) {
	fake.getSpaceQuotaByNameMutex.Lock()
	defer fake.getSpaceQuotaByNameMutex.Unlock()
	fake.GetSpaceQuotaByNameStub = stub
}

func (fake *FakeSpaceQuotaActor) GetSpaceQuotaByNameArgsForCall(i int) (string, string) {
	fake.getSpaceQuotaByNameMutex.RLock()
	defer fake.getSpaceQuotaByNameMutex.RUnlock()
	argsForCall := fake.getSpaceQuotaByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpaceQuotaActor) GetSpaceQuotaByNameReturns(result1 v7action.SpaceQuota, result2 v7action.Warnings, result3 error) {
	fake.getSpaceQuotaByNameMutex.Lock()
	defer fake.getSpaceQuotaByNameMutex.Unlock()
	fake.GetSpaceQuotaByNameStub = nil
	fake.getSpaceQuotaByNameReturns = struct {
		result1 v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceQuotaActor) GetSpaceQuotaByNameReturnsOnCall(i int, result1 v7action.SpaceQuota, result2 v7action.Warnings, result3 error) {
	fake.getSpaceQuotaByNameMutex.Lock()
	defer fake.getSpaceQuotaByNameMutex.Unlock()
	fake.GetSpaceQuotaByNameStub = nil
	if fake.getSpaceQuotaByNameReturnsOnCall == nil {
		fake.getSpaceQuotaByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.SpaceQuota
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceQuotaByNameReturnsOnCall[i] = struct {
		result1 v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceQuotaActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSpaceQuotaByNameMutex.RLock()
	defer fake.getSpaceQuotaByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceQuotaActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v7.SpaceQuotaActor = new(FakeSpaceQuotaActor)