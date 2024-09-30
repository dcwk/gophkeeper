package cli

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
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

	authRequest := &gophkeeper.AuthRequest{
		Login:    login,
		Password: password,
	}

	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := gophkeeper.NewGophkeeperClient(conn)
	resp, err := client.Auth(context.Background(), authRequest)
	if err != nil {
		log.Fatalf("Cannot auth user: %v", err)
	}
	log.Printf("User's jwt token", resp.JwtToken)
}
