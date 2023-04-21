package options

import "encoding/json"

// String 返回并打印api server的选项
func (o *ApiServerOptions) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
