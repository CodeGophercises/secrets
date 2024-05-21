package vault

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

type Vault map[string]string

var vaultPath string

func init() {
	dir, _ := homedir.Dir()
	vaultPath = filepath.Join(dir, ".secrets")
}
