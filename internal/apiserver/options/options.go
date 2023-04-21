package options

import (
	generaloptions "github.com/chang144/ciam/internal/pkg/options"
	"github.com/chang144/golunzi/cli"
)

type ApiServerOptions struct {
	MySQLOptions *generaloptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewOptions returns a new ApiServerOptions
func NewOptions() *ApiServerOptions {
	o := &ApiServerOptions{
		MySQLOptions: generaloptions.NewMySQLOptions(),
	}
	return o
}

// Flags 为apiserver添加flag
func (o *ApiServerOptions) Flags() (fss cli.AppFlagSets) {
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	return
}
