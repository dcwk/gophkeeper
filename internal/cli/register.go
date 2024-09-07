package cli

import (
	"log"

	"github.com/spf13/cobra"
)

var registerUser = &cobra.Command{
	Use:   "register",
	Short: "Register a user",
	Run:   doRegister,
}

func initRegister(rootCmd *cobra.Command) {
	rootCmd.AddCommand(registerUser)

	registerUser.Flags().StringP("login", "l", "", "User's login")
	if err := registerUser.MarkFlagRequired("login"); err != nil {
		log.Fatalf("unable to mark 'login' flag as required")
	}

	registerUser.Flags().StringP("password", "p", "", "User's password")
	if err := registerUser.MarkFlagRequired("password"); err != nil {
		log.Fatalf("unable to mark 'password' flag as required")
	}
}

func doRegister(cmd *cobra.Command, args []string) {
	login, err := cmd.Flags().GetString("login")
	if err != nil {
		log.Fatalf("unable to get 'login' flag: %v", err)
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		log.Fatalf("unable to get 'password' flag: %v", err)
	}

	log.Printf("Registering user with login: %s, pass: %s", login, password)
}
