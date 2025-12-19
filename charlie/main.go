package main

import (
	friendsFacade "github.com/Vinicamilotti/charlie/cmd/friends/application"
	friendsHandler "github.com/Vinicamilotti/charlie/cmd/friends/inbounds"
	friendsRepository "github.com/Vinicamilotti/charlie/cmd/friends/outbounds"
	"github.com/Vinicamilotti/charlie/cmd/shared/api"
	"github.com/Vinicamilotti/charlie/cmd/shared/store"
	"github.com/Vinicamilotti/charlie/secrets"
	"github.com/joho/godotenv"
)

func bootstrap() {
	godotenv.Load()
	store.InitializeDatabase()
	err := secrets.LoadScrets()
	if err != nil {
		panic(err)
	}

}

func main() {
	bootstrap()
	friendsFacade := friendsFacade.NewFriendsFacade(*friendsRepository.NewFriendsRepository())
	friendsHandler := friendsHandler.NewFriendsHandler(friendsFacade)

	api := api.NewApi("0.0.0.0", 8000)

	api.AddHandler(friendsHandler)
	api.Start()
}
