package commands

type leftovers interface {
	Delete(filter string, regex bool) error
	DeleteByType(filter, rType string, regex bool) error
	List(filter string, regex bool)
	ListByType(filter, rType string, regex bool)
	Types()
}
