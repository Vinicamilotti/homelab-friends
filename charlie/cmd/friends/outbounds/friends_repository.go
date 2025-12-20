package outbounds

import (
	"fmt"

	"github.com/Vinicamilotti/charlie/cmd/friends/domain"
	"github.com/Vinicamilotti/charlie/store"
)

type FriendsRepository struct {
}

func NewFriendsRepository() *FriendsRepository {
	return &FriendsRepository{}
}

func (r *FriendsRepository) GetFriend(dns string) (domain.Friend, error) {
	conn := store.NewSqliteConn()
	sql := "SELECT id, dns, friend_name, friend_key FROM friends WHERE dns = ?"
	row, err := conn.DB().Query(sql, dns)
	if err != nil {
		return domain.Friend{}, err
	}
	defer row.Close()

	var friend domain.Friend
	for row.Next() {
		err := row.Scan(&friend.Id, &friend.Dns, &friend.FriendName, &friend.FriendKey)
		if err != nil {
			return domain.Friend{}, err
		}
	}

	return friend, nil
}

func (r *FriendsRepository) GetFriends() ([]domain.Friend, error) {
	conn := store.NewSqliteConn()
	sql := "SELECT id, dns, friend_name, friend_key FROM friends"
	db := conn.DB()
	defer db.Close()
	rows, err := db.Query(sql)

	if err != nil {
		return []domain.Friend{}, err
	}
	defer rows.Close()

	friends := []domain.Friend{}
	for rows.Next() {
		var friend domain.Friend
		err := rows.Scan(&friend.Id, &friend.Dns, &friend.FriendName, &friend.FriendKey)
		if err != nil {
			return []domain.Friend{}, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}

func (r *FriendsRepository) AddFriendInvitation(invitation domain.FriendRequest) error {
	conn := store.NewSqliteConn()
	db := conn.DB()
	defer db.Close()
	helper := store.NewDBHelper(db)
	err := helper.Insert("friend_requests", map[string]any{
		"id":              invitation.Id,
		"dns":             invitation.Dns,
		"request_message": invitation.RequestMessage,
		"request_status":  domain.StatusPending,
		"friend_name":     invitation.FriendName,
		"friend_key":      invitation.FriendKey,
	})
	return err
}

func (r *FriendsRepository) GetFriendInvitation(id string) (domain.FriendRequest, error) {
	getReq, err := r.GetFriendInvitations(domain.FriendRequest{
		Id: id,
	})

	if err != nil {
		return domain.FriendRequest{}, err
	}

	if len(getReq) == 0 {
		return domain.FriendRequest{}, fmt.Errorf("not found")
	}

	return getReq[0], nil

}

func (r *FriendsRepository) GetFriendInvitations(req domain.FriendRequest) ([]domain.FriendRequest, error) {
	sql := "SELECT id, friend_name, dns, request_message, request_status, friend_key FROM friend_requests WHERE 1=1"
	conn := store.NewSqliteConn()
	params := []interface{}{}
	if req.RequestStatus != "" {
		sql += " AND status = ?"
		params = append(params, req.RequestStatus)
	}

	if req.Dns != "" {
		sql += " AND dns = ?"
		params = append(params, req.Dns)
	}

	if req.Id != "" {
		sql += " AND id = ?"
		params = append(params, req.Id)
	}
	db := conn.DB()
	defer db.Close()
	rows, err := db.Query(sql, params...)

	if err != nil {
		return []domain.FriendRequest{}, err
	}
	defer rows.Close()

	invitations := []domain.FriendRequest{}
	for rows.Next() {
		var invitation domain.FriendRequest
		err := rows.Scan(&invitation.Id, &invitation.FriendName, &invitation.Dns, &invitation.RequestMessage, &invitation.RequestStatus, &invitation.FriendKey)
		if err != nil {
			return []domain.FriendRequest{}, err
		}
		invitations = append(invitations, invitation)
	}
	return invitations, nil
}

func (r *FriendsRepository) AcceptFriendInvitation(dns string) error {
	conn := store.NewSqliteConn()
	db := conn.DB()
	defer db.Close()
	sql := "UPDATE friend_requests SET status = ? WHERE dns = ?"
	_, err := db.Exec(sql, domain.StatusAccepted, dns)

	if err != nil {
		return err
	}

	sql = "INSERT INTO friends (id, dns, friend_name, friend_key) VALUES (SELECT id, dns, friend_name, friend_key FROM friend_requests WHERE dns = ?)"
	_, err = db.Exec(sql, dns)
	return err
}
