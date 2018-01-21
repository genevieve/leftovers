package leftovers

type Deleter interface {
	// Delete deletes every resource named with environment,
	// return an error if any happens.
	Delete(environment string) error
}

