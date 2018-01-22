package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type UrlMapsClient struct {
	ListUrlMapsCall struct {
		CallCount int
		Returns   struct {
			Output *gcpcompute.UrlMapList
			Error  error
		}
	}

	DeleteUrlMapCall struct {
		CallCount int
		Receives  struct {
			UrlMap string
		}
		Returns struct {
			Error error
		}
	}
}

func (u *UrlMapsClient) ListUrlMaps() (*gcpcompute.UrlMapList, error) {
	u.ListUrlMapsCall.CallCount++

	return u.ListUrlMapsCall.Returns.Output, u.ListUrlMapsCall.Returns.Error
}

func (u *UrlMapsClient) DeleteUrlMap(urlMap string) error {
	u.DeleteUrlMapCall.CallCount++
	u.DeleteUrlMapCall.Receives.UrlMap = urlMap

	return u.DeleteUrlMapCall.Returns.Error
}
