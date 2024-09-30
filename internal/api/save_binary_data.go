package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func (c *Controller) SaveBinaryData(stream grpc.ClientStreamingServer[gophkeeper.SaveBinaryDataRequest, gophkeeper.SaveBinaryDataResponse]) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive binary info")
	}

	//binaryType := req.GetInfo().GetType()
	binaryExt := req.GetInfo().GetExt()

	binaryData := bytes.Buffer{}
	binarySize := 0

	for {
		log.Print("waiting to receive a binary data")

		req, err = stream.Recv()
		if err == io.EOF {
			log.Print("finished receiving binary data")
			break
		}

		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		binarySize += size
		_, err := binaryData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data %v", err)
		}
	}

	imagePath := fmt.Sprintf("%s/%s.%s", "/Users/ruslan.golovizin/Projects/practicum/gophkeeper/storage", "test", binaryExt)
	file, err := os.Create(imagePath)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot create file: %v", err)
	}

	_, err = binaryData.WriteTo(file)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot write data: %v", err)
	}

	resp := &gophkeeper.SaveBinaryDataResponse{FileId: 1, Size: int64(binarySize)}
	err = stream.SendAndClose(resp)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot send response: %v", err)
	}

	return nil
}
