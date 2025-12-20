package inbounds

import (
	"io"
	"os"

	"github.com/Vinicamilotti/charlie/cmd/friends/application"
	"github.com/Vinicamilotti/charlie/cmd/friends/domain"
	"github.com/Vinicamilotti/charlie/cmd/shared/lib"
	"github.com/Vinicamilotti/charlie/cmd/shared/middleware"
	"github.com/gin-gonic/gin"
)

type FriendsHandler struct {
	FriendsFacade *application.FriendsFacade
}

func NewFriendsHandler(friendsFacade *application.FriendsFacade) *FriendsHandler {
	return &FriendsHandler{FriendsFacade: friendsFacade}
}

func (h *FriendsHandler) GetFriends(c *gin.Context) {
	friends, err := h.FriendsFacade.GetFriends()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, map[string]any{"friends": friends})
}

func (h *FriendsHandler) SendFriendInvitation(c *gin.Context) {
	var body io.Reader = c.Request.Body
	invitationRequest, err := lib.ReadBody[domain.FriendRequest](body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid friend request body"})
		return
	}
	err = h.FriendsFacade.SendFriendInvitation(invitationRequest.Dns, invitationRequest.RequestMessage)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Friend invitation sent successfully"})
}

func (h *FriendsHandler) InviteMe(c *gin.Context) {
	var body io.Reader = c.Request.Body
	invitationRequest, err := lib.ReadBody[domain.FriendRequest](body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid friend request body"})
		return
	}
	err = h.FriendsFacade.ReciveFriendInvitation(invitationRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp := domain.FriendRequestRecived{
		MyNameIs: os.Getenv("MY_NAME"),
	}

	c.JSON(200, resp)
}

func (h *FriendsHandler) GetFriendRequests(c *gin.Context) {
	friendRequests, err := h.FriendsFacade.GetFriendInvitations()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, friendRequests)
}

func (h *FriendsHandler) RegisterRoutes(app *gin.Engine) {
	app.GET("/friends", middleware.AuthOwner, h.GetFriends)
	app.POST("/friends/invite", middleware.AuthOwner, h.SendFriendInvitation)
	app.POST("/friends/request", h.InviteMe)
	app.GET("/friends/request", middleware.AuthOwner, h.GetFriendRequests)
}
