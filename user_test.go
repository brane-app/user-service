package main

import (
	"github.com/brane-app/database-library"
	"github.com/brane-app/types-library"

	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

const (
	email    = "foo@bar.com"
	password = "foobar2000"
	nick     = "foobar"
)

var (
	user types.User
)

func mustMarshal(it interface{}) (data []byte) {
	var err error
	if data, err = json.Marshal(it); err != nil {
		panic(err)
	}

	return
}

func setup(main *testing.M) {
	user = types.NewUser(nick, "", email)

	var err error
	if err = database.WriteUser(user.Map()); err != nil {
		panic(err)
	}
}

func Test_PostUser_badrequest(test *testing.T) {
	var set []byte
	var sets [][]byte = [][]byte{
		mustMarshal(map[string]string{
			"nick":  nick,
			"email": email,
		}),
		mustMarshal(map[string]string{
			"nick":     nick,
			"password": password,
		}),
		mustMarshal(map[string]string{
			"email":    email,
			"password": password,
		}),
		mustMarshal(map[string]string{
			"email": email,
		}),
		mustMarshal(map[string]string{
			"nick": nick,
		}),
		mustMarshal(map[string]string{
			"password": password,
		}),
		mustMarshal(map[string]string{
			"nick":     nick,
			"email":    "bad_email",
			"password": password,
		}),
		[]byte("benis lol"),
	}

	var request *http.Request

	var code int
	var err error
	for _, set = range sets {
		if request, err = http.NewRequest("POST", "https://imonke/user", bytes.NewReader(set)); err != nil {
			test.Fatal(err)
		}

		if code, _, err = postUser(request); err != nil {
			test.Fatal(err)
		}

		if code != 400 {
			test.Errorf("%s got wrong code %d", string(set), code)
		}
	}
}

func Test_PostUser_Conflicts(test *testing.T) {
	var key string
	var set []byte
	var sets map[string][]byte = map[string][]byte{
		"email": mustMarshal(map[string]string{
			"email":    email,
			"nick":     "unused",
			"password": "longer_unused",
		}),
		"nick": mustMarshal(map[string]string{
			"email":    "unused@bar.com",
			"nick":     nick,
			"password": "longer_unused",
		}),
	}

	var request *http.Request
	var code int
	var r_map map[string]interface{}
	var err error
	for key, set = range sets {
		if request, err = http.NewRequest("POST", "https://imonke.io/", bytes.NewReader(set)); err != nil {
			test.Fatal(err)
		}

		if code, r_map, err = postUser(request); err != nil {
			test.Fatal(err)
		}

		if code != 409 {
			test.Errorf("got code %d", code)
		}

		if r_map["key"] == nil || r_map["key"].(string) != key {
			test.Errorf("did not get conflict %s in %#v", key, r_map)
		}
	}
}

func Test_PostUser(test *testing.T) {
	var new_nick, new_email string = "longer_new", "new@bar.com"
	var data []byte = mustMarshal(map[string]interface{}{
		"nick":     new_nick,
		"email":    new_email,
		"password": password,
	})

	var request *http.Request
	var err error
	if request, err = http.NewRequest("POST", "https://imonke.io/", bytes.NewReader(data)); err != nil {
		test.Fatal(err)
	}

	var code int
	var r_map map[string]interface{}
	if code, r_map, err = postUser(request); err != nil {
		test.Fatal(err)
	}

	if code != 200 {
		test.Errorf("got code %d", code)
	}

	var id string = r_map["user"].(map[string]interface{})["id"].(string)
	defer database.DeleteUser(id)

	var exists bool
	if _, exists, err = database.ReadSingleUser(id); err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("user %s does not exist", id)
	}
}
