package vault

import (
	"bytes"
	"encoding/gob"
	"io"
	"os"

	"github.com/CodeGophercises/secrets/encrypt"
)

func truncateFile(f *os.File) error {
	return f.Truncate(0)
}

func loadSecretsFile(masterPass string) (*bytes.Buffer, error) {
	_, err := os.Stat(vaultPath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(vaultPath)
			if err != nil {
				return nil, err
			}
			var empty []byte
			return bytes.NewBuffer(empty), nil

		} else {
			return nil, err
		}
	}

	f, err := os.Open(vaultPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return encrypt.DecryptContent(f, masterPass)
}

func StoreInVault(key, value, masterPass string) error {
	buf, err := loadSecretsFile(masterPass)
	if err != nil {
		return err
	}
	vault := make(Vault)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&vault)
	if err != nil && err != io.EOF {
		return err
	}
	vault[key] = value
	f, err := os.OpenFile(vaultPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	err = truncateFile(f)
	if err != nil {
		return err
	}
	var nbuf bytes.Buffer
	enc := gob.NewEncoder(&nbuf)
	err = enc.Encode(vault)
	if err != nil {
		return err
	}
	encrypt.EncryptContent(f, &nbuf, masterPass)
	if err != nil {
		return err
	}
	return nil
}
