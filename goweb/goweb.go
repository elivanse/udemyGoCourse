package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(" metodo " + r.Method)
	fmt.Fprintln(rw, "hola")
}

func PageNf(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}
func main() {

	http.HandleFunc("/", Hola)
	http.HandleFunc("/page", Hola)

	fmt.Println("corriendo en 6666")
	fmt.Println("run server: http://localhost:6666/")
	log.Fatal(http.ListenAndServe("127.0.0.1:6666", nil))

}
