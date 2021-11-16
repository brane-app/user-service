package main

import (
	"github.com/brane-app/user-service/handlers"

	"github.com/gastrodon/groudon/v2"

	"os"
)

const (
	UUID_PATTERN = "[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}"
)

var (
	prefix = os.Getenv("PATH_PREFIX")

	routeRoot = "^" + prefix + "/?$"
	routeId   = "^" + prefix + "/id/" + UUID_PATTERN + "/?$"
	routeNick = "^" + prefix + "/nick/.+/?$"
)

func register_handlers() {
	groudon.AddHandler("POST", routeRoot, handlers.PostUser)
	groudon.AddHandler("GET", routeId, handlers.GetUserID)
	groudon.AddHandler("GET", routeNick, handlers.GetUserNick)
}
