package main

import (
	"github.com/brane-app/database-library"
	"github.com/brane-app/types-library"
	"github.com/gastrodon/groudon/v2"

	"net/http"
	"regexp"
)

var (
	email_regex *regexp.Regexp = regexp.MustCompile(`^[\w\+]+@[\w]+\.[\w\.]+$`)
)

func checkConflicts(body CreateUserBody) (conflicts bool, key string, err error) {
	if _, conflicts, err = database.ReadSingleUserNick(body.Nick); conflicts || err != nil {
		key = "nick"
		return
	}

	if _, conflicts, err = database.ReadSingleUserEmail(body.Email); conflicts || err != nil {
		key = "email"
		return
	}

	return
}

func postUser(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var body CreateUserBody
	var external error
	err, external = groudon.SerializeBody(request.Body, &body)
	switch {
	case err != nil:
		return
	case external != nil, !email_regex.MatchString(body.Email):
		code = 400
		return
	}

	var key string
	var conflicts bool
	if conflicts, key, err = checkConflicts(body); conflicts || err != nil {
		code = 409
		r_map = map[string]interface{}{
			"error": "conflict",
			"key":   key,
		}

		return
	}

	var created map[string]interface{} = types.NewUser(body.Nick, body.Bio, body.Email).Map()
	if err = database.WriteUser(created); err != nil {
		return
	}
	if err = database.SetPassword(created["id"].(string), body.Password); err != nil {
		return
	}

	code = 200
	r_map = map[string]interface{}{
		"user": created,
	}
	return
}
