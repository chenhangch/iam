package options

import (
	"time"

	"github.com/chenhangch/iam/pkg/db"
	"github.com/spf13/pflag"
	"gorm.io/gorm"
)

type MySQLOptions struct {
	Host                  string        `json:"host,omitempty" mapstructure:"host"`
	Username              string        `json:"username,omitempty" mapstructure:"username"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	LogLevel              int           `json:"log-level" mapstructure:"log-lecel"`
}

// NewMySQLOptions 创建一个空值的MySQLOptions实例
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:                  "127.0.0.1:3306",
		Username:              "",
		Password:              "",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
		LogLevel:              1,
	}
}

func (o *MySQLOptions) Validate() []error {
	var errs []error

	return errs
}

// AddFlags 添加mysql的命令行参数
func (o *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", "localhost", "mysql host")
	fs.StringVar(&o.Username, "mysql.username", o.Username, "mysql username")
	fs.StringVar(&o.Password, "mysql.password", o.Password, "mysql password")
	fs.StringVar(&o.Database, "mysql.database", o.Database, "mysql database")
	fs.IntVar(&o.MaxIdleConnections, "mysql.max-idle-connections", o.MaxIdleConnections, "mysql max idle connections")
	fs.IntVar(&o.MaxOpenConnections, "mysql-max-open-connections", o.MaxOpenConnections, "mysql max open connections")
	fs.DurationVar(&o.MaxConnectionLifeTime, "mysql-max-connection-life-time", o.MaxConnectionLifeTime, "mysql max connection life time")
	fs.IntVar(&o.LogLevel, "mysql.log-mode", o.LogLevel, "Specify gorm log level")
}

// NewClient 创建一个gorm的db实例
func (o *MySQLOptions) NewClient() (*gorm.DB, error) {
	opts := &db.DatabaseOptions{
		Host:                  o.Host,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		LogLevel:              o.LogLevel,
		Logger:                nil,
	}

	return db.New(opts)
}
