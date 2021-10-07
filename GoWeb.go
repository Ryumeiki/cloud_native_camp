package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type User struct {
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Age         int         `json:"age"`
	Hobbies     []string    `json:"hobbies"`
}

var ss string

func headerHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Fprintf(w, "%s: %v\n", key, value)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
<html>
    <head>
        <title>Go Web 编程之 request</title>
    </head>
    <body>
        <form method="post" action="/healthz">
            <label for="username">用户名：</label>
            <input type="text" id="username" name="username">
            <label for="email">邮箱：</label>
            <input type="text" id="email" name="email">
            <button type="submit">提交</button>
        </form>
    </body></html>`)
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, r.ContentLength)
	r.Body.Read(data) // 忽略错误处理
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	for key, value := range r.Header {
		ss = strings.Join(value, "")
		//fmt.Fprintf(w, "%s\n", ss)
		w.Header().Set(key, ss)
		//fmt.Fprintf(w, "%s: %v\n", key, value)
	}
	w.WriteHeader(200)
	w.Write([]byte("Thanks for your visit :)\n"))

}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web 编程之 响应</title></head>
<body><h1>直接使用 Write 方法<h1></body>
</html>`
	w.Write([]byte(str))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Name", "my name is smallsoup")
	u := &User {
		FirstName:  "Meiki",
		LastName:   "Ryu",
		Age:        40,
		Hobbies:    []string{"Reading", "Cooking"},
	}
	data, _ := json.Marshal(u)
	w.Write(data)
}

func main() {
	mux := http.NewServeMux()
	// 注册
	mux.HandleFunc("/header", headerHandler)
	mux.HandleFunc("/write", writeHandler)
	mux.HandleFunc("/json", jsonHandler)
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/healthz", bodyHandler)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}