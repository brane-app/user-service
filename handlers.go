package main

import (
	"github.com/gastrodon/groudon/v2"

	"os"
)

const (
	uuid_regex = `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`
)

var (
	prefix = os.Getenv("PATH_PREFIX")

	routeId   = "^" + prefix + "/id/" + uuid_regex + "/?$"
	routeNick = "^" + prefix + "/nick/.+/?$"
)

func register_handlers() {
	groudon.AddHandler("GET", routeId, getUserID)
	groudon.AddHandler("GET", routeNick, getUserNick)
}
