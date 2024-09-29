package api

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func (c *Controller) SaveBinaryData(ctx context.Context, req *gophkeeper.SaveBinaryDataRequest) (*gophkeeper.SaveBinaryDataResponse, error) {
	binaryData := bytes.Buffer{}
	binarySize := 0

	for {
		chunk := req.GetChunkData()
		size := len(chunk)

		binarySize += size
		_, err := binaryData.Write(chunk)
		if err != nil {
			return nil, err
		}
	}

	imagePath := fmt.Sprintf("./%s/%s.%s", "img", "test", req.GetInfo().Ext)
	file, err := os.Create(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot create file: %w", err)
	}

	_, err = binaryData.WriteTo(file)
	if err != nil {
		return nil, fmt.Errorf("cannot write data: %w", err)
	}

	return &gophkeeper.SaveBinaryDataResponse{FileId: 1, Size: int64(binarySize)}, nil
}
