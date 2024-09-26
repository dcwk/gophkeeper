package cli

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func initCreateAuthPair(rootCmd *cobra.Command) {
	authUser := &cobra.Command{
		Use:   "create-auth-pair",
		Short: "Create user's auth pair",
		Run:   doCreateAuthPair,
	}
	rootCmd.AddCommand(authUser)

	authUser.Flags().StringP("name", "n", "", "Auth pair name")
	if err := authUser.MarkFlagRequired("name"); err != nil {
		log.Fatalf("unable to mark 'name' flag as required")
	}

	authUser.Flags().StringP("login", "l", "", "User's login")
	if err := authUser.MarkFlagRequired("login"); err != nil {
		log.Fatalf("unable to mark 'login' flag as required")
	}

	authUser.Flags().StringP("password", "p", "", "User's password")
	if err := authUser.MarkFlagRequired("password"); err != nil {
		log.Fatalf("unable to mark 'password' flag as required")
	}

	authUser.Flags().StringP("meta", "m", "", "Secret's metadata format key=value")
}

func doCreateAuthPair(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatalf("unable to get 'name' flag: %v", err)
	}

	login, err := cmd.Flags().GetString("login")
	if err != nil {
		log.Fatalf("unable to get 'login' flag: %v", err)
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		log.Fatalf("unable to get 'password' flag: %v", err)
	}

	metadata, err := cmd.Flags().GetString("meta")
	if err != nil {
		log.Fatalf("unable to get 'meta' flag: %v", err)
	}

	saveAuthPair := &gophkeeper.SaveAuthPairRequest{
		SecretName: name,
		Login:      login,
		Password:   password,
		Metadata:   metadata,
	}

	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := gophkeeper.NewGophkeeperClient(conn)
	resp, err := client.SaveAuthPair(context.Background(), saveAuthPair)
	if err != nil {
		log.Fatalf("Cannot register user: %v", err)
	}

	log.Printf("Success saved auth pair with name: %s id: %v", saveAuthPair, resp.SecretId)
}
