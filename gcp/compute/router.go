package compute

type Router struct {
	routersClient routersClient
}

func NewRouter(routersClient routersClient) Router {
	return Router{
		routersClient: routersClient,
	}
}

func (r Router) Delete() error {
	return nil
}

func (r Router) Type() string {
	return ""
}

func (r Router) Name() string {
	return ""
}
