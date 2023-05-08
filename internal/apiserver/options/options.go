package options

import (
	generaloptions "github.com/chang144/ciam/internal/pkg/options"
	"github.com/chang144/ciam/internal/pkg/options/serverOptions"
	"github.com/chang144/ciam/pkg/log"
	"github.com/chang144/golunzi/cli"
)

// ApiServerOptions api server所需的所有功能选项
type ApiServerOptions struct {
	// server
	GenericServerRunOptions *serverOptions.ServerRunOptions       `json:"server" mapstructure:"server"`
	InsecureServing         *serverOptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *serverOptions.SecureServingOptions   `json:"secure" mapstructure:"secure"`

	// database
	MySQLOptions *generaloptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`

	Log *log.Options `json:"log" mapstructure:"log"`
}

// NewOptions returns a new ApiServerOptions
func NewOptions() *ApiServerOptions {
	o := &ApiServerOptions{
		MySQLOptions: generaloptions.NewMySQLOptions(),

		Log: log.NewOptions(),
	}
	return o
}

// Flags 为apiserver添加flag
func (o *ApiServerOptions) Flags() (fss cli.AppFlagSets) {
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.Log.AddFlags(fss.FlagSet("log"))
	return
}
