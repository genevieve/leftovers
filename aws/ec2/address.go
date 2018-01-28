package ec2

type Address struct {
	client       addressesClient
	publicIp     *string
	allocationId *string
	identifier   string
}

func NewAddress(client addressesClient, publicIp, allocationId *string) Address {
	return Address{
		client:       client,
		publicIp:     publicIp,
		allocationId: allocationId,
		identifier:   *publicIp,
	}
}
