package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()
	fmt.Println("Executando api!")
	log.Fatal(http.ListenAndServe(":5000", r))
}
