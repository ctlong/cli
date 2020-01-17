// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/util/manifestparser"
	"github.com/cloudfoundry/bosh-cli/director/template"
)

type FakeManifestParser struct {
	InterpolateAndParseStub        func(string, []string, []template.VarKV) (manifestparser.Manifest, error)
	interpolateAndParseMutex       sync.RWMutex
	interpolateAndParseArgsForCall []struct {
		arg1 string
		arg2 []string
		arg3 []template.VarKV
	}
	interpolateAndParseReturns struct {
		result1 manifestparser.Manifest
		result2 error
	}
	interpolateAndParseReturnsOnCall map[int]struct {
		result1 manifestparser.Manifest
		result2 error
	}
	MarshalManifestStub        func(manifestparser.Manifest) ([]byte, error)
	marshalManifestMutex       sync.RWMutex
	marshalManifestArgsForCall []struct {
		arg1 manifestparser.Manifest
	}
	marshalManifestReturns struct {
		result1 []byte
		result2 error
	}
	marshalManifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestParser) InterpolateAndParse(arg1 string, arg2 []string, arg3 []template.VarKV) (manifestparser.Manifest, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []template.VarKV
	if arg3 != nil {
		arg3Copy = make([]template.VarKV, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.interpolateAndParseMutex.Lock()
	ret, specificReturn := fake.interpolateAndParseReturnsOnCall[len(fake.interpolateAndParseArgsForCall)]
	fake.interpolateAndParseArgsForCall = append(fake.interpolateAndParseArgsForCall, struct {
		arg1 string
		arg2 []string
		arg3 []template.VarKV
	}{arg1, arg2Copy, arg3Copy})
	fake.recordInvocation("InterpolateAndParse", []interface{}{arg1, arg2Copy, arg3Copy})
	fake.interpolateAndParseMutex.Unlock()
	if fake.InterpolateAndParseStub != nil {
		return fake.InterpolateAndParseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.interpolateAndParseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifestParser) InterpolateAndParseCallCount() int {
	fake.interpolateAndParseMutex.RLock()
	defer fake.interpolateAndParseMutex.RUnlock()
	return len(fake.interpolateAndParseArgsForCall)
}

func (fake *FakeManifestParser) InterpolateAndParseCalls(stub func(string, []string, []template.VarKV) (manifestparser.Manifest, error)) {
	fake.interpolateAndParseMutex.Lock()
	defer fake.interpolateAndParseMutex.Unlock()
	fake.InterpolateAndParseStub = stub
}

func (fake *FakeManifestParser) InterpolateAndParseArgsForCall(i int) (string, []string, []template.VarKV) {
	fake.interpolateAndParseMutex.RLock()
	defer fake.interpolateAndParseMutex.RUnlock()
	argsForCall := fake.interpolateAndParseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeManifestParser) InterpolateAndParseReturns(result1 manifestparser.Manifest, result2 error) {
	fake.interpolateAndParseMutex.Lock()
	defer fake.interpolateAndParseMutex.Unlock()
	fake.InterpolateAndParseStub = nil
	fake.interpolateAndParseReturns = struct {
		result1 manifestparser.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) InterpolateAndParseReturnsOnCall(i int, result1 manifestparser.Manifest, result2 error) {
	fake.interpolateAndParseMutex.Lock()
	defer fake.interpolateAndParseMutex.Unlock()
	fake.InterpolateAndParseStub = nil
	if fake.interpolateAndParseReturnsOnCall == nil {
		fake.interpolateAndParseReturnsOnCall = make(map[int]struct {
			result1 manifestparser.Manifest
			result2 error
		})
	}
	fake.interpolateAndParseReturnsOnCall[i] = struct {
		result1 manifestparser.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) MarshalManifest(arg1 manifestparser.Manifest) ([]byte, error) {
	fake.marshalManifestMutex.Lock()
	ret, specificReturn := fake.marshalManifestReturnsOnCall[len(fake.marshalManifestArgsForCall)]
	fake.marshalManifestArgsForCall = append(fake.marshalManifestArgsForCall, struct {
		arg1 manifestparser.Manifest
	}{arg1})
	fake.recordInvocation("MarshalManifest", []interface{}{arg1})
	fake.marshalManifestMutex.Unlock()
	if fake.MarshalManifestStub != nil {
		return fake.MarshalManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.marshalManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifestParser) MarshalManifestCallCount() int {
	fake.marshalManifestMutex.RLock()
	defer fake.marshalManifestMutex.RUnlock()
	return len(fake.marshalManifestArgsForCall)
}

func (fake *FakeManifestParser) MarshalManifestCalls(stub func(manifestparser.Manifest) ([]byte, error)) {
	fake.marshalManifestMutex.Lock()
	defer fake.marshalManifestMutex.Unlock()
	fake.MarshalManifestStub = stub
}

func (fake *FakeManifestParser) MarshalManifestArgsForCall(i int) manifestparser.Manifest {
	fake.marshalManifestMutex.RLock()
	defer fake.marshalManifestMutex.RUnlock()
	argsForCall := fake.marshalManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifestParser) MarshalManifestReturns(result1 []byte, result2 error) {
	fake.marshalManifestMutex.Lock()
	defer fake.marshalManifestMutex.Unlock()
	fake.MarshalManifestStub = nil
	fake.marshalManifestReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) MarshalManifestReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.marshalManifestMutex.Lock()
	defer fake.marshalManifestMutex.Unlock()
	fake.MarshalManifestStub = nil
	if fake.marshalManifestReturnsOnCall == nil {
		fake.marshalManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalManifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpolateAndParseMutex.RLock()
	defer fake.interpolateAndParseMutex.RUnlock()
	fake.marshalManifestMutex.RLock()
	defer fake.marshalManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifestParser) recordInvocation(key string, args []interface{}) {
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

var _ v7.ManifestParser = new(FakeManifestParser)