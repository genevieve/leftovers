package leftovers

type Deleter interface {
	// Delete filters and deletes every resource,
	// return an error if any happens.
	Delete(filter string) error
}
