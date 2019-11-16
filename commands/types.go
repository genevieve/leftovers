package commands

import "github.com/genevieve/leftovers/app"

type Types struct {
	leftovers leftovers
}

func NewTypes(l leftovers) Types {
	return Types{
		leftovers: l,
	}
}

func (t Types) Execute(o app.Options) error {
	t.leftovers.Types()
	return nil
}
