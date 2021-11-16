package handlers

import (
	"github.com/brane-app/database-library"
	"github.com/brane-app/tools-library"
	"github.com/brane-app/types-library"

	"net/http"
)

var (
	no_such_user map[string]interface{} = map[string]interface{}{"error": "no_such_user"}
)

func GetUserID(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = getUserKey("id", request)
	return
}

func GetUserNick(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = getUserKey("nick", request)
	return
}

func getUserKey(key string, request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var parts []string = tools.SplitPath(request.URL.Path)
	var query string = parts[len(parts)-1]

	var fetched types.User
	var exists bool
	switch key {
	case "id":
		fetched, exists, err = database.ReadSingleUser(query)
	case "nick":
		fetched, exists, err = database.ReadSingleUserNick(query)
	}

	if err != nil {
		return
	}

	if !exists {
		code = 404
		r_map = no_such_user
		return
	}

	code = 200
	r_map = map[string]interface{}{"user": fetched.PublicMap()}
	return
}
