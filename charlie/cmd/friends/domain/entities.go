package domain

const (
	StatusPending  = "PENDING"
	StatusAccepted = "ACCEPTED"
	StatusSent     = "SENT"
)

type FriendRequest struct {
	Id        string `json:"id,omitempty"`
	Dns       string `json:"dns"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	FriendKey string `json:"friend_key"`
	Status    string `json:"status,omitempty"`
}

type Friend struct {
	Id        string `json:"id"`
	Dns       string `json:"dns"`
	Name      string `json:"name"`
	FriendKey string `json:"friend_key"`
}

type FriendRequestRecived struct {
	MyNameIs string `json:"my_name_is"`
}
