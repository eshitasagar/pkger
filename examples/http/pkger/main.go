package main

import (
	"log"
	"net/http"

	"github.com/markbates/pkger"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dir, err := pkger.HTTP("/public")
	if err != nil {
		return err
	}
	return http.ListenAndServe(":3000", dir)
}
