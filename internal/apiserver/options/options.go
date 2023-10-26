package options

import (
	"github.com/chang144/golunzi/cli"
	generaloptions "github.com/chang144/iam/internal/pkg/options"
	"github.com/chang144/iam/internal/pkg/options/serverOptions"
	"github.com/chang144/iam/pkg/log"
)

// ApiServerOptions api server所需的所有功能选项
type ApiServerOptions struct {
	// logicServer
	GenericServerRunOptions *serverOptions.ServerRunOptions       `json:"logicServer" mapstructure:"logicServer"`
	InsecureServing         *serverOptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *serverOptions.SecureServingOptions   `json:"secure" mapstructure:"secure"`

	// database
	MySQLOptions *generaloptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`

	Log *log.Options `json:"log" mapstructure:"log"`
}

// NewOptions returns a new ApiServerOptions
func NewOptions() *ApiServerOptions {
	o := &ApiServerOptions{
		GenericServerRunOptions: serverOptions.NewServerRunOptions(),
		InsecureServing:         serverOptions.NewInsecureServingOptions(),
		SecureServing:           serverOptions.NewSecureServingOptions(),

		MySQLOptions: generaloptions.NewMySQLOptions(),
		Log:          log.NewOptions(),
	}
	return o
}

// Flags 为apiserver添加flag
func (o *ApiServerOptions) Flags() (fss cli.AppFlagSets) {
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))

	o.GenericServerRunOptions.AddFlags(fss.FlagSet("logicServer"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure"))
	o.SecureServing.AddFlags(fss.FlagSet("secure"))

	o.Log.AddFlags(fss.FlagSet("log"))
	return
}
