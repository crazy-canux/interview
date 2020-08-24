package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Msg struct {
	User string
	Text string
}

func MessageHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		msg, _ := ioutil.ReadAll(r.Body)
		f, _ := os.OpenFile("msg.txt",  os.O_APPEND, 0666)
		defer f.Close()
		_, err := f.Write(msg)
		var rsp = map[string]bool{}
		if err != nil {
			rsp["ok"] = true
		} else {
			rsp["ok"] = false
		}
		rspJson, _ := json.Marshal(rsp)
		rw.Write(rspJson)
	} else if r.Method == "GET" {
		var msgList = []Msg{}
		f, _ := os.Open("msg.txt")
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			msgList.append(scanner.Text())
		}
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
