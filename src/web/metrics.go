package web

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func Metrics(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		allocatedBefore := memStats.Alloc

		startTime := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Println("video load took", time.Since(startTime))
		runtime.ReadMemStats(&memStats)
		fmt.Println((memStats.Alloc-allocatedBefore)/1024, "kB allocated")
	}
}
