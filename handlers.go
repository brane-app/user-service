package main

import (
	"github.com/brane-app/user-service/handlers"

	"github.com/gastrodon/groudon/v2"

	"os"
)

const (
	UUID_PATTERN = "[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89AB][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}"
)

var (
	prefix = os.Getenv("PATH_PREFIX")

	routeRoot = "^" + prefix + "/?$"
	routeId   = "^" + prefix + "/id/" + UUID_PATTERN + "/?$"
	routeNick = "^" + prefix + "/nick/.+/?$"
	route404  = "^" + prefix + "/(?:nick|id)/[^/]+/?$"
)

func register_handlers() {
	groudon.AddHandler("POST", routeRoot, handlers.PostUser)

	groudon.AddHandler("GET", routeId, handlers.GetUserID)
	groudon.AddHandler("GET", routeNick, handlers.GetUserNick)
	groudon.AddHandler("GET", route404, handlers.GetUser404)
}
