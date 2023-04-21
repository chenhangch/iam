package mysql

import (
	"github.com/chang144/ciam/internal/apiserver/store"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Screats() store.SecretStore {
	return newSecrets(ds)
}
