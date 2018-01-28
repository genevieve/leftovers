package ec2

type KeyPair struct {
	client     keyPairsClient
	name       *string
	identifier string
}

func NewKeyPair(client keyPairsClient, name *string) KeyPair {
	return KeyPair{
		client:     client,
		name:       name,
		identifier: *name,
	}
}
