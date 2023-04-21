package server

import (
	"log"
	"net/http"
)

// ServeHealthCheck 运行一个http服务器，用来提供一个api来检查泵的健康状态。
func ServeHealthCheck(healthPath string, healthAddress string) {
	http.HandleFunc("/"+healthAddress, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`"status": "OK"`))
	})

	if err := http.ListenAndServe(healthAddress, nil); err != nil {
		log.Fatalf("error serving health check endpoint: %s", err.Error())
	}
}
