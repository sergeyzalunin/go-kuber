package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sergeyzalunin/go-kuber/version"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	home(version.BuildTime, version.Commit, version.Release)(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	have := string(greeting)
	want := "{\"buildTime\":\"unset\",\"commit\":\"unset\",\"release\":\"unset\"}"

	if have != want {
		t.Errorf("The greeting is wrong. Have: %s, want: %s.", have, want)
	}
}
