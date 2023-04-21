package store

type UserStore interface {
	Create() error
	List() error
}
