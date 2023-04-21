package mysql

import "gorm.io/gorm"

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{db: ds.db}
}

func (u *users) Create() error {
	return nil
}

func (u users) List() error {
	return nil
}
