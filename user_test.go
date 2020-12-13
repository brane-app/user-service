package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monketype"

	"net/http"
	"os"
	"testing"
)

const (
	nick  = "foobar"
	email = "foo@bar.com"
	bio   = "mmm monke"
)

var (
	user monketype.User = monketype.NewUser(nick, bio, email)
)

func userOK(test *testing.T, fetched map[string]interface{}, target monketype.User) {
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

func TestMain(main *testing.M) {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))

	var err error
	if monkebase.WriteUser(user.Map()); err != nil {
		panic(err)
	}

	var result int = main.Run()
	monkebase.DeleteUser(user.ID)
	os.Exit(result)
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

	var stale monketype.User
	var exists bool
	var err error
	if stale, exists, err = monkebase.ReadSingleUserNick(nobody); err != nil {
		test.Fatal(err)
	}

	if exists {
		monkebase.DeleteUser(stale.ID)
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
