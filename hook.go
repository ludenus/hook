package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret(secret()))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(".")
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				fmt.Println("ok event wasn;t one of the ones asked to be parsed")
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
	http.ListenAndServe(addr(), nil)
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

	fmt.Printf("%s", addr)
	return addr
}

func secret() string {

	secret := os.Getenv("HOOK_SECRET")

	if secret == "" {
		panic("ERROR: HOOK_SECRET environment variable is not set")
	}

	return secret
}