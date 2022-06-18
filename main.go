package main

import "net/http"

func main(){
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9091", nil)
}

func hello(w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("Home!"))
}