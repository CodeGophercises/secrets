package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/argon2"
)

var nonce, salt []byte

func init() {
	//TODO: Do not reuse nonce, its called nonce for a reason
	nonce, _ = hex.DecodeString("64a9433eae7ccceee2fc0eda")
	salt = []byte("Rand0m3n0ugh$alt#")
}

func getKeyFromPassphrase(passw string) []byte {
	return argon2.IDKey([]byte(passw), salt, 1, 64*1024, 4, 32)
}

func DecryptContent(f *os.File, masterPass string) (*bytes.Buffer, error) {
	gobData, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	key := getKeyFromPassphrase(masterPass)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, gobData, nil)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(plaintext)
	return buf, nil

}

func EncryptContent(f *os.File, buf *bytes.Buffer, masterPass string) error {
	key := getKeyFromPassphrase(masterPass)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	cipherText := aesgcm.Seal(nil, nonce, buf.Bytes(), nil)
	_, err = f.Write(cipherText)
	return err
}
