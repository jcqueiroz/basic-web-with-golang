package main

import (
	"ALURA/jcqueiroz/basic-web-with-golang/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
