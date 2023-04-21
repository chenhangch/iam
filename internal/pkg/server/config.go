package server

import (
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"time"
)

const (
	// RecommendedHomeDir defines the default directory used to place all iam service configurations.
	RecommendedHomeDir = ".ciam"

	// RecommendedEnvPrefix defines the ENV prefix used by all iam service.
	RecommendedEnvPrefix = "CIAM"
)

// Config 是一个用于配置genericapisserver的结构。
// 它的成员大致按组件的重要性排序。
type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

// CertKey 包含证书相关的配置项。
type CertKey struct {
	// CertFile 是一个包含pem编码证书的文件，可能包含完整的证书链
	CertFile string
	// KeyFile 是一个包含pem编码的证书私钥的文件，由CertFile指定
	KeyFile string
}

// SecureServingInfo 保存TLS服务器的配置。
type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey
}

// Address 将主机IP地址和端口号组合成一个地址字符串，如:0.0.0.0:8443。
func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

// InsecureServingInfo 保存不安全http服务器的配置。
type InsecureServingInfo struct {
	Address string
}

// JwtInfo 定义用于创建jwt认证中间件的jwt字段。
type JwtInfo struct {
	// defaults to "iam jwt"
	Realm string
	// defaults to empty
	Key string
	// defaults to one hour
	Timeout time.Duration
	// defaults to zero
	MaxRefresh time.Duration
}

// NewConfig 返回一个带着默认值的 Config 的实例
func NewConfig() *Config {
	return &Config{
		SecureServing:   nil,
		InsecureServing: nil,
		Jwt: &JwtInfo{
			Realm:      "ciam jwt",
			Key:        "",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
		},
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		Healthz:         true,
		EnableProfiling: true,
		EnableMetrics:   true,
	}
}
