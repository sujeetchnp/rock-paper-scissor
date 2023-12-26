package main

import (
	service "github.com/sujeetchnp/rock-paper-scissor/service"
)

func main() {
	cacheService := service.NewCacheServiceImpl()
	gameService := service.NewGameServiceImpl()

	rps := NewRockPaperScissor(cacheService, gameService)
	rps.Run()
}
