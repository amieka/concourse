// This file was generated by counterfeiter
package githubfakes

import (
	"net/http"
	"sync"

	"github.com/concourse/atc/auth/github"
)

type FakeClient struct {
	CurrentUserStub        func(*http.Client) (string, error)
	currentUserMutex       sync.RWMutex
	currentUserArgsForCall []struct {
		arg1 *http.Client
	}
	currentUserReturns struct {
		result1 string
		result2 error
	}
	currentUserReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	OrganizationsStub        func(*http.Client) ([]string, error)
	organizationsMutex       sync.RWMutex
	organizationsArgsForCall []struct {
		arg1 *http.Client
	}
	organizationsReturns struct {
		result1 []string
		result2 error
	}
	organizationsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	TeamsStub        func(*http.Client) (github.OrganizationTeams, error)
	teamsMutex       sync.RWMutex
	teamsArgsForCall []struct {
		arg1 *http.Client
	}
	teamsReturns struct {
		result1 github.OrganizationTeams
		result2 error
	}
	teamsReturnsOnCall map[int]struct {
		result1 github.OrganizationTeams
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CurrentUser(arg1 *http.Client) (string, error) {
	fake.currentUserMutex.Lock()
	ret, specificReturn := fake.currentUserReturnsOnCall[len(fake.currentUserArgsForCall)]
	fake.currentUserArgsForCall = append(fake.currentUserArgsForCall, struct {
		arg1 *http.Client
	}{arg1})
	fake.recordInvocation("CurrentUser", []interface{}{arg1})
	fake.currentUserMutex.Unlock()
	if fake.CurrentUserStub != nil {
		return fake.CurrentUserStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.currentUserReturns.result1, fake.currentUserReturns.result2
}

func (fake *FakeClient) CurrentUserCallCount() int {
	fake.currentUserMutex.RLock()
	defer fake.currentUserMutex.RUnlock()
	return len(fake.currentUserArgsForCall)
}

func (fake *FakeClient) CurrentUserArgsForCall(i int) *http.Client {
	fake.currentUserMutex.RLock()
	defer fake.currentUserMutex.RUnlock()
	return fake.currentUserArgsForCall[i].arg1
}

func (fake *FakeClient) CurrentUserReturns(result1 string, result2 error) {
	fake.CurrentUserStub = nil
	fake.currentUserReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CurrentUserReturnsOnCall(i int, result1 string, result2 error) {
	fake.CurrentUserStub = nil
	if fake.currentUserReturnsOnCall == nil {
		fake.currentUserReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.currentUserReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Organizations(arg1 *http.Client) ([]string, error) {
	fake.organizationsMutex.Lock()
	ret, specificReturn := fake.organizationsReturnsOnCall[len(fake.organizationsArgsForCall)]
	fake.organizationsArgsForCall = append(fake.organizationsArgsForCall, struct {
		arg1 *http.Client
	}{arg1})
	fake.recordInvocation("Organizations", []interface{}{arg1})
	fake.organizationsMutex.Unlock()
	if fake.OrganizationsStub != nil {
		return fake.OrganizationsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.organizationsReturns.result1, fake.organizationsReturns.result2
}

func (fake *FakeClient) OrganizationsCallCount() int {
	fake.organizationsMutex.RLock()
	defer fake.organizationsMutex.RUnlock()
	return len(fake.organizationsArgsForCall)
}

func (fake *FakeClient) OrganizationsArgsForCall(i int) *http.Client {
	fake.organizationsMutex.RLock()
	defer fake.organizationsMutex.RUnlock()
	return fake.organizationsArgsForCall[i].arg1
}

func (fake *FakeClient) OrganizationsReturns(result1 []string, result2 error) {
	fake.OrganizationsStub = nil
	fake.organizationsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) OrganizationsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.OrganizationsStub = nil
	if fake.organizationsReturnsOnCall == nil {
		fake.organizationsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.organizationsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Teams(arg1 *http.Client) (github.OrganizationTeams, error) {
	fake.teamsMutex.Lock()
	ret, specificReturn := fake.teamsReturnsOnCall[len(fake.teamsArgsForCall)]
	fake.teamsArgsForCall = append(fake.teamsArgsForCall, struct {
		arg1 *http.Client
	}{arg1})
	fake.recordInvocation("Teams", []interface{}{arg1})
	fake.teamsMutex.Unlock()
	if fake.TeamsStub != nil {
		return fake.TeamsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.teamsReturns.result1, fake.teamsReturns.result2
}

func (fake *FakeClient) TeamsCallCount() int {
	fake.teamsMutex.RLock()
	defer fake.teamsMutex.RUnlock()
	return len(fake.teamsArgsForCall)
}

func (fake *FakeClient) TeamsArgsForCall(i int) *http.Client {
	fake.teamsMutex.RLock()
	defer fake.teamsMutex.RUnlock()
	return fake.teamsArgsForCall[i].arg1
}

func (fake *FakeClient) TeamsReturns(result1 github.OrganizationTeams, result2 error) {
	fake.TeamsStub = nil
	fake.teamsReturns = struct {
		result1 github.OrganizationTeams
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TeamsReturnsOnCall(i int, result1 github.OrganizationTeams, result2 error) {
	fake.TeamsStub = nil
	if fake.teamsReturnsOnCall == nil {
		fake.teamsReturnsOnCall = make(map[int]struct {
			result1 github.OrganizationTeams
			result2 error
		})
	}
	fake.teamsReturnsOnCall[i] = struct {
		result1 github.OrganizationTeams
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.currentUserMutex.RLock()
	defer fake.currentUserMutex.RUnlock()
	fake.organizationsMutex.RLock()
	defer fake.organizationsMutex.RUnlock()
	fake.teamsMutex.RLock()
	defer fake.teamsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ github.Client = new(FakeClient)
