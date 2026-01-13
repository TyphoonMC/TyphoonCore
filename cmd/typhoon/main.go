package main

import (
	typhoon "github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := typhoon.Init()
	core.SetBrand("typhoon")

	core.On(func(e *typhoon.PlayerJoinEvent) {
		e.Player.SendMessage(typhoon.ChatMessage("Welcome to Typhoon Server!"))
	})

	core.Start()
}
