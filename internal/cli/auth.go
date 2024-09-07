package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func initAuth(rootCmd *cobra.Command) {
	authUser := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate user",
		Run:   doAuth,
	}
	rootCmd.AddCommand(authUser)

	authUser.Flags().StringP("login", "l", "", "User's login")
	if err := authUser.MarkFlagRequired("login"); err != nil {
		log.Fatalf("unable to mark 'login' flag as required")
	}

	authUser.Flags().StringP("password", "p", "", "User's password")
	if err := authUser.MarkFlagRequired("password"); err != nil {
		log.Fatalf("unable to mark 'password' flag as required")
	}
}

func doAuth(cmd *cobra.Command, args []string) {
	login, err := cmd.Flags().GetString("login")
	if err != nil {
		log.Fatalf("unable to get 'login' flag: %v", err)
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		log.Fatalf("unable to get 'password' flag: %v", err)
	}

	log.Printf("Authenticated user with login: %s, pass: %s", login, password)
}
