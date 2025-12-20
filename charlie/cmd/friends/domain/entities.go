package domain

const (
	StatusPending  = "PENDING"
	StatusAccepted = "ACCEPTED"
	StatusSent     = "SENT"
)

type FriendRequest struct {
	Id             string `json:"id,omitempty"`
	Dns            string `json:"dns"`
	FriendName     string `json:"friend_name"`
	RequestMessage string `json:"request_message"`
	FriendKey      string `json:"friend_key"`
	RequestStatus  string `json:"request_status,omitempty"`
}

type Friend struct {
	Id         string `json:"id"`
	Dns        string `json:"dns"`
	FriendName string `json:"name"`
	FriendKey  string `json:"friend_key"`
}

type FriendRequestRecived struct {
	MyNameIs string `json:"my_name_is"`
}
