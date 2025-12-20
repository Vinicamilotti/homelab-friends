package secrets

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type Secrets struct {
	AuthHash   string `json:"auth_hash"`
	PublicHash string `json:"public_hash"`
}

var secrets Secrets

func LoadScrets() error {
	file, err := os.ReadFile("secrets.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &secrets)

	return EnsureSecretsFile(false)

}

func EnsureSecretsFile(recreate bool) error {

	if secrets.AuthHash == "" || recreate {
		secrets.AuthHash = uuid.NewString()
	}

	if secrets.PublicHash == "" || recreate {
		secrets.PublicHash = uuid.NewString()
	}

	err := os.Remove("./secrets.json")
	if err != nil {
		return err
	}

	newSecretsContent, err := json.Marshal(secrets)

	return os.WriteFile("./secrets.json", newSecretsContent, os.ModeType)
}

func GetSecrets() Secrets {
	return secrets
}
