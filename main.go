package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib"
	"github.com/gastrodon/groudon"

	"log"
	"net/http"
	"os"
)

const (
	uuid_regex = `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`
)

func main() {
	monkebase.Connect(os.Getenv("DATABASE_CONNECTION"))
	groudon.RegisterHandler("GET", "^/id/"+uuid_regex+"/?$", getUserID)
	groudon.RegisterHandler("GET", "^/nick/"+monkelib.NICK_PATTERN+"/?$", getUserNick)
	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
