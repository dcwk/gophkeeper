package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func initGetUsersSecretsList() {
	getUsersSecretsList := &cobra.Command{
		Use:   "get-users-secrets-list",
		Short: "Get users secrets list",
		Run:   doGetUsersSecretsList,
	}
	rootCmd.AddCommand(getUsersSecretsList)
}

func doGetUsersSecretsList(_ *cobra.Command, _ []string) {
	log.Printf("User's secrets list")
}
