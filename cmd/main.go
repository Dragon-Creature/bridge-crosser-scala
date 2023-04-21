package main

import "git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/api"

func main() {
	err := api.Start()
	if err != nil {
		panic(err)
	}
}
