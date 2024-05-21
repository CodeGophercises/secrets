package cmd

import (
	"fmt"

	"github.com/CodeGophercises/secrets/vault"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get secret from vault.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		val, err := vault.FetchFromVault(key, masterPass)
		if err != nil {
			return err
		}
		fmt.Printf("Fetched secret is: %s\n", val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
