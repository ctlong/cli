// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeCreateServiceBrokerActor struct {
	CreateServiceBrokerStub        func(ccv3.ServiceBrokerModel) (v7action.Warnings, error)
	createServiceBrokerMutex       sync.RWMutex
	createServiceBrokerArgsForCall []struct {
		arg1 ccv3.ServiceBrokerModel
	}
	createServiceBrokerReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	createServiceBrokerReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBroker(arg1 ccv3.ServiceBrokerModel) (v7action.Warnings, error) {
	fake.createServiceBrokerMutex.Lock()
	ret, specificReturn := fake.createServiceBrokerReturnsOnCall[len(fake.createServiceBrokerArgsForCall)]
	fake.createServiceBrokerArgsForCall = append(fake.createServiceBrokerArgsForCall, struct {
		arg1 ccv3.ServiceBrokerModel
	}{arg1})
	fake.recordInvocation("CreateServiceBroker", []interface{}{arg1})
	fake.createServiceBrokerMutex.Unlock()
	if fake.CreateServiceBrokerStub != nil {
		return fake.CreateServiceBrokerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createServiceBrokerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerCallCount() int {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	return len(fake.createServiceBrokerArgsForCall)
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerCalls(stub func(ccv3.ServiceBrokerModel) (v7action.Warnings, error)) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = stub
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerArgsForCall(i int) ccv3.ServiceBrokerModel {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	argsForCall := fake.createServiceBrokerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerReturns(result1 v7action.Warnings, result2 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	fake.createServiceBrokerReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	if fake.createServiceBrokerReturnsOnCall == nil {
		fake.createServiceBrokerReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.createServiceBrokerReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateServiceBrokerActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateServiceBrokerActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.CreateServiceBrokerActor = new(FakeCreateServiceBrokerActor)
