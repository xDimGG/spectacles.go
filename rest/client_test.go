package rest_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/spec-tacles/spectacles/rest"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestMake(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/users/@me", nil)
	if err != nil {
		t.Error(err)
		return
	}

	client := rest.NewClient(os.Getenv("DISCORD_TOKEN"))
	res, err := client.Make(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)
	if res.StatusCode >= 400 {
		t.Fail()
	}
}
