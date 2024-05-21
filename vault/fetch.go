package vault

import (
	"encoding/gob"
	"errors"
	"os"

	"github.com/CodeGophercises/secrets/encrypt"
)

func FetchFromVault(key, masterPass string) (string, error) {
	f, err := os.Open(vaultPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buf, err := encrypt.DecryptContent(f, masterPass)
	if err != nil {
		return "", err
	}
	dec := gob.NewDecoder(buf)
	vault := make(Vault)
	err = dec.Decode(&vault)
	if err != nil {
		return "", err
	}
	val, ok := vault[key]
	if !ok {
		return "", errors.New("no such key in vault")
	}
	return val, nil
}
