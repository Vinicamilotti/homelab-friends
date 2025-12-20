package secrets

import (
	"encoding/json"
	"os"
)

type Secrets struct {
	AuthHash string `json:"auth_hash"`
}

var secrets Secrets

func LoadScrets() error {
	file, err := os.ReadFile("secrets.json")
	if err != nil {
		panic(err)
	}

	secrets = GetSecrets()

	err = json.Unmarshal(file, secrets)
	return err

}

func GetSecrets() Secrets {
	return secrets
}
