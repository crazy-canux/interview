package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Msg struct {
	TimeStamp float64
	User string
	Text string
}

func MessageHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		msg, _ := ioutil.ReadAll(r.Body)
		result := true
		f, err := os.OpenFile("msg.txt",  os.O_RDWR | os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("open file error: %v", err)
			result = false
		} else{
			defer f.Close()
			_, err := f.Write(msg)
			if err != nil {
				fmt.Printf("write msg error: %v", err)
				result = false
			}
		}
		var rsp = map[string]bool{}
		if result {
			rsp["ok"] = true
		} else {
			rsp["ok"] = false
		}
		rspJson, _ := json.Marshal(rsp)
		rw.Write(rspJson)
	} else if r.Method == "GET" {
		msgs, err := ioutil.ReadFile("msg.txt")
		if err != nil {
			fmt.Printf("read msg error: %v", err)
		}
		var msgList []Msg
		json.Unmarshal(msgs, &msgList)
		rsp := map[string][]Msg{
			"messages": msgList,
		}
		rspJson, _ := json.Marshal(rsp)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(rspJson)
	}
}

func main() {
	http.HandleFunc("/message", MessageHandler)
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
