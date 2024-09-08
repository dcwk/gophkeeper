package cli

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	grpcclient "github.com/dcwk/gophkeeper/internal/client/grpc"
	"github.com/dcwk/gophkeeper/proto"
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

	registerRequest := &proto.RegisterRequest{
		Login:    login,
		Password: password,
	}

	connection := grpcclient.OpenGrpcClientConnection()
	client := proto.NewGophkeeperClient(connection)
	resp, err := client.Register(context.Background(), registerRequest)
	if err != nil {
		log.Fatalf("Cannot register user: %v", err)
	}

	log.Printf("User successfully registered with id %s", resp.UserId)
}
