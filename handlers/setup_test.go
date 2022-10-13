package handlers

import (
	"github.com/brane-app/librane/database"
	"github.com/brane-app/librane/types"

	"os"
	"testing"
)

const (
	nick     = "foobar"
	email    = "foo@bar.com"
	bio      = "mmm monke"
	password = "foobar2000"
)

var (
	user types.User = types.NewUser(nick, bio, email)
)

func TestMain(main *testing.M) {
	database.Connect(os.Getenv("DATABASE_CONNECTION"))
	database.Create()
	os.Exit(main.Run())
}
