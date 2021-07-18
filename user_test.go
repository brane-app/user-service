package main

import (
	"github.com/brane-app/database-library"
	"github.com/brane-app/types-library"

	"net/http"
	"testing"
)

const (
	nick  = "foobar"
	email = "foo@bar.com"
	bio   = "mmm monke"
)

var (
	user types.User = types.NewUser(nick, bio, email)
)

func userOK(test *testing.T, fetched map[string]interface{}, target types.User) {
	if fetched["id"].(string) != target.ID {
		test.Errorf("id mismatch! have: %s, want: %s", fetched["id"], target.ID)
	}

	if fetched["nick"].(string) != target.Nick {
		test.Errorf("nick mismatch! have: %s, want: %s", fetched["nick"], target.Nick)
	}

	if fetched["email"] != nil {
		test.Errorf("got public email %s!", fetched["email"])
	}

	if fetched["created"] != nil {
		test.Errorf("got public created %d!", fetched["created"])
	}
}

func setup(main *testing.M) {
	var err error
	if database.WriteUser(user.Map()); err != nil {
		panic(err)
	}
}

func Test_getUserID(test *testing.T) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest("GET", "http://imonke.co/user/id/"+user.ID, nil); err != nil {
		test.Fatal(err)
	}

	var code int
	var r_map map[string]interface{}
	if code, r_map, err = getUserID(request); err != nil {
		test.Fatal(err)
	}

	if code != 200 {
		test.Errorf("got code %d", code)
	}

	userOK(test, r_map["user"].(map[string]interface{}), user)
}

func Test_getUserID_noSuchUser(test *testing.T) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest("GET", "http://imonke.co/user/id/foobar", nil); err != nil {
		test.Fatal(err)
	}

	var code int
	var r_map map[string]interface{}
	if code, r_map, err = getUserID(request); err != nil {
		test.Fatal(err)
	}

	if code != 404 {
		test.Errorf("got code %d", code)
	}

	if r_map["error"] == nil || r_map["error"].(string) != "no_such_user" {
		test.Errorf("%#v", r_map)
	}
}

func Test_getUserNick(test *testing.T) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest("GET", "http://imonke.co/user/nick/"+user.Nick, nil); err != nil {
		test.Fatal(err)
	}

	var code int
	var r_map map[string]interface{}
	if code, r_map, err = getUserNick(request); err != nil {
		test.Fatal(err)
	}

	if code != 200 {
		test.Errorf("got code %d", code)
	}

	userOK(test, r_map["user"].(map[string]interface{}), user)
}

func Test_getUserNick_noSuchUser(test *testing.T) {
	var nobody string = "foobar"

	var stale types.User
	var exists bool
	var err error
	if stale, exists, err = database.ReadSingleUserNick(nobody); err != nil {
		test.Fatal(err)
	}

	if exists {
		database.DeleteUser(stale.ID)
	}

	var request *http.Request
	if request, err = http.NewRequest("GET", "http://imonke.co/user/nick/"+nobody, nil); err != nil {
		test.Fatal(err)
	}

	var code int
	var r_map map[string]interface{}
	if code, r_map, err = getUserNick(request); err != nil {
		test.Fatal(err)
	}

	if code != 404 {
		test.Errorf("got code %d", code)
	}

	if r_map["error"] == nil || r_map["error"].(string) != "no_such_user" {
		test.Errorf("%#v", r_map)
	}
}
