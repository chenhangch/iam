package mysql

import (
	"context"

	v1 "github.com/chenhangch/go-api-module/iam/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func (u *users) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *users) List(ctx context.Context, options *metav1.ListOptions) (*v1.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func newUsers(ds *datastore) *users {
	return &users{db: ds.db}
}
