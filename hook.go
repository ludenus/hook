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
		
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				fmt.Printf("received event is NOT to be parsed: %+v\n", r)
			} else {
				fmt.Printf("ERROR: %+v\n", err)
			}
		}

		switch payload.(type) {
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			fmt.Printf("release: %+v\n", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			fmt.Printf("pull request: %+v\n", pullRequest)

		case github.PushPayload:
			pullRequest := payload.(github.PushPayload)
			fmt.Printf("push: %+v\n", pullRequest)
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