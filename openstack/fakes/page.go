package fakes

type Page struct {
	Name string
}

func (page Page) NextPageURL() (string, error) {
	return "", nil
}

func (page Page) IsEmpty() (bool, error) {
	return false, nil
}

func (page Page) GetBody() interface{} {
	return nil
}
