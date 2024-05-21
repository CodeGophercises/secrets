`secrets` stores your secrets ( key-value pairs ) securely and conveniently through simple CLI commands. 

Internally uses a map to store all secrets, serilaised using `gob` and persisted to `secrets.gob` in your home dir.

The file is encrypted using AES-GCM for confidentiality and integrity.
