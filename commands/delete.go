package commands

import "github.com/genevieve/leftovers/app"

type Delete struct {
	leftovers leftovers
}

func NewDelete(l leftovers) Delete {
	return Delete{
		leftovers: l,
	}
}

func (d Delete) Execute(o app.Options) error {
	if o.Type == "" {
		return d.leftovers.Delete(o.Filter, o.RegexFiltered)
	}

	return d.leftovers.DeleteByType(o.Filter, o.Type, o.RegexFiltered)
}
