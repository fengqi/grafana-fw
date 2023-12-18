package main

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AlertGotifyHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	webhook := &GrafanaBody{}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		panic(err)
	}

	gotifyReq := &GotifyMessage{
		Title:    webhook.Title,
		Message:  webhook.Message,
		Priority: 100,
	}
	reqBytes, err := json.Marshal(gotifyReq)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(gotifyServer+"/message?token="+gotifyKey, "application/json", bytes2.NewReader(reqBytes))
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func AlertHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	webhook := &GrafanaBody{}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		panic(err)
	}

	barkReq := &BarkRequest{
		Title:     webhook.Title,
		Body:      webhook.Message,
		IsArchive: false,
	}
	reqBytes, err := json.Marshal(barkReq)
	if err != nil {
		//
	}

	resp, err := http.Post(server+"/"+key, "application/json", bytes2.NewReader(reqBytes))
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func DefaultHandle(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hi"))
}
