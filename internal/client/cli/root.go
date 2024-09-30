package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gophkeeper",
	Short: "gophkeeper is a CLI tool for managing user's secret data like a pay card data, login pass pair, etc",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initRegister(rootCmd)
	initAuth(rootCmd)
	initGetUsersSecretsList()
	initCreateAuthPair(rootCmd)
	initSaveBinaryData(rootCmd)
}
