package iam

type ServerCertificate struct {
	client     serverCertificatesClient
	name       *string
	identifier string
}

func NewServerCertificate(client serverCertificatesClient, name *string) ServerCertificate {
	return ServerCertificate{
		client:     client,
		name:       name,
		identifier: *name,
	}
}
