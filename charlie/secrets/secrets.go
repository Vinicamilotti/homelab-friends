package secrets

import (
	"encoding/json"
	"log"
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
		log.Print("Secrets file not found. Creating ...")
		return EnsureSecretsFile(true)
	}

	err = json.Unmarshal(file, &secrets)
	if err != nil {
		log.Print("Malformed secrets file, recreating ...")
		return EnsureSecretsFile(true)
	}

	return EnsureSecretsFile(false)

}

func EnsureSecretsFile(newFile bool) error {

	if secrets.AuthHash == "" || newFile {
		secrets.AuthHash = uuid.NewString()
	}

	if secrets.PublicHash == "" || newFile {
		secrets.PublicHash = uuid.NewString()
	}

	err := os.Remove("./secrets.json")
	if err != nil && !newFile {
		return err
	}

	newSecretsContent, err := json.Marshal(secrets)

	return os.WriteFile("./secrets.json", newSecretsContent, os.ModeType)
}

func GetSecrets() Secrets {
	return secrets
}
