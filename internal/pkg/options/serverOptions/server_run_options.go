package serverOptions

import (
	"github.com/chang144/iam/internal/pkg/server"
	"github.com/spf13/pflag"
)

// ServerRunOptions contains the options while a generic api logicServer is running
type ServerRunOptions struct {
	Mode        string   `json:"mode" mapstructure:"mode"`
	Healthz     bool     `json:"healthz" mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	defaultServerOptions := server.NewNilConfig()

	return &ServerRunOptions{
		Mode:        defaultServerOptions.Mode,
		Healthz:     defaultServerOptions.Healthz,
		Middlewares: defaultServerOptions.Middlewares,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Mode = s.Mode
	c.Healthz = s.Healthz
	c.Middlewares = s.Middlewares

	return nil
}

// Validate checks validation of ServerRunOptions.
func (s *ServerRunOptions) Validate() []error {
	errors := []error{}

	return errors
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet.
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs.StringVar(&s.Mode, "logicServer.mode", s.Mode, ""+
		"Start the logicServer in a specified logicServer mode. Supported logicServer mode: debug, test, release.")

	fs.BoolVar(&s.Healthz, "logicServer.healthz", s.Healthz, ""+
		"Add self readiness check and install /healthz router.")

	fs.StringSliceVar(&s.Middlewares, "logicServer.middlewares", s.Middlewares, ""+
		"List of allowed middlewares for logicServer, comma separated. If this list is empty default middlewares will be used.")
}
