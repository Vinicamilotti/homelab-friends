package outbounds

import (
	"github.com/Vinicamilotti/charlie/cmd/friends/domain"
	"github.com/Vinicamilotti/charlie/cmd/shared/store"
	"github.com/google/uuid"
)

type FriendsRepository struct {
	Store store.SqliteConn
}

func NewFriendsRepository(store store.SqliteConn) *FriendsRepository {
	return &FriendsRepository{Store: store}
}

func (r *FriendsRepository) GetFriend(dns string) (domain.Friend, error) {
	sql := "SELECT id, dns, name FROM friends WHERE dns = ?"
	row, err := r.Store.DB().Query(sql, dns)
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
	sql := "SELECT id, dns, name FROM friends"
	rows, err := r.Store.DB().Query(sql)

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
	sql := "INSERT INTO friend_requests (id, name, dns, status) VALUES (?, ?, ?, ?)"
	_, err := r.Store.DB().Exec(sql, uuid.New().String(), invitation.Dns, invitation.Name, invitation.Status)
	return err
}

func (r *FriendsRepository) GetFriendInvitations(req domain.FriendRequest) ([]domain.FriendRequest, error) {
	sql := "SELECT id, name, message, status FROM friend_requests WHERE 1=1"

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

	rows, err := r.Store.DB().Query(sql, params...)

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
	sql := "UPDATE friend_requests SET status = ? WHERE dns = ?"
	_, err := r.Store.DB().Exec(sql, domain.StatusAccepted, dns)

	if err != nil {
		return err
	}

	sql = "INSERT INTO friends (id, dns, name) SELECT id, dns, name FROM friend_requests WHERE dns = ?"
	_, err = r.Store.DB().Exec(sql, dns)
	return err
}
