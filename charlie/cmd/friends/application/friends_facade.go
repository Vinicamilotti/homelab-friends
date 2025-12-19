package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Vinicamilotti/charlie/cmd/friends/domain"
	"github.com/Vinicamilotti/charlie/cmd/friends/outbounds"
	"github.com/google/uuid"
)

type FriendsFacade struct {
	FriendsRepository outbounds.FriendsRepository
	Client            *http.Client
}

func NewFriendsFacade(friendsRepository outbounds.FriendsRepository) *FriendsFacade {
	return &FriendsFacade{FriendsRepository: friendsRepository, Client: new(http.Client)}
}

func (f *FriendsFacade) GetFriends() ([]domain.Friend, error) {
	return f.FriendsRepository.GetFriends()
}

func (f *FriendsFacade) SendFriendInvitation(friendURL string, message string) error {
	hasFriend, err := f.FriendsRepository.GetFriend(friendURL)
	if err != nil {
		return err
	}
	if hasFriend.Id != "" {
		return fmt.Errorf("friend already exists")
	}
	friendUrl := fmt.Sprintf("%s/friends/request", friendURL)

	req := domain.FriendRequest{
		Dns:     os.Getenv("MY_DNS"),
		Name:    os.Getenv("MY_NAME"),
		Message: message,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := f.Client.Post(friendURL, "application/json", bytes.NewReader(body))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send friend invitation, status code: %d", res.StatusCode)
	}
	var friendReqRecived domain.FriendResquestRecived
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(resBody, &friendReqRecived); err != nil {
		return err
	}

	friendName := friendReqRecived.MyNameIs
	return f.FriendsRepository.AddFriendInvitation(domain.FriendRequest{
		Dns:     friendUrl,
		Name:    friendName,
		Message: message,
		Status:  domain.StatusSent,
	})
}

func (f *FriendsFacade) GetFriendInvitations() ([]domain.FriendRequest, error) {
	return f.FriendsRepository.GetFriendInvitations(domain.FriendRequest{
		Status: domain.StatusPending,
	})
}

func (f *FriendsFacade) ReciveFriendInvitation(req domain.FriendRequest) error {

	if sentResquet, err := f.FriendsRepository.GetFriendInvitations(domain.FriendRequest{
		Dns:    req.Dns,
		Status: domain.StatusSent}); err == nil && len(sentResquet) > 0 {
		// auto accept friend request if we have sent one to this dns
		return f.FriendsRepository.AcceptFriendInvitation(req.Dns)
	}
	err := f.FriendsRepository.AddFriendInvitation(domain.FriendRequest{
		Id:     uuid.NewString(),
		Dns:    req.Dns,
		Name:   req.Name,
		Status: domain.StatusPending,
	})

	return err
}
