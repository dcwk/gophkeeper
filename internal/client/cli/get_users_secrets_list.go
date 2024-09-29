package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
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
	GetUserSecretsList := gophkeeper.GetUserSecretsRequest{}
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := gophkeeper.NewGophkeeperClient(conn)
	resp, err := client.GetUserSecretsList(context.Background(), &GetUserSecretsList)
	if err != nil {
		log.Fatalf("Cannot get secrets list: %v", err)
	}

	fmt.Println("Secrets list:")
	for _, secret := range resp.Secrets {
		fmt.Println(secret)
	}
}
