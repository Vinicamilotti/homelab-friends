package outbounds

import (
	"github.com/Vinicamilotti/charlie/cmd/friends/domain"
	"github.com/Vinicamilotti/charlie/cmd/shared/store"
)

type FriendsRepository struct {
}

func NewFriendsRepository() *FriendsRepository {
	return &FriendsRepository{}
}

func (r *FriendsRepository) GetFriend(dns string) (domain.Friend, error) {
	conn := store.NewSqliteConn()
	sql := "SELECT id, dns, name FROM friends WHERE dns = ?"
	row, err := conn.DB().Query(sql, dns)
	if err != nil {
		return domain.Friend{}, err
	}
	defer row.Close()

	var friend domain.Friend
	for row.Next() {
		err := row.Scan(&friend.Id, &friend.Dns, &friend.Name)
		if err != nil {
			return domain.Friend{}, err
		}
	}

	return friend, nil
}

func (r *FriendsRepository) GetFriends() ([]domain.Friend, error) {
	conn := store.NewSqliteConn()
	sql := "SELECT id, dns, name FROM friends"
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
		err := rows.Scan(&friend.Id, &friend.Dns, &friend.Name)
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
		"id":      invitation.Id,
		"dns":     invitation.Dns,
		"name":    invitation.Name,
		"message": invitation.Message,
		"status":  domain.StatusPending,
	})
	return err
}

func (r *FriendsRepository) GetFriendInvitations(req domain.FriendRequest) ([]domain.FriendRequest, error) {
	sql := "SELECT id, name, message, status FROM friend_requests WHERE 1=1"
	conn := store.NewSqliteConn()
	params := []interface{}{}
	if req.Status != "" {
		sql += " AND status = ?"
		params = append(params, req.Status)
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
		err := rows.Scan(&invitation.Id, &invitation.Dns, &invitation.Name, &invitation.Message, &invitation.Status)
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

	sql = "INSERT INTO friends (id, dns, name) SELECT id, dns, name FROM friend_requests WHERE dns = ?"
	_, err = db.Exec(sql, dns)
	return err
}
