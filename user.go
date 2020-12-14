package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib"
	"git.gastrodon.io/imonke/monketype"

	"net/http"
)

var (
	no_such_user map[string]interface{} = map[string]interface{}{"error": "no_such_user"}
)

func getUserID(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = getUserKey("id", request)
	return
}

func getUserNick(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = getUserKey("nick", request)
	return
}

func getUserKey(key string, request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var parts []string = monkelib.SplitPath(request.URL.Path)
	var query string = parts[len(parts)-1]

	var fetched monketype.User
	var exists bool
	switch key {
	case "id":
		fetched, exists, err = monkebase.ReadSingleUser(query)
	case "nick":
		fetched, exists, err = monkebase.ReadSingleUserNick(query)
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
