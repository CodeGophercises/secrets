package vault

import (
	"encoding/gob"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/CodeGophercises/secrets/encrypt"
)

func FetchFromVault(key, masterPass string) (string, error) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
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
		if err == io.EOF {
			return "", errors.New("there are no secrets stored.")
		}
		return "", err
	}
	val, ok := vault[key]
	if !ok {
		return "", errors.New("no such key in vault")
	}
	return val, nil
}
