package s3

type Bucket struct {
	client     bucketsClient
	name       *string
	identifier string
}

func NewBucket(client bucketsClient, name *string) Bucket {
	return Bucket{
		client:     client,
		name:       name,
		identifier: *name,
	}
}
