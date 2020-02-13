// Code generated by counterfeiter. DO NOT EDIT.
package backendfakes

import (
	"sync"

	"github.com/concourse/concourse/worker/backend"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeRootfsManager struct {
	SetupCwdStub        func(*specs.Spec, string) error
	setupCwdMutex       sync.RWMutex
	setupCwdArgsForCall []struct {
		arg1 *specs.Spec
		arg2 string
	}
	setupCwdReturns struct {
		result1 error
	}
	setupCwdReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRootfsManager) SetupCwd(arg1 *specs.Spec, arg2 string) error {
	fake.setupCwdMutex.Lock()
	ret, specificReturn := fake.setupCwdReturnsOnCall[len(fake.setupCwdArgsForCall)]
	fake.setupCwdArgsForCall = append(fake.setupCwdArgsForCall, struct {
		arg1 *specs.Spec
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetupCwd", []interface{}{arg1, arg2})
	fake.setupCwdMutex.Unlock()
	if fake.SetupCwdStub != nil {
		return fake.SetupCwdStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setupCwdReturns
	return fakeReturns.result1
}

func (fake *FakeRootfsManager) SetupCwdCallCount() int {
	fake.setupCwdMutex.RLock()
	defer fake.setupCwdMutex.RUnlock()
	return len(fake.setupCwdArgsForCall)
}

func (fake *FakeRootfsManager) SetupCwdCalls(stub func(*specs.Spec, string) error) {
	fake.setupCwdMutex.Lock()
	defer fake.setupCwdMutex.Unlock()
	fake.SetupCwdStub = stub
}

func (fake *FakeRootfsManager) SetupCwdArgsForCall(i int) (*specs.Spec, string) {
	fake.setupCwdMutex.RLock()
	defer fake.setupCwdMutex.RUnlock()
	argsForCall := fake.setupCwdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRootfsManager) SetupCwdReturns(result1 error) {
	fake.setupCwdMutex.Lock()
	defer fake.setupCwdMutex.Unlock()
	fake.SetupCwdStub = nil
	fake.setupCwdReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRootfsManager) SetupCwdReturnsOnCall(i int, result1 error) {
	fake.setupCwdMutex.Lock()
	defer fake.setupCwdMutex.Unlock()
	fake.SetupCwdStub = nil
	if fake.setupCwdReturnsOnCall == nil {
		fake.setupCwdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupCwdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRootfsManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setupCwdMutex.RLock()
	defer fake.setupCwdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRootfsManager) recordInvocation(key string, args []interface{}) {
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

var _ backend.RootfsManager = new(FakeRootfsManager)
