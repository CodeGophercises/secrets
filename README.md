`secrets` stores your secrets ( key-value pairs ) securely and conveniently through simple CLI commands. 

Internally uses a map to store all secrets, serilaised using `gob` and persisted to `.secrets` in your home dir.

The file is encrypted using AES-GCM for confidentiality and integrity.

Safe for concurrent use by multiple goroutines ( although, why would you? )