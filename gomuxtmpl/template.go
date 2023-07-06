package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Curso struct {
	Nombre string
}

type Usuarios struct {
	User   string
	Edad   int
	Activo bool
	Admin  bool
	Cursos []Curso
}

// destinos de mux
func Index(rw http.ResponseWriter, r *http.Request) {
	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"Javascript"}
	c5 := Curso{"Inteligencia Artificial"}
	// parse files devuelve dos variable un tmpl y un err
	template, err := template.ParseFiles("index.html")
	cursos := []Curso{c1, c2, c3, c4, c5}
	usuario := Usuarios{"Milo", 5, true, false, cursos}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, usuario)
	}

}

func main() {

	//mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	//servidoprt
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Server ranin")
	log.Fatal(server.ListenAndServe())

}
