package config

import "github.com/chang144/ciam/internal/apiserver/options"

type Config struct {
	*options.ApiServerOptions
}

// CreateConfigFormOptions 基于给定的IAM泵命令行或配置文件选项创建一个正在运行的配置实例。
func CreateConfigFormOptions(opts *options.ApiServerOptions) (*Config, error) {
	return &Config{opts}, nil
}
