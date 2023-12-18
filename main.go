package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	server       string
	key          string
	gotifyServer string
	gotifyKey    string
)

func init() {
	flag.StringVar(&server, "server", "https://api.day.app", "bark server")
	flag.StringVar(&key, "key", "", "bark key")

	flag.StringVar(&gotifyServer, "gotify-server", "http://gotify.fengqi.io", "gotify server")
	flag.StringVar(&gotifyKey, "gotify-key", "", "gotify key")
}

func main() {
	flag.Parse()
	if key == "" {
		flag.Usage()
		return
	}

	http.HandleFunc("/", DefaultHandle)
	http.HandleFunc("/alert", AlertHandle)
	http.HandleFunc("/alert-gotify", AlertGotifyHandle)

	fmt.Println("serve run at :8080")
	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}
}
