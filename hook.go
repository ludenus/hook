package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kataras/iris"
	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

func main() {
	hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecrect...?"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	http.ListenAndServe(":3000", nil)
}

func main2() {

	addr := addr()

	app := iris.Default()

	app.Logger().SetLevel("debug")

	// https://github.com/go-playground/webhooks

	app.Post("/", func(ctx iris.Context) {

		app.Logger().Debug("~~~~~~~~~~~~~~~~~~~~ ctx.Path: " + ctx.Path())

		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			app.Logger().Error("1111111111111111111111")
			app.Logger().Error(err.Error())
		}

		t := time.Now().Format("20060102150405")

		ctx.JSON(iris.Map{
			"pong": t,
			"user": user,
		})
	})

	app.Run(iris.Addr(addr))
}

func addr() string {

	addr := ":80"
	fromEnv := os.Getenv("HOOK_LISTENING_ADDRESS")

	if fromEnv != "" {
		addr = fromEnv
	}

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	return addr
}
