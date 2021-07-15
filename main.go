package main

import (
	"github.com/gastrodon/groudon"
	"git.gastrodon.io/imonke/monkebase"

	"log"
	"net/http"
	"os"
)

func main() {
	monkebase.Connect(os.Getenv("DATABASE_CONNECTION"))
	groudon.RegisterHandler("POST", "^/$", postUser)
	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
