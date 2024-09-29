package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func initSaveBinaryDataCmd(rootCmd *cobra.Command) {
	saveBinaryData := &cobra.Command{
		Use:   "save-binary-data",
		Short: "Save user's binary data",
		Run:   doSaveBinaryData,
	}
	rootCmd.AddCommand(saveBinaryData)

	saveBinaryData.Flags().StringP("name", "n", "", "Secret name")
	//if err := saveBinaryData.MarkFlagRequired("name"); err != nil {
	//	log.Fatalf("unable to mark 'name' flag as required")
	//}

	saveBinaryData.Flags().StringP("file", "f", "", "Path to the binary data file")
	if err := saveBinaryData.MarkFlagRequired("name"); err != nil {
		log.Fatalf("unable to mark 'name' flag as required")
	}

	saveBinaryData.Flags().StringP("metadata", "m", "", "Metadata for secret")
}

func doSaveBinaryData(cmd *cobra.Command, args []string) {
	//name, err := cmd.Flags().GetString("name")
	//if err != nil {
	//	log.Fatalf("unable to get 'name' flag: %v", err)
	//}

	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		log.Fatalf("unable to get 'name' flag: %v", err)
	}

	//metadata, err := cmd.Flags().GetString("meta")
	//if err != nil {
	//	log.Fatalf("unable to get 'meta' flag: %v", err)
	//}

	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal("cannot open file", err)
	}
	defer file.Close()

	client := gophkeeper.NewGophkeeperClient(conn)
	stream, err := client.SaveBinaryData(context.Background())
	if err != nil {
		log.Fatalf("Cannot get secrets list: %v", err)
	}

	req := &gophkeeper.SaveBinaryDataRequest{
		Data: &gophkeeper.SaveBinaryDataRequest_Info{
			Info: &gophkeeper.BinaryInfo{
				Type: gophkeeper.BinaryType_IMAGE_FILE,
				Ext:  "png",
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatalf("Cannot send binary data info to server: %v", err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &gophkeeper.SaveBinaryDataRequest{
			Data: &gophkeeper.SaveBinaryDataRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatalf("Cannot send chunk to server: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	fmt.Println(fmt.Sprintf("successfully saved binary data with id %v and size %v", resp.FileId, resp.Size))
}
