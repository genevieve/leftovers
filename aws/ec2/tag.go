package ec2

type Tag struct {
	client     tagsClient
	key        *string
	value      *string
	resourceId *string
	identifier string
}

func NewTag(client tagsClient, key, value, resourceId *string) Tag {
	return Tag{
		client:     client,
		key:        key,
		value:      value,
		resourceId: resourceId,
		identifier: *value,
	}
}
