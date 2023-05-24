package fakes

import (
	"sync"

	"google.golang.org/api/artifactregistry/v1"
)

type RepositoriesClient struct {
	DeleteRepositoryCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Cluster string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListRepositoriesCall struct {
		mutex     sync.Mutex
		CallCount int
		Returns   struct {
			ListRepositoriesResponse *artifactregistry.ListRepositoriesResponse
			Error                    error
		}
		Stub func() (*artifactregistry.ListRepositoriesResponse, error)
	}
}

func (f *RepositoriesClient) DeleteRepository(param1 string) error {
	f.DeleteRepositoryCall.mutex.Lock()
	defer f.DeleteRepositoryCall.mutex.Unlock()
	f.DeleteRepositoryCall.CallCount++
	f.DeleteRepositoryCall.Receives.Cluster = param1
	if f.DeleteRepositoryCall.Stub != nil {
		return f.DeleteRepositoryCall.Stub(param1)
	}
	return f.DeleteRepositoryCall.Returns.Error
}
func (f *RepositoriesClient) ListRepositories() (*artifactregistry.ListRepositoriesResponse, error) {
	f.ListRepositoriesCall.mutex.Lock()
	defer f.ListRepositoriesCall.mutex.Unlock()
	f.ListRepositoriesCall.CallCount++
	if f.ListRepositoriesCall.Stub != nil {
		return f.ListRepositoriesCall.Stub()
	}
	return f.ListRepositoriesCall.Returns.ListRepositoriesResponse, f.ListRepositoriesCall.Returns.Error
}
