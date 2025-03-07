// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"kakio/GoStudy/Study/Start/ch1/lissajous"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/image", imageGif)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
// handler 是一个HTTP请求处理函数，用于处理传入的HTTP请求并生成响应。
func handler(w http.ResponseWriter, r *http.Request) {
	// 使用fmt.Fprintf将请求的方法、URL和协议版本写入响应。
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// 遍历请求头中的所有键值对，并将它们写入响应。
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	// 将请求的主机名写入响应。
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	// 将请求的远程地址写入响应。
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// 解析请求的表单数据。如果解析过程中出现错误，则打印错误信息。
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// 遍历解析后的表单数据中的所有键值对，并将它们写入响应。
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func imageGif(w http.ResponseWriter, r *http.Request) {
	cyclesStr := r.FormValue("cycles")
	cycles, err := strconv.Atoi(cyclesStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	lissajous.Lissajous(w, cycles)
}
