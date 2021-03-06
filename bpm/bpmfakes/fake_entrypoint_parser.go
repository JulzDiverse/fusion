// Code generated by counterfeiter. DO NOT EDIT.
package bpmfakes

import (
	"io"
	"sync"

	"github.com/JulzDiverse/fusion/bpm"
	"github.com/JulzDiverse/fusion/docker"
)

type FakeEntrypointParser struct {
	ParseDockerfileEntrypointStub        func(io.Reader) (docker.Entrypoint, error)
	parseDockerfileEntrypointMutex       sync.RWMutex
	parseDockerfileEntrypointArgsForCall []struct {
		arg1 io.Reader
	}
	parseDockerfileEntrypointReturns struct {
		result1 docker.Entrypoint
		result2 error
	}
	parseDockerfileEntrypointReturnsOnCall map[int]struct {
		result1 docker.Entrypoint
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEntrypointParser) ParseDockerfileEntrypoint(arg1 io.Reader) (docker.Entrypoint, error) {
	fake.parseDockerfileEntrypointMutex.Lock()
	ret, specificReturn := fake.parseDockerfileEntrypointReturnsOnCall[len(fake.parseDockerfileEntrypointArgsForCall)]
	fake.parseDockerfileEntrypointArgsForCall = append(fake.parseDockerfileEntrypointArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	fake.recordInvocation("ParseDockerfileEntrypoint", []interface{}{arg1})
	fake.parseDockerfileEntrypointMutex.Unlock()
	if fake.ParseDockerfileEntrypointStub != nil {
		return fake.ParseDockerfileEntrypointStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.parseDockerfileEntrypointReturns.result1, fake.parseDockerfileEntrypointReturns.result2
}

func (fake *FakeEntrypointParser) ParseDockerfileEntrypointCallCount() int {
	fake.parseDockerfileEntrypointMutex.RLock()
	defer fake.parseDockerfileEntrypointMutex.RUnlock()
	return len(fake.parseDockerfileEntrypointArgsForCall)
}

func (fake *FakeEntrypointParser) ParseDockerfileEntrypointArgsForCall(i int) io.Reader {
	fake.parseDockerfileEntrypointMutex.RLock()
	defer fake.parseDockerfileEntrypointMutex.RUnlock()
	return fake.parseDockerfileEntrypointArgsForCall[i].arg1
}

func (fake *FakeEntrypointParser) ParseDockerfileEntrypointReturns(result1 docker.Entrypoint, result2 error) {
	fake.ParseDockerfileEntrypointStub = nil
	fake.parseDockerfileEntrypointReturns = struct {
		result1 docker.Entrypoint
		result2 error
	}{result1, result2}
}

func (fake *FakeEntrypointParser) ParseDockerfileEntrypointReturnsOnCall(i int, result1 docker.Entrypoint, result2 error) {
	fake.ParseDockerfileEntrypointStub = nil
	if fake.parseDockerfileEntrypointReturnsOnCall == nil {
		fake.parseDockerfileEntrypointReturnsOnCall = make(map[int]struct {
			result1 docker.Entrypoint
			result2 error
		})
	}
	fake.parseDockerfileEntrypointReturnsOnCall[i] = struct {
		result1 docker.Entrypoint
		result2 error
	}{result1, result2}
}

func (fake *FakeEntrypointParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseDockerfileEntrypointMutex.RLock()
	defer fake.parseDockerfileEntrypointMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEntrypointParser) recordInvocation(key string, args []interface{}) {
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

var _ bpm.EntrypointParser = new(FakeEntrypointParser)
