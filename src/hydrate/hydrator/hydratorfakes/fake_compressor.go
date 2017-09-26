// Code generated by counterfeiter. DO NOT EDIT.
package hydratorfakes

import (
	"hydrate/hydrator"
	"sync"
)

type FakeCompressor struct {
	WriteTgzStub        func(string, string) error
	writeTgzMutex       sync.RWMutex
	writeTgzArgsForCall []struct {
		arg1 string
		arg2 string
	}
	writeTgzReturns struct {
		result1 error
	}
	writeTgzReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCompressor) WriteTgz(arg1 string, arg2 string) error {
	fake.writeTgzMutex.Lock()
	ret, specificReturn := fake.writeTgzReturnsOnCall[len(fake.writeTgzArgsForCall)]
	fake.writeTgzArgsForCall = append(fake.writeTgzArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("WriteTgz", []interface{}{arg1, arg2})
	fake.writeTgzMutex.Unlock()
	if fake.WriteTgzStub != nil {
		return fake.WriteTgzStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writeTgzReturns.result1
}

func (fake *FakeCompressor) WriteTgzCallCount() int {
	fake.writeTgzMutex.RLock()
	defer fake.writeTgzMutex.RUnlock()
	return len(fake.writeTgzArgsForCall)
}

func (fake *FakeCompressor) WriteTgzArgsForCall(i int) (string, string) {
	fake.writeTgzMutex.RLock()
	defer fake.writeTgzMutex.RUnlock()
	return fake.writeTgzArgsForCall[i].arg1, fake.writeTgzArgsForCall[i].arg2
}

func (fake *FakeCompressor) WriteTgzReturns(result1 error) {
	fake.WriteTgzStub = nil
	fake.writeTgzReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCompressor) WriteTgzReturnsOnCall(i int, result1 error) {
	fake.WriteTgzStub = nil
	if fake.writeTgzReturnsOnCall == nil {
		fake.writeTgzReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeTgzReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCompressor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeTgzMutex.RLock()
	defer fake.writeTgzMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCompressor) recordInvocation(key string, args []interface{}) {
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

var _ hydrator.Compressor = new(FakeCompressor)