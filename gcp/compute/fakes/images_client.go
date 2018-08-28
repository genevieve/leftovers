package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type ImagesClient struct {
	ListImagesCall struct {
		CallCount int
		Returns   struct {
			Output []*gcpcompute.Image
			Error  error
		}
	}

	DeleteImageCall struct {
		CallCount int
		Receives  struct {
			Image string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *ImagesClient) ListImages() ([]*gcpcompute.Image, error) {
	n.ListImagesCall.CallCount++

	return n.ListImagesCall.Returns.Output, n.ListImagesCall.Returns.Error
}

func (n *ImagesClient) DeleteImage(image string) error {
	n.DeleteImageCall.CallCount++
	n.DeleteImageCall.Receives.Image = image

	return n.DeleteImageCall.Returns.Error
}
