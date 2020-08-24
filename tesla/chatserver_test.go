package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPostMsg(t *testing.T) {
	msg := map[string]string{
		"user": "batman",
		"text": "hello",
    }
    msgJson, _ := json.Marshal(msg)
	res, err := http.Post("http://127.0.0.1:8081/message", "application/json", bytes.NewBuffer(msgJson))
	if err != nil {
		t.Error("failed")
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(body)
}

