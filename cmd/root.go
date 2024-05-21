package cmd

import "github.com/spf13/cobra"

var masterPass string

var rootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Store secrets securely",
}

func init() {
	// Bind global flags to root command
	rootCmd.PersistentFlags().StringVarP(&masterPass, "pass", "p", "", "master password for vault")
	rootCmd.MarkPersistentFlagRequired("pass")
}

func Execute() error {
	return rootCmd.Execute()
}
