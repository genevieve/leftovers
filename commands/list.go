package commands

import "github.com/genevieve/leftovers/app"

type List struct {
	leftovers leftovers
}

func NewList(l leftovers) List {
	return List{
		leftovers: l,
	}
}

func (l List) Execute(o app.Options) error {
	if o.Type == "" {
		l.leftovers.List(o.Filter)
	} else {
		l.leftovers.ListByType(o.Filter, o.Type)
	}
	return nil
}
