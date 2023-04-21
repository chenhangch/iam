package server

import (
	"github.com/chang144/ciam/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/util/homedir"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

// CompletedConfig 是GenericAPIServer的完整配置。
type CompletedConfig struct {
	*Config
}

// Complete 完整填写任何未设置的字段，这些字段需要有有效的数据，并且可以派生
// 从其他字段。如果你要去ApplyOptions，先做那个。它让接收器发生变异。
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// New 从给定的配置返回一个GenericAPIServer的新实例。
func (c CompletedConfig) New() (*GenericAPIServer, error) {
	// setMode before gin.New()
	gin.SetMode(c.Mode)

	s := &GenericAPIServer{
		SecureServingInfo:   c.SecureServing,
		InsecureServingInfo: c.InsecureServing,
		healthz:             c.Healthz,
		enableMetrics:       c.EnableMetrics,
		enableProfiling:     c.EnableProfiling,
		middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	initGenericAPIServer(s)

	return s, nil
}

// LoadConfig 读取配置文件和ENV变量，如果设置了。.
func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Join(homedir.HomeDir(), RecommendedHomeDir))
		viper.AddConfigPath("/etc/iam")
		viper.SetConfigName(defaultName)
	}

	// Use config file from the flag.
	viper.SetConfigType("yaml")              // set the type of the configuration to yaml.
	viper.AutomaticEnv()                     // read in environment variables that match.
	viper.SetEnvPrefix(RecommendedEnvPrefix) // set ENVIRONMENT variables prefix to IAM.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("WARNING: viper failed to discover and load the configuration file: %s", err.Error())
	}
}
