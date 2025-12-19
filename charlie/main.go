package main

import (
	friendsFacade "github.com/Vinicamilotti/charlie/cmd/friends/application"
	friendsHandler "github.com/Vinicamilotti/charlie/cmd/friends/inbounds"
	friendsRepository "github.com/Vinicamilotti/charlie/cmd/friends/outbounds"
	"github.com/Vinicamilotti/charlie/cmd/shared/api"
	"github.com/Vinicamilotti/charlie/cmd/shared/store"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	store.InitializeDatabase()
	friendsFacade := friendsFacade.NewFriendsFacade(*friendsRepository.NewFriendsRepository())
	friendsHandler := friendsHandler.NewFriendsHandler(friendsFacade)

	api := api.NewApi("0.0.0.0", 8000)

	api.AddHandler(friendsHandler)
	api.Start()
}
