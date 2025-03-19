package main

import (
	"kakio/ch3/surface"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	width := parseQueryParam(r, "width", 600)   // 默认宽度 600
	height := parseQueryParam(r, "height", 320) // 默认高度 320
	color := r.URL.Query().Get("color")         // 默认颜色为空，使用 surface 包的默认颜色逻辑

	// 设置 Content-Type 头部
	w.Header().Set("Content-Type", "image/svg+xml")

	// 调用 surface 包的绘制函数
	surface.DrawSurfaceWithParams(w, width, height, color)
}

// parseQueryParam 解析请求参数并返回整数值，如果参数不存在或无效，则返回默认值
func parseQueryParam(r *http.Request, key string, defaultValue int) int {
	valStr := r.URL.Query().Get(key)
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return val
}
