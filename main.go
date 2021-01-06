package main

import (
	"fmt"

	"github.com/task4233/techtrain-mission/gameapi"
)

func main() {
	api := gameapi.NewGameAPI()
	if err := api.Run(); err != nil {
		panic(fmt.Sprintf("failed Run API: %v", err))
	}
}
