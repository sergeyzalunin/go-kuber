package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHome(t *testing.T) {
	req, _ := http.NewRequest("GET", "/home", nil)
	w := httptest.NewRecorder()

	r := gin.New()
	r.GET("/home", home("", "", ""))

	r.ServeHTTP(w, req)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	var res InfoResponse
	err = json.Unmarshal(greeting, &res)
	if err != nil {
		t.Fatal(err)
	}

	have := string(greeting)
	want := "{\"buildTime\":\"\",\"commit\":\"\",\"release\":\"\"}"

	if have != want {
		t.Errorf("The greeting is wrong. Have: %s, want: %s.", have, want)
	}
}
