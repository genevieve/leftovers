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
		Receives  struct {
			Region string
		}
		Returns struct {
			ListRepositoriesResponse []*artifactregistry.Repository
			Error                    error
		}
		Stub func(string) ([]*artifactregistry.Repository, error)
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
func (f *RepositoriesClient) ListRepositories(param1 string) ([]*artifactregistry.Repository, error) {
	f.ListRepositoriesCall.mutex.Lock()
	defer f.ListRepositoriesCall.mutex.Unlock()
	f.ListRepositoriesCall.CallCount++
	f.ListRepositoriesCall.Receives.Region = param1
	if f.ListRepositoriesCall.Stub != nil {
		return f.ListRepositoriesCall.Stub(param1)
	}
	return f.ListRepositoriesCall.Returns.ListRepositoriesResponse, f.ListRepositoriesCall.Returns.Error
}
