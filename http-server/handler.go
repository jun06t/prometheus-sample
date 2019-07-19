package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func alive(w http.ResponseWriter, _ *http.Request) {
	dur := rand.Intn(1000)
	time.Sleep(time.Duration(dur) * time.Millisecond) // 処理を表現するためのsleep
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func hello(w http.ResponseWriter, _ *http.Request) {
	dur := rand.Intn(1000)
	time.Sleep(time.Duration(dur) * time.Millisecond) // 処理を表現するためのsleep
	n := rand.Intn(4)                                 // エラーレスポンスを返すためのランダム値
	switch n {
	case 0:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello World")
	case 1:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not Found")
	case 2:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
	case 3:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}
}
