package handlers

import (
	"github.com/brane-app/user-service/types"

	"github.com/brane-app/database-library"
	library_types "github.com/brane-app/types-library"
	"github.com/gastrodon/groudon/v2"

	"net/http"
)

func checkConflicts(body types.CreateUserBody) (conflicts bool, key string, err error) {
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

func PostUser(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var body types.CreateUserBody
	var external error
	err, external = groudon.SerializeBody(request.Body, &body)
	switch {
	case err != nil:
		return
	case external != nil:
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

	var created map[string]interface{} = library_types.NewUser(body.Nick, body.Bio, body.Email).Map()
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
