package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

const (
	Author  = "Francis Bouvier <francis.bouvier@gmail.com>"
	VERSION = "0.1.0"
)

var portFlag = cli.StringFlag{
	Name:  "port, p",
	Value: "8085",
	Usage: "Port to run the server on.",
}

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := r.FormValue("code")
	scope := r.FormValue("scope")
	state := r.FormValue("state")
	if code == "" || scope == "" || state == "" {
		http.Error(w, "Not enough parameters", http.StatusBadRequest)
		return
	}
	data := struct {
		Code  string `json:"code"`
		Scope string `json:"scope"`
		State string `json:"state"`
	}{
		Code:  code,
		Scope: scope,
		State: state,
	}
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}

func main() {
	app := cli.NewApp()
	app.Author = Author
	app.Version = VERSION
	app.Usage = "HubiC oauth app server"
	app.Flags = []cli.Flag{portFlag}
	app.Action = func(c *cli.Context) {
		port := c.String("port")
		fmt.Printf("Running server on :%s ...\n", port)
		http.HandleFunc("/", oauthHandler)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
	app.Run(os.Args)
}
