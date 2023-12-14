package mysql

import (
	"fmt"
	"sync"

	v1 "github.com/chenhangch/go-api-module/iam/apiserver/v1"
	"github.com/chenhangch/golunzi/errors"
	"github.com/chenhangch/iam/internal/pkg/logger"

	"github.com/chenhangch/iam/internal/apiserver/store"
	"github.com/chenhangch/iam/internal/pkg/options"
	"github.com/chenhangch/iam/pkg/db"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Secrets() store.SecretStore {
	return newSecrets(ds)

}

// Close closes database
func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts *options.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store factory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.DatabaseOptions{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = db.New(options)

		mysqlFactory = &datastore{dbIns}
	})
	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store factory")
	}

	return mysqlFactory, nil
}

// cleanDatabase clean database
func cleanDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&v1.User{}); err != nil {
		return errors.Wrap(err, "drop user table failed")
	}

	return nil
}

// migrateDatabase auto migration database
func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&v1.User{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}

	return nil
}
