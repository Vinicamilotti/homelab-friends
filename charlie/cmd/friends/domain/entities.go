package domain

const (
	StatusPending  = "PENDING"
	StatusAccepted = "ACCEPTED"
	StatusSent     = "SENT"
)

type FriendRequest struct {
	Id      string `json:"id,omitempty"`
	Dns     string `json:"dns"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Status  string `json:"status,omitempty"`
}

type Friend struct {
	Id   string `json:"id"`
	Dns  string `json:"dns"`
	Name string `json:"name"`
}

type FriendResquestRecived struct {
	MyNameIs string `json:"my_name_is"`
}
