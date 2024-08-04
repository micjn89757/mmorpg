package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)


func TestLogin(t *testing.T) {
	type Login struct {
		Username string	
		Password string
	}

	packet := &Login{
		Username: "ddd",
		Password: "ddd",
	}
	data, err := json.Marshal(packet)
	if err != nil {
		t.Error(err)
	}

	res, err := http.Post("http://127.0.0.1:3000/login", "application/json", bytes.NewBuffer(data))
	
	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

}