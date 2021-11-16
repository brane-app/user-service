package types

import (
	"github.com/brane-app/tools-library"
	"github.com/gastrodon/groudon/v2"
)

func ValidNick(it interface{}) (ok bool, _ error) {
	var nick string
	if nick, ok = it.(string); !ok {
		return
	}

	ok = tools.NickRegex.MatchString(nick)
	return
}

func ValidPass(it interface{}) (ok bool, _ error) {
	var pass string
	if pass, ok = it.(string); !ok {
		return
	}

	ok = len(pass) >= 8
	return
}

func ValidBio(it interface{}) (ok bool, _ error) {
	var bio string
	if bio, ok = it.(string); !ok {
		return
	}

	ok = len(bio) <= 63
	return
}

type CreateUserBody struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

func (_ CreateUserBody) Validators() (values map[string]func(interface{}) (bool, error)) {
	values = map[string]func(interface{}) (bool, error){
		"nick":     ValidNick,
		"password": ValidPass,
		"email":    groudon.ValidEmail,
		"bio":      ValidBio,
	}

	return
}

func (_ CreateUserBody) Defaults() (values map[string]interface{}) {
	values = map[string]interface{}{
		"bio": "",
	}

	return
}
